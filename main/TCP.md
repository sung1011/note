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

流程

- C -->> S: `SYN` seq=j // 在吗, 我要联机
- C <<-- S: `ACK` ack=j+1, `SYN` seq=k // ok, 在
  - S: syn_recv
- C -->> S: `ACK` ack=k+1 // ok, 建立连接吧
  - C: syn_send
  - C: established
  - S: established

意义

- 当第一次握手出现网络延迟，第二次握手依然会正常响应，这时第三次握手就会被抛弃，服务端没收到第三次握手就不会建立连接。
- 反之，若无第三次握手，服务端在第二次握手就开始建立客户端已经认为过期抛弃的连接，浪费资源。

### 数据传输

- C -->> S: GET /foo HTTP/1.1
- C <<-- S: `ACK`
- C <<-- S: HTTP/1.1 200 OK
- C -->> S: `ACK`

### 四次挥手

流程:

- C -->> S: `FIN` seq=i // 我要关了
- C <<-- S: `ACK` ack=i+1 // ok收到, 稍等, 我还有返回没发完呢
  - S: close_wait
  - C: fin_wait1
- C <<-- S: `FIN` seq=j // (发完了)我关了, 你也关吧
  - S: last_wait
  - C: fin_wait2
- C -->> S: `ACK` ack=j+1 // ok, 我也关了, 并且轮询几下检测是否真断了。
  - C: time_wait
  - C: closed
  - S: closed
  - C: 轮询

意义

- 第二，三次挥手的意义在于，服务端可能不会立即关闭socket，故只能先回复一个ACK确认。  
- 第四次挥手会进入time_wait状态， 重发几次以确认服务端是否已经断开。

## 拥塞
