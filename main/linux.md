# linux  
  
## 架构  
### 用户模式:  
应用程序（sh、vi、OpenOffice.org等）  
复杂库（KDE、glib等）  
简单库（opendbm、sin等）  
C库（open、fopen、socket、exec、calloc等）  
### 内核模式:  
系统中断、调用、错误等软硬件消息  
内核（驱动程序、进程、网络、内存管理等）  
硬件（处理器、内存、各种设备）  
  
## 内存
### 存储节点 node
### 管理区 zone
### 页面 page

## 进程  
### 概念  
进程是具有一定独立功能的程序关于某个数据集合上的一次运行活动,进程是系统进行资源分配和调度的一个独立单位。每个进程都有自己的独立内存空间，不同进程通过进程间通信来通信。由于进程比较重量，占据独立的内存，所以上下文进程间的切换开销（栈、寄存器、虚拟内存、文件句柄等）比较大，但相对比较稳定安全。  
### 内容: 进程中拥有的数据  
那个程序的可执行机器代码的一个在存储器的映像。  
分配到的存储器（通常是虚拟的一个存储器区域）。存储器的内容包括可执行代码、特定于进程的数据（输入、输出）、调用堆栈、堆栈（用于保存运行时运输中途产生的数据）。  
分配给该进程的资源的操作系统描述符，诸如文件描述符（Unix术语）或文件句柄（Windows）、数据源和数据终端。  
安全特性，诸如进程拥有者和进程的权限集（可以容许的操作）。  
处理器状态（内文），诸如寄存器内容、物理存储器定址等。当进程正在运行时，状态通常存储在寄存器，其他情况在存储器。  
### 状态: 进程当前的动作  
新生（new）：进程新产生中。  
运行（running）：正在运行。  
等待（waiting）：等待某事发生，例如等待用户输入完成。亦称“阻塞”（blocked）  
就绪（ready）：排班中，等待CPU。  
结束（terminated）：完成运行。  
### 调度: 将任务分配至资源的过程  
依序循环调度（RR，Round-robin scheduling）  
最短先做排班（SJN，Shortest job next）  
Shortest remaining time  
Weighted round-robin scheduling  
速率单调 (RMS)  
Deadline-monotonic scheduling (DMS)  
Earliest deadline first scheduling (EDF)  
Two-level scheduling  
先进先出  
LIFO  
Fair-share scheduling  
完全公平调度 (CFS)  
Least slack time scheduling (LST)  
Multilevel Feedback Queue  
Take scheduling  
Gang scheduling  
Least-connection scheduling  
Weighted least-connection scheduling  
Shortest expected delay scheduling  
Never queue scheduling  
List scheduling  
Genetic Anticipatory  
Lottery Scheduling  
关键路径  
  
### 进程间通信 IPC Inter-Process Communication  
  
#### 目的  
数据传输: 一个进程需要将它的数据发送给另一个进程，发送的数据量在一个字节到几M字节之间  
共享数据: 多个进程想要操作共享数据，一个进程对共享数据  
通知事: 一个进程需要向另一个或一组进程发送消息，通知它（它们）发生了某种事件（如进程终止时要通知父进程）。  
资源共享: 多个进程之间共享同样的资源。为了作到这一点，需要内核提供锁和同步机制。  
进程控制: 有些进程希望完全控制另一个进程的执行（如Debug进程），此时控制进程希望能够拦截另一个进程的所有陷入和异常，并能够及时知道它的状态改变。  
  
#### 途径  
信号 singal: 不能传递复杂消息，只能用来同步  
消息队列 Message queue: 容量受到系统限制，且要注意第一次读的时候，要考虑上一次没有读完数据的问题  
管道 Pipeline / 流管道 s_pipe: 速度慢，容量有限，只有父子进程能通讯  
命名管道 FIFO: 任何进程间都能通讯，但速度慢  
共享内存 shared memory: 能够很容易控制容量，速度快，但要保持同步，比如一个进程在写的时候，另一个进程要注意读写的问题，相当于线程中的线程安全，当然，共享内存区同样可以用作线程间通讯，不过没这个必要，线程间本来就已经共享了同一进程内的一块内存  
信号量 Semaphore / 互斥锁 Mutex  
Message passing  
Memory-mapped file  
文件  
套接字 socket  
  
  
## 线程 thread  
### 概念  
线程是进程的一个实体,是CPU调度和分派的基本单位,它是比进程更小的能独立运行的基本单位.线程自己基本上不拥有系统资源,只拥有一点在运行中必不可少的资源(如程序计数器,一组寄存器和栈),但是它可与同属一个进程的其他的线程共享进程所拥有的全部资源。线程间通信主要通过共享内存，上下文切换很快，资源开销较少，但相比进程不够稳定容易丢失数据。   
  
