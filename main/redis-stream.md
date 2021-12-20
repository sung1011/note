# redis stream

@@redis

    stream主要用于消息队列(MQ,Message Queue),Redis本身是有一个Redis发布订阅(pub/sub)来实现消息队列的功能,但它有个缺点就是消息无法持久化,如果出现网络断开、Redis宕机等,消息就会被丢弃.

## struct

TODO

## feature

- 消息ID的序列化生成
- 消息遍历
- 消息的阻塞和非阻塞读取
- 消息的分组消费
- 未完成消息的处理
- 消息队列监控

## cmd

- `XADD` 添加消息到末尾(生产消息)
- `XTRIM` 对流进行修剪,限制长度
- `XDEL` 删除消息
- `XLEN` 获取流包含的元素数量,及消息长度
- `XRANGE` 获取消息列表,会自动过滤已经删除的消息
- `XREVERANGE` 反向获取消息列表,ID从大到小
- `XREAD` 以阻塞或非阻塞方式获取消息列表(消费消息)

- `XGROUP CREATE` 创建消费者组
- `XREADGROUP GROUP` 读取消费者组中的消息
- `XACK` 将消息标记为“已处理”
- `XGROUP SETID` 为消费者组设置新的最后递送消息ID
- `XGROUP DELCONSUMER` 删除消费者
- `XGROUP DESTROY` 删除消费者组
- `XPENDING` 显示待处理消息的相关信息
- `XCLAIM` 转移消费者组的相关信息
- `XINFO GROUPS` 打印消费者组的信息
- `XINFO STREAM` 打印流信息
