# nginx

## [概述](nginx-overview.md)

## [编译安装](nginx-compile.md)

## [目录结构](nginx-file.md)

## [配置语法](nginx-config-grammer.md)

## [信号 signal](nginx-signal.md)

## [架构](nginx-arch.md)

## [流程](nginx-process.md)

## [模块 modules](nginx-modules.md)

## [内存池](nginx-pool.md)

## [进程间通信](nginx-process-communicate.md)

## [容器](nginx-container.md)

## [openresty](nginx-openresty.md)

## 实战

### 为什么不多线程

每个worker采用单线程来异步非阻塞处理请求(epoll)，不会为每个请求分配cpu和内存资源，节省资源，同时也减少CPU的上下文切换。

### [负载均衡](load-balance.md)

### C10k问题

(webserver apache)当创建的进程或线程多了，数据拷贝频繁（缓存I/O、内核将数据拷贝到用户进程空间、阻塞，进程/线程上下文切换消耗大）。

## ref

[nginx平台初探(100%)](http://tengine.taobao.org/book/chapter_02.html#nginx)
[深入NGINX：我们如何设计它的性能和扩展性](https://www.cnblogs.com/chenjfblog/p/8715580.html)
[理解nginx工作原理](https://www.jianshu.com/p/6215e5d24553)
