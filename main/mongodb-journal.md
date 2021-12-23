# MongoDB Journal

    MongoDB执行命令前, 预先将命令写入Journal日志到磁盘, 保障数据安全.

> [WiredTiger](mongodb-wiredtiger.md)的checkpoints提供磁盘数据的一致性视图, 并允许MongoDB从最后一个checkpoints恢复. checkpoints之间的数据(大约1min), 可用Journal恢复(大约100ms).

## 流程

1. 命令写入buffer, buffer最多128kb
2. buffer写入Journal磁盘
    - 条件
      - write的参数`j: true`
      - 配置中`writeConcernMajorityJournalDefault=true`
      - 每100ms (配置`storage.journal.commitIntervalMs`)
      - replset
        - (Secondary)每次批量处理oplog之后
        - 某些oplog写入后

> mongod异常挂掉, buffer中数据会消失

> severStatus的wiredTiger.log是Journal相关统计

## 存储

### path

    {dbPath}/journal/WiredTigerLog.<sequence>