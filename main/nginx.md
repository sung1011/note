# nginx

## [ 概述 ](nginx-overview.md)

## [ 编译安装 ](nginx-compile.md)

## [ 目录结构 ](nginx-file.md)

## [ 配置语法 ](nginx-configGrammer.md)

## [ signal ](nginx-signal.md) 

## [ 架构 ](nginx-arch.md)

## [ modules ](nginx-modules.md)

## [ openresty ](nginx-openresty.md)

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

## 处理请求的流程
nginx在启动时，会解析配置文件，得到需要监听的端口与ip地址，然后在nginx的master进程里面  
初始化好这个监控的socket，再进行listen  
fork出多个子进程(worker)出来,  worker会竞争accept_mutex新的连接  
客户端与nginx(worker)进行三次握手建立连接  
当某一个worker accept成功，会创建nginx对连接的封装，即ngx_connection_t结构体  
根据事件调用相应的事件处理模块，如http模块与客户端进行数据的交换  
nginx或客户端来主动关掉连接。(一般客户端主动， 超时的话nginx主动)  

## 惊群
- 同一个时刻只能有唯一一个worker子进程监听web端口，此时新连接事件只能唤醒唯一正在监听端口的worker子进程。采用锁，互斥量实现。

## 实战
### https
`./configure ... --with-http_ssl_module` 必要SSL模块  
`certbot --nginx --nginx-server-root=< nginx.conf.path >` certbot协助配置证书  

## ref
[ nginx平台初探(100%) ](http://tengine.taobao.org/book/chapter_02.html#nginx)
[ 深入NGINX：我们如何设计它的性能和扩展性 ](https://www.cnblogs.com/chenjfblog/p/8715580.html)
[ 理解nginx工作原理 ](https://www.jianshu.com/p/6215e5d24553)