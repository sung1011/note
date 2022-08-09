# redis编码

@@redis

- `OBJ_ENCODING_RAW` 0
- `OBJ_ENCODING_INT` 1
- `OBJ_ENCODING_HT` 2
- `OBJ_ENCODING_ZIPMAP` 3
- `OBJ_ENCODING_LINKEDLIST` 4
- `OBJ_ENCODING_ZIPLIST` 5
- `OBJ_ENCODING_INTSET` 6
- `OBJ_ENCODING_SKIPLIST` 7  
- `OBJ_ENCODING_EMBSTR` 8
- `OBJ_ENCODING_QUICKLIST` 9
- `OBJ_ENCODING_STREAM` 10

## [整形 INT](_) TODO

    int

## [简单动态字符串 SDS (embstr && raw)](redis-encoding-sds.md)

    len, alloc, type, data

## [链表 linkedlist](redis-encoding-linkedlist.md)

    prev, next, data

## [压缩列表 ziplist](redis-encoding-ziplist.md)

    entries的数据存储在连续内存, 数据可压缩

## [快速列表 quicklist](redis-encoding-quicklist.md)

    linkedlist + ziplist

## [字典 hashtable](redis-encoding-hashtable.md)

    hash表

## [跳跃表 skiplist](redis-encoding-skiplist.md)

    member, score 包含多重索引的双向链表

## [整数集合 intset](redis-encoding-intset.md)

    有序 去重

## [stream](_) TODO
