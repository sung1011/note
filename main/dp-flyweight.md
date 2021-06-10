# 享元模式

    创建和复用只读的对象, 即共享元数据

## 特点

- `节省内存` 复用
- `对象级别的复用` 复用(只读)对象和对象的字段
- `创建对象` 可以创建多个对象

## 实现

- `Map工厂` 在工厂中, 通过一个Map来缓存已经创建过的对象

## code

- [go](../script/go/dp/flyweight.go)

- [php](src/php_design_patterns/flyweight/flyweight.php)

## 场景

- 复用只读对象