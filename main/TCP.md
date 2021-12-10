# TCP

## 标志位  

- **SYN** synchronous 建立联机
- **ACK** acknowledgement(确认) 向对方发送确认
- **PSH** push(传送) 接收端尽快传送数据到应用层, 不必等缓冲区满再发送。
- **FIN** finish(结束) 结束连接
- **RST** reset(重置) 告知对方可以重置状态（清空端口值数据）
- **URG** urgent(紧急)

- **seq** Sequence number(顺序号码)
- **ack** Acknowledge number(确认号码)

## 通讯流程

### 三次握手

```bash
    C -------------------------> S: `SYN` seq=j # 在吗, 我要联机
  SYN_SEND                    SYN_RCVD


    C <------------------------- S: `ACK` ack=j+1, `SYN` seq=k # ok, 在
  ESTABLISHED


    C -------------------------> S: `ACK` ack=k+1 # ok, 建立连接吧
                            ESTABLISHED
```

> 当第一次握手出现网络延迟, 第二次握手依然会正常响应(但C端已认为过期), 这时第三次握手就会被抛弃, S端没收到第三次握手就不会建立连接。  

> 反之, 若无第三次握手, S端会在第二次握手就开始建立C端已经认为过期抛弃的连接, 浪费资源。

> SYN_RCVD [syn_ack flood](ddos.md)

### 四次挥手

```bash
    C --------------------> S: `FIN` seq=i # 我要关了
 FIN_WAIT1              CLOSE_WAIT


    C <-------------------- S: `ACK` ack=i+1 # ok收到, 稍等, 我还有返回没发完呢
 FIN_WAIT2 (半关闭连接。C端进入只收不发状态)


    C <-------------------- S: `FIN` seq=j # (发完了)我关了, 你也关吧
 TIME_WAIT              LAST_ACK (压测时C端突然断开, S端很多LAST_ACK


    C --------------------> S: `ACK` ack=j+1 # ok, 我也关了, 并且轮询几下检测是否真断了。
  CLOSED
                        CLOSED
```

> 第二, 三次挥手的意义在于, 服务端可能不会立即关闭socket, 故只能先回复一个ACK确认。  

> 第四次挥手会进入time_wait状态, 重发几次以确认服务端是否已经断开。

## 拥塞

TODO

## ref

[tcp状态](https://blog.csdn.net/wuji0447/article/details/78356875)
[理解TCP](https://www.jianshu.com/p/ca64764e4a26)
