# 微服务 监控

## 对象

- `功能接口监控` c-s 客户端对服务接口的调用  
- `服务调用监控` s-s 业务依赖的服务的RPC调用  
- `依赖服务监控` s-s' 被调用的redis, mysql等服务组件  
- `基础资源监控` >s< CPU,内存,磁盘,带宽等基础资源  

## 指标

- `RT`  
- `请求量`  
- `错误率`  

## 维度

- 全局
- 机房
- 单机
- 时间
- 核心

## 原理

- 采集
- 传输
- 处理
- 展示

## 分类

- `集中式` ELK(Elasticsearch + Logstash + Kibana + Beats)
- `时序式` Graphite、TICK、Prometheus

## 实例

- `ELK`

  - `Logstash` 收集和传输；收集数据、过滤、分析、格式化、聚合、存储.
  - `Elasticsearch` 处理；分布式搜索和分析引擎；可伸缩、高可靠、易管理.
  - `Kibana` 展示；图形化展示.
  - `Beats` 数据收集；轻量.
  
- `Graphite`

  - Carbon 处理
  - Whisper 存储；时序数据库
  - Graphite-Web 展示

- `TICK`

  - Telegraf 收集
  - InfluxDB 存储
  - Chronograf 展示
  - Kapacitor 报警

- `Prometheus`

  - Prometheus Server 存储
  - Jobs/exporters 收集
  - Pushgateway 短期收集
  - Alertmanager 报警
  - web UI 展示
