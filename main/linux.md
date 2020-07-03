# linux  
  
## [cpu](linux-cpu.md)

## [memeroy](linux-mem.md)

## disk

## [network](linux-network.md)

## [架构](linux-arch.md)  

## [进程](linux-process.md)  

## [线程 thread](linux-thread.md)

## [协程](linux-coroutine.md)

## [io模型](linux-io.md)
  
## crontab  

## Makefile  

## [cmd](linux-cmd.md)

## [hugepage](ref/hugepage.md)  
  
## 实战  

### coredump
  
### 内存分配和多线程原理  

### 协程，线程，进程的区别

进程多与线程比较  

1) 地址空间:线程是进程内的一个执行单元，进程内至少有一个线程，它们共享进程的地址空间，而进程有自己独立的地址空间  
2) 资源拥有:进程是资源分配和拥有的单位,同一个进程内的线程共享进程的资源  
3) 线程是处理器调度的基本单位,但进程不是  
4) 二者均可并发执行  
5) 每个独立的线程有一个程序运行的入口、顺序执行序列和程序的出口，但是线程不能够独立执行，必须依存在应用程序中，由应用程序提供多个线程执行控制  
协程多与线程进行比较  
6) 一个线程可以多个协程，一个进程也可以单独拥有多个协程，这样能使用多核CPU。  
7) 线程进程都是同步机制，而协程则是异步  
8) 协程能保留上一次调用时的状态，每次过程重入时，就相当于进入上一次调用的状态  
  
### 互斥锁，读写锁，死锁问题是怎么解决  
  
## ref  

[awk](ref/awk.md)  
[文件描述符](ref/文件描述符.md)  
[关于网络IO中的同步、异步、阻塞、非阻塞](ref/关于网络IO中的同步、异步、阻塞、非阻塞.md)  
[select、poll、epoll、同步、异步之间的区别总结](https://blog.csdn.net/lsgqjh/article/details/65629609)  
