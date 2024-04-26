# redis pubsub

发布订阅模式 >=2.8.0

## 用法

发布订阅模式是一种消息通信模式, 用于解耦消息的发送者和接收者。在这种模式中, 消息发送者称为发布者（publisher）, 而消息接收者称为订阅者（subscriber）。发布者将消息发送到频道（channel）, 订阅者可以订阅一个或多个频道。当发布者向频道发送消息时, 系统会实时地将消息发送给订阅频道的所有订阅者。

### 发布消息

```bash
redis-cli PUBLISH channel message
```

### 订阅消息

```bash
redis-cli SUBSCRIBE channel
```

## 例子

```bash
# terminal 1
redis-cli SUBSCRIBE channel

# terminal 2
redis-cli PUBLISH channel message
```

## 实现

发布订阅模式的实现原理是通过`redis`的`PUBLISH`和`SUBSCRIBE`命令实现的。`redis`服务器会为每个订阅者维护一个订阅列表, 当有消息发布到频道时, `redis`服务器会遍历订阅列表, 将消息发送给所有订阅者。

## 优点

1. 解耦消息的发送者和接收者。
2. 实时性高, 消息发布后, 订阅者会立即收到消息。

## 缺点

1. 订阅者不能接收历史消息, 只能接收订阅后发布的消息。
2. 订阅者不能选择接收特定消息, 只能接收所有消息。
3. 发布者和订阅者之间没有消息确认机制, 消息为fire-and-forget模式, 消息可能会丢失。

## 适用场景

1. 实时性要求高的消息通信。
2. 消息发送者和接收者解耦。

## 参考

- [Redis发布订阅模式](https://redis.io/topics/pubsub)

## 相关命令

- [PUBLISH](https://redis.io/commands/publish)
- [SUBSCRIBE](https://redis.io/commands/subscribe)