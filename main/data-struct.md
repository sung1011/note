# 数据结构

## 分类

- [浮点数 Float](float.md)

      浮点数是指用符号、尾数、基数和指数这四部分来表示的小数

- [数组 Array](ds-array.md)  

      有限个相同类型元素组成的有序(连续内存空间存储)集合

- [链表 LinkedList](ds-linkedlist.md)  

      非连续, 非顺序(随机离散存储)的有若干节点组成.易增删

- [栈 Stack](ds-stack.md)  

      先进后出的数组或链表

- [队列 Queue](ds-queue.md)

      先进先出的数组或链表

- [哈希 Hash](ds-hash.md)

      对象转为数字, 作为数组下标
      可以转bitmap节约空间

- [位图 Bitmap](ds-bitmap.md)

      数据映射为bit数组, 只存0或1

- [堆 PriorityQueue/heap](ds-heap.md)

      数组形式的二叉树
      最大堆: 父节点 > 子节点, 子节点无序
      最小堆: 父节点 < 子节点, 子节点无序

- [跳表 SkipList](ds-skiplist.md)

      多重索引的有序链表

- [倒排索引 inverted index](ds-inverted-index.md)

      根据关键内容(tag)找key

- [树 Tree](ds-tree.md)

      由多个节点组成, 具有层次关系

- [二叉树 Binary Tree](ds-binary-tree.md)

      每个节点拥有0~2个子节点

- [二叉检索树 Binary Search Tree](ds-binary-search-tree.md)

      各节点值不同 且 有序 (左 < 根 < 右)

- [平衡二叉树 AVL Tree](ds-AVL-tree.md)

      自平衡(左右层级差不会大于1)

- [B-树 平衡多路查找树 B-Tree](ds-b-tree.md)

      每个节点都有key(索引)和value(数据), 搜索可能在中间节点结束

- [B+树 B+Tree](ds-b+tree.md)

      中间节点不存储value(数据), 只存储key(索引), value都在叶子节点里
      每个叶子节点都有兄弟叶子的指针

- [LSM Log Structured Merge](ds-LSM.md)

      先写内存, 数据写满后, 逐层滚动地归并 顺序的 写入磁盘
      写效率高

- [红黑树](ds-rbtree.md)  

      不严格控制左、右子树层级差, 但再平衡时效率更高
      比AVL的增删改效率高

- [字典树](trie-tree.md)

      利用字符串的公共前缀来减少查询时间, 减少无谓的字符串比较,
      空间效率 遍历效率 插入删除效率 比哈希表高, 且不会冲突

## ref

- [数据结构](ref/data-struct.md)

- [IO wiki](https://oi-wiki.org/ds/)
