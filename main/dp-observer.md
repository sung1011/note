# 观察者模式 (发布订阅模式)

    在对象之间定义一对多的依赖, 当对象状态改变, 所有依赖对象都收到通知.

## 特点

- 被观察者持有了集合存放观察者 (收通知的为观察者)

## 实现

- `同步阻塞`
- `异步非阻塞` 协程
- `同进程`
- `跨进程` 消息队列

> 订阅者a, b, c都实现了recv方法, 被观察者main统一调用

## code

- [go](src/go/dp/observer.go)

- [php](src/php_design_patterns/observer/observer.php)

## 场景

- 报纸订阅, 报社为被观察者, 订阅的人为观察者; 新报纸印出来就通知所有订阅者
- MVC 模式, 当 model 改变时, View 视图会自动改变, model 为被观察者, View 为观察者