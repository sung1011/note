# 门面模式 (外观模式)

    定义一组高层接口,令子系统更易用
    将底层接口包装成一个新接口

## 特点

- `高易用性`
- `减少请求` 一次请求干多件事儿

## 实现

- 将底层接口包装成一个新接口

## issue

- 不符合开闭原则，如果要改东西很麻烦，继承重写都不合适。

## code

- [go](../script/go/dp/facade.go)

- [php](src/php_design_patterns/facade/facade.php)

## 场景

- 为复杂的模块或子系统提供外界访问的模块