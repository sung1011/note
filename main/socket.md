# socket

    传输层实现端到端的通信,  其端点为socket

## 实质

![img](res/socket-layer0.jpeg)
![img](res/socket-layer.jpeg)

> Socket是应用层与TCP/IP协议族通信的中间软件抽象层, 它是一组接口.  

> 在设计模式中, Socket其实就是一个门面模式, 它把复杂的TCP/IP协议族隐藏在Socket接口后面, 对用户来说, 一组简单的接口就是全部.  

> socket = IP : prot  

## 分类

### 流式套接字

TCP

### 数据报套接字

UDP

### 原始套接字

对较低层协议(如IP或ICMP)进行直接访问, 常用于网络协议分析

## 流程

![img](res/socket-process.jpeg)

### 建立连接

- S 初始化socket(), 绑定端口bind(), 监听端口lisen(), 阻塞等待C请求accept().  
- C 初始化socket(), 连接S connect().  

### 处理请求

- C 写入数据write(), 发送给S.  
- S 读取请求的数据read(), 解析并处理, 之后写入数据write(), 返回给C  
- C 读取返回的数据read().  

### 关闭连接

- C 主动请求并关闭连接close().  
- S 读取到关闭连接的请求read(), 进行关闭连接close().  

## 异常处理

### 链路存活检测

- C定时发送心跳检测(一般是ping), S连续n次心跳没回应, 则链路失效, C重新与S建立连接connect().

### 断连重试

- S宕机, C断线重连时需要等待固定时间再发起重连, 避免S回收连接不及时, 而C瞬间大量重连把连接数占满.