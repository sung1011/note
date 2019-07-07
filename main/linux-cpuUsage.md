# cpu 使用率

## CPU使用率

1 - (空闲时间 / 总CPU时间)

## 指标

user (us): 用户态CPU时间。 它不包含nice时间，但包含guest时间。  
nice (ni): 低优先级用户态CPU时间，进程的nice值被调为1~19之间。值越大，优先级反而低。可取值范围-20~19。  
system (sys): 内核态CPU时间。  
idle (id): 空闲时间。不包含iowait。  
iowait (wa): 等待IO的CPU时间。  
irq (hi): 处理硬中断的CPU时间  
softirq (si): 处理软中断的CPU时间。  
steal (st): 当系统运行在虚拟机中的时候，被其他虚拟机占用的CPU时间。  
guest (guest): 通过虚拟化运行其他操作系统的时间，即运行虚拟机的时间。  
guest_nice (gnice): 低优先级运行虚拟机的时间。  

## 实战

### perf `perf top -g -p < pid >`

### perf record -g; 等待x秒后ctrl+c ; perf report

### /boot/config -> CONFIG_HZ 节拍率:每秒触发中断次数

### pstree 查看父进程

### execsnoop 分析短时进程

> https://github.com/brendangregg/perf-tools

### CPU使用率高，但找不到高CPU的进程

#### 场景

CPU总占用很高，但用top, pidstat等常规工具找不到高cpu应用

#### 思路

进程在不停崩溃重启。
进程是短时进程，如:在其他应用内部调用exec。
