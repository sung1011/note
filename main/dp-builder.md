# 建造者模式

    当一个类比较复杂(多个属性), 可以用不同的builder实例化出不同的主体对象

## 特点

- builder独立, 易于扩展

## 实现

- 主体对象 + builder对象
  - 不同的builder, 构建出不同的主体对象

## code

- [go 建造者 传统](src/go/dp/builder-simple.go)
- [go 建造者 改进](src/go/dp/builder-opt.go)

- [php 建造者](src/php_design_patterns/builder/builder.php)

## 场景

- 造车都需要组装轮子 座椅 车门 方向盘 发动机..., builder不同 区分了不同的车型
