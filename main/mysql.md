# mysql

## [cmd](mysql-sql.md)

## [架构](mysql-arch.md)

## [redo log](mysql-redolog.md)

## binlog

## 复制

### 流程

1. 主库写入binlog
2. 主库更新数据
3. 主库返回响应
4. 主库把binlog复制到从库
5. 从库回放binlog, 更新数据

> 也支持同步复制到从库后再返回响应

| HA方案                 | 高可用 | 可靠性 | 性能 | 说明                                       |
| ---------------------- | ------ | ------ | ---- |
| 1M1S 异步复制,手动切换 | x      | o      | o    |
| 1M1S 异步复制,自动切换 | o      | x      | o    |
| 1M1S 同步复制,自动切换 | o      | x      | x    |
| 1M2S 同步复制,自动切换 | o      | o      | x    | 同步到任一S, 相比1M1S预防了某一S挂掉的情况 |

> HA即high availability 高可用

## [index 索引](mysql-index.md)

## [dict](mysql-dict.md)

## ref

[mysql事务](https://www.cnblogs.com/xuwangqi/p/11389964.html)
