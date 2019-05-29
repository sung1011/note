# redis   
  
## 数据结构  
kv, hash, set, zset, list, stream  
  
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
  
## pipeline  
  
## 事务相关操作命令  
multi  
exec  
discard  
watch  
  
## 持久化方案  

### RBD  

#### sync (主从)  
意义  
- 全量同步快照  

流程  
- 主服务器需要执行BGSAVE命令来生成RDB文件，这个生成操作会耗费主服务器大量的CPU、内存和磁盘I/O资源；   
- 主服务器需要将自己生成的RDB文件发送给从服务器，这个发送操作会耗费主从服务器大量的网络资源（带宽和流量），并对主服务器响应命令请求的时间产生影响；   
- 接收到RDB文件的从服务器需要载入主服务器发来的RDB文件，并且在载入期间，从服务器会因为阻塞而没办法处理命令请求。  

#### psync1 (>=2.8) (主从)  
意义  
- 尝试进行部分重同步。   

流程:   
1. 当redis复制中断后，slave会尝试采用psync, slave上报原master runid + 当前已同步master的offset,   
2. 若runid与master的一致 && 复制偏移量小于master的复制积压缓冲区, 则进行部分重同步, 否则进行全量同步（同sync）。  

名词:  
- 复制积压缓冲区 replication backlog: master保留一个1M大小的偏移量缓冲区(所有slave共享)  
- 服务器的随机标识符 runid  
- 复制偏移量replication offset  

#### psync2 (>=4.0) (主从)  
意义: psync1需要满足runid && offset双重条件， 因而在 1.slave因故重启，master runid和offset都丢失时， psync1失效。 2. 故障切换后，新的slave需进行全量重同步。psync2以上问题做了优化。  

流程:  
1. redis关闭时，把复制信息作为辅助字段(AUX Fields)存储在RDB文件中；以实现同步信息持久化；  
2. redis启动加载RDB文件时，会把复制信息赋给相关字段；  
3. redis重新同步时，会上报repl-id和repl-offset同步信息，如果和主实例匹配，且offset还在主实例的复制积压缓冲区内，则只进行部分重新同步。  

名词:  
- master_replid: 复制ID1(后文简称：replid1)，一个长度为41个字节(40个随机串+’0’)的字符串。redis实例都有，和runid没有直接关联，但和runid生成规则相同，都是由getRandomHexChars函数生成。当实例变为从实例后，自己的replid1会被主实例的replid1覆盖。  
- master_replid2：复制ID2(后文简称:replid2),默认初始化为全0，用于存储上次主实例的replid1  
- master_repl_offset: master偏移量  
- second_repl_offset: 上次主实例repid1和复制偏移量；用于兄弟实例或级联复制，主库故障切换psync.  

#### 快照 save SNAPSHOTTING (被动)  
配置  
- save 900 1 #在900秒(15分钟)之后，如果至少有1个key发生变化，则dump内存快照。  
- save 300 10 #在300秒(5分钟)之后，如果至少有10个key发生变化，则dump内存快照。  
- save 60 10000 #在60秒(1分钟)之后，如果至少有10000个key发生变化，则dump内存快照。  

### AOF  

#### append of file (被动)  
意义  
- 追加写命令  

流程  
1. redis会将每一个收到的写命令都通过write函数追加到文件中(默认是 appendonly.aof)。  

配置  
- appendonly yes              //启用aof持久化方式  
- #appendfsync always      //每次收到写命令就立即强制写入磁盘，最慢的，但是保证完全的持久化，不推荐使用  
- appendfsync everysec     //每秒钟强制写入磁盘一次，在性能和持久化方面做了很好的折中，推荐  
- #appendfsync no    //完全依赖os，性能最好,持久化没保证  

优化  
- bgrewriteaof  
    - 意义  
        - 为了压缩aof的持久化文件(aof文件是可读的 + 保存了全部写操作 所以体积会很大)。  
    - 流程  
        1. redis调用fork ，现在有父子两个进程  
        2. 子进程根据内存中的数据库快照，往临时文件中写入重建数据库状态的命令.`注意：这里是重写了aof文件， 并没有读取旧aof`  
        3. 父进程继续处理client请求，除了把写命令写入到原来的aof文件中。同时把收到的写命令缓存起来。这样就能保证如果子进程重写失败的话并不会出问题。  
        4. 当子进程把快照内容写入已命令方式写到临时文件中后，子进程发信号通知父进程。然后父进程把缓存的写命令也写入到临时文件。  
        5. 现在父进程可以使用临时文件替换老的aof文件，并重命名，后面收到的写命令也开始往新的aof文件中追加。  
  
## 哪几种数据淘汰策略  
noeviction:返回错误当内存限制达到并且客户端尝试执行会让更多内存被使用的命令（大部分的写入指令，但DEL和几个例外）  
allkeys-lru: 尝试回收最少使用的键（LRU），使得新添加的数据有空间存放。  
volatile-lru: 尝试回收最少使用的键（LRU），但仅限于在过期集合的键,使得新添加的数据有空间存放。  
allkeys-random: 回收随机的键使得新添加的数据有空间存放。  
volatile-random: 回收随机的键使得新添加的数据有空间存放，但仅限于在过期集合的键。  
volatile-ttl: 回收在过期集合的键，并且优先回收存活时间（TTL）较短的键,使得新添加的数据有空间存放。  
  
## 集群方案  
cluster  
twemproxy  
- 理论最大节点, (hash槽个数)  
  
## 集群优点  
轻量快速可报警  
支持redis,memcache;支持pipelinling;支持多个哈希算法  
实质上就是分片，单机内存变为多机内存。  
  
## 集群缺点  
涉及多个key的操作通常不会被支持。  
不支持事务  
数据处理变得困难，如备份时必须从不同redis实例和主机同时收集RDB/AOF文件  
自身单端口实例的压力  
节点变更后，数据无法平滑迁移到新节点（扩缩容）  
  
## Sentinel 哨兵机制  
  
## 调优  
  
## 经验  
key名尽量短，但也要保障可读性  
Master最好不要做任何持久化工作，如RDB内存快照和AOF日志文件  
如果数据比较重要，某个Slave开启AOF备份数据，策略设置为每秒同步一次  
为了主从复制的速度和连接的稳定性，Master和Slave最好在同一个局域网内  
主从复制不要用图状结构，用单向链表结构更为稳定，即：Master <- Slave1 <- Slave2 <- Slave3... 这样的结构方便解决单点故障问题，实现Slave对Master的替换。如果Master挂了，可以立刻启用(目前只支持手动切换)Slave1做Master，其他不变。  
  
## cmd 常用命令  
scan所有  
删除匹配到的key  