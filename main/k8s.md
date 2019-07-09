# kubernetes  

## [API](k8sAPI.md)

## [API yaml](k8sYAML.md)

## 优点

简化应用部署  
高效利用硬件  
健康检查自我修复  
自动扩容  

## 架构  

![img](res/k8sarch.png)

### 控制面板 master  

![img](res/k8smaster.png)

kubernetes api  
controller manager  
scheduler  
etcd  

### 工作节点 node  

![img](res/k8snode.png)

kubelet  
kube-proxy  
容器  

### 核心组件

etcd保存了整个集群的状态；  
apiserver提供了资源操作的唯一入口，并提供认证、授权、访问控制、API注册和发现等机制；  
controller manager负责维护集群的状态，比如故障检测、自动扩展、滚动更新等；  
scheduler负责资源的调度，按照预定的调度策略将Pod调度到相应的机器上；  
kubelet负责维护容器的生命周期，同时也负责Volume（CVI）和网络（CNI）的管理；  
Container runtime负责镜像管理以及Pod和容器的真正运行（CRI）；  
kube-proxy负责为Service提供cluster内部的服务发现和负载均衡；  

### Add-ons组件

kube-dns负责为整个集群提供DNS服务  
Ingress Controller为服务提供外网入口  
Heapster提供资源监控  
Dashboard提供GUI  
Federation提供跨可用区的集群  
Fluentd-elasticsearch提供集群日志采集、存储与查询  

### 分层架构

核心层：Kubernetes最核心的功能，对外提供API构建高层的应用，对内提供插件式应用执行环境  
应用层：部署（无状态应用、有状态应用、批处理任务、集群应用等）和路由（服务发现、DNS解析等）  
管理层：系统度量（如基础设施、容器和网络的度量），自动化（如自动扩展、动态Provision等）以及策略管理（RBAC、Quota、PSP、NetworkPolicy等）  
接口层：kubectl命令行工具、客户端SDK以及集群联邦  
生态系统：在接口层之上的庞大容器集群管理调度的生态系统，可以划分为两个范畴  

- Kubernetes外部：日志、监控、配置管理、CI、CD、Workflow、FaaS、OTS应用、ChatOps等
- Kubernetes内部：CRI、CNI、CVI、镜像仓库、Cloud Provider、集群自身的配置和管理等

![k8slayer](res/k8slayer.jpg)

## Pod

Pod是Kubernetes创建或部署的最小/最简单的基本单位，一个Pod代表集群上正在运行的一个进程。  

### 共享资源  

PID命名空间：Pod中的不同应用程序可以看到其他应用程序的进程ID；  
网络命名空间：Pod中的多个容器能够访问同一个IP和端口范围；  
IPC命名空间：Pod中的多个容器能够使用SystemV IPC或者POSIX消息队列进行通信；  
UTS命名空间：Pod中的多个容器共享一个主机名；  
Volumes（共享存储卷）：Pod中的各个容器可以访问在Pod级别定义的Volumes。  

### 存活探针 LivenessProbe

判断容器是否存活（running）

#### 类型

http  
TCP  
exec  

### 自动调度 Replication Controller (rc)

创建启动rc
(通过修改pod label)rc中移出pod, rc自动新建pod
缩放rc
删除rc(可选是否删除其pod)

### 自动调度 ReplicaSet (rs)

rs是rc的进化，加强了label表达

### 守护调度 DaemonSet

每个node运行一个pod

### 定向调度 NodeSelector

(通过label)指定某个Pod需要指定在某个Node上

### 批处理调度 Job

单次  
并行  
串行  
限时  
定时  

## Service

### 创建service

expose
yaml

- 会话亲和性（spec.sessionAffinity: clintIP）

服务暴露多端口
命名端口

### 服务发现

Env
DNS

- K8s集群会内置一个dns服务器，service创建成功后，会在dns服务器里导入一些记录，想要访问某个服务，通过dns服务器解析出对应的ip和port，从而实现服务访问

### endpoint

endpoint是k8s集群中的一个资源对象，存储在etcd中，用来记录一个service对应的所有pod的访问地址。  

- service配置selector: endpoint controller会自动创建对应的endpoint对象。

- service不配置selector: endpoint controller不会生成endpoint对象. 可手动创建endpoint(name必须与svc name相同)

endpoint controller: 是k8s集群控制器的其中一个组件，其功能如下：  

- 负责生成和维护所有endpoint对象的控制器
- 负责监听service和对应pod的变化
- 监听到service被删除，则删除和该service同名的endpoint对象
- 监听到新的service被创建，则根据新建service信息获取相关pod列表，然后创建对应endpoint对象
- 监听到service被更新，则根据更新后的service信息获取相关pod列表，然后更新对应endpoint对象
- 监听到pod事件，则更新对应的service的endpoint对象，将podIp记录到endpoint中

### 暴露服务 - ClusterIp

默认模式，只能在集群内部访问

### 暴露服务 - NodePort

通过每个 Node 上的 IP 和静态端口（NodePort）暴露服务。NodePort 服务会路由到 ClusterIP 服务，这个 ClusterIP 服务会自动创建。通过请求 <  NodeIP >:< NodePort >，可以从集群的外部访问一个 NodePort 服务。

### 暴露服务 - LoadBalancer

使用云提供商的负载局衡器，可以向外部暴露服务。外部的负载均衡器可以路由到 NodePort 服务和 ClusterIP 服务。

### 暴露服务 - ExternalName

通过返回 CNAME 和它的值，可以将服务映射到 externalName 字段的内容（例如， foo.bar.example.com）。 没有任何类型代理被创建，这只有 Kubernetes 1.7 或更高版本的 kube-dns 才支持。

### Ingress路由到服务

### 就绪探针 ReadinessProbe

### headless发现独立pod

### 排除服务故障

## Volume

## Nodes

## Namespaces

## Labels & Selectors

## Annoatations

## 对象

## 组件

## Names

## Master-Node通信

## Deployment

## StatefulSets 部署有状态的多副本应用

## GC
