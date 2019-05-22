# 平衡负载
处于可运行状态(running/runnable) 和 不可中断状态(uninterruptible sleep/disk sleep)的进程数
即包含`正在使用cpu`的进程， `等待cpu计算`的进程， `等待I/O`的进程

## 平衡负载 & CPU使用率
| status | 平衡负载 | CPU使用率 |
| --- | --- | --- |
|CPU密集 | uuu | uuu |
|I/O密集 | uuu | _ |
|大量CPU等待调度| uuu | u |

## 调试
uptime  
iostat  
mpstat  查看每个cpu性能指标，平均指标
pidstat  进程性能分析工具。查看进程cpu,mem,io,ctx


