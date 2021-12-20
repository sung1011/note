# redis 线程

## 模型

Redis 基于 Reactor 模式开发了自己的网络事件处理器：文件事件处理器(file event handler)

1. 文件事件处理器使用 I/O 多路复用(multiplexing)程序来同时监听多个套接字,  并根据套接字目前执行的任务来为套接字关联不同的事件处理器.
2. 当被监听的套接字准备好执行连接应答(accept)、读取(read)、写入(write)、关闭(close)等操作时,  与操作相对应的文件事件就会产生,  这时文件事件处理器就会调用套接字之前关联好的事件处理器来处理这些事件.

## 文件事件处理器

![img](res/redis-file-event-handler.png)

- 套接字 socket
- I/O多路复用
- 文件事件分派器 dispatcher
- 事件处理器

![img](res/redis-dispatch-event-via-queue.png)

尽管多个文件事件可能会并发地出现,  但 I/O 多路复用程序总是会将所有产生事件的套接字都入队到一个队列里面,  通过这个队列,  以`有序(sequentially)`、`同步(synchronously)`、`每次一个套接字` 的方式向文件事件分派器传送套接字： 当上一个套接字产生的事件被处理完毕之后(该套接字为事件所关联的事件处理器执行完毕),  I/O 多路复用程序才会继续向文件事件分派器传送下一个套接字

## 事件的类型

### AE_READABLE

当套接字变得可读时(客户端对套接字执行 write 操作, 或者执行 close 操作)  
或有新的可应答(acceptable)套接字出现时(客户端对服务器的监听套接字执行 connect 操作)

### AE_WRITABLE

当套接字变得可写时(客户端对套接字执行 read 操作)

## 实战

### 为何单进程单线程, 如何提高cpu利用率

非计算密集型io,  单进程简单易用  
规避多线程或多进程的上下文切换消耗cpu  
规避锁与死锁导致的性能损耗  
多路复用io模型  
开多个实例  

## ref

- Redis线程模型 <https://www.cnblogs.com/barrywxx/p/8570821.html>
