# 跳表 skip list

![img](res/ds-skiplist.png)

    多重索引的有序链表

## 场景

    优化链表查询速度
    数据(member)和其数值(score), 善于对score范围查找

## 特征

- `有序` score的范围查找
- `数据` 包含成员 分数 排名

## 结构

```js
    SKIP_LIST {
        NODE head; # 表头
        int level; # 层级
    }

    NODE {
        data; # 数据

        # 每个元素表示该层下一个node的地址
        # next[1]是第1层的下一个node地址, next[2]是第2层的下一个node地址...
        # null表示末尾
        NODE []next; # 节点
    }
```

### 层

    加快访问其他节点的速度, 一般层越多越快, 层高在1~32随机的

### 跨度

    节点间的距离, 即排位.

> 前进节点=NULL则跨度=0; 前进节点有值则至少跨度=1  

### 前进指针

    用于遍历

### 后退指针

    用于逆向遍历

### 分值和成员

    成员互斥, 分值可相同. 分值相等时, 按成员排序.

## 操作

- `find O(log n)`

- `insert O(log n)`

- `delete O(log n)`

## 实战

- redis zset
