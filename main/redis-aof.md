# redis AOF append-only file  

@@redis @@WAL

    主进程写内存后追加log文件持久化数据

## 优点

    - 间隔较短, 足够耐久耐用(durable)  
    - 容灾性好(redis-check-aof修复)  
    - rewrite  
    - 易读易用, 可编辑  

## 缺点

    - 文件较大  
    - 重写很消耗CPU和内存

## 流程

1. 执行命令
2. 执行成功的命令写入内存 aof_buf
3. 由后台线程将aof_buf内存 fsync写入文件

> `先执行` redis不语法检查, 所以先执行确认返回值是正确, 才写文件; 有可能丢失和阻塞

## 配置

- `appendonly yes`          //启用aof持久化方式  
- `appendfsync everysec`    //每秒钟强制写入磁盘一次, 在性能和持久化方面做了很好的折中, 推荐  
- `appendfsync always`      //每次收到写命令就立即强制写入磁盘, 最慢的, 但是保证完全的持久化, 不推荐使用  
- `appendfsync no`          //完全依赖os, 性能最好, 持久化没保证  

## rewrite

      子进程rewrite 用内存中数据重写aof文件

1. redis调用fork(), 现在有父子两个进程.  
2. 子进程根据`内存`中的数据库快照, 往临时aof文件写入数据库状态. `注意: 这里是重写了aof文件, 并没有读取旧aof`  
3. 父进程继续处理client请求写旧aof, 这样保证如果rewrite失败的话并不会出问题
4. 新收到的写命令触及的内存页copy到 aof_rewrite_buf_blocks, 然后新老数据页都更新.
   - 默认在rewrite时也会fsync, 但消耗CPU, 推荐关闭
5. 临时aof文件生成完毕后, 缓存aof_rewrite_buf_blocks也append在后面
6. 父进程使用临时aof替换老的aof

> `no-appendfsync-on-rewrite no` 配置默认no, 推荐为yes 即rewrite时对新写的操作暂定fsync到旧aof, 仅存在buf

> 如果操作时bigKey, 那copy内存页到 aof_rewrite_buf_blocks 的CPU消耗会很大
  
> `bgrewriteaof` >2.2手动 >2.4自动

## 容灾修复

1. 备份AOF.
2. `redis-check-aof –fix`进行修复.
3. `diff -u`对比修复前后差异.
4. 重启redis-server, 等待其载入修复后的aof.

## RDB + AOF  

- AOF的`BGREWRITEAOF`, RBD的`BGSAVE`无法同时使用(防止IO过高).
- AOF优先级更高.

## 过期

    slaves不会独立处理过期, 会等到master执行DEL命令.

## ref

- <http://www.redis.cn/topics/persistence.html>
- 事故 <https://cloud.tencent.com/developer/article/1633077>