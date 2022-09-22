# redis 分布式锁

    分布式服务器获取锁

## 最低保障的分布式锁特性

1. 安全属性(Safety property): 独享(相互排斥). 在任意一个时刻, 只有一个客户端持有锁.
2. 活性A(Liveness property A): 无死锁. 即便持有锁的客户端崩溃(crashed)或者网络被分裂(gets partitioned), 锁仍然可以被获取.
3. 活性B(Liveness property B): 容错. 只要大部分Redis节点都活着, 客户端就可以获取和释放锁.

## 模式

### 加锁set解锁del

```bash
    # 加锁 
       set {key} {unique_random} EX 5 NX
       # 设置 过期时间 + 若key不存在 + 校验唯一随机值

    # 解锁
        # redis
        del {key}

        # lua
        if redis.call("get",KEYS[1]) == ARGV[1] then
          return redis.call("del",KEYS[1])
        else
          return 0
        end
```

> issue: 线程a抢到锁后 若执行时间过长, 可能错误的删除线程b的锁; 需要[compare && del](redis-distlock.md#compare&&del解锁)解决

### compare&&del解锁

```bash
    # 加锁
       set {key} {unique_random} EX 5 NX
    
    # 解锁
        # redis
        get
        del

        # lua
        // 加锁
        String uuid = UUID.randomUUID().toString().replaceAll("-","");
        SET key uuid NX EX 30
        // 解锁
        if (redis.call('get', KEYS[1]) == ARGV[1])
            then return redis.call('del', KEYS[1])
        else return 0
        end
```

> issue: a执行时间长, b获取锁执行, 此时a, b在同时执行.

### 守护线程

```bash
    # 启动守护线程, 在锁过期之前给锁"续命"
    # a获得锁执行很慢, 锁即将超时, 守护线程给a锁的过期时间延长n秒
```

### 可重入锁

```bash
    # 若需要线程在持有锁的情况下再次请求锁, 就是可重入锁
    # hincrby 计数每个线程的锁情况 加锁+=1, 解锁-=1
     // 如果 lock_key 不存在
     if (redis.call('exists', KEYS[1]) == 0)
     then
         // 设置 lock_key 线程标识 1 进行加锁
         redis.call('hset', KEYS[1], ARGV[2], 1);
         // 设置过期时间
         redis.call('pexpire', KEYS[1], ARGV[1]);
         return nil;
         end;
     // 如果 lock_key 存在且线程标识是当前欲加锁的线程标识
     if (redis.call('hexists', KEYS[1], ARGV[2]) == 1)
         // 自增
         then redis.call('hincrby', KEYS[1], ARGV[2], 1);
         // 重置过期时间
         redis.call('pexpire', KEYS[1], ARGV[1]);
         return nil;
         end;
     // 如果加锁失败，返回锁剩余时间
     return redis.call('pttl', KEYS[1]);
```

### 发布订阅阻塞等待unlock

    其他竞争的线程不轮询, 而是通过订阅发布功能去抢夺锁

```bash
    # 1. 线程a加锁成功
    # 2. 线程bcd加锁失败, 订阅"锁释放"消息, 阻塞等待
    # 3. a解锁, 发布"锁释放"消息, 到mq
    # 4. bcd从mq获得"锁释放"消息 并 加锁
```

### 集群故障迁移时

1. a线程从master获取到锁
2. 在master将锁同步到slave之前, master宕掉了.
3. slave节点被晋级为master节点
4. b线程获得本该a线程占用的锁, ab同时执行任务

## RedLock算法

### 流程

1. 获取当前Unix时间为开始获取锁的时间.`start_t`
2. 依次尝试从N个实例(如5个实例), 使用相同的key和随机值获取锁.
3. 计算获取锁所用时间 `cost_t = now - start_t`.
4. 当大多数实例(如5分之3实例成功)获取锁且cost_t小于锁过期时间, 则获取成功.
5. 若获取锁失败(如5分之2实例成功), 则进行解锁(5个实例都解锁, 不论是否加锁成功).
6. 无脑向所有实例执行释放锁

> `决定因素` 获取锁 + 耗时

> `key的真正过期时间` = 过期时间 - cost_t

### 是否异步

### 失败重试

### 安全性

### 活性争议

### 性能、崩溃、同步

### 锁的扩展

## ref

- <http://redis.cn/topics/distlock.html> redis分布式锁

- <https://xiaomi-info.github.io/2019/12/17/redis-distributed-lock/>