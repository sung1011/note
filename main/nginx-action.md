# nginx 实战

## 为什么不多线程

    每个worker采用单线程来异步非阻塞处理请求(epoll), 不会为每个请求分配cpu和内存资源, 节省资源, 同时也减少CPU的上下文切换.

## [LVS, nginx, haproxy 负载均衡](load-balance.md)

## C10k问题

    (apache)当创建的进程或线程多了, 会数据拷贝频繁(缓存I/O、内核将数据拷贝到用户进程空间、阻塞, 进程/线程上下文切换消耗大).

> 每个进程/线程处理1个连接

> 每个进程/线程处理多个连接 即IO多路复用 select, poll, epoll, 异步IO
