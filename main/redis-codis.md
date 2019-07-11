# redis - codis

![img](res/redis-codis.png)

## 概述

    Codis是一个代理中间件，通过内存保存着槽位和实例节点之间的映射关系,槽位间的信息同步交给ZooKeeper来管理。
    不支持事务和官方的某些命令(原因就是分布多个的Redis实例没有回滚机制和WAL,所以是不支持的)。

## 优点

    开发简单，对应用几乎透明
    Go在多核cpu性能比Twemproxy好
    有图形化界面，扩容容易(仅影响正在迁移key的写操作)，运维方便

## 缺点

    代理依旧影响性能
    组件过多，需要很多机器资源
    修改了redis代码，导致和官方无法同步，新特性跟进缓慢

## 组成

### ZooKeeper

    存放路由表和代理节点元数据
    分发Codis-Config的命令

### Codis-Config

    集成管理工具，有web界面

### Codis-Proxy

    无状态代理，兼容Redis协议
    对业务透明

### Codis-Redis

    基于2.8版本，二次开发
    加入slot支持和迁移命令

## ref

[深入浅出百亿请求高可用Redis(codis)分布式集群揭秘](https://www.jianshu.com/p/6d56e5a229f6)
