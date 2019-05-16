# kubernetes  

  
## 架构  
  
### 控制面板 master  
kubernetes api  
controller manager  
scheduler  
etcd  
  
### 工作节点 node  
kubelet  
kube-proxy  
容器  

## 好处
简化应用部署
高效利用硬件
健康检查自我修复
自动扩容

## 概念
### Pod
Pod是Kubernetes创建或部署的最小/最简单的基本单位，一个Pod代表集群上正在运行的一个进程。

### Nodes
### Namespaces
### Labels & Selectors
### Annoatations

### 对象
### 组件
### Names
### Volume
### Master-Node通信
### Replication Controller (rc)
### Replica Sets (rs)
### Service
### Deployment
### StatefulSets
### Ingress
### GC