# cache

## 使用方式

### Read/Write Through

        流程: 
                            命中cache --------------
                           /                        \
            读请求 -> 读cache -> (未命中cache)读DB -> 更新cache -> 返回

                                   未命中cache ---------------
                                 /                            \
            写请求 -> 写DB -> 读cache -> (命中cache)更新cache -> 返回

        优点: 直观的方案
        缺点: 有概率写入脏cache

> 脏数据出现流程: 并发读和写, a读到旧值, b写新值并更新cache为新值, a更新cache为旧值
>
> 脏数据解决方案: 版本号, 写锁

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

### 版本号

        流程: 数据加版本号, 写DB时自增, 更新cache时只允许高版本数据覆盖低版本数据

## 常见问题

### 穿透

查询不存在的key

### 雪崩

缓存设置相同的过期时间, 或重新生成cache时间较久

### 击穿, 热点

某key访问量过大, 一旦过期
