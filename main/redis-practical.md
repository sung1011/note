# redis 实战

## 经验

- key名尽量短, 但也要保障可读性  
- 尽可能使用散列表(hashes)节约内存  
- Master最好不要做任何持久化工作, 如RDB内存快照和AOF日志文件  
- 如果数据比较重要, 某个Slave开启AOF备份数据, 策略设置为每秒同步一次  
- 为了主从复制的速度和连接的稳定性, Master和Slave最好在同一个局域网内  
- 主从复制不要用图状结构, 用单向链表结构更为稳定, 即: Master <- Slave1 <- Slave2 <- Slave3... 这样的结构方便解决单点故障问题, 实现Slave对Master的替换.如果Master挂了, 可以立刻启用(目前只支持手动切换)Slave1做Master, 其他不变.  

## 值拷贝

`restore srckey 0 "{dump dstkey}"` // 伪代码

## 大key

1. redis-cli --bigkeys  
2. RDB

## 删除匹配到的key  

```bash
   redis-cli keys *something* | xargs redis-cli del  
```

## 从文件中执行命令(大量) --pipe

```bash
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

```bash
   - 主从复制
   - 哨兵模式
   - 集群
```
