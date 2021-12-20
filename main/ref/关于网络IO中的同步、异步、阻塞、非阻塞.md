# 关于网络IO中的同步、异步、阻塞、非阻塞

====================

在高并发编程当中, 我们经常会遇到一些异步、非阻塞等一些概念, 一些常用的技术比如异步的httpclient、netty nio、nginx、node.js等, 它们的原理大都跟异步、非阻塞有关.特别是在服务器开发中, 并发的请求处理是个大问题, 阻塞式的函数会导致资源浪费和时间延迟.通过事件注册、异步函数, 开发人员可以提高资源的利用率, 性能也会改善.其nginx和node.js处理并发都是采用的事件驱动异步非阻塞模式.其中nginx中处理并发用的是epoll, poll,queue等方式, node.js使用的是libev, 它们对大规模的HTTP请求处理的都很好.

那么到底什么是异步、非阻塞, 它们的原理是什么, 它们之间又有什么区别呢？其实在很多情况下, 异步与非阻塞(同步与阻塞)表示的是同一个意思, 但是在特定的上下文环境中, 它们含义又十分不同.再具体讲它们的区别之前, 先介绍一下上下文背景.

## 一、上下文背景

我们所遇到的这些场景大部分都是当用户进程(或线程)在进行网络IO时即进行Socket读写时遇到的, 所以本文讨论的上下文背景是基于Linux环境下的network IO.先介绍一下其中我们最常见的五种IO：

    1.  blocking IO
    2.  nonblocking IO
    3.  IO multiplexing
    4.  signal driven IO
    5.  asynchronous IO

由于signal driven IO在实际中并不常用, 所以我这只提及剩下的四种IO Model.

再说一下IO发生时涉及的对象和步骤.对于一个network IO (这里我们以read举例), 它会涉及到两个系统对象, 一个是调用这个IO的进程(或线程), 另一个就是系统内核(kernel).当一个read操作发生时, 它会经历两个阶段：

1. 等待数据准备`(Waiting for the data to be ready)`
2. 将数据从内核拷贝到进程中 `(Copying the data from the kernel to the process)`

记住这两点很重要, 因为这些IO Model的区别就是在两个阶段上各有不同的情况.

## 二、各种IO介绍

### 2.1 blocking IO

