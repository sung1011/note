# MongoDB wiredtiger

## concurrency

    doc级别的并发

## snapshot & checkpoints

    MultiVersion并发控制, 操作开始时, WiredTiger为操作提供数据的时间点快照. 快照提供了内存数据的一致视图.
    写入磁盘时, 将修改后的快照覆盖写入磁盘. 最新写入的快照充当checkpoints, 可以此进行数据恢复.

## [journal](mongodb-journal.md)

## compression

    压缩所有集合和索引, 但会增加CPU开销

- zlib
- zstd | >= 4.2

> 配置项 storage.wiredTiger.collectionConfig.blockCompressor

## memory use

    max( 0.5 * ( RAM-1024mb ), 256mb )

> 配置 hostInfo.system.memLimitMB

> 配置 storage.wiredTiger.engineConfig.cacheSizeGB

> 配置 wiredTigerCacheSizeGB