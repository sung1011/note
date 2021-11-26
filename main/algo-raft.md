# raft

## 选举 leader-election

### 状态

- `Leader领导人` 处理所有客户端请求, 通常系统中Leader是唯一的.

- `Follower跟随者` 不会主动发送请求, 只简单响应`Leader`或`Candidate`的请求; 客户端对`Follower`请求, 会被`Follower`重定向到Leader.

- `Candidate候选人` 若`Follower`接收`Leader`的心跳超时, 它就会变成`Candidate`并发起一次选举(新任期Term), 获得半数以上选票的`Candidate`会成为新`Leader`; 选票平分或都不超半数则重选.
  - `PreCandidate` 假设网络原因当某个`Follower`无法收到心跳, 它将不断自增term并发起选举. 为避免此类无效选举, etcd3.4引入`PreVote参数`(默认false), 令`Follower`转`Candidate`前先进入`PreCandidate`状态, 不自增term发起预投票, 大多数节点认可才真正开始选举流程

### 流程

1. `Follower`将自己Term加1, 状态转换为`Candidate`, 发送`RequestVote RPC`给其他节点
2. 当选/落选/重选

   - `当选` 当获得多数选票, 状态转换为`Leader`, 向其他节点发送`AppendEntries RPC`以确立权威 和 阻止新选举

   - `落选` 收到其他`Candidate`的`AppendEntries RPC`并且其`Term`大于自己的, 则承认其`Leader`的合法性并自己状态转换回`Follower`; 若其`Term`小于自己的, 会拒绝这次RPC并且保持`Candidate`状态

   - `重选` 多个`Candidate`时, 可能选票无法超越半数, 此时Term加1并`RequestVote RPC`重新选举, 故不会出现多个`Leader`

> 当选条件/投票原则: `Follower`会拒绝日志没有自己新(先对比term, 后对比lastLogIndex)的`RequestVote RPC`; 即 `Candidate`至少要比大多数新才能当选

> 选举时间: 每个选举的时间都是随机的, 以减小出现多个`Candidate`同时出现的概率

> `PreVote`: 假设网络原因当某个`Follower`无法收到心跳, 它将不断自增term并发起选举. 为避免此类无效选举, etcd3.4引入`PreVote参数`(默认false), 令`Follower`转`Candidate`前先进入`PreCandidate`状态, 不自增term发起预投票, 大多数节点认可才真正开始选举流程

---

## 节点间通信 RPC

### `AppendEntries RPC` 日志的复制,覆盖和充当心跳 L->F/C

| 参数         | 含义                     | 意义                                |
| ------------ | ------------------------ | ----------------------------------- |
| term         | Leader的任期号           |
| leaderId     | LeaderId                 | 用于F重定向C的请求                  |
| prevLogIndex | 上一个log条目的索引      | 用来校验请求合法性                  |
| prevLogTerm  | 上一个log条目的任期号    | 用来校验请求合法性                  |
| entries[]    | 准备存储的log数据        | 心跳则为空;数组形式一次通讯可发多个 |
| leaderCommit | Leader已经提交的日志索引 | 用于覆盖Follower                    |

### `RequestVote RPC` Candidate发起选举请求 C->F

| 参数          | 含义                              | 意义                                                    |
| ------------- | --------------------------------- | ------------------------------------------------------- |
| term          | Candidate的任期号                 |
| candidateId   | 发起选举的ID                      |
| lastLogTerm   | Candidate最后一个日志条目的任期号 | `Follower`会拒绝比自己term小的`RequestVote`             |
| lastLongIndex | Candidate最后一个日志条目的索引值 | term相同时,`Follower`会拒绝比自己index小的`RequestVote` |


### `InstallSnapshot RPC` 分块日志快照给太落后的节点进行覆盖 L->F

分块的日志快照

---

## Log Replication

### 流程

1. WAL `Leader`将客户端发来的请求命令附加(append)到日志中
2. 并行的向其他节点广播`AppendEntries RPC`
3. 其他节点收到RPC后进行持久化
4. 若超过半数节点持久化成功, 则该日志标记为已提交(Committed)
5. 响应`Leader`返回值
6. etcdserver模块异步从Raft模块获取已提交的日志, 更新到状态机(boltdb)

> 此过程出现超时或报错时(崩溃,运行缓慢,网络丢包等), `Leader`会不断重试

> `返回时机` 日志条目已提交(commited)后, Leader响应C返回值

### 已提交 committed

- `AppendEntries RPC`被复制到过半数节点, 该日志就会被提交.
- 同时Leader日志中该日志前所有日志也会被提交, 包括由其他Leader创建的日志

### 一致性

`Leader` `Follower` 冲突的日志会被`Leader`的日志覆盖

1. `Leader`针对每个`Follower`都维护了一个`nextIndex`字段(下一个需要发送给该Follower的日志索引值).
2. 当`Leader`刚获得权力时,他初始化所有`nextIndex`作为自己最后一条日志的索引值+1
3. 若一个`Follower`日志和`Leader`不一致, 那么下一次`AppendEntries RPC`的一致性检查就会失败
4. 失败后`Leader`会减小`nextIndex`并重试, 直到`nextIndex`会在某个位置双方达成一致
5. 在该位置后执行日志覆盖, 使`Leader` `Follower`的日志保持一致

---

## 成员变更

TODO

---

## 网络分区

TODO

---

## 日志压缩

1. 确定日志中某一时间点
2. 创建快照,压缩,持久化
3. 删除时间点前的日志

---

## vs paxos

- 相对简洁易懂
- 简单易用的REST API
- 确保安全性(包括网络延迟,分区,数据包丢失,重复,乱序时不会返回不正确结果)
- 只要超过一半服务可运行, 一致性算法就可用
- 不依赖于时序来确保日志的一致性, 避免错误的时钟和极端消息延迟导致的问题

## ref

- [共识算法：Raft](https://www.jianshu.com/p/8e4bbe7e276c)
- [动画演示](http://thesecretlivesofdata.com/raft/)
- [分布式一致性协议Raft原理](https://wingsxdu.com/post/algorithms/raft/)
- [深度解析raft分布式一致性协议](https://blog.csdn.net/z69183787/article/details/112168120)