在linux中, 默认情况下所有的socket都是blocking, 也就是说我们的一个进程在进行IO操作时如果没有数据达到, 这个进程是被阻塞的.一个典型的读操作流程大概是这样：  
![IMG](https://img-blog.csdn.net/20161010174613299)

当用户进程调用了recvfrom这个系统调用, kernel就开始了IO的第一个阶段：准备数据.对于network io来说, 很多时候数据在一开始还没有到达(比如, 还没有收到一个完整的UDP包), 这个时候kernel就要等待足够的数据到来.而在用户进程这边, 整个进程会被阻塞.当kernel一直等到数据准备好了, 它就会将数据从kernel中拷贝到用户内存, 然后kernel返回结果, 用户进程才解除block的状态, 重新运行起来.**所以, blocking IO的特点就是在IO执行的wait和copy两个阶段都被block了**.

在这种block IO的情况下, 如果请求的连接比较多, 但其中大部分都是阻塞的.因为cpu的核数是有限的, 所以一般的解决方案就是每个cpu启用多个线程来处理多个连接.这种解决方案有很大的缺陷：

    1. 线程是有内存开销的, 1个线程可能需要512K(或2M)存放栈, 那么1000个线程就要512M(或2G)内存
    2. 线程的切换开销和很大, 因为线程切换时需要保持当前线程上下文信息, 当大量时间花在上下文切换的时候, 分配给真正的操作的CPU就要少很多
    3. 一个cpu所支持的线程数量时有限的(因为上面两个原因), 一般来说线程的数量级在几百个左右就已经很大了

为了解决block IO存在的问题, 就引入了no-blocking IO概念.

### 2.2 non-blocking IO

no-blocking IO很简单, 通过将socket设为非阻塞模式, 这时, 当你调用read时, 如果有数据就绪, 就返回数据, 如果没有数据就绪, 就立刻返回一个错误, 如EWOULDBLOCK.这样是不会阻塞线程了, 但是你还是要不断的轮询来读取或写入.当对一个non-blocking socket执行读操作时, 流程是这个样子：  
![IMG](https://img-blog.csdn.net/20161010175002093)  
从图中可以看出, 当用户进程发出read操作时, 如果kernel中的数据还没有准备好, 那么它并不会block用户进程, 而是立刻返回一个error.从用户进程角度讲 , 它发起一个read操作后, 并不需要等待, 而是马上就得到了一个结果.用户进程判断结果是一个error时, 它就知道数据还没有准备好, 于是它可以再次发送read操作.一旦kernel中的数据准备好了, 并且又再次收到了用户进程的system call, 那么它马上就将数据拷贝到了用户内存, 然后返回.

从上面介绍可以看到, **blocking IO的特点就是在IO执行的wait阶段是非阻塞的, 但是copy阶段还是阻塞的**.

但是no-blocking IO也存在很大的缺陷, 就是IO线程还是要不断的轮询socket来读取或写入,  于是, 我们又引入了IO多路复用.

### 2.3 IO multiplexing(IO多路复用)

IO multiplexing即IO多路复用, 有些地方也称这种IO方式为event driven IO(事件驱动IO).**它的基本原理就是用通过操作系统提供的select/epoll等这些函数不断的轮询所负责的所有socket, 而不是让用户进程自己去轮询, 注意这个socket必须先设成异步的socket, 当某个socket有数据到达了, 就通知用户进程**.它的流程如图：  
![IMG](https://img-blog.csdn.net/20161010175052063)  
当用户进程调用了select, 那么整个进程会被block, 而同时, kernel会“监视”所有select负责的socket, 当任何一个socket中的数据准备好了, select就会返回.这个时候用户进程再调用read操作, 将数据从kernel拷贝到用户进程.这个图和blocking IO的图其实并没有太大的不同, 事实上, 还更差一些.因为这里需要使用两个system call (select 和 recvfrom), 而blocking IO只调用了一个system call (recvfrom).但是, 用select的优势在于它可以同时处理多个connection.

由上面的图示可知, 采用多路模型会多一次系统调用select, 如果处理的连接数不是很高的话, 使用select/epoll的web server不一定比使用multi-threading + blocking IO的web server性能更好, 可能延迟还更大.select/epoll的优势并不是对于单个连接能处理得更快, 而是在于能处理更多的连接.

那么IO多路复用的优势在哪里呢, 其实就是在”多路复用”这个词上.上面也讲到了多路复用是指使用一个线程来检查多个Socket(也成文件描述符 )的就绪状态, 比如调用select和epoll函数, 传入多个文件描述符, 如果有一个文件描述符就绪, 则返回, 否则阻塞直到超时.所以, 在高并发的场景中, 比如要处理10000个连接, 只需要1个线程监控就绪状态, 对就绪的每个连接开一个线程处理或者直接丢到线程池处理, 当然也可以用当前线程处理, 那么这个IO线程可以同时管理多个连接, 也就是多路复用了.

### 2.4 Asynchronous I/O

linux下的asynchronous IO其实用得很少.先看一下它的流程：  
![IMG](https://img-blog.csdn.net/20161010175157114)  
用户进程发起read操作之后, 立刻就可以开始去做其它的事.而另一方面, 从kernel的角度, 当它受到一个asynchronous read之后, 首先它会立刻返回, 所以不会对用户进程产生任何block.然后, kernel会等待数据准备完成, 然后将数据拷贝到用户内存, 当这一切都完成之后, kernel会给用户进程发送一个signal, 告诉它read操作完成了.

## 三、各种IO之间的区别

到目前为止, 已经将四个IO Model都介绍完了.现在回过头来回答最初的那几个问题：blocking和non-blocking的区别在哪, synchronous IO和asynchronous IO的区别在哪.

blocking vs non-blocking, 这个问题很简单, 前面的介绍中其实已经很明确的说明了这两者的区别：

    1. blocking IO 会在wait和copy阶段都会阻塞进程
    2. non-blocking IO 在wait阶段会立即返回不会阻塞进程, 而在copy阶段仍会阻塞进程copy数据

在说明synchronous IO和asynchronous IO的区别之前, 需要先给出两者的定义.Stevens给出的定义(其实是POSIX的定义)是这样子的：

    1. synchronous I/O：IO操作过程中进程会被阻塞, 直到IO操作完成
    2. asynchronous I/O：IO操作过程中进程不会被阻塞, 操作系统帮你完成IO操作之后直接返回给你

按照这个定义, 在网络IO层面, 同步异步相对于阻塞非阻塞是一个更加宏观的概念, 之前所述的阻塞IO, 非阻塞IO, IO多路复用都属于同步IO, 因为它们在内核copy数据阶段都会阻塞进程.而异步IO则不一样, 当进程发起IO 操作之后, 就直接返回再也不理睬了, 直到操作系统内核发送一个信号, 告诉进程说操作系统IO已经完成, 在这整个过程中, 进程完全没有被阻塞.

各个IO Model的比较如图所示：  
![IMG](https://img-blog.csdn.net/20161010175306457)

经过上面的介绍, 会发现非阻塞IO和异步IO的区别还是很明显的.在非阻塞 IO中, 虽然进程大部分时间都不会被block, 但是它仍然要求进程去主动的check, 并且当数据准备完成以后, 也需要进程主动的再次调用recvfrom来将数据拷贝到用户内存.而异步 IO则完全不同, 它就像是用户进程将整个IO操作交给了操作系统(内核)完成, 然后操作系统做完后发信号通知.在此期间, 用户进程不需要去检查IO操作的状态, 也不需要主动的去拷贝数据.

参考文档：[http://blog.csdn.net/historyasamirror/article/details/5778378](http://blog.csdn.net/historyasamirror/article/details/5778378)
