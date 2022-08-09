# redis 分布式锁

    分布式服务器获取锁

## 最低保障的分布式锁特性

1. 安全属性(Safety property): 独享(相互排斥). 在任意一个时刻, 只有一个客户端持有锁.
2. 活性A(Liveness property A): 无死锁. 即便持有锁的客户端崩溃(crashed)或者网络被分裂(gets partitioned), 锁仍然可以被获取.
3. 活性B(Liveness property B): 容错. 只要大部分Redis节点都活着, 客户端就可以获取和释放锁.

## 故障迁移时

1. 客户端A从master获取到锁
2. 在master将锁同步到slave之前, master宕掉了.
3. slave节点被晋级为master节点
4. 客户端B取得了同一个资源被客户端A已经获取到的另外一个锁. 安全失效!

## 单实例锁

`SET resource_name random_value NX PX 30000`

- NX key不存在时才能执行成功
- PX 设置30s过期
- random_value 为了更安全的释放锁, 值需要随机性


```lua
-- 安全释放锁
if redis.call("get",KEYS[1]) == ARGV[1] then
    return redis.call("del",KEYS[1])
else
    return 0
end
```

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
