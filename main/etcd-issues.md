# etcd issues

## 原理

- 为何适合读多写少?
- 线性读和串行读各自适用于什么场景?
- 如何判定etcd是否适合你的业务场景?
- 为什么Follower日志会和Leader日志冲突?
- 日志冲突时Follower的WAL日志如何删除已持久化但冲突的日志?
- etcd watch机制能否保证事件不丢失?
- 误删除一个key能否找回?

## 稳定性和性能

- 哪些因素导致Leader切换?
- 为何官方读取QPS10+w但你的业务集群QPS几百就超时?
- etcd能跨地域部署吗?
- 如何优化提升性能和稳定性?

## 一致性

- 为什么etcd会出现数据不一致?
- 为什么节点磁盘IO很低, 写请求依然会超时?
- 为什么你只存储一个几百kb的KV, etcd进程却消耗数GB内存?
- 如何分析etcd内存和延时异常背后的原因?

## db大小

- 为什么删除了大量数据, db大小不减少?
- 为什么etcd社区建议db大小不要超过8G?
- 哪些因素会导致db大小增加?

## k8s

- k8s创建pod后etcd是如何工作的?
- etcd如何为k8s控制器编程模型提供支撑?
- APIServer的"too old resource version"错误跟etcd有什么关系?

## 最佳实践

- 但你在一个namespace下创建了数万个pod/CRD资源时,同时通过标签去潮汛指定pod/CRD资源, 为什么ApiServer和etcd会扛不住?
- 快速增长的业务如何避免etcd集群出现性能瓶颈?
- 如何创建安全,可靠的etcd集群运维体系?