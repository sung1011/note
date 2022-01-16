# 中介者模式

    用一个中介对象来封装一系列对象交互

## 特点

- 降低类的复杂度

## 实现

- 中介对象由多个组件组合

## issue

- 中介者对象越来越庞大
- 不应在职责混乱时候使用

## code

- [go](src/go/dp/mediator.go)

- [php](src/php_design_patterns/mediator/mediator.php)

## 场景

- 对象之间存在比较复杂的引用关系, 导致依赖关系结构混乱而难以复用
- 想通过一个中间类封装多个类中的行为, 又不想生成太多子类