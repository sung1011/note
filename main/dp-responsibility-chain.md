# 责任链模式

    将请求的发送者和接受者解耦, 使得多个对象依次处理这个请求, 沿着这个链传递这个请求

## 特点

- 复用
- 扩展

## 实现

- `链表`
- `数组`

> 实现了同一接口的实例挂在一个链表上, 顺序执行

## code

- [go](src/go/dp/responsibility-chain.go)

- [php](src/php_design_patterns/responsibility_chain/responsibility_chain.php)

## 场景

- 过滤器
- 中间件
