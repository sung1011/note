# redis replication sync

    将RDB文件复制到slave并载入内存, 以实现数据同步

## Sync

      全量同步快照文件(RDB)

### 流程

1. `生成` 消耗资源; 主服务器执行BGSAVE命令来生成RDB文件, 这个生成操作会耗费主服务器大量的CPU、内存和磁盘I/O资源
2. `同步` 占用带宽; 主服务器需要将自己生成的RDB文件发送给从服务器,这个发送操作会耗费主从服务器大量的网络资源(带宽和流量),并对主服务器响应命令请求的时间产生影响
3. `加载` 阻塞载入RDB; 载入RDB文件, 因为阻塞而没办法处理命令请求

> `阻塞` save阻塞主线程, bgsave不阻塞主线程

## psync1 (>=2.8)

    尝试进行部分重同步.

### 流程

1. 每个实例都有唯一的repl-id
2. 当redis复制中断, slave会psync上报 repl-id 和 当前已同步的repl-offset
3. redis重新同步时, 若repl-id与master的一致 且 复制偏移量小于master的复制积压缓冲区, 则进行部分重同步, 否则进行全量同步 (同sync).  

> 复制积压缓冲区 replication backlog: master保留一个1M大小的偏移量缓冲区(所有slave共享)  

> 复制偏移量 replication offset  

### issues

- psync1需要满足 repl-id && offset 双重条件
- slave因故重启, repl-id 和 offset 都丢失时, psync1失效 只能全量同步
- 故障切换后, 新的slave需进行全量重同步

## psync2 (>=4.0)

    在psync基础上, 将复制信息(repl-id, repl-offset)记录在了RDB中以持久化.

### 流程

1. redis关闭时, 把复制信息作为辅助字段存储在RDB文件中; 重启时重新赋值给相关字段
2. 其他流程同psync1

> `查看相关参数` info命令的replication

## 实战

### 处理key的过期

    slave不处理过期的key, master发现key过期后, 会发送DEL到slave
    若master无法及时提供DEL, slave会通过自己的逻辑时钟判定是否key已经过期, 避免读取时返回已过期的key
    slave当master时自然处理过期

### master压力

    master进行复制是很消耗资源的, >2.8.18支持无磁盘复制
    正常流程: 
        master磁盘创建RDB
        RDB加载到内存
    无磁盘复制:
        master直接发送RDB文件给slave, 不需要磁盘作为中间介质

> 也可以不在master进行RDB, 只在slave做

## ref

- redis复制官方文档<http://www.redis.cn/topics/replication.html>