# redis - twemproxy

![img](res/redis-twemproxy.png)

## 模式

    代理

## 优点  

    支持自动分区, 如果其代理的某节点不可用, 自动排除该节点 (数据丢失, 并将改变keys-instance映射关系, 所以只能用于充当缓存的场景).  
    支持redis, memcache; 支持pipelinling;
    支持多个哈希算法包括DHT
    hash_tag

## 缺点  

    redis扩容非常麻烦
    twitter内部已放弃使用该方案, 新使用的架构未开源

> 扩容可以采用预分片, 停服务后迁移RDB以扩容

## 架构

```js
# 读写分离
                      ----> m ----> s
write ----> twemproxy ----> m ----> s
                      ----> m ----> s

           twemproxy ----> s
read ---->
                     ----> s
           twemproxy ----> s
```

## ref

- 无感知扩容 <https://www.cnblogs.com/dodng/p/7744330.html>
- 特点介绍 <https://developer.aliyun.com/article/763439>
- 读写分离 <https://www.csdn.net/tags/NtjacgwsNjU5MS1ibG9n.html>