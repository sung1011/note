# 对象存储 Object-Storage-Service

## 特点

- 无限容量
- 超高读写性能
- 搞可靠性 节点磁盘损坏不会丢失数据
- 高可用性 节点宕机不会影响服务

## 缺点

- 文件必须整存整取, 不能直接修改

## 架构

```go
          Client
            |
            | 1. 请求Key对应的数据
            |
----------------------------
|           网关            |
----------------------------
    |                |
    |2. 获取节点信息   |
    |                | 3. 访问多个node的容器数据,并整合
    |                |
---------   -----------------
|       |   |  node         |       ----------
|       |   |           node+------| 容器      |
| META  |   |               |      | 容器      |
|       |   |     node      |      | 容器      |
|       |   |               |      | ...      |
|       |   |               |       ----------
---------   -----------------

```

- `网关 Gateway` 接收请求, 获取节点信息, 访问节点数据; 服务无状态
- `元数据 META` 数据与Key的映射关系
- `数据节点 Node` 包含多个容器
- `容器` 数据被分片到不同node的不同容器中

## C端概念

```go
----------------
| obj          |       -------------------
|         obj--+------|key, metadata, data|
|              |       -------------------
|     obj      |
|              |
----------------
    bucket
```

- `桶 bucket` 存储对象的容器, 桶中的所有对象处于同一逻辑层级(区别于树结构), 具有以下属性
  - `存储类别`
    - `标准存储` 热点数据, 小文件(<1M);
    - `低频访问` 不频繁访问(12次/年); 成本更低; 拉取时间长
    - `归档存储` 很少访问(1次/年); 成本最低; 拉取时间最长
  - `访问权限`
  - `所属区域`

- `对象 object` 文件与属性信息的集合
  - `Key` 对象名, 桶内唯一
  - `Metadata` 元数据
    - 系统 Creation-time, Content-length, Last-modify, Content-Encoding ...
    - 自定义
  - `Data` 文件的(部分)数据

> `桶` 是扁平的, 没有树结构, 但一般会用 / 分隔符模拟树结构的文件夹进行管理.如: {bucket}/v2/foo/abc.png

## 使用场景

- `标准存储` 移动应用, 大型网站, 图片分享, 热点视音频
- `低频访问` 网盘应用, 监控数据, 企业数据, 移动设备
- `归档存储` 档案数据, 医疗影像, 视频素材