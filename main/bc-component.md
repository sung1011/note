# 区块链核心组成

## P2P网络协议

        网络传输和广播、节点发现和维护.

### 网络连接

        网络连接基于TCP

### 拓扑结构

```js
        全分布式拓扑结构 (交易从某个节点产生, 广播到邻近节点传播)

        P2P网络拓扑结构分类:
            中心化拓扑 (Centralized Topology); 
            全分布式非结构化拓扑 (Decentralized Unstructured Topology); 
            全分布式结构化拓扑 (Decentralized Structured Topology, 也称作DHT网络); 
            半分布式拓扑 (Partially Decentralized Topology).
```

### 节点发现

        初始节点发现

- DNS
- 硬编码种子节点

         启动节点发现

- 比特币: 拷贝邻近节点的对等节点列表(peer list)
- 以太坊: KAD(Kademlia)

### 局域网穿透

1. NAT
2. UPnP(Universal Plug and Play 通用即插即用)

### 节点交互协议

1. 获取可用节点列表 getaddr
2. 数据传输 inv
3. 同步
   1. HeaderFirst 先同步头 再同步体
   2. BlockFirst 同步整体

## 共识机制(分布式一致性算法)

[PoW 工作量证明算法](algo-pow.md)

[PoS 权益证明算法](algo-pos.md)

[DPos 代表权益证明算法](algo-dpos.md)

## 加密签名算法

## 账户与存储模型
