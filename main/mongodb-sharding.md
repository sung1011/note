# sharding

## 场景

- 数据量巨大
  - 访问性能降低
  - 故障恢复极慢
- 并发量巨大
- 地理分布数据

## 分布策略

- 范围
- 哈希
- 自定义zone
  - 地域

## 组成

- `config`  

      用来存储sharded集群的元数据和配置信息

> 轻存储, 可配置小硬盘
  
- `router` (mongos)  

      路由就是mongos的实例, 客户端直接连接mongos, 由mongos把读写请求路由到指定的Shard上去; 可以配置多个独立的mongos减轻路由压力.

> 重CPU, 处理链接用

> 可以每个app-server配置一个mongos来减轻路由压力

> monogs不存储数据, 所以不要配置为ReplSet

> 不要在mongos上层部署负载均衡. 驱动会无法探测哪些是存活节点, 从而无法自动故障恢复; 驱动无法判定游标是哪个节点创建的, 从而遍历游标时出错.


- `shard`  

      用来保存分片数据, 保证数据的高可用性和一致性. 可以是一个单独的mongod实例, 也可以是一个副本集 | >= 3.2; 

> 最多个1024分片

> 重硬盘, 重内存

## 名词

- `片键 shard key` 文档中的一或多个字段
  - 取值基数尽量大 cardinality --- 尽量大(如 ID),避免很大的块出现.可考虑组合片键增加基数(如 uid+time)
  - 分散写集中读 --- 写分散分摊压力, 定向性读(一次查询到某一片键, 而非多个片键)
  - 取值分布 --- 尽量均匀
  - 避免单调递增减的片键
- `文档 doc` 包含shard key的一行数据
- `块 chunk` n个doc 一chunk约64M, 集群间以chunk为单位均衡
- `分片 shard` n个chunk, 主动增减分片, 自动迁移chunk
- `集合 cluster` n个shard
  
## 分片键 Shard keys  

### 优点  

    读写: sharding将读写负载均匀到各个shard, 且workload上限可以通过水平扩展来增加.  
    扩容: 每个shard保存一部分数据, 可以通过增加shards来扩容(动态扩容, 无需下线).  
    高可用: 即便某个shard不可用了, 整个集群也可以对外提供服务, 只不过访问down掉的shard会报"Connection refused"的错误.

### 缺点  

    额外消耗: router与shards节点间消耗; 读写多个分片数据
    管理复杂

### 额外  

    sharding集群不支持一些常规的单实例方法, 如group(), 可以使用mapReduce()或者aggregate()中的group来替代, 因此建议从一开始学习就直接使用aggregate(),这种写法较为简单明了, 且统一化易于识别.  
    对于没有用到shard key的查询, 路由进行全集群广播(broadcast operation), 对每个shard都查一遍进行scatter/gather, 此时效率会很低.  
    单片不要超过3TB, 保持在2TB较好

### 规划

#### 分片数量

    存储总量 / 单服容量 --- 8TB / 4TB = 4
    缓存总量 / 单服mongo内存用量 --- 400GB / (256G * 0.6) = 3 MongoDB使用物理机内存的60%
    并发总量 / 单服并发量 --- 30000 / ( 9000 * 0.7 ) = 6 额外开销系数0.7

> 分片数量 = max(a, b, c) = 6

#### 架构规划

    是否跨机房
    是否需要容灾
    对高可用的要求
