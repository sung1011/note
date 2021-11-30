# redis Remote Dictionary Server

## [概述](redis-overview.md)

## [对象](redis-obj.md)

## [编码](redis-encoding.md)

## [命令](redis-cmd.md)

## 管道 pipeline

## 发布订阅 pub sub

## [stream](redis-stream.md)

## lua脚本

## [内存优化](redis-mem-optimization.md)

## [过期](redis-expire.md)

## [LRU 缓存淘汰策略](redis-cache-eliminate.md)

## 事务

## [分区/集群](redis-partitioning.md)

## [分布式锁 distribution lock](redis-distlock.md)

## key事件通知 Redis keyspace notifications

## [创建二级索引 Creating secondary indexes with Redis](redis-secondary-index.md)

## 配置

## [RDB](redis-rdb.md)

## [AOF](redis-aof.md)

## [线程](redis-thread.md)

## 管理

## 安全

## 加密

## 信号处理

## 连接处理

## 高可用

## 延迟监控

## 基准
  
## 定时器 `redis.c/serverCron`

- 更新服务器的各类统计信息，比如时间、内存占用、数据库占用情况等。  
- 清理数据库中的过期键值对。  
- 对不合理的数据库进行大小调整。  
- 关闭和清理连接失效的客户端。  
- 尝试进行 AOF 或 RDB 持久化操作。  
- 主节点对附属节点进行定期同步。  
- 如果处于集群模式，对集群进行定期同步和连接测试。  

## [实战](redis-practical.md)

## ref

- [doc](http://www.redis.cn/documentation.html)  
- [redisbook](http://redisbook.com)
- [Redis进阶不得不了解的内存优化细节](https://blog.csdn.net/belalds/article/details/81106853)
- [redis主体流程](https://www.jianshu.com/p/427cf97d7951)
