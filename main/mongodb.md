# mongodb  

## [概述](mongodb-overview.md)
  
## [命令 cmd](mongodb-cmd.md)

## [索引index](mongodb-index.md)

## sync 同步

initial sync 全量同步  
oplog 增量同步  

- Primary上的写操作完成后，会向特殊的local.oplog.rs特殊集合写入一条oplog，Secondary不断的从Primary取新的oplog并应用  
  
## [replica set 复制集](mongodb-relicset.md)

## [mongos shard 分片集群 (分片的复制集)](mongodb-mongos.md)

## [transaction 事务](mongodb-transaction.md)

## capped collection

Capped Collection是性能出色的有着固定大小的集合，以LRU（Least Recently Used，最近最少使用）规则和插入顺序执行age-out（老化移出）处理，自动维护集合中对象的插入顺序。多用以日志归档。  

### 注意
  
- 如果写比读多，最好不要在上面创建索引
- 使用natual ordering可以有效地检索最近插入的元素，因为capped collection能够保证自然排序就是插入的顺序。  
- capped collection不能被shard.  
- 可以在创建capped collection时指定collection中能够存放的最大文档数。  

### [change stream (变更追踪、触发器)](mongodb-changestream.md)

### [设计模式](mongodb-design.md)
