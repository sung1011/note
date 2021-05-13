# mysql

## [cmd](mysql-sql.md)

## [架构](mysql-arch.md)

## [redo log](mysql-redolog.md)

## binlog

## 复制

### 流程.异步复制

1. 主库写入binlog
2. 主库更新数据,提交事务
3. 主库返回响应给客户端
4. 从库拉取主库的binlog的增量数据(与步骤2,3异步)
5. 从库回放binlog,更新数据(与步骤4异步;独立线程)
6. 从库返回响应给主库

> 也支持同步复制到从库后再返回响应

| HA方案                   | 高可用 | 可靠性 | 性能 | 说明                                                                                  |
| ------------------------ | ------ | ------ | ---- | ------------------------------------------------------------------------------------- |
| 1M1S 异步复制,手动切换   | x      | o      | o    | 手动主从切换,手动将M的binlog拿到S,进行恢复和切换S为M                                  |
| 1M1S 异步复制,自动切换   | o      | x      | o    | 自动主从切换,以S的数据切换S为新M,老M没来得及传过来的数据会丢失,重新切回老M会有问题      |
| 1M1S 同步复制,自动切换   | o      | x      | x    | 同步异步都是由S请求M响应.区别在于同步:提交事务->复制到所有S->返回给C端.同步到所有S    |
| 1M5S 半同步复制,自动切换 | o      | x      | x-o  | (>=5.7)同步到部分S(1/5)就返回C端,相比1M1S预防了某一S挂掉的情况,性能是同步和异步的折中 |

> 异步自动和版同步自动 比较常用

> HA即high availability 高可用

> 配置项`rpl_semi_sync_master_wait_no_slave`: master至少等待复制到几个slave再返回C端.一般配1

> 配置项`rpl_semi_sync_master_wait_point`: 主库执行事务 `AFTER_SYNC`(*)先复制后提交 / `AFTER_COMMIT`先提交后复制.前者可靠,后者高效.  
> AFTER_SYNC时,若复制过程超时,M依然会提交.

> 配置项`rpl_semi_sync_master_wait_no_slave`: 当`rpl_semi_sync_master_wait_point`为AFTER_SYNC,但复制过程超时,Mysql会自动降级为异步复制模式,知道足够多的S追上M.

## [索引 index](mysql-index.md)

## [分表](mysql-split.md)

## [删除](mysql-delete.md)

## ref

[mysql事务](https://www.cnblogs.com/xuwangqi/p/11389964.html)
