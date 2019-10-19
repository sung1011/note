# 微服务 服务发现

## 概念

## 组成

![img](res/ms-registry.png)

### 服务端 RPC server

外提供后台服务，将自己的服务信息注册到注册中心

### 客户端 RPC client

注册中心获取远程服务的注册信息，然后进行远程调用

### 注册中心 registry

服务端注册远程服务以及客户端发现服务

#### API

- 注册: server向registry注册
- 反注册: server向registry撤销注册
- 心跳: server向registry发送心跳来完成节点存活状态上报。
- 订阅: client向registry获取可⽤的server节点列表
- 变更: client向registry获取最新可用server节点列表
- 后台管理
  - 查询: 查询registry当前注册了哪些服务
  - 修改: 修改registry某一服务信息

#### 部署

一般采用集群部署保障高可用性。如zookeeper

#### 健康检查

#### 服务状态变更通知

#### 白名单

只有registry白名单中存在的server才能注册。防止测试server节点意外跑到线上。
