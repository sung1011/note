# 经验

## proj

- 服务器架构
  - 接入层
    - 对外协议
    - 内部协议
    - 版本校验
    - 聊天与广播
    - session(jwt)
    - 加密
    - 压缩
    - pb
  - 逻辑层
  - 缓存层
  - 持久化层
    - 一致性
    - 高可用
    - 优化
    - 扩容缩容
  - 网络链路
    - 全球同服
    - 分区同服
- 代码框架
  - 定时任务
    - cron
    - 秒级
  - 充值
    - 订单
    - 验证
    - 回调发奖
  - 运行模式
    - cli
    - api
  - 路由
  - db
  - cache
  - hook
  - log
    - 收集
    - 分级
    - 查询
    - 归档
    - 清除
  - 监控报告
    - 较大数据
    - 异常数值
    - 短期报错
    - 近期活动
    - 在线登录
    - 注册新增
    - 留存
  - 打点统计
    - 消耗
    - 充值
  - 异常处理
  - 单元测试
  - 性能优化
  - 压测
  - 优雅重启
- 功能设计
  - 成就(历史/短期/活动行为)
  - 排行
  - 活动
  - 过期数据清除
  - 队列排队
  - 物品统一增减
  - 领奖邮箱
  - VIP系统
  - 伪随机
  - 匹配
  - 通知
- 配置
  - 导出 (如: excel导出json, 启动server时载入内存)
    - 指定单表
    - 历史差异(新增新改)
  - 格式
    - string
    - int
    - array
  - 分表
  - 合并行
- 开发者工具
  - 临时修改配置
    - 调服务器时间
    - 调配置数值
  - 查改玩家数据
  - 查删缓存
  - 报错信息
  - API文档
  - 导入导出备份数据
- 多语言翻译系统
  - 权限
  - 工作流
- 发布系统
  - git流
    - 分支管理
    - tag
    - 回滚
    - 合并与向上合并
  - 版本
    - 各模块版本
    - 差异资源
  - 资源上传
    - 差异上传
    - rebase
    - 定期删除
  - 发布
    - 灰度
      - 按地理
      - 按逻辑服
      - 按渠道
      - 按uid
      - 预热cache防止穿透
    - 独立审核
  - 切维护
    - 单服
    - 全服
  - 白名单(开发者账号)
    - uid
    - 公司IP
  - 黑名单
    - IP
    - uid
  - hotfix
- 运营
  - 在线
  - 收入
  - 活动
  - 公告
  - 发补偿
    - 范围
      - 单人
      - 全服
      - uid
    - 权限
    - 操作记录
- 文档

## crontab

- 集群分布式锁竞争压力是否过大
- 扩容缩容等机器变更时, 是否影响crontab执行
- 定时任务测试脚本(检测缩容后是否正确地全区执行)
- 输出必须重定向

## code

- 关联数据.如:  英雄上标记穿着啥装备, 装备上标记着是否被穿.
- 各种运营log, 需要有独立的命名, 以便区分(不要都叫snap, datalog).
- 充分考虑跨服合服需求.
- 刚上线的一段时间, 注意db实例放到不同机器.(replica分片(没用mongos)机制)
- 关注数据库连接数与数据库连接池的数量匹配.(如: 每台物理机的php-fpm worker太多, 连接数也会很多)
- db数据提前考虑数据过期(如活动数据)
- 分包处理比batch处理更对log友好
- 分级记录bug list
- 透明大页Huge page
- 排行榜的vip等级(或充值数)纠察高排名的非R作弊嫌疑

## 上线前准备

- 压力测试
- 是否需要新增机器
- 是否需要机器特殊配置
- 是否准备好运维脚本
- 是否需要做初始化工作
- 环境部署后检查

## QA

- 如何切维护踢人
- 登录其他设备如何踢人
- 如何统计在线时间

## issues

- 海外登录异常, 如何复现和定位
