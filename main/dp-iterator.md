# 迭代器模式

    一种遍历访问容器对象中各个元素的方法, 不暴露该对象内部结构

## 特点

- 封装 不暴露内部结构, 直接使用迭代器进行迭代
- 单一职责 将对象的遍历操作从对象的类拆分出来, 放到迭代器类中
- 开闭原则 可以添加新的迭代算法

## 实现

- `容器` 通过iterator()方法创建迭代器
- `容器迭代器`
  - 需要`hasNext()`, `currentItem()`, `next()`三个基本方法
  - 待遍历的容器对象通过依赖注入传递到迭代器中

## code

- [go](src/go/dp/iterator.go)

- [php](src/php_design_patterns/iterator/iterator.php)

## 场景