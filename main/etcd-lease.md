# etcd lease

lease是etcd中的一个重要概念, 它用来维持一个key在etcd中的存在时间。lease的一个重要作用是和租约相关的API中, 比如`Put`, `KeepAlive`等。

## 为什么需要lease

etcd中的key默认是永久存在的, 但是在实际的业务场景中, 有时候我们需要对key设置一个过期时间, 比如锁的自动释放。这时候lease就派上用场了。

## lease的操作

### 创建lease

```shell
etcdctl lease grant 5
```

### 关联key和lease

```shell
etcdctl put --lease=LEASE_ID key value
```

### 续约lease

```shell
etcdctl lease keep-alive LEASE_ID
```

### 释放lease

```shell
etcdctl lease revoke LEASE_ID
```

## lease的应用

### 自动续约

我们可以通过`KeepAlive`接口来维持一个key的存活时间, keepAlive是基于gRPC的一种心跳机制, 客户端会定期向服务器发送ping消息。如果服务器在一定时间内没有响应, 客户端会认为连接已经断开, 并尝试重新连接。

### 锁

通过lease, 我们可以实现一个简单的分布式锁, 通过`Put`接口的`lease`参数, 我们可以设置key的过期时间, 这样就可以实现一个简单的锁。

### 事务

在事务中, 我们可以通过`If`接口来判断一个key是否存在, 如果key不存在, 我们可以通过`Then`接口来创建一个key, 并且设置一个lease, 这样就可以实现一个简单的事务。

### 队列

通过lease, 我们可以实现一个简单的队列, 通过`Put`接口的`lease`参数, 我们可以设置key的过期时间, 这样就可以实现一个简单的队列。

## 总结

lease是etcd中一个非常重要的概念, 它可以帮助我们实现一些复杂的功能, 比如分布式锁, 事务, 队列等。