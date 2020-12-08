# todo

- stdout stdin stderr 重定向

## proj

- 服务器架构
  - 接入层
    - 对外协议
    - 内部协议
    - 版本校验
    - 聊天与广播
    - session
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
  - 成就（历史/短期/活动行为）
  - 排行
  - 活动
  - 过期数据清除
  - 队列排队
  - 物品统一增减
  - 领奖邮箱
  - VIP系统
  - 随机
  - 匹配
  - 通知
- 配置
  - 导出
    - 指定表
    - 差异（新增新改）
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
    - 独立审核
  - 切维护
    - 单服
    - 全服
  - 白名单（开发者账号）
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

## issues

- 某DB实例压力大
- 某玩家请求多
- trace一次请求的所有log(不跨服务、跨服务)
- 慢请求
- pay记录过多如何处理
- 支付校验都有啥
- DNS nameserver
- 如何管道标准错误（而非标准输出） `http://www.dovov.com/854.html`
- git ref `https://www.php.cn/manual/view/35099.html`
- why the design `https://draveness.me/whys-the-design/`
- ssh `http://www.ruanyifeng.com/blog/2011/12/ssh_port_forwarding.html`

## code

- mq  
- CURL
- gRPC  
- protobuff  
- etcd  
- nginx源码  
- redis源码  

## geektime

- 05 从零开始学架构  
- 07 微服务核心架构20讲  
- 14 深入浅出gRPC-李林峰  
- 20 如何设计一个秒杀系统  
- 29 趣谈网络协议音频修复版  
- 31 从零开始学微服务  
- 32 深入剖析Kubernetes  
- 33 算法面试通关40讲  
- 35 Go语言核心36讲  
- 44 MySQL实战45讲  
- 46 数据结构与算法之美  
- 50 Linux性能优化实战  
- 53 程序员的数学基础课  
- 54 Nginx核心知识100讲  
- 57 10x程序员工作法 (已完结)  
- 76 趣谈Linux操作系统  
- 79 web协议详解与抓包实战  
- 80 深入浅出计算机组成原理  

## tool

- tk  
- 批量文件改名 前缀 后缀 匹配  
- dotfile  
- ssh 秘钥管理 IP列表管理  

## other

```go
// matcher为副本
matcher, exists := matchers[feed.Type]

// v为副本
for _, v := range feeds {

}

// chan传递副本

// 不在函数中做错误处理, 而是用返回err代替throw
func foo() (string, error) {
  v, err := bar();
  if err != nil {
    return "", err
  }
  return v, nil;
}

// 如果一个接口类型只包含一个方法，那类型名需要以er结尾 ex: Reader, Writer, Matcher...
type Matcher interface {
  Search(feed *Feed, searchTerm string) ([]*Result, error)
}

```
