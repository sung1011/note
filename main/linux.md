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
阻塞
- 调用后等待结果返回

非阻塞
- 调用后立刻返回, 之后各自轮询获取结果

select 多路复用
- 集中轮询文件描述符上的事件状态
- 采用一个1024数组存储状态，最多同时检查1024。

poll
- 采用链表避免数组长度限制，其他同select

epoll  
- 事件通知，执行回调
 
kqueue  
- FreeBSD系统，其他同kqueue

异步

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