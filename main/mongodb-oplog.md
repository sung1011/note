# MongoDB oplog

    replset主从同步操作, 来实现数据同步

## 流程

1. Primary上的写操作完成后, 会向local.oplog.rs特殊集合写入一条oplog
2. Secondary异步的不断的从Primary取新的oplog并幂等地写入

## 大小

```js
   - 默认
     - WiredTiger 空闲磁盘5%
     - In-Memory 物理内存5%
   - 配置 oplogSizeMB
   - 命令
     - 调整 replSetResizeOplog
     - 查看 rs.printReplicationInfo()
```

## 速度

- 配置延迟 flowControlTargetLagSeconds | >=4.2
- 命令
  - 查看 db.getReplicationInfo()

## slow-oplog

    Secondary会记录oplog中应用时间超过阈值的满操作. | >= 4.2

## ref

- docs <https://docs.mongodb.com/manual/core/replica-set-oplog/>