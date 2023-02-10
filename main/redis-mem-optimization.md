# redis内存调优  

## 占用预估

```js
    # 同等数据量, 不同结构对比
    `hash` 最优 (必须是值短的 用ziplist)
    `list`, `set` 次之
    `string` 容易全局hash碰撞, 占用较多(碎片多)
    `zset` 由于skiplist结构复杂, 占用最大(索引复杂)
```

- 内存占用估算 <http://www.redis.cn/redis_memory/>

## 内存压缩

    以CPU换取内存

```js
# redis.conf
hash-max-zipmap-entries 64 (2.6以上使用hash-max-ziplist-entries)
hash-max-zipmap-value 512  (2.6以上使用hash-max-ziplist-value)
list-max-ziplist-entries 512
list-max-ziplist-value 64
zset-max-ziplist-entries 128
zset-max-ziplist-value 64
set-max-intset-entries 512
```

## 内存分配

    redis缓存被删除后, 并不会立刻把内存归还给OS, 即磁盘碎片, 这是系统函数malloc()的特性.
    但内存分配器是智能的, 新建的key可复用已经被释放的内存.

## 数据结构

### 位级别和字级别的操作

    setbit  
    getbit  
    bitop  多个key进行AND,OR,XOR,NOT运算, 生成结果于destkey
    bitpos  查询第一个bit(0 / 1)的位置
    bitcount  bit位为1的个数
    bitfield  

