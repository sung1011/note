# redis内存调优  

## 内存压缩

以CPU换取内存

```conf
//redis.conf
hash-max-zipmap-entries 64 (2.6以上使用hash-max-ziplist-entries)
hash-max-zipmap-value 512  (2.6以上使用hash-max-ziplist-value)
list-max-ziplist-entries 512
list-max-ziplist-value 64
zset-max-ziplist-entries 128
zset-max-ziplist-value 64
set-max-intset-entries 512
```

## 32位redis

每个key内存占用更少，但整个实例限制4G以下。  
AOF，RDB不区分32位和64位（可以互相备份恢复）。

## 位级别和字级别的操作

setbit  
getbit  
bitop  多个key进行AND,OR,XOR,NOT运算，生成结果于destkey
bitpos  查询第一个bit（0 / 1）的位置
bitcount  bit位为1的个数
bitfield  

## 内存分配

当某些缓存被删除后Redis并不是总是立即将内存归还给操作系统。这并不是redis所特有的，而是函数malloc()的特性。例如你缓存了5G的数据，然后删除了2G数据，从操作系统看，redis可能仍然占用了5G的内存（这个内存叫RSS,后面会用到这个概念），即使redis已经明确声明只使用了3G的空间。这是因为redis使用的底层内存分配器不会这么简单的就把内存归还给操作系统，可能是因为已经删除的key和没有删除的key在同一个页面（page）,这样就不能把完整的一页归还给操作系统.  
上面的一点意味着，你应该基于你可能会用到的 最大内存 来指定redis的最大内存。如果你的程序时不时的需要10G内存，即便在大多数情况是使用5G内存，你也需要指定最大内存为10G.  
内存分配器是智能的，可以复用用户已经释放的内存。所以当使用的内存从5G降低到3G时，你可以重新添加更多的key，而不需要再向操作系统申请内存。分配器将复用之前已经释放的2G内存.  
因为这些，当redis的peak内存非常高于平时的内存使用时，碎片所占可用内存的比例就会波动很大。当前使用的内存除以实际使用的物理内存（RSS）就是fragmentation；因为RSS就是peak memory，所以当大部分key被释放的时候，此时内存的mem_used / RSS就比较高.  
