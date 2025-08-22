# algorithm

## 分类

- [递归]

- [排序]

  - 冒泡排序(Bubble Sort)
  - 插入排序(Insertion Sort)
  - 希尔排序(Shell Sort)
  - 选择排序(Selection Sort)
  - 快速排序(Quick Sort)
  - [归并排序(Merge Sort)](algo-sort-merge.md)
  - 堆排序(Heap Sort)
  - 计数排序(Counting Sort)
  - 桶排序(Bucket Sort)
  - 基数排序(Radix Sort)

- [二分查找]

- [搜索]

- [贪心]

- [分治]

- [回溯]

- [动态规划]

- [字符串匹配算法]

- [雪花 snowflake](algo-snowflake.md)

    分布式全局唯一ID

- [位图 bitmap](ds-bitmap.md)  

    数据映射为bit数组, 只存0或1

- [布隆 BloomFilter](algo-bloomfilter.md)  

    将一个元素带入多个哈希算法, 结果存入bitmap

- [RBM RoaringBitmap](algo-roaringbitmap.md)

    将大块的bitmap分成各个小块，其中每个小块在需要存储数据的时候才会存在
    解决bitmap存稀松数字的却需要长len结构的问题

- [哈希 hash](algo-hash.md)

    将给定数据转化为`固定长度的不规则值`的函数

- [一致性哈希 DHT](algo-DHT.md)  

    一致性哈希算法是一种分布式哈希算法，主要用于负载均衡和缓存分布式存储

- [共识算法 raft](algo-raft.md)

    一种分布式一致性算法，用于保证分布式系统中各个节点之间的数据一致性

- [海盗分金币](algo-pirate-gold.md)

    有n个海盗，分金币，规则是先提出分配方案的海盗，如果有超过半数的海盗同意，则方案通过，否则提出方案的海盗被扔下海

- [工作量证明算法 PoW](algo-pow.md)

    一种防止DDoS攻击的算法，通过计算机的计算能力来证明自己的身份

- [权益证明算法 PoS](algo-pos.md)

    一种防止DDoS攻击的算法，通过持有的货币数量来证明自己的身份

- [代表权益证明算法 DPos](algo-dpos.md)

    一种防止DDoS攻击的算法，通过持有的货币数量来证明自己的身份

- [漏桶](algo-leaky-bucket.md)

    一种限流算法，用于控制数据传输速率

- [令牌桶](algo-token-bucket.md)

   一种限流算法，用于控制数据传输速率

- [广度优先算法 Breadth First Search (BFS)](algo-bfs.md)

    一种图搜索算法，用于寻找图中的最短路径

- [深度优先算法 Deep First Search (DFS)](algo-dfs.md)

    一种图搜索算法，用于寻找图中的全部路径

- [A*算法](algo-a-star.md)

    一种启发式搜索算法，用于寻找图中的最短路径

- [极小极大搜索](algo-minimax.md)

    一种博弈树搜索算法，用于寻找最优解

- [鸟群算法](algo-boid.md)

    一种基于鸟群行为的优化算法，用于群体移动

## ref

[算法](https://cloud.tencent.com/developer/article/1101517)
