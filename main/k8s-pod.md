# k8s pod

    Pod是Kubernetes创建或部署的最小/最简单的基本单位, 表达一个应用实例
    每个container是一个进程, 而Pod是一组容器的集合, 它们共享存储/网络, 并在一个节点上运行.

> 如 nginx+php+mysql, golang+mongodb

## 共享资源  

    PID命名空间: Pod中的不同应用程序可以看到其他应用程序的进程ID
    网络命名空间: Pod中的多个容器能够访问同一个IP和端口范围
    IPC命名空间: Pod中的多个容器能够使用SystemV IPC或者POSIX消息队列进行通信
    UTS命名空间: Pod中的多个容器共享一个主机名
    Volumes(共享存储卷): Pod中的各个容器可以访问在Pod级别定义的Volumes.  

## 存活探针 LivenessProbe

    判断容器是否存活(running)

### 类型

    http  
    TCP  
    exec  

## 自动调度 Replication Controller (rc)

    创建启动rc
    (通过修改pod label)rc中移出pod, rc自动新建pod
    缩放rc
    删除rc(可选是否删除其pod)

## 自动调度 ReplicaSet (rs)

    rs是rc的进化, 加强了label表达

## 守护调度 DaemonSet

    每个node运行一个pod

## 定向调度 NodeSelector

    (通过label)指定某个Pod需要指定在某个Node上

## 批处理调度 Job

单次执行, 执行后退出

- 并行
- 串行
- 限时  
- 定时  

