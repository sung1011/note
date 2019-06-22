# redis持久化

## RBD  

### sync (主从)  
意义  
- 全量同步快照  

流程  
- 主服务器需要执行BGSAVE命令来生成RDB文件，这个生成操作会耗费主服务器大量的CPU、内存和磁盘I/O资源；   
- 主服务器需要将自己生成的RDB文件发送给从服务器，这个发送操作会耗费主从服务器大量的网络资源（带宽和流量），并对主服务器响应命令请求的时间产生影响；   
- 接收到RDB文件的从服务器需要载入主服务器发来的RDB文件，并且在载入期间，从服务器会因为阻塞而没办法处理命令请求。  

### psync1 (>=2.8) (主从)  
意义  
- 尝试进行部分重同步。   

流程:   
1. 当redis复制中断后，slave会尝试采用psync, slave上报原master runid + 当前已同步master的offset,   
2. 若runid与master的一致 && 复制偏移量小于master的复制积压缓冲区, 则进行部分重同步, 否则进行全量同步（同sync）。  

名词:  
- 复制积压缓冲区 replication backlog: master保留一个1M大小的偏移量缓冲区(所有slave共享)  
- 服务器的随机标识符 runid  
- 复制偏移量replication offset  

### psync2 (>=4.0) (主从)  
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

### 快照 save SNAPSHOTTING (被动)  
配置  
- save 900 1 #在900秒(15分钟)之后，如果至少有1个key发生变化，则dump内存快照。  
- save 300 10 #在300秒(5分钟)之后，如果至少有10个key发生变化，则dump内存快照。  
- save 60 10000 #在60秒(1分钟)之后，如果至少有10000个key发生变化，则dump内存快照。  

## AOF  

### append of file (被动)  
#### 意义  
追加写命令  

#### 流程  
redis会将每一个收到的写命令都通过write函数追加到文件中(默认是 appendonly.aof)。  

#### 配置  
appendonly yes              //启用aof持久化方式  
#appendfsync always      //每次收到写命令就立即强制写入磁盘，最慢的，但是保证完全的持久化，不推荐使用  
appendfsync everysec     //每秒钟强制写入磁盘一次，在性能和持久化方面做了很好的折中，推荐  
#appendfsync no    //完全依赖os，性能最好,持久化没保证  

#### AOF rewrite  
bgrewriteaof  
- 意义  
    - 为了压缩aof的持久化文件(aof文件是可读的 + 保存了全部写操作 所以体积会很大)。  
- 流程  
    1. redis调用fork ，现在有父子两个进程  
    2. 子进程根据内存中的数据库快照，往临时文件中写入重建数据库状态的命令.`注意：这里是重写了aof文件， 并没有读取旧aof`  
    3. 父进程继续处理client请求，除了把写命令写入到原来的aof文件中。同时把收到的写命令缓存起来。这样就能保证如果子进程重写失败的话并不会出问题。  
    4. 当子进程把快照内容写入已命令方式写到临时文件中后，子进程发信号通知父进程。然后父进程把缓存的写命令也写入到临时文件。  
    5. 现在父进程可以使用临时文件替换老的aof文件，并重命名，后面收到的写命令也开始往新的aof文件中追加。  
  
