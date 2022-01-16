# 工厂模式

    集中创建对象的方法

## 特点

- 封装变化 统一实例化规则
- 职责单一 将实例化抽离出来, 降低复杂度

## 实现

- `简单工厂` 根据参数, 生产不同的产品类

- `工厂方法` 根据参数, 生产不同的工厂类, 工厂实现了生产产品的接口

- `抽象工厂` 不常用

## code

- [go 简单工厂](src/go/dp/factory-simple.go)
- [go 工厂方法](src/go/dp/factory-method.go)
- [go 抽象工厂](src/go/dp/factory-abstract.go)

- [php 简单工厂](src/php_design_patterns/simple_factory/simple_factory.php)
- [php 工厂方法](src/php_design_patterns/factory_method/factory_method.php)
- [php 抽象工厂](src/php_design_patterns/abstract_factory/abstract_factory.php)
