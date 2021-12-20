# mongodb 副本集 (复制+选举)

## 意义

1. 服务高可用
2. 数据分发 全球跨区域复制, 用户访问临近节点
3. 读写分离
4. 异地容灾

## 复制

- v 是否有选举权
- priority 优先级 主数据中心调高, 备份中心调低
- hidden 隐藏 备份时用
- slaveDelay 延迟 防止误操作删除数据

## 选举

1. 具有`投票权`的节点之间两两互发心跳
2. 当5次心跳未收到则判定为节点失联
   1. primary失联, 则secondary发起选举
   2. secondary失联, 无选举
3. 选举基于[raft](algo-raft.md)一致性算法, 选举成功必要条件是大多数投票节点存货.
4. 最多存在50个节点, 但`投票权`的节点最多7个

## 节点  

Primary  
Secondary  

- Secondary: 普通从节点, 可被选为主节点, 以下都是特殊从节点.  
- Arbiter(不推荐): arbiter节点只参与投票, 不能被选为Primary, 并且不从Primary同步数据  
- Priority0: priority0节点的选举优先级为0, 不会被选举为Primary  
- Vote0: vote0节点不参与投票 (复制集成员最多50个, 参与Primary选举投票的成员最多7个,  其他成员都是Vote0)  
- Hidden: hidden节点不能被选为主(Priority为0), 并且对Driver不可见.(可使用Hidden节点做一些数据备份、离线计算的任务)  
- Delayed: delayed节点必须是Hidden节点, 并且其数据落后与Primary一段时间  
  
| 节点类型   | 可读 | 可写 | 投票 | oplog操作 | 当选primary | 否决 | 备注               |
| :--------- | :--- | :--- | :--- | :-------- | :---------- | :--- | :----------------- |
| primary    | O    | O    | O    | 生成      | -           | O    | -                  |
| secondary  | O    | X    | O    | 同步      | O           | O    | 常规的seconday     |
| Priority=0 | O    | X    | O    | 同步      | X           | O    | -                  |
| Hidden     | X    | X    | O    | 同步      | X           | O    | Priority=0, 不可见 |
| Delayed    | X    | X    | O    | 同步      | X           | O    | 为Hidden, 延迟同步 |
| Arbiter    | X    | X    | O    | X         | X           | O    | Priority=0, 无数据 |
| vote=0     | O    | X    | X    | 同步      | O           | O    | 不能投票           |

> 客户端一般会保持连接多个实例(主从从从选...都有连接), 以确保主挂后可以从其他实例拿到最新的副本集状态, 进而连接到新的主节点.(若只连接主(readPref = Primary), 主跪了, 客户端便不能得到任何服务)  
  
## 选举因素

健康监测  

- 节点间心跳  

节点优先级  

- 投票给优先级最高的节点  
- 优先级为0的节点不会主动发起选举  
- 当Primary发现有优先级更高Secondary, 并且该Secondary的数据落后在10s内, 则Primary会主动降级, 让优先级更高的Secondary有成为Primary的机会.  

拥有较新的oplog  

- 拥有最新optime(最近一条oplog的时间戳)的节点才能被选为主.  

多数派连接  

- 一个member要成为primary, 它必须与“多数派”的其他members建立连接, 如果未能与足够多的member建立连接, 事实上它本身也无法被选举为primary；多数派参考的是“总票数”, 而不是member的个数, 因为我们可以给每个member设定不同的“票数”.假设复制集内投票成员数量为N, 则大多数为 N/2 + 1.  

## [事务](mongodb-transaction.md)
