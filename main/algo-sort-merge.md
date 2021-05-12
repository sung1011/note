# 归并排序

## 概念

      将链表划分为n个子链表
      `子链表` 间 `逐个元素对比排序` 后合并
      递归地进行上一步操作, 最终得到一个排序后的链表

> `子链表` 子链表内永远是有序的  
>
> `逐个元素对比排序` 两个链表长度分别为 m 和 n .若无序, 两层遍历进行排序O(m*n)；若有序, 逐个元素对比, 将较小的元素push到新链表 `O(m+n)`  
>
> 常用于两/多个有序链表合并

## 特征

- `速度稳定`

- `适用于总体无序, 各子项有序`

## 时间

- `optimal` O(k*log(n))
- `worst` O(k*log(n))
- `avg` O(k*log(n))

## 空间

O(n)

## 操作

  ![img](res/inverted-index-step.png)

## 优化归并

- `双链表` 普通归并

      链表 + 链表: O(a+b)

- `有序链表 + 跳跃表` 其一比较长时

  ![img](res/inverted-skiplist.png)
  ![img](res/inverted-skiplist2.png)

      a <1, 500, 888> (链表)
      b <1, 2, 3, ..., 999> (跳表)

      跳跃表a + 链表b = 在a中遍历 O(b) + 利用b跳表的索引 O(log(a)) 

- `跳跃表 + 跳跃表 互相二分查找` a, b差不多都挺长时

   ![img](res/inverted-binary.png)

      a <1, 3, 4, 6, ..., 999> (跳表)
      b <1, 2, 5, 6, ..., 987> (跳表)

      跳跃表 + 跳跃表 = O(log(a)) + O(log(b))
      两跳跃表互相二分查找: 当元素a0 < b0, 以b元素做key在a中利用跳表的索引快速往前跳, 反之亦然. O(log(min(a, b)))

- `哈希+有序链表` 其一比较短时, 长链中包含短链时

   ![img](res/inverted-hash.png)

      a <1,2,3,4, ..., 999> (哈希)
      b <1, 500, 888> (链表)

      哈希a + 链表b = 遍历b中每个元素去a的hash中进行查询, 由于hash查询是O(1), 所以整体是O(b)

- `bitmap + bitmap` a, b间隔稠密, 长度有限

      a <1,2,3,4, ..., 999>
      b <1, 500, 888>

      bitmap + bitmap: 求交、并、补 O(1).

## 优化 求交、并、补

- `调整顺序`

      多个集合进行求交、并、差（多路归并）（联合查询）时, 当abc的个数有较大差异, (a ∩ b) ∩ c 和 a ∩ (b ∩ c) 速度有差异, 但都是O(n).

- `快速多路归并`

      利用跳跃表的特性加快多路归并的效率
      如a, b, c, d四个有序链表求交集
      1. 将4个链表第一个元素取出, 其中最大一个元素记为max变量
      2. 从a开始, 如果当前位置的值小于max, 则用`跳表法`快速将指针调整到a中第一个 >= max 的元素位置.新位置元素若 > max, 则更新max.
      3. b, c, d依次执行上一步操作.

    ![img](res/merge-sort-col1.png)
    ![img](res/merge-sort-col2.png)
    ![img](res/merge-sort-col3.png)

- `预先联合`

      若 a ∩ b ∩ c 经常查询, 可以预先将a_b_c作为 [倒排索引](ds-inverted-index.md) 的key, 交集的结果作为值记录在posting list, 查询时直接查询.

    ![img](res/merge-sort-op3.png)

- `缓存加速`

      多个`预先联合`+LRU （least recently used）
      一个简单LRU: 链表head插入缓存元素, tail淘汰元素；为了快速查询, 在向链表head插入元素的同时, 也向哈希插入key, 对应的value是链表中这个节点的地址.

    ![img](res/merge-sort-op4.png)
