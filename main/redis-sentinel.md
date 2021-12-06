# redis sentinel

- `监控 Monitoring` 不断检查所有server
- `自动故障迁移 Automatic-failover` 主从切换
- `提醒 Notification` 被监控server故障时, 通过API通知管理员

## 主观下线 & 客观下线

- `主观下线 Subjectivity Down` 单个Sentinel对S判定下线

- `客观下线 Objectively Down` 多个Sentinel对S判定下线, 并且通过一段时间的相互通信, 都得到下线的结论

> master-down-after-milliseconds 该时间内互相通信都是主观下线, 则被判定为客观下线.

## 自动发现 sentinel & slave

    sentinel通过对master进行订阅和发布, 实现自动发现

## API

## cmd

## 故障转移

## sentinel状态的持久化