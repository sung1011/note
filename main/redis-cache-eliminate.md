# redis缓存淘汰

## maxmemory配置触发删除

### 配置

redis.conf -> maxmemory

#### LRU实现

redis-server维护24bit全局时钟， 每隔一段时间更新这个时钟。  
新增自定义的key对象时会将全局时钟赋值给key对象时钟(redisObject.lru:24)。  
LRU时会对比两个时钟，淘汰时钟相差最久的。  
24bit按秒为单位最多存储`194天`，所以可能出现key时钟大于全局时钟情况。这时就不是两个相加而是相减来求最久了。

### LRU策略

1. noeviction: 返回错误当内存限制达到并且客户端尝试执行会让更多内存被使用的命令（大部分的写入指令，但DEL和几个例外）  
2. allkeys-lru: 尝试回收最少使用的键（LRU），使得新添加的数据有空间存放。  
3. volatile-lru: 尝试回收最少使用的键（LRU），但仅限于在过期集合的键，使得新添加的数据有空间存放。  
4. allkeys-random: 回收随机的键使得新添加的数据有空间存放。  
5. volatile-random: 回收随机的键使得新添加的数据有空间存放，但仅限于在过期集合的键。  
6. volatile-ttl: 回收在过期集合的键，并且优先回收存活时间（TTL）较短的键，使得新添加的数据有空间存放。  

#### LRU问题

|处淘汰，LRU会将A淘汰，但A使用频率比较高。redis4.0+的LFU策略解决这一问题。

A~~A~~A~~A~~A~~A~~A~~A~~A~~A~~~|
B~~~~~B~~~~~B~~~~~B~~~~~~~~~~~B|

### LFU策略 (Least Frequently Used)

#### LFU实现

LFU将LRU的24bit key对象时钟分为 `16bit时钟` + `8bit计数器`

- 前16bit 依然表示时间，但单位为小时（2^16=65536h=2730d）
- 后8bit 表示当前key访问频率，是一个非线性的表达。

```code
+--------+------------+------------+------------+------------+------------+
| factor | 100 hits   | 1000 hits  | 100K hits  | 1M hits    | 10M hits   |
+--------+------------+------------+------------+------------+------------+
| 0      | 104        | 255        | 255        | 255        | 255        |
+--------+------------+------------+------------+------------+------------+
| 1      | 18         | 49         | 255        | 255        | 255        |
+--------+------------+------------+------------+------------+------------+
| 10     | 10         | 18         | 142        | 255        | 255        |
+--------+------------+------------+------------+------------+------------+
| 100    | 8          | 11         | 49         | 143        | 255        |
+--------+------------+------------+------------+------------+------------+
```

#### LFU策略

lfu-log-factor 10 //key命中因子
lfu-decay-time 1 //控制后8位访问频率递减的时间
