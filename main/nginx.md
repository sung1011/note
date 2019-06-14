# nginx

## [ 概述 ](nginx-overview.md)

## [ 编译安装 ](nginx-compile.md)

## [ 目录结构 ](nginx-file.md)

## [ 配置语法 ](nginx-config-grammer.md)

## [ signal ](nginx-signal.md) 

## [ 架构 ](nginx-arch.md)

## [ modules ](nginx-modules.md)

## [ openresty ](nginx-openresty.md)

## [ 流程 ](nginx-process.md)

## [ 进程间通信 ](nginx-process-communicate.md)

## nginx负载均衡的算法
轮询（默认）  
weight  
ip_hash  
fair (第三方) 按响应速度  
url_hash (第三方)  

## 事件处理机制
多进程单线程异步非阻塞事件处理机制：运用了io多路复用epoll模型

## 为什么不多线程
每个worker采用单线程来异步非阻塞处理请求(epoll)，不会为每个请求分配cpu和内存资源，节省资源，同时也减少CPU的上下文切换。

## 实战
### https
`./configure ... --with-http_ssl_module` 必要SSL模块  
`certbot --nginx --nginx-server-root=< nginx.conf.path >` certbot协助配置证书  

## ref
[ nginx平台初探(100%) ](http://tengine.taobao.org/book/chapter_02.html#nginx)
[ 深入NGINX：我们如何设计它的性能和扩展性 ](https://www.cnblogs.com/chenjfblog/p/8715580.html)
[ 理解nginx工作原理 ](https://www.jianshu.com/p/6215e5d24553)