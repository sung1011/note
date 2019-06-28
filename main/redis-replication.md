# redis复制

## Sync (主从)

意义  

- 全量同步快照

流程  

- 主服务器需要执行BGSAVE命令来生成RDB文件，这个生成操作会耗费主服务器大量的CPU、内存和磁盘I/O资源；
- 主服务器需要将自己生成的RDB文件发送给从服务器，这个发送操作会耗费主从服务器大量的网络资源（带宽和流量），并对主服务器响应命令请求的时间产生影响；
- 接收到RDB文件的从服务器需要载入主服务器发来的RDB文件，并且在载入期间，从服务器会因为阻塞而没办法处理命令请求。  

## psync1 (>=2.8) (主从)  

意义  

- 尝试进行部分重同步。

流程:

1. 当redis复制中断后，slave会尝试采用psync, slave上报原master runid + 当前已同步master的offset
2. 若runid与master的一致 && 复制偏移量小于master的复制积压缓冲区, 则进行部分重同步, 否则进行全量同步（同sync）。  

名词:  

- 复制积压缓冲区 replication backlog: master保留一个1M大小的偏移量缓冲区(所有slave共享)  
- 服务器的随机标识符 runid  
- 复制偏移量replication offset  

## psync2 (>=4.0) (主从)  

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

## ref

[redis系列--主从复制以及redis复制演进](https://www.cnblogs.com/wdliu/p/9407179.html)
