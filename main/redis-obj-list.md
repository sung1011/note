# redis list对象

OBJ_LIST

## encoding

OBJ_ENCODING_QUICKLIST  
<!-- OBJ_ENCODING_ZIPLIST   -->
<!-- OBJ_ENCODING_LINKEDLIST   -->

### 实例

```c
// ZIPLIST
redisObject {
    type: REDIS_LIST;
    encoding: REDIS_ENCODING_ZIPLIST;
    ...
    *ptr:  &ZIPLIST{TODO};
};

// LINKEDLIST
redisObject {
    type: REDIS_LIST;
    encoding: REDIS_ENCODING_LINKEDLIST;
    ...
    *ptr:  &LINKEDLIST{TODO};
};
```

### 转换

| encoding   | 条件                                |
| ---------- | ----------------------------------- |
| linkedlist | 元素值字符长度 > 64B 或 len > 512   |
| ziplist    | 元素值字符长度 <= 64B && len <= 512 |

> 条件可通过配置修改 **list-max-ziplist-value**、**list-max-ziplist-entries**

<!-- ## 实现

| cmd     | ziplist                                                      | linkedlist                                                    |
| ------- | ------------------------------------------------------------ | ------------------------------------------------------------- |
| LPUSH   | 调用ziplistPush                                              | 调用listAddNodeHead                                           |
| RPUSH   | 调用ziplistPush                                              | 调用listAddNodeTail                                           |
| LPOP    | 调用ziplistIndex定位表头，调用ziplistDelete删除表头。        | 调用listFirst定位表头，调用listDelNode删除表头。              |
| RPOP    | 调用ziplistIndex定位表尾，调用ziplistDelete删除表尾。        | 调用listLast定位表尾，调用listDelNode删除表尾。               |
| LINDEX  | 调用ziplistIndex定位节点，返回节点所保存的元素。             | 调用listIndex定位节点，然后返回节点所保存的元素。             |
| LLEN    | 调用ziplistLen返回压缩列表的长度。                           | 调用listLength返回双端链表的长度。                            |
| LINSERT | 插入表头表尾时ziplistPush；插入其他ziplistInsert。           | 调用listInsertNode，将新节点插入到双端链表的指定位置。        |
| LREM    | 遍历节点，并调用ziplistDelete删除包含了给定元素的节点。      | 遍历双端链表节点，并调用listDelNode删除包含了给定元素的节点。 |
| LTRIM   | 调用ziplistDeleteRange，删除不在索引范围内的节点。           | 遍历双端链表节点，并调用listDelNode删除不在索引范围内的节点。 |
| LSET    | ziplistDelete删除指定节点，然后调用ziplistInsert插入新节点。 | listIndex定位节点，然后赋值更新节点的值。                     | --> |
