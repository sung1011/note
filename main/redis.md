# redis Remote Dictionary Server

## [概述](redis-overview.md)

    各个版本发布特性

## [对象](redis-obj.md)

    string, list, hash, set, zset

## [编码](redis-encoding.md)

    int, sds, linkedlist, quicklist, hashtable, skiplist, intset, ziplist, stream

## [命令](redis-cmd.md)

    commands

## 管道 pipeline

## 发布订阅 pub sub

## [stream](redis-stream.md)

    可持久化的消息队列

## lua脚本

## [内存优化](redis-mem-optimization.md)

    内存调优: 压缩, 分配, 数据结构

## [过期](redis-expire.md)

    数据过期删除策略

## [缓存淘汰策略](redis-cache-eliminate.md)

    LRU Least-Recently-Used 按时间过期删除
    LFU Least-Frequently-Used 按时间过期 + 使用频率删除

## [事务](redis-transaction.md)

        隔离：事务中的所有命令都会序列化、按顺序地执行。事务在执行的过程中，不会被其他客户端发送来的命令请求所打断。
        原子：事务中的命令要么全部被执行，要么全部都不执行。
        但不会报错回滚

## [分区/集群](redis-partitioning.md)

    将数据分发到不同实例, 每个实例是所有key的一个子集.

## [分布式锁 distribution lock](redis-distlock.md)

    分布式服务器获取锁

## key事件通知 Redis keyspace notifications

## 创建二级索引 Creating secondary indexes with Redis

## 配置

## [复制](redis-replication.md) todo

## [持久化RDB](redis-rdb.md)

    生成紧凑的快照文件来持久化当前的数据

## [持久化AOF](redis-aof.md)

    追加log文件持久化数据

## 哨兵

    当master宕机, 不需要手动, 而是自动的恢复. 即 监控和自动容灾

## [线程](redis-thread.md)

    单线程

## 管理

## 安全

## 加密

## 信号处理

## 连接处理

## [高可用](redis-sentinel.md)

    Redis-Sentinel是Redis官方的高可用性解决方案

## 延迟监控

## 基准
  
## [实战](redis-practical.md)

## ref

- [doc](http://www.redis.cn/documentation.html)  
- [redisbook](http://redisbook.com)
- [Redis进阶不得不了解的内存优化细节](https://blog.csdn.net/belalds/article/details/81106853)
- [redis主体流程](https://www.jianshu.com/p/427cf97d7951)