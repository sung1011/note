# redis 有序集合对象

## encoding

ziplist, skiplist

实例

```c
typedef struct zset {
    dict *dict;
    zskiplist *zsl;
} zset;
```
