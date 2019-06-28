# redis持久化

## RBD  

快照文件

### rdb优点

数据紧凑  
备份灵活  
恢复快捷  

### rdb缺点

间隔较长, 不够耐久耐用(durable)  
fork耗时  

### rdb流程

1. fork子进程，子进程将数据集快照写入临时rdb中。(copy-on-write 写时复制)
2. 完成后替换旧的rdb。

### rdb配置  

save 900 1 #在900秒(15分钟)之后，如果至少有1个key发生变化，则dump内存快照。  
save 300 10 #在300秒(5分钟)之后，如果至少有10个key发生变化，则dump内存快照。  
save 60 10000 #在60秒(1分钟)之后，如果至少有10000个key发生变化，则dump内存快照。  

## AOF append-only file  

追加写命令  

### aof优点

间隔较短, 足够耐久耐用(durable)
容灾性好(redis-check-aof修复)  
重写  
易读易用  

### aof缺点

文件较大  

### aof流程

redis会将每一个收到的写命令都通过write函数追加到文件中(默认是 appendonly.aof)。  

### aof配置

appendonly yes              //启用aof持久化方式  
appendfsync everysec     //每秒钟强制写入磁盘一次，在性能和持久化方面做了很好的折中，推荐  
appendfsync always      //每次收到写命令就立即强制写入磁盘，最慢的，但是保证完全的持久化，不推荐使用  
appendfsync no    //完全依赖os，性能最好,持久化没保证  

### aof rewrite

为了压缩aof的持久化文件(aof文件是可读的 + 保存了全部写操作 所以体积会很大)。  

#### aof rewrite 命令

手动 bgrewriteaof >2.2
自动 bgrewriteaof >2.4

#### aof rewrite 流程

1. redis调用fork ，现在有父子两个进程  
2. 子进程根据内存中的数据库快照，往临时文件中写入重建数据库状态的命令.`注意：这里是重写了aof文件， 并没有读取旧aof`  
3. 父进程继续处理client请求，除了把写命令写入到原来的aof文件中。同时把收到的写命令缓存起来。这样就能保证如果子进程重写失败的话并不会出问题。  
4. 当子进程把快照内容写入已命令方式写到临时文件中后，子进程发信号通知父进程。父进程把缓存的写命令也写入到临时文件。  
5. 现在父进程可以使用临时文件替换老的aof文件，并重命名，后面收到的写命令也开始往新的aof文件中追加。  
  
### aof 容灾修复

1. 备份AOF。
2. `redis-check-aof –fix`进行修复。
3. `diff -u`对比修复前后差异。
4. 重启redis-server，等待其载入修复后的aof。

## RDB + AOF  

1. `BGREWRITEAOF`, `BGSAVE`无法同时使用。
2. RBD与AOF同时生效时，AOF优先级更高。
