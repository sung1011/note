# dstat

## 概况

dstat 是一个可以取代vmstat, iostat, netstat和ifstat这些命令的多功能产品.dstat可以让你实时地看到所有系统资源.

## 特性

结合了vmstat, iostat, ifstat, netstat以及更多的信息
实时显示统计情况
在分析和排障时可以通过启用监控项并排序
模块化设计
使用python编写的, 更方便扩展现有的工作任务
容易扩展和添加你的计数器(请为此做出贡献)
包含的许多扩展插件充分说明了增加新的监控项目是很方便的
可以分组统计块设备/网络设备, 并给出总数
可以显示每台设备的当前状态
极准确的时间精度, 即便是系统负荷较高也不会延迟显示
显示准确地单位和和限制转换误差范围
用不同的颜色显示不同的单位
显示中间结果延时小于1秒
支持输出CSV格式报表, 并能导入到Gnumeric和Excel以生成图形

## 输出

1. --total-cpu-usage---- CPU使用率  
usr: 用户空间的程序所占百分比;   
sys: 系统空间程序所占百分比;   
idel: 空闲百分比;   
wai: 等待磁盘I/O所消耗的百分比;   
hiq: 硬中断次数;   
siq: 软中断次数;   

2. -dsk/total-磁盘统计  
read: 读总数  
writ: 写总数  

3. -net/total- 网络统计  
recv: 网络收包总数  
send: 网络发包总数  

4. ---paging-- 内存分页统计  
in:  pagein(换入)  
out: page out(换出)  

5. --system--系统信息  
int: 中断次数  
csw: 上下文切换  

## 常用参数

-l : 显示负载统计量  
-m : 显示内存使用率(包括used, buffer, cache, free值)  
-r : 显示I/O统计  
-s : 显示交换分区使用情况  
-t : 将当前时间显示在第一行  
–fs : 显示文件系统统计数据(包括文件总数量和inodes值)  
–nocolor : 不显示颜色(有时候有用)  
–socket : 显示网络统计数据  
–tcp : 显示常用的TCP统计  
–udp : 显示监听的UDP接口及其当前用量的一些动态数据  

## 实战

`dstat -g -l -m -s --top-mem` 查看全部内存都有谁在占用  
`dstat -c -y -l --proc-count --top-cpu` cpu资源损耗数据  
`dstat –output /tmp/sampleoutput.csv -cdn` 输出csv  