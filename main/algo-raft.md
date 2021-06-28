# raft

## 节点间通信 RPC

### `AppendEntries RPC` 日志复制和心跳 L->F/C

| 参数         | 含义                                     |
| ------------ | ---------------------------------------- |
| term         | Leader的任期号                           |
| leaderId     | LeaderId,用于F重定向C的请求              |
| prevLogIndex | 上一个log条目的索引,用来校验请求合法性   |
| prevLogTerm  | 上一个log条目的任期号,用来校验请求合法性 |
| entries[]    | 准备存储的log(心跳为空;一次通讯可发多个) |
| leaderCommit | Leader已经提交的日志索引                 |

### `RequestVote RPC` Candidate发起选举请求 C->F

| 参数          | 含义                              |
| ------------- | --------------------------------- |
| term          | Candidate的任期号                 |
| candidateId   | 发起选举的ID                      |
| lastLongIndex | Candidate最后一个日志条目的索引值 |
| lastLogTerm   | Candidate最后一个日志条目的任期号 |


### `InstallSnapshot RPC`

TODO

---

## 选举 leader-election

### 状态

- `Leader领导人` 处理所有客户端请求, 通常系统中只有一个Leader.

- `Follower跟随者` 不会主动发送请求, 只简单响应`Leader`或`Candidate`的请求; 客户端对`Follower`请求, 会被`Follower`重定向到Leader.

- `Candidate候选人` 若`Follower`接收不到`Leader`的消息, 它就会变成`Candidate`并发起一次选举(新任期Term), 获得半数以上选票的`Candidate`会成为新`Leader`; 选票平分或都不超半数则重选

### 流程

1. `Follower`将自己Term加1, 状态转换为`Candidate`, 发送`RequestVote RPC`给其他节点
2. 当选/落选/重选

   - `当选` 当获得多数选票, 状态转换为`Leader`, 向其他节点发送`AppendEntries RPC`以确立权威 和 阻止新选举

   - `落选` 收到其他`Candidate`的`AppendEntries RPC`并且其`Term`大于自己的, 则承认其`Leader`的合法性并自己状态转换回`Follower`; 若其`Term`小于自己的, 会拒绝这次RPC并且保持`Candidate`状态

   - `重选` 多个`Candidate`时, 可能选票无法超越半数, 此时Term加1并`RequestVote RPC`重新选举

> 投票原则 `先到先得`, 先收到谁的`RequestVote RPC`就把选票给谁

> 选举时间 每个选举的时间都是随机的, 以减小出现多个`Candidate`同时出现的概率

---

## Log Replication

TODO

## vs paxos

- 相对简洁易懂
- 确保安全性(包括网络延迟,分区,数据包丢失,重复,乱序时不会返回不正确结果)
- 只要超过一半服务可运行, 一致性算法就可用
- 不依赖于时序来确保日志的一致性, 避免错误的时钟和极端消息延迟导致的问题

## ref

- [共识算法：Raft](https://www.jianshu.com/p/8e4bbe7e276c)
- [动画演示](http://thesecretlivesofdata.com/raft/)
- [分布式一致性协议Raft原理](https://wingsxdu.com/post/algorithms/raft/)
