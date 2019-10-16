# redis 集合对象

## encoding

- [OBJ_ENCODING_INTSET](redis-encoding-intset.md)
- [OBJ_ENCODING_HT](redis-encoding-hashtable.md)

### 实例

```c
// INTSET
redisObject {
    type: REDIS_SET,
    encoding: REDIS_ENCODING_INTSET,
    ...
    ptr:  &INTSET{TODO}
}

// HT
redisObject {
    type: REDIS_SET,
    encoding: REDIS_ENCODING_HT,
    ...
    ptr:  &DICT{TODO}
}
```

### 转换

| encoding | 条件                       |
| -------- | -------------------------- |
| ht       | 元素值非整数 或 len > 512  |
| intset   | 元素值为整数 && len <= 512 |

> 条件可通过配置修改**list-max-intset-entries**
