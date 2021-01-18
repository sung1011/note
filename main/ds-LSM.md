# LSM Log Structured Merge

![img](res/ds-lsm.png)

    先写内存，数据写满后，逐层滚动地归并合并写入磁盘。
    内存数据 memtable，磁盘数据(多层) SStable(sorted string table)，内存与磁盘中不可更改的数据 immutable。
    牺牲少量读速度，换取更优的写入速度。

> SStable一共7层（L0~L6），每层比上一层大10倍。

## 特性

- `批量写入` 多次单页随机写 变成 单次多页顺序写，提高写效率。

- `存储延迟` 可异步归并；WAL 保障可用性

## 操作

- `find O (k*log n)`
  
    先查内存，若没查到这个key，则逆序遍历SStable文件。
    SStable的内容有序，所以二分法log n，逐个遍历SStable文件，所以k*log n。(可加bloom filter优化)

    1. check MEM L0
    2. check DIST L1
    3. check DIST L2
    4. check DIST L3
    ...

> 优化: bloom filter快速得到是否在当前层，替代二分查找。
> 读放大: 用户读取1k数据，通过上述流程实际可能读取10k数据，造成一些IO压力。

- `insert O (k*log n)`

    1. 追加写前日志 WAL。
    2. 数据插入L0。
    3. L0占内存较多时，C0与C1归并排序生成new-C1，替换old-C1。 -- 这个过程称为 Compaction，可稍后异步完成。
    4. L1占磁盘较多时，C1与C2归并排序生成new-C2，替换old-C2。
    5. L2占磁盘较多时...

> 写放大: 用户插入1k数据，通过上述流程实际可能插入10k数据，造成一些IO压力。

- `delete`

## ref

> `https://blog.csdn.net/u014774781/article/details/52105708`
> `https://www.jianshu.com/p/5c846e205f5f`
