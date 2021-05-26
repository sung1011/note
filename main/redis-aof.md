# redis AOF append-only file  

@@redis @@WAL

    追加log文件持久化数据

## 优点

- 间隔较短, 足够耐久耐用(durable)  
- 容灾性好(redis-check-aof修复)  
- 重写  
- 易读易用  

## 缺点

- 文件较大  
- 重写很消耗CPU和内存

## 流程

1. reids会将收到的每一个写命令以一定格式通过write()函数写入到内存缓存区aof_buf
2. 后台线程flushAppendOnlyFile()将aof_buf内容写入到文件

## 配置

- `appendonly yes`          //启用aof持久化方式  
- `appendfsync everysec`    //每秒钟强制写入磁盘一次, 在性能和持久化方面做了很好的折中, 推荐  
- `appendfsync always`      //每次收到写命令就立即强制写入磁盘, 最慢的, 但是保证完全的持久化, 不推荐使用  
- `appendfsync no`          //完全依赖os, 性能最好,持久化没保证  

## rewrite

      为了压缩aof的持久化文件(aof文件是可读的 + 保存了全部写操作 所以体积会很大)。  

1. redis调用fork(), 现在有父子两个进程。  
2. 子进程根据`内存`中的数据库快照, 往临时文件中写入重建数据库状态的命令.`注意：这里是重写了aof文件,  并没有读取旧aof`  
3. 父进程继续处理client请求, 除了把写命令写入到原来的aof文件中。同时把收到的写命令缓存起来。这样就能保证如果子进程重写失败的话并不会出问题。  
4. 当子进程把快照内容以命令方式写到临时文件中后, 子进程发信号通知父进程。父进程把缓存的写命令也写入到临时文件。  
5. 现在父进程可以使用临时文件替换老的aof文件, 并重命名, 后面收到的写命令也开始往新的aof文件中追加。  
  
> 手动 bgrewriteaof >2.2  
> 自动 bgrewriteaof >2.4

## 容灾修复

1. 备份AOF。
2. `redis-check-aof –fix`进行修复。
3. `diff -u`对比修复前后差异。
4. 重启redis-server, 等待其载入修复后的aof。

## RDB + AOF  

1. `BGREWRITEAOF`, `BGSAVE`无法同时使用。
2. RBD与AOF同时生效时, AOF优先级更高。

## 过期

slaves不会独立处理过期, 会等到master执行DEL命令。
