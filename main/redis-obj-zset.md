# redis 有序集合对象

## encoding

OBJ_ENCODING_ZIPLIST
OBJ_ENCODING_SKIPLIST

### 实例

```c
typedef struct zset {
    dict *dict;
    zskiplist *zsl;
} zset;

// ZIPLIST
redisObject {
    type: REDIS_ZSET,
    encoding: REDIS_ENCODING_ZIPLIST,
    ...
    ptr:  &ZIPLIST{TODO}
}

// SKIPLIST
redisObject {
    type: REDIS_ZSET,
    encoding: REDIS_ENCODING_SKIPLIST,
    ...
    ptr:  &SKIPLIST{TODO}
}
```

### 转换

| encoding | 条件                                |
| -------- | ----------------------------------- |
| skiplist | 元素值字符长度 > 64b 或 len > 128   |
| ziplist  | 元素值字符长度 <= 64b && len <= 128 |

> 条件可通过配置修改 **list-max-ziplist-value**、**list-max-ziplist-entries**
