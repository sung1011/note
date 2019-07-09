# docker  
  
## 隔离  

### 命名空间 namespace: 独立的系统视图(文件 进程 网络接口 主机名等)  

Monut (mnt)  

- 提供磁盘挂载点和文件系统的隔离能力

Process Id (pid)  

- 提供进程隔离能力  

Network (net)  

- 提供网络隔离能力  

Inter-process communication (IPC)  

- 提供进程间通信的隔离能力  

UTS  

- 提供主机名隔离能力  

User Id (user)  

- 提供用户隔离能力  

### 控制组 cgroups: 限制资源量(cpu 内存 网络带宽)  

cpu  
memory  
network  
  
## 组成  

镜像 image: 构建  
镜像仓库 image repository: 分发  
容器 container: 运行  
  
## docker vs VM  

每个虚拟机会独立虚拟化CPU与kernel  
