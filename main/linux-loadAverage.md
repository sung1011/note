# 平均负载 load average

平均活跃进程数

- 处于可运行状态(running/runnable)的平均进程数, 即: `正被cpu计算`和`等待cpu计算`的进程.
- 不可中断状态(uninterruptible sleep/disk sleep)的平均进程数, 即: `等待I/O`的进程

## 取值

值大于cpu个数时超载,  超载不应超过100%.(如: 2核 平均负载4)

## 平均负载 & CPU使用率

| status          | 平均负载 | CPU使用率 | 指标                                                                  |
| --------------- | -------- | --------- | --------------------------------------------------------------------- |
| CPU密集         | uuu      | uuu       | `load average high` && `%usr 100%` && `%iowait 0%`                    |
| 大量CPU等待调度 | uuu      | uu        | `load average high` && `%usr 100%` && `%iowait 80%` && `running high` |
| I/O密集         | uuu      | u         | `load average high` && `%usr 40%` && `%iowait 80%`                    |

## 实战

### uptime  

### sysstat套件

### iostat  

### mpstat: 查看每个cpu性能指标, 平均指标 `mpstat -P ALL 5`  

### [pidstat](src/cmd/pidstat.md): 进程性能分析工具.查看进程cpu,mem,io,ctx switch `pidstat -u 5`  

### cpu核数

- 物理cpu数 (physical id): 主板上CPU数量 `grep 'physical id' /proc/cpuinfo|sort|uniq|wc -l`
- cpu核数 (cpu cores): 单块物理CPU芯片组的数量 `grep 'cpu cores' /proc/cpuinfo|uniq|awk -F ':' '{print $2}'`
- 逻辑cpu数(processer 0-n): 物理CPU数×每颗核数 `cat /proc/cpuinfo| grep "processor"|wc -l`

### stress

--cpu cpu压测选项, 
-i io压测选项, 
-c 进程数压测选项, 
--timeout 执行时间
> 安装 `yum install -y epel-release; yum install -y stress`
