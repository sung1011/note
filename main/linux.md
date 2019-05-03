# linux

## 内存分配和多线程原理

## 协程，线程，进程的区别。

## 互斥锁，读写锁，死锁问题是怎么解决

## 进程间通信方式
管道

信号量

信号

消息队列

套接字

## io模型
### 概念
同步
- 请求发起方对消息结果的获取是主动发起

异步
- 请求发起方对消息结果的获取是被动通知

阻塞
- 调用(针对io)后线程挂起等待结果

非阻塞
- 调用(针对io)后线程立刻返回

### 组合
同步阻塞
- 请求方发起，线程一直等待结果

同步非阻塞
- 请求方发起，线程不等待结果直接返回，通过主动轮询不断查看结果

异步阻塞 (没意义)
- 请求方发起后，挂起等待服务方通知

异步非阻塞
- 请求方发起后，返回做其他任务，等待服务方通知

### 实例
select多路复用 [同步非阻塞]
- 主动轮询遍历fd集合的事件状态
- 采用一个1024数组存储状态，最多同时检查1024。

poll多路复用 [同步非阻塞]
- 主动轮询链表(采用链表避免数组长度限制),其他同select

epoll多路复用 [同步非阻塞]  
- 主动轮询链表，设备就绪后执行回调将就绪的fd事件放入链表。
- 只要判断就绪链表是否空，而不用遍历整个fb，提高cpu性能。
- kqueue: FreeBSD系统的epoll

信号驱动 (SIGIO) [同步非阻塞]
- 主动轮询链表，设备就绪后信号通知将就绪fd事件放入链表。

aio [异步非阻塞]
- io完成后，信号通知或者回调线程


## crontab

## Makefile

## cmd
- [awk](related/awk.md)
    - 某列求和 `awk -F "\t" '{sum+=$1}END{print sum}'`

- [grep](related/grep.md)

- du
    - 文件大小 `du -sh * | sort -n`

- lsof
    - 端口占用排序排重 `lsof -n | awk '{print $2}'|sort|uniq -c |sort -nr`

- ulimit
    - 每个进程可打开的文件数 `ulimit -n`

- fuser
    - 查看占用端口的进程 `fuser -n tcp 9000`

### HugePage 透明大页
- [hugepage](related/hugepage.md)

### 其他
- grep, awk, top, ps, tar, ln
- telnet, lsof, scp, ssh, vim

## related
[文件描述符](related/文件描述符.md)  
[关于网络IO中的同步、异步、阻塞、非阻塞](related/关于网络IO中的同步、异步、阻塞、非阻塞.md)
[select、poll、epoll、同步、异步之间的区别总结](https://blog.csdn.net/lsgqjh/article/details/65629609)