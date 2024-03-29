# 微服务 容器化

    开发人员负责开发、测试、上线全生命周期, 减轻测试和运维负担.

## 镜像仓库

    docker镜像存储和拉取

- `权限控制`
  - 登录访问
  - 仓库权限
- `镜像同步` 大量节点拉取带宽受限
  - 主从
  - P2P
- `高可用性`
  - 多IDC

## 集群资源调度

    docker镜像分发到哪些机器, 机器从哪里来.

- `物理机集群`
- `虚拟机集群`
- `公有云集群`

## 调用策略

- 主机过滤
  - 存活过滤
  - 硬件过滤

- 调度策略
  - 方案
    - spread 为了主机负载平均, 选择资源消耗最少的主机
    - binpack 为了节省资源, 选择资源消耗最多的主机
  - 场景
    - 各主机配置相同 可随机选择
    - 不同资源消耗混合 如: cpu密集与磁盘适合混合
    - 划定资源限制 如: cpu密集与io密集都会抢占CPU, 不适合混合

## 服务编排

- 服务依赖
- 服务发现
- 自动扩缩容
