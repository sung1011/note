# etcd basic

## etcd

unix 的`/etc`文件夹 + d (distribute) = etcd 用于`存储分布式配置的信息存储服务`

## 架构

```go
--------------------------------------------------------
| `Client层`     etcdctl/clientv3                       |
--------------------------------------------------------
--------------------------------------------------------
| `API层`                                               |
| Raft-HTTP                                            |
| gRPC-API                                             |
| HTTP-API(v2/v3)                                      |
--------------------------------------------------------
--------------------   ----------------------------------
| `Raft层`          |  | `Logic层`                        |
| Leader-Election  |   |                                 |
|                  |   | KVServer                        |
| Log-Replication  |   |                                 |
|                  |   | Quota                           |
| Membership       |   |                                 |
|                  |   | Maintenance                     |
| Read-Index       |   |                                 |
|                  |   | Apply                           |
| Learner          |   |                                 |
|                  |   | Auth                            |
|                  |   |                                 |
|                  |   | Compactor                       |
|                  |   |                                 |
|                  |   | TreeIndex                       |
|                  |   |                                 |
|                  |   | Lease                           |
|                  |   |                                 |
|                  |   |                                 |
--------------------    ---------------------------------
----------------------------------------------------------
| `存储层`                                                 |
| WAL                                                     |
| Snapshot                                                |
| boltDB                                                  |
----------------------------------------------------------
```

- `Client层` v2/v3两大版本的API客户端库

- `API层` C访问S/S访问S的通讯协议

- `Raft层` raft算法实现层
  - Leader选举
  - 日志复制
  - ReadIndex
  - 成员切换

- `Logic层` 特性实现层
  - KVServer模块
  - MVCC模块
    - treeIndex模块
    - boltdb模块(状态机)
  - Auth鉴权模块
  - Lease租约模块
  - Compactor压缩模块

- `存储层`
  - 预写日志WAL模块 保障了数据持久化
  - 快照Snapshot模块 差异较大时Leader覆盖同步Follower数据
  - boltdb模块 集群元数据和用户写入数据

## 场景

- 存放非频繁更新的数据
- 配置管理
- 服务注册发现
- 选主
- 应用调度
- 分布式队列
- 分布式锁

## ref

- `https://github.com/mattn/goreman`