### 内容  
同一进程中的多条线程将共享该进程中的全部系统资源，如虚拟地址空间，文件描述符和信号处理等等。    
但同一进程中的多个线程有各自的调用栈（call stack），自己的寄存器环境（register context），自己的线程本地存储（thread-local storage）。    
  
### 状态  
产生（spawn）  
中断（block）  
非中断（unblock）  
结束（finish  
  
## 协程  
### 概念  
协程是一种用户态的轻量级线程，协程的调度完全由用户控制。协程拥有自己的寄存器上下文和栈。协程调度切换时，将寄存器上下文和栈保存到其他地方，在切回来的时候，恢复先前保存的寄存器上下文和栈，直接操作栈则基本没有内核切换的开销，可以不加锁的访问全局变量，所以上下文的切换非常快。  
  
  
## 互斥锁，读写锁，死锁问题是怎么解决  
  
## 进程间通信方式  
管道  
  
信号量  
  
信号  
  
消息队列  
  
套接字  
  
## [io模型](linux-io.md)
  
## crontab  

## Makefile  

## cmd  
[awk](src/cmd/awk.md)  
- 某列求和 `awk -F "\t" '{sum+=$1}END{print sum}'`  
  
[grep](ref/grep.md)  
  
du  
- 文件大小 `du -sh * | sort -n`  
  
lsof  
- 端口占用排序排重 `lsof -n | awk '{print $2}'|sort|uniq -c |sort -nr`  
  
ulimit  
- 每个进程可打开的文件数 `ulimit -n`  
  
fuser  
- 查看占用端口的进程 `fuser -n tcp 9000`  
  
### HugePage 透明大页  
- [hugepage](ref/hugepage.md)  
### 其他  
- grep, awk, top, ps, tar, ln  
- telnet, lsof, scp, ssh, vim  
  
## 实战  
### [平衡负载](linux-loadAverage.md)  
### [ctx switch - CPU上下文切换](linux-ctxSwitch.md)
### [CPU使用率](linux-cpuUsage.md)
### [不可中断和僵尸进程](linux-uninterrupt.md)
### [中断](linux-interrupt.md)

### coredump
  
### 内存分配和多线程原理  
### 协程，线程，进程的区别?  
进程多与线程比较    
1) 地址空间:线程是进程内的一个执行单元，进程内至少有一个线程，它们共享进程的地址空间，而进程有自己独立的地址空间  
2) 资源拥有:进程是资源分配和拥有的单位,同一个进程内的线程共享进程的资源  
3) 线程是处理器调度的基本单位,但进程不是  
4) 二者均可并发执行  
5) 每个独立的线程有一个程序运行的入口、顺序执行序列和程序的出口，但是线程不能够独立执行，必须依存在应用程序中，由应用程序提供多个线程执行控制  
协程多与线程进行比较    
1) 一个线程可以多个协程，一个进程也可以单独拥有多个协程，这样能使用多核CPU。  
2) 线程进程都是同步机制，而协程则是异步  
3) 协程能保留上一次调用时的状态，每次过程重入时，就相当于进入上一次调用的状态  
  
  
## ref  
[awk](ref/awk.md)  
[文件描述符](ref/文件描述符.md)    
[关于网络IO中的同步、异步、阻塞、非阻塞](ref/关于网络IO中的同步、异步、阻塞、非阻塞.md)  
[select、poll、epoll、同步、异步之间的区别总结](https://blog.csdn.net/lsgqjh/article/details/65629609)  