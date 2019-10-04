# reids对象

## 数据结构

```js
typedef struct redisObject {
    unsigned type:4; // 类型
    unsigned encoding:4; // 编码
    unsigned lru:LRU_BITS;
    int refcount;
    void *ptr; // 指向底[层数据结构](redis-data-struct.md)的指针
} robj;
```

## [string对象](redis-obj-string.md)

## [list对象](redis-obj-list.md)

## [hash对象](redis-obj-hash.md)

## [set对象](redis-obj-set.md)

## [zset对象](redis-obj-zset.md)

## 对象回收

## 对象共享
