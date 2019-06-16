# nginx 哈希表

## 场景
静态不变的内容 即通常不会插入，删除

## 源码
```
src/core/ngx_hash.h
src/core/ngx_hash.c
```

## hash_max_size
最大hash bucket个数

## hash_bucket_size
与cpu cache向上对齐(64b)


