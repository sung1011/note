# redis - cluster

![img](res/redis-cluster.png)

## 优点

- 自动故障转移、Slot迁移中数据可用  
- 节点增减，每个节点匀一小部分数据
- 组件all-in-box，部署简单，节约机器资源  
- 性能比proxy模式好  
- 官方方案，更新与支持有保障  

## 缺点

- 一致性较弱, 脑裂会丢失数据
- 多键操作支持有限, 如: 不建议进行集合对比(不巧是两个key在不同node, 传输与计算是瓶颈)
- 为了性能提升，客户端需要缓存node:key路由表信息  
- 节点发现、resharding操作不够自动化  

## 分片

    HASH_SLOT = CRC16(key) mod 16384

    通过哈希槽(hash-slot), 而非一致性哈希
    每个node负责一部分哈希槽 (nodeA: 0~5500, nodeB: 5501~11000, nodeC: 11001~16384)
    每个key通过CRC16校验后对16384个哈希槽取模, 决定放在哪个槽位.

## 实现

- `通信` 所有node之间通过 TCP连接 和 二进制集群连接(cluster bus) 建立通信.
- `协议` gossip.
  - 发现新node
  - ping健康检查心跳
  - 发数据消息
- `重定向` node不会转发请求, 只会将正确的路由返回给客户端, 由客户端重新发起

## 主从

    master挂调后, 通过选举, 将slave节点升级为master节点.

## 弱一致性

### 异步复制

    master处理并返回后, 异步的复制到slave.
    当尚未复制时master挂了, 数据就丢失.

### 脑裂

    master: A, B, C 
    slave: A1, B1, C1
    客户端: Z1
    
    发生网络分区, 集群被分为两方
    B, Z1
    A, B, C, A1, C1

    Z1仍然可以向B写入请求, 但达到节点超时时间(node timeout)后, B1会被选举成master节点, 这期间B新增的数据会丢失.

## 选举

    TODO

## resharding

    增删node后, 数据会自动迁移

## ref

- redis.io-tutorial <https://redis.io/topics/cluster-tutorial>

- redis.cn-tutorial <http://redis.cn/topics/cluster-tutorial>

- redis.io-cluster-spec <https://redis.io/topics/cluster-spec>

- redis.cn-cluster-spec <http://redis.cn/topics/cluster-spec>