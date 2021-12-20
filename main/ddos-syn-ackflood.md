# DDOS syn_ack flood

    TCP三次握手过程中, 客户端故意不发起第三次握手.  
    S端在发出SYN+ACK应答报文后无法收到C端的ACK报文
    S端一般会重试并等待一段时间后才丢弃这个连接.

```bash
    C -------------------------> S: `SYN` seq=j # 在吗, 我要联机
  SYN_SEND                    SYN_RCVD


    C <------------------------- S: `ACK` ack=j+1, `SYN` seq=k # ok, 在
  ESTABLISHED


    # S waiting:
    # C -------------------------> S: `ACK` ack=k+1 # ok, 建立连接吧
    #                        ESTABLISHED

```

## 影响

1. 未完成的连接占用TCP连接数.
2. s端维护一个非常大的无效连接列表.
3. s端进入分钟级(30s~2min)等待与syn+ack重试.

## 防御

- `源认证方式` 单独向源(c端)发送SYN, 能在限定时间返回说明源提供服务, 加入白名单, 多次失败加入黑名单
