# 适配器模式

    多个类接口不兼容时, 将一个类的接口转换成另一个统一接口

## 特点

- 改变原有接口

> 当需要适配的接口多且需要大量修改时, 适合使用`对象适配器`

## 实现

- `类适配器` 继承
  - a类crawl方法 b类flee方法, 新建adapter类实现a, b的'跑'方法到同一个接口run中

- `对象适配器` 组合

## code

- [go](src/go/dp/adapter.go)

- [php](src/php_design_patterns/adapter/adapter.php)

## 场景

- 阿里云 和 AWS 的SDK提供的'创建主机'接口不同, 适配器模式将其统一

- 不同数据库的创建数据库连接接口不同