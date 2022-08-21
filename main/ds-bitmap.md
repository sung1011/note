# 位图 bitmap

    数据映射为bit数组, 只存0或1

## 场景

    解决巨量不重复数字/ID/字符时的优化 交并补 问题

> 解决问题时也经常将多位看作一个元素 如两位 01 00 10 11; 三位 101 001...

## 结构

```bash
bitmap1 : 0 0 1 0 0 0 1 0

bitmap2 : 0 0 0 0 1 0 1 0

result  : 0 0 0 0 0 0 1 0

# 1byte = 8bit
```

## 特征

- `空间占用小` 算出存一个大数字(字符)需要多少 bit, 使用位运算仅占用1bit, 即效率提升.  

- `运算效率高` 位运算 O(1)

- `去重 判重`

- `排序`

## 缺点

- `可读性差`

- `冲突` 数字需要有唯一性, 字符串散列到bitmap可能有冲突(不唯一). 解决方案 [Bloom Filter](algo-bloomfilter.md)

- `稀疏` 存少量的大数字却需要 len 很长的bitmap. 解决方案 [Roaring BitMap](algo-roaringbitmap.md)

## 操作

- `find O(1)`

- `delete O(1)`

- `insert` 预分配

## 压缩

- RBM roaring bitmap
- RLE run length encoding
- WAH Word Aligned Hybrid Compression Scheme
- Concise Compressed ‘n’ Composable Integer Set
- gzip
- lzo

## 排序 (去重的)

```bash
    # 原理
      一个数字用int32存储即4byte, 按位看, 可以保存32个数字

        如: {0, 3, 5}
        0000...000101001 # 32个数字, 省略左右的一堆0

      00~31 存在第一个int32中
      32~63 存在第一个int32中...
      如此100个数字, 需要 (100/32) + 1 = 4个int32就可以表达

    # 内存估算
      n * 4 / 1024 / 1024 / 32 # int32=4byte;
      如 1e = 10^8 个数字, 占用12M内存

    # 如果是int64
      n * 8 / 1024 / 1024 / 64
```

## 实战

- URL是否已经被爬虫抓取
- [布隆 bloom filter](algo-bloomfilter.md)
- `某个数字x是否在40e个不重复整数中(数字最大不超40e)`

       40e整数需要160e字节约16G内存(2^32=42.9e 每个数字4字节=32bit 则(40e*4*8)bit /8/1024/1024=16G)
       而每个数字映射bitmap上的1位则需要5e字节约500M内存(每个数字需要1bit 则(40e)bit /8/1024/1024=476.8M)
       进行位运算得到结果

- `多个数字找出重复/不重复的个数`

      1个bitmap每2bit表达一个数字, 00=不存在 01出现一次 10出现多次 11无意义
        遍历拿到结果

      2个bitmap每个bit表达一个数字, 00=不存在 01出现一次 10出现多次 11无意义
        第一个bitmap存储是否出现
        第二个bitmap存储是否再次出现
        两个bitmap做位运算得到结果

- linux inode
- linux 磁盘块
- id是否存在
- 用户标签  
- 用户标签交集 -- 位运算

## ref

[什么是bitmap](https://www.jianshu.com/p/6e2285c85295)