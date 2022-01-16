# 桥接模式

    抽象和实现解耦, 使其可以独立变化
    一个类存在多个独立变化的维度, 通过组合令其可以独立扩展

## 特点

- `抽象和实现分离`
- `高扩展性`

## 实现

- `接口` 接口嵌套

## code

- [go](src/go/dp/bridge.go)

- [php](src/php_design_patterns/bridge/bridge.php)

## 场景

- 不同的通知类型, 有不同的报警级别(即两个维度)
