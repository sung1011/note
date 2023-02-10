# cache

## 使用方式

### Read/Write Through

```js
        流程: 
                            命中cache --------------------------
                           /                                    \
            读请求 -> 读cache -> (未命中cache)读DB -> 更新cache -> 返回

                                   未命中cache ---------------
                                 /                            \
            写请求 -> 写DB -> 读cache -> (命中cache)更新cache -> 返回

        优点: 直观的方案
        缺点: 有概率写入脏cache

# 脏数据: 并发读和写, a读到旧值, b写新值并更新cache为新值, a更新cache为旧值

# 脏数据解决: 版本号, 写锁
```

### Cache Aside

    流程:
        读请求 ->(同Read/Write Through)

        写请求 -> 写DB -> 删除cache -> 返回

    优点: 相比 Read/Write-Through 不易读到脏数据
    缺点: 相比 Read/Write-Through 稍微增大了cache穿透的可能性

### 异步生成

    流程: 所有读请求都读cache, 定期异步从DB中更新cache
    优点: 高性能
    缺点: 低时效

> 启动一个更新订单缓存的服务, 订单变更发消息给MQ, MQ来更新cache.

> 解析binlog来更新缓存; 开源方案 `Canal`

### 版本号

    数据加版本号, 写DB时自增, 更新cache时只允许高版本数据覆盖低版本数据

## FAQ

### 穿透

```js
    Q:
        查询不存在DB的key
    A:
        1. 缓存空对象, 并加过期时间
        2. bloomfilter
```

### 雪崩

```js
    Q:
        缓存挂掉了, 所有请求打到DB
    A:
        1. 高可用 (sentinel, cluster)
        2. 本地内存缓存
```

### 击穿, 热点

```js
    Q:
        缓存设置相同的过期时间, 或重新生成cache时间较久
    A:
        1. 互斥锁, 使之只有一个请求发向DB
        2. 不设置过期, 而是将过期时间置在value中, 异步删除
```

### 缓存预热

```js
    Q:
        server启动后, 用户无cache情况下高并发重建cache
    A:
        server启动时, 就把cache加载好
```
