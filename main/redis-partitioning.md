# redis 分区

    将数据分发到不同实例, 每个实例是所有key的一个子集.

## 模式

### 1. 客户端分区  

    客户端就已经决定数据会被存储到哪个redis节点或者从哪个redis节点读取 (webserver-config) 

### 2. 代理分区  

    意味着客户端将请求发送给代理, 然后代理决定去哪个节点写数据或者读数据 (twemproxy, codis)  

### 3. 查询路由(Query routing)  

    客户端随机地请求任意一个redis实例, 然后由Redis将请求转发给key所在的Redis节点 (cluster)  

    - 涉及多个key的操作通常不会被支持 (如事务)  
    - 数据处理变得困难, 如备份时必须从不同redis实例和主机同时收集RDB/AOF文件  
    - 节点变更后, 数据无法平滑迁移到新节点 (扩缩容)  

## 逻辑实现

### range范围 如: id范围

```js
    1~1000 -> R0,  1001~2000 -> R1,  2001~3000 -> R2
    3001~4000 -> R0,  4001~5000 -> R1,  5001~6000 -> R2
    ...
```

### hash函数 如: crc32

```js
    slot_num = 1024                 # 槽
    hash = crc32(key)               # key = foobar, hash = 93024922
    slot_index = hash % slot_num    # slot_index = 666
    redis = slots[slot_index].redis # redis = 666槽位对应的redis实例
    redis.do(command)
```

> 如 redis-cluster

### [一致性哈希 DHT](algo-DHT.md)

```js
    NODE_HASH_SLOT = CRC16(node) mod 16384 # 实例所在的slot
    KEY_HASH_SLOT = CRC16(key) mod 16384 # key所在的slot
```

## 成熟方案

### [twemproxy](redis-twemproxy.md)

### [codis](redis-codis.md)

### [cluster](redis-cluster.md)

## 方案对比

| en                                    | cn                  | codis | twemproxy | cluster |
| ------------------------------------- | ------------------- | :---: | :-------: | :-----: |
| resharding without restarting cluster | 平滑分片            |   o   |     x     |    o    |
| pipeline                              | 管道                |   o   |     o     |    x    |
| hash tags for multi-key operations    | 以hash tag操作多key |   o   |     o     |    o    |
| multi-key operations while resharding | 分片时多key操作     |   o   |     x     |    x    |
| Redis clients supporting              | 客户端支持          |   o   |     o     |    o    |

## hash tag

    通过key的一部分进行hash, 该部分称为hash tag  
    识别`{}`作为reshard的key, 如: `a{bar}1, sun{bar}, {bar}xie`, 将以上key分片到同一个实例.

## 预分片

    为防止以后的扩容, 最好的办法就是一开始就启动较多实例(即便只有一台机器).

## ref

- [分区: 怎样将数据分布到多个redis实例](http://www.redis.cn/topics/partitioning.html)