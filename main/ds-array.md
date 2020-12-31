# 数组

## 概念

有限个相同类型元素组成的有序（连续内存空间存储）集合。

## 特征

- 相同类型

- 有限数量 -- 需要预先分配内存空间; 动态扩容

- 连续内存空间 -- CPU缓存命中率高; 可随机访问

- 快 find O(1)

- 慢 insert delete O(n)

## 结构

```bash
head  
len  
cap  
data  
```

## 操作

### 查 O(1)

<details>
  <summary>details</summary>
<pre>
    head + n * size
</pre>
</details>

### 改 O(1)

### 增 O(n)

<details>
  <summary>details</summary>
<pre>
    尾部追加 O(1)
    中间插入 向右移 O(n)
      1. i及其右边的元素(从右向左逐个)向右移
      2. 新元素插入i到位置
    中间替换 被替换的插入尾部 O(1)
    超范围插入 先扩容 O(n)
      1. 创建新数组，cap为原先的2倍
      2. 旧数组元素复制到新数组
      3. 进行插入或追加操作
</pre>
</details>

### 删 O(n)

<details>
  <summary>details</summary>
<pre>
    1. 删除元素i
    2. i右侧元素逐个向左移
</pre>
</details>
