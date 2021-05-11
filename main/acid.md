# 事务 (transactions)  

> 可以想象成 余额`B`是订单`O`的冗余数据, 同时修改余额和订单, 需要事务

## ACID 四大特性  

- `A 原子性` 即不可分割性,事务要么全部被执行,要么就全部不被执行 -- 行为
- `C 一致性` 数据库从一种正确状态转换成另一种正确状态 -- 状态
- `I 隔离性` 在事务正确提交之前,不允许把该事务对数据的任何改变提供给任何其他事务 -- 独立
- `D 持久性` 事务正确提交后,其结果将永久保存在数据库中,即使在事务提交后有了其他故障,事务的处理结果也会得到保存 -- 保存

> ACI强调的角度不同, 但概念有点杂糅

## 隔离级别  

| 隔离级别    | 脏读 DR| 不可重复读 NR | 幻读 PR |
| ----------- | ---- | ---------- | ---- |
| 读未提交 RU | o    | o          | o    |
| 读已提交 RC | x    | o          | o    |
| 可重复读 RR | x    | x          | o    |
| 串行 S      | x    | x          | x    |

### RC读已提交 vs RR可重复读

```sql
# (session)A看到初始状态为一笔订单(id:3) 和 100余额, 并且A`未commit事务`
$A begin;
$A select * from orders; # [ id:3, amount:100 ]
$A select * from balance; # blance: 100

    # (session)B在A`未commit事务`的情况下, commit了一个事务, 插入一笔订单(id:4),同时增加50余额
    $B begin;
    $B insert into orders values(4, 50);
    $B update balance set balance = balance + 50;
    $B commit;

# 若RC 读已提交: A`未commit事务`的情况重复读取数据, order和balance`会读到B的提交`, 即被其他事务干扰
$A select * from orders; # [ id:3, amount:100 ], [ id:4, amount:50 ]
$A select * from balance; # blance: 150
$A ...
$A commit;

# 若RR 可重复读: A`未commit事务`的情况重复读取数据, order和balance`不会读到B的提交`, 即体现了事务的隔离性
$A select * from orders; # [ id:3, amount:100 ]
$A select * from balance; # blance: 100
$A ...
$A commit;
```

> `读未提交` RU read-uncommitted -- 不常见; O已经生成, 读到尚未修改的b的值
>
> `读已提交` RC read-committed -- 非幂等读取; 读O, B的一致状态;
>
> `可重复读` RR repeatable-read -- 幂等读取; 读O, B的一致状态;
>
> `串行化` serializable -- 不常见; 无并发,性能差,同步锁,无需隔离
>
> `脏读` DR dirty-read
>
> `不可重复读读` NR nonrepeatable-read
>
> `幻读` PR phantom-read -- 不常见; 一般无恶劣影响; RR情况下, session A, B同时进行插入都认为id可以自增到101, 由于AB隔离级别是RR, 所以读不到对方新插入的记录, 导致主键冲突 (像幻觉一样读不到对方的101, 但会冲突)

## CAP  

- `P` (Partition tolerance): 分区容错 --- 指网络故障时两个分区节点是否相同 --- 一般认为P总是成立，剩下的C, A无法同时做到。  
- `C` (Consistency): 一致性 --- 两个实例保证相同状态。
- `A` (Availability): 可用性 --- 两个实例挂掉一个另一个还能提供服务。
- `CP型` 牺牲可用性，保证强一致性。
- `AP型` 牺牲一致性，保证高可用性。

## 分布式事务

### 2PC 二阶段提交

#### 成员

- 1个协调者节点 `C` coordinator
- N个参与者节点 `P` partcipant

#### 阶段

![img](res/2pc0.jpeg)

![img](res/2pc1.jpeg)

##### 阶段1: 提出 propose

  1. `C` -> `P` 请求事务询问
  2. `P` 本地执行事务但不commit
  3. `P` -> `C` 响应询问结果

##### 阶段2: 提交/执行 commit

- 当响应询问结果全部ok
  1. `C` -> `P` 请求事务commit
  2. `P` 事务commit

- 当响应询问结果任一异常 or 等待超时
  1. `C` -> `P` 发送rollback请求
  2. `P` 本地执行事务rollback

#### 缺点

##### 性能

##### 单点`C`

- `C`正常, `P`挂了

```bash
  现象: C阻塞等待P响应  
  解决: 超时机制
```

- `C`挂了, `P`正常

```bash
  现象: 由于C是单点, 所以无论处于哪个阶段, C挂后未提交的操作都会阻塞
  解决: C改为主备两个节点, 并记录oplog. C主挂后由C备取代, 并读取oplog
```

- `C`, `P`都挂了

```bash
  TODO
```

### 3PC

### TCC

### Saga

[2PC和3PC原理](https://www.cnblogs.com/qdhxhz/p/11167025.html)
