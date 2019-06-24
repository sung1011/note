# redis分区

## 概念
将数据分发到不同实例，每个实例是所有key的一个子集。
  
## 实现 - 算法
- 范围 如：id范围
```
1~1000 -> R0， 1001~2000 -> R1， 2001~3000 -> R2
3001~4000 -> R0， 4001~5000 -> R1， 5001~6000 -> R2
...
```
- 散列 如：crc32
```
foobar -> 93024922 by crc32
93024922 -> 1 by 93024922%3
foobar -> R1
```
- [ 一致性哈希 DHT ](algo-DHT.md)

## 实现
### 1. 客户端分区  
客户端就已经决定数据会被存储到哪个redis节点或者从哪个redis节点读取。  
#### 客户端
https://redis.io/clients  

### 2. 代理分区  
意味着客户端将请求发送给代理，然后代理决定去哪个节点写数据或者读数据。(twemproxy)  
#### twemproxy
`https://github.com/twitter/twemproxy`
##### twemproxy优点  
轻量快速可报警  
支持redis,memcache;支持pipelinling;支持多个哈希算法  
支持自动分区，如果其代理的某节点不可用，自动排除该节点（这将改变keys-instance映射关系，所以只能用于充当缓存的场景）。  
非单点  
##### twemproxy缺点  
自身单端口实例的压力  
(同普遍缺点)

### 3. 查询路由(Query routing)  
客户端随机地请求任意一个redis实例，然后由Redis将请求转发给正确的Redis节点。  
#### redis cluster
http://www.redis.cn/topics/cluster-tutorial.html  
http://www.redis.cn/topics/cluster-spec.html  

## 分区的普遍缺点
涉及多个key的操作通常不会被支持。  
不支持事务  
粒度是key，故一个key很大时也无法被分片
数据处理变得困难，如备份时必须从不同redis实例和主机同时收集RDB/AOF文件  
节点变更后，数据无法平滑迁移到新节点（扩缩容）  

## 实战
### 预分片
为防止以后的扩容，最好的办法就是一开始就启动较多实例(即便只有一台机器)。

## ref
[ 分区：怎样将数据分布到多个redis实例 ](http://www.redis.cn/topics/partitioning.html)