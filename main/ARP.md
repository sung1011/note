# ARP 地址解析协议 Address Resolution Protocol

根据对方IP获取对方MAC地址.  

## ARP数据包

```ARP
//响应
{
    目标MAC (6 byte),
    源MAC (6 byte),
    类型 (2 byte),
    ARP报文: {
        硬件类型:
        协议类型: IP
        硬件地址长度: 6
        协议地址长度: 4
        操作代码: 1 request / 2 reply
        发送者MAC:
        发送者IP:
        目标MAC:
        目标IP:
    }
}
```

## 工作原理

![IMG](res/arp-process.jpg)

1. 查看本地ARP表
2. 本地ARP表无信息, 则广播ARP请求
3. 目标主机收到请求后, 进行ARP应答
4. 本地缓存该应答的IP-MAC映射(arp table cache)

## 缺点

1. ARP缓存表导致的ARP欺骗攻击 - 网络不可用或中间人攻击.  
   - 范围: 内网相同网段
   - 原理: 在内网中存在一个攻击者主机, 不断发出伪造的ARP响应包, 以更改目标主机(或网关)ARP缓存中的IP-MAC映射表, 造成网络不可用或中间人攻击.
   - 解决: 排查与排除攻击者主机.

2. 服务器切换时新主机无法访问.  
   - 范围: 服务器切换时新主机无法被访问到
   - 原理: 新主机MAC不在其他主机的ARP的缓存IP-MAC映射中
   - 解决: 清除缓存, 重新广播通知, 更新IP-MAC映射.
