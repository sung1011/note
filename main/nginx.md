# nginx

## 用途：
- 静态资源
- API服务
- 反向代理(用于HTTP、HTTPS、SMTP、POP3和IMAP协议)

## 优缺点:
优点：
- 跨平台、配置简单：在核心代码都使用了与操作系统无关的代码实现
- 高性能：非阻塞、高并发连接：处理2-3万并发连接数，官方监测能支持5万并发
- 高可用：复杂均衡, 内置健康检查
- 节约内存：3万并发连接下，开启10个nginx才占150M内存，Nginx采取了分阶段资源分配技术
- 节约宽带：支持GZIP压缩，可以添加浏览器本地缓存
- 稳定性高：用于反向代理，宕机的概率微乎其微
- nginx处理静态文件好,耗费内存少
- 扩展性高
- 热部署
- BSD许可证

## 版本
### version
mainline 当前开发。如：nginx-1.17.0 (奇数)
stable 稳定。如：nginx-1.16.0 (偶数)
legacy 遗产。如：nignx-1.14.0, nginx-1.12.0, nginx-1.10.0 (即，非最新的stable)

### changes
feature 特性
bugfix 修复
change 重构

## [ 编译安装 ](nginxCompile.md)

## [ 组成 ](nginxFile.md)

## [ 配置语法 ](nginxConfigGrammer.md)

## [ signal ](nginxSignal.md) 


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
- nginx在启动时，会解析配置文件，得到需要监听的端口与ip地址，然后在nginx的master进程里面
- 初始化好这个监控的socket，再进行listen
- fork出多个子进程(worker)出来,  子进程会竞争accept新的连接
- 客户端与nginx(worker)进行三次握手建立连接
- 当某一个worker accept成功，会创建nginx对连接的封装，即ngx_connection_t结构体
- 根据事件调用相应的事件处理模块，如http模块与客户端进行数据的交换
- nginx或客户端来主动关掉连接。(一般客户端主动， 超时的话nginx主动)

## 惊群
- 同一个时刻只能有唯一一个worker子进程监听web端口，此时新连接事件只能唤醒唯一正在监听端口的worker子进程。采用锁，互斥量实现。

## 实战
### [ modules ](nginxModules.md)