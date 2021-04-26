# redis 哈希对象

## encoding

- OBJ_ENCODING_ZIPLIST  
- OBJ_ENCODING_HT  

### 实例

```c
// ZIPLIST
redisObject {
    type: REDIS_HASH,
    encoding: REDIS_ENCODING_ZIPLIST,
    ...
    ptr:  &ZIPLIST{TODO}
}

// HT
redisObject {
    type: REDIS_HASH,
    encoding: REDIS_ENCODING_HT,
    ...
    ptr:  &DICT{TODO}
}
```

### 转换

| encoding | 条件                                |
| -------- | ----------------------------------- |
| ht       | 元素值字符长度 > 64B 或 len > 512   |
| ziplist  | 元素值字符长度 <= 64B && len <= 512 |

> 条件可通过配置修改 `list-max-ziplist-value`、`list-max-ziplist-entries`

<!-- ## 实现 -->
