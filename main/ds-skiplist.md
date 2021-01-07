# 跳表 skip list

![img](res/ds-skiplist.png)

## 概念

    多重索引的有序链表

## 特征

- `有序` 范围查找

## 结构

    SKIP_LIST {
        NODE head; # 表头
        int level; # 层级
    }

    NODE {
        data; # 数据

        # 每个元素表示该层下一个node的地址
        # next[1]是第1层的下一个node地址，next[2]是第2层的下一个node地址...
        # null表示末尾
        NODE []next; # 节点
    }

## 操作

- `查 O(log n)`

- `增 O(log n)`

- `删 O(log n)`

## ref

- [深入理解Redis跳跃表的基本实现和特性](https://juejin.cn/post/6893072817206591496)
