# redisLRU缓存

## maxmemory配置触发删除

### 配置

redis.conf -> maxmemory

### 策略

noeviction:返回错误当内存限制达到并且客户端尝试执行会让更多内存被使用的命令（大部分的写入指令，但DEL和几个例外）  
allkeys-lru: 尝试回收最少使用的键（LRU），使得新添加的数据有空间存放。  
volatile-lru: 尝试回收最少使用的键（LRU），但仅限于在过期集合的键,使得新添加的数据有空间存放。  
allkeys-random: 回收随机的键使得新添加的数据有空间存放。  
volatile-random: 回收随机的键使得新添加的数据有空间存放，但仅限于在过期集合的键。  
volatile-ttl: 回收在过期集合的键，并且优先回收存活时间（TTL）较短的键,使得新添加的数据有空间存放。  
