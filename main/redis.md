# redis   

## 版本
## [ 命令 ](redis-cmd.md)
## 管道 pipeline
## 发布订阅 pub sub
## lua脚本
## [ 内存优化 ](redis-mem-optimization.md)
## [ 过期 ](redis-expire.md)
## [ LRU缓存 ](redis-lru-cache.md)
## 事务
## [ 分区/集群 ](redis-partitioning.md)
## 分布式锁 distribution lock
## key事件通知 Redis keyspace notifications
## 创建二级索引 Creating secondary indexes with Redis
## 配置
## [ 复制 ](redis-replication.md)
## [ 持久化 ](redis-persistence.md)
## 管理
## 安全
## 加密
## 信号处理
## 连接处理
## 高可用
## 延迟监控
## 基准

  
## 应用场景  
会话缓存  
页面缓存  
消息队列  
计数器  
排行榜  
秒杀  
分布式锁  
  
## 为何单进程单线程,如何提高cpu利用率  
非计算密集型io， 单进程简单易用  
规避多线程或多进程的上下文切换消耗cpu  
规避锁与死锁导致的性能损耗  
多路复用io模型  
开多个实例  
  
## 定时器 `redis.c/serverCron`
更新服务器的各类统计信息，比如时间、内存占用、数据库占用情况等。  
清理数据库中的过期键值对。  
对不合理的数据库进行大小调整。  
关闭和清理连接失效的客户端。  
尝试进行 AOF 或 RDB 持久化操作。  
如果服务器是主节点的话，对附属节点进行定期同步。  
如果处于集群模式的话，对集群进行定期同步和连接测试。  

## redis vs memcache
redis:
- 数据结构丰富
- 主从备份
- 持久化

memcache:
- 100k以上数据，内存使用率更高。
- 适合做缓存
  
## 经验（待归类）  
key名尽量短，但也要保障可读性  
尽可能使用散列表（hashes）节约内存  
Master最好不要做任何持久化工作，如RDB内存快照和AOF日志文件  
如果数据比较重要，某个Slave开启AOF备份数据，策略设置为每秒同步一次  
为了主从复制的速度和连接的稳定性，Master和Slave最好在同一个局域网内  
主从复制不要用图状结构，用单向链表结构更为稳定，即：Master <- Slave1 <- Slave2 <- Slave3... 这样的结构方便解决单点故障问题，实现Slave对Master的替换。如果Master挂了，可以立刻启用(目前只支持手动切换)Slave1做Master，其他不变。  
  
## 实战
### scan所有  
### 删除匹配到的key  
`redis-cli keys *something* | xargs redis-cli del`  
### hotkey
### 从文件中执行命令(大量)
1. 创建文件
```
set myk12 v1
zadd zset12 0 a 1 b 3 c
sadd sset12 e f g hh
set myk22 v2
hset myset12 k1 v1
hmset myset22 k2 v2 k3 v3 k4 v4
set myk32 v3
```
2. 转码
`unix2dos < file >`
3. 导入
`cat < file > | redis-cli --pipe`

## ref
[ doc ](http://www.redis.cn/documentation.html)
[ redisbook ](http://redisbook.com)