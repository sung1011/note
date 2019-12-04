# TCP

## 标志位  

- **SYN** synchronous(建立联机)
- **ACK** acknowledgement(确认) 向对方发送确认
- **PSH** push(传送) 接收端尽快传送数据到应用层，不必等缓冲区满再发送。
- **FIN** finish(结束) c告知s我要结束连接
- **RST** reset(重置) 告知对方可以重置状态（清空端口值数据）
- **URG** urgent(紧急)
- **seq** Sequence number(顺序号码)
- **ack** Acknowledge number(确认号码)

## 通讯流程

### 三次握手

- C -->> S: `SYN` seq=j // 在吗, 我要联机
  - C: SYN_SEND
  - S: **SYN_RCVD** --- [syn_ack flood](ddos.md)

- C <<-- S: `ACK` ack=j+1, `SYN` seq=k // ok, 在
  - C: ESTABLISHED

- C -->> S: `ACK` ack=k+1 // ok, 建立连接吧
  - S: **ESTABLISHED**

> 当第一次握手出现网络延迟，第二次握手依然会正常响应，这时第三次握手就会被抛弃，服务端没收到第三次握手就不会建立连接。  
> 反之，若无第三次握手，服务端在第二次握手就开始建立客户端已经认为过期抛弃的连接，浪费资源。

### 数据传输

- C -->> S: GET /foo HTTP/1.1

- C <<-- S: `ACK`

- C <<-- S: HTTP/1.1 200 OK

- C -->> S: `ACK`

### 四次挥手

- C -->> S: `FIN` seq=i // 我要关了
  - C: FIN_WAIT1 --- s端重启，即无法响应fin时出现FIN_WAIT1
  - S: **CLOSE_WAIT**

- C <<-- S: `ACK` ack=i+1 // ok收到, 稍等, 我还有返回没发完呢
  - C: FIN_WAIT2 --- 半关闭连接。c端进入只收不发状态

- C <<-- S: `FIN` seq=j // (发完了)我关了, 你也关吧
  - S: **LAST_ACK** --- 压测时c端突然断开，s端很多LAST_ACK
  - C: TIME_WAIT

- C -->> S: `ACK` ack=j+1 // ok, 我也关了, 并且轮询几下检测是否真断了。
  - C: CLOSED
  - S: **CLOSED**
  - C: 轮询

> 第二，三次挥手的意义在于，服务端可能不会立即关闭socket，故只能先回复一个ACK确认。  
> 第四次挥手会进入time_wait状态， 重发几次以确认服务端是否已经断开。

## 拥塞

## ref

[tcp状态](https://blog.csdn.net/wuji0447/article/details/78356875)
[理解TCP](https://www.jianshu.com/p/ca64764e4a26)
