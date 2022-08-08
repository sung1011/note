# redis 编码 quicklist 快速列表

    OBJ_ENCODING_QUICKLIST

## 作用于

    OBJ_LIST

## souce code

    src/quicklist

## 特性

    quicklist 是linkedlist + ziplist的组合
        双端队列, 操作队首队尾都是O(1)
        ziplist越小而多, 造成内存碎片. 极端情况 每个ziplist只有1个数据entry, 即退化为linkedlist
        ziplist越大而少, 造成分配ziplist连续内存困难. 极端情况 quicklist只剩1个大的ziplist, 即退化为ziplist

> quicklist重要配置参数 `每个node的字节大小限制 list-max-ziplist-value`、`每个node的entries数量 list-max-ziplist-entries`

## 数据结构

![img](res/redis-encoding-quicklist.png)

```c
// 快速列表
typedef struct quicklist {
    quicklistNode *head;
    quicklistNode *tail;
    unsigned long count;        /* 所有node中的元素总和; total count of all entries in all ziplists */
    unsigned long len;          /* node的个数; number of quicklistNodes */
    int fill : 16;              /* ; fill factor for individual nodes */
    unsigned int compress : 16; /* 两端各有n个节点不压缩; depth of end nodes not to compress;0=off */
} quicklist;

// 快速列表节点
typedef struct quicklistNode {
    struct quicklistNode *prev;
    struct quicklistNode *next;
    unsigned char *zl;           /* node指向的ziplist */
    unsigned int sz;             /* zl的字节大小; ziplist size in bytes */
    unsigned int count : 16;     /* zl中的元素个数; count of items in ziplist */
    unsigned int encoding : 2;   /* 编码格式; RAW==1 or LZF==2 */
    unsigned int container : 2;  /* 存储方式; NONE==1 or ZIPLIST==2 */
    unsigned int recompress : 1; /* 是否被压缩; was this node previous compressed? */
    unsigned int attempted_compress : 1; /* 是否能被压缩 node can't compress; too small */
    unsigned int extra : 10; /* 预留的bit位; more bits to steal for future usage */
} quicklistNode;
```

## API

TODO
