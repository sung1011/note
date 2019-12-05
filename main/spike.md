# 秒杀

## 设计原则

1. 数据少 --- 减少转码时间、压缩时间、传输时间，
2. 请求少 --- 减少多个请求的消耗
3. 环节少 --- 减少多环节不确定性，减少序列化时间、减少网络传输时间
4. 依赖少 --- 减少不必要的依赖模块（商品信息，用户信息，优惠信息，订单信息，支付...）

## 动静分离

### 静态化处理

- 客户端cache
- CDN
- 服务器cache

### 分离方式

- URL唯一化
- 登陆身份信息动态请求

## 热点数据

- 优化 缓存
- 限制 根据请求做hash分桶，每个分桶设置请求队列，限制热点数据占用过多资源。
- 隔离
  - 业务 热点产品单独报名
  - 系统 请求落到不同的集群
  - 数据 单独的cache或者db

## 流量削峰

### 排队

- 消息队列
- 线程池加锁等待
- 内存排队
- 请求序列化到文件，顺序的读取文件

### 答题/验证码

### 过滤

- 缓存请求
- 过滤过期失效的请求
- 写请求做限流
- 写数据强校验

## 减库存

### 特征

短时间，大流量，商品少，付款率高

### 方式

#### 下单时减库存

问题 竞争对手挤兑

#### 付款时减库存

问题 库存超卖

#### 预扣库存，延时恢复

问题 竞争对手重复挤兑

### 业务解决

1. 经常下单不购买的用户标记flag
2. 每人最多买3件限制
3. 同商品重复下单不购买操作次数进行限制

### 技术解决

1. 缓存数据库（redis）
2. 数据库事务（mysql）
3. 乐观锁 set foo=bar where num > 1
4. 单机（应用层面）队列限制
5. 全局（数据库层面）队列限制
6. 隔离

## 高可用

### 架构阶段

- 异地容灾
- 异步化
- 分组隔离
- 避免单点

### 编码阶段

- 限流
- 超时
- 异步
- 异常

### 测试阶段

- beta
- 自动化对比

### 发布阶段

- 分批发布
- 多版本发布

### 运行阶段

- 数据对账
- 自动降级
- 过载保护
- 实时监控

### 故障发生

- 故障定位
- 快速恢复