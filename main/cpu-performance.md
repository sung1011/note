# cpu 性能

## 指标

### 使用率

- 用户cpu
  - user: 用户态CPU使用率
  - nice: 低优先级用户态CPU使用率
- 系统cpu (sys): CPU在内核态运行的时间百分比（不包含中断）
- 等待io (iowait): 等待io的时间百分比
- 软中断 (softirq): 内核调用软中断的时间百分比
- 硬中断 (hardirq): 内核调用硬中断的时间百分比
- 窃取CPU (steal): 其他虚拟机占用的CPU百分比
- 客户CPU (guest): 运行客户虚拟机的CPU百分比

### 平均负载

- 平均负载 (load average): 平均活跃进程数

### 上下文切换

- 系统调用
- 进程间切换

### 缓存命中率

- cpu缓存 (cache memory)
  - L1: 64k
  - L2: 1M
  - L3: 16M 多核共享

## 工具

- 平均负载 `uptime`, `top`
- 系统整体CPU使用率 `vmstat`, `mpstat`, `top`, `sar`, `/proc/stat`
- 进程CPU使用率 `top`, `pidstat`, `ps`, `htop`, `atop`
- 系统上下文 `vmstat`
- 进程上下文 `pidstat`
- 软中断 `top`, `/proc/softirqs`, `mpstat`
- 硬中断 `vmstat`, `/proc/interrupts`
- 网络 `dstat`, `sar`, `tcpdump`
- io `dstat`, `sar`
- cpu个数 `/proc/cpuinfo`, `lscpu`
- 事件剖析 `perf`, `execsnoop`
