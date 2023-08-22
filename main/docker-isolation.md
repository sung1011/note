# docker 隔离

## 命名空间 namespace

    系统全局资源 与 容器内部资源隔离
    
    如: 文件, 进程, 网络接口, 主机名等

- `Monut (mnt)` 磁盘挂载点和文件系统的隔离

- `Process Id (pid)` 进程的隔离

- `Network (net)` 网络的隔离
  - bridge(default) 为每一个容器分配、设置IP等, 并将容器连接到一个docker0虚拟网桥, 通过docker0网桥以及Iptables nat表配置与宿主机通信.
    - 优点 独立的网卡, 独立的IP, 独立的端口, 不易冲突
    - 缺点 速度不如host模式
  - host 使用宿主机的IP和端口. 容器将不会虚拟出自己的网卡, 配置自己的IP等
    - 优点 速度比bridge模式快
    - 缺点 容器与宿主机共享IP, 端口范围, 容易冲突
  - none 无网络
    - 由第三方插件提供网络
  - container 创建的容器不会创建自己的网卡, 配置自己的IP, 而是和一个指定的容器共享IP、端口范围.
  
- `Inter-process communication (IPC)` 进程间通信的隔离
  
- `UTS` 主机名的隔离
  
- `User Id (user)` 用户的隔离

## 控制组 cgroups

    限制资源量

    如: cpu, 内存, 网络带宽

- cpu  
- memory  
- network  

## chroot

    更改进程的根目录, 限制文件系统的访问范围

## 对比

### docker vs VM  

    每个虚拟机会独立虚拟化cpu, kernel
    每个docker会共享宿主机的cpu, kernel