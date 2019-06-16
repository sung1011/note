# nginx处理流程

## 处理请求的流程
1. nginx在启动时，会解析配置文件，得到需要监听的端口与ip地址，然后在nginx的master进程里面  
2. 初始化好这个监控的socket，再进行listen  
3. fork出多个子进程(worker)出来,  worker会竞争accept_mutex新的连接  
4. 客户端与nginx(worker)进行三次握手建立连接  
5. 当某一个worker accept成功，会创建nginx对连接的封装，即ngx_connection_t结构体  
6. 根据事件调用相应的事件处理模块，如http模块与客户端进行数据的交换  
7. nginx或客户端来主动关掉连接。(一般客户端主动， 超时的话nginx主动)  

## 惊群
同一个时刻只能有唯一一个worker子进程监听web端口，此时新连接事件只能唤醒唯一正在监听端口的worker子进程。采用锁，互斥量实现。

## 互斥锁
实现1: 原子操作 + 信号量  
实现2: 文件锁封装

## 锁占用时间过长

## 处理请求header的流程
TODO