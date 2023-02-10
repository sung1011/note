# redis 实战

## 实践

- key名尽量短, 但也要保障可读性  
- 尽可能使用hashe节约内存 (ziplist)
- Master最好不要做任何持久化工作, 如RDB和AOF
- 如果数据比较重要, 某个Slave开启AOF备份数据, 策略设置为每秒同步一次  
- 为了主从复制的速度和连接的稳定性, Master和Slave最好在同一个局域网内  
- 主从复制不要用图状结构, 用单向链表结构更为稳定, 即: Master <- Slave1 <- Slave2 <- Slave3... 这样的结构方便解决单点故障问题, 实现Slave对Master的替换.如果Master挂了, 可以立刻启用(目前只支持手动切换)Slave1做Master, 其他不变.  

## 值拷贝

`restore srckey 0 "{dump dstkey}"` // 伪代码

## 大key

1. redis-cli --bigkeys  
2. RDB

## 删除匹配到的key  

```js
   redis-cli keys *something* | xargs redis-cli del  
```

## 从文件中执行命令(大量) --pipe

```js
   # 1. 创建文件  
   set myk12 v1  
   zadd zset12 0 a 1 b 3 c  
   sadd sset12 e f g hh  
   # 2. 转码
   unix2dos < file >
   # 3. 导入
   cat < file > | redis-cli --pipe
```

## 高可用

```js
   - 主从复制
   - 哨兵模式
   - 集群
```

## keys *

```js
    不要使用keys *, 会阻塞主进程
    可以使用scan进行迭代
```

## 异步队列

todo

## 延迟队列

```js
    # zset方案
    生产者用zadd, score存时间戳
    消费者用zrangebyscore获取 并 zremrangebyscore消费.
      原子性考虑: 1. redis事务 2. 消息队列
```