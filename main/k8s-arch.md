# k8s arch

## arch

### client 管理员客户端

- `kubectl` 命令行工具
- `dashboard` web管理界面
- `sdks` 客户端SDK

### Master-node 管理节点  

![img](res/k8smaster.png)

- `etcd` 保存了整个集群的状态, 只与apiserver通信; 分布式数据库
- `apiserver` 提供了资源操作的唯一入口, 并提供认证、授权、访问控制、API注册和发现等机制; 作为etcd的代理, 无状态, 所有数据都保存在etcd中; 通过订阅通知的方式与其他组件保持通信
- `scheduler` 负责调度pod到可用的node上, 确保pod副本数符合配置
- `controller-manager` 负责维护集群的状态, 监控异常并调节恢复, 自动扩展, 滚动更新等

### Worker-node 工作节点

![img](res/k8snode.png)

- `kubelet` 订阅apiserver的状态数据, 汇报节点数据给apiserver; node上的Pod创建、启停等的管理工作
- `container-runtime` 镜像和容器的运行时, 在kubelet指挥下执行创建, 启停等实际工作; 可以是docker, 也可以是其他实现了CRI接口的容器技术
- `kube-proxy` node的网络通讯代理, 负责为node提供服务发现和负载均衡; 网络通讯的代理转发

### Add-ons组件

- `Dashboard` 提供GUI
- `kube-DNS` 负责为整个集群提供DNS服务
- `Ingress-Controller` 为服务提供外网入口
- `Heapster` 提供资源监控
- `Federation` 提供跨可用区的集群
- `Fluentd-elasticsearch` 提供集群日志采集、存储与查询

### 分层架构

核心层: Kubernetes最核心的功能, 对外提供API构建高层的应用, 对内提供插件式应用执行环境  
应用层: 部署(无状态应用、有状态应用、批处理任务、集群应用等)和路由(服务发现、DNS解析等)  
管理层: 系统度量(如基础设施、容器和网络的度量), 自动化(如自动扩展、动态Provision等)以及策略管理(RBAC、Quota、PSP、NetworkPolicy等)  
接口层: kubectl命令行工具、客户端SDK以及集群联邦  
生态系统: 在接口层之上的庞大容器集群管理调度的生态系统, 可以划分为两个范畴  

- k8s外部: 日志、监控、配置管理、CI、CD、Workflow、FaaS、OTS应用、ChatOps等
- k8s内部: CRI、CNI、CVI、镜像仓库、Cloud Provider、集群自身的配置和管理等

![k8slayer](res/k8slayer.jpg)