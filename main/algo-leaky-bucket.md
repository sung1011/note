# 漏桶 leaky-bucket

    突发流量会进入到一个漏桶，漏桶会按照我们定义的速率依次处理请求
    如果水流过大也就是突发流量过大就会直接溢出，则多余的请求会被拒绝
    所以漏桶算法能控制数据的传输速率。

## 结构

```js
============        
           ||       # token恒定速率生成
           T
           T
    |------T---------|
    |                |    # token超过桶容量会停止生成
    |T     T   T   T |  
    |  T  T    T T   |
    --------T---------
---request--->          # 消耗token, 得到处理请求的资格
```