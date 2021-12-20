# huge page

## 关于HugePage

操作系统默认的内存是以4KB分页的, 而虚拟地址和内存地址需要转换,  而这个转换要查表, CPU为了加速这个查表过程会内建TLB(Translation Lookaside Buffer). 显然, 如果虚拟页越小, 表里的条目数也就越多, 而TLB大小是有限的, 条目数越多TLB的Cache Miss也就会越高,  所以如果我们能启用大内存页就能间接降低这个TLB Cache Miss.

## PHP7与HugePage

PHP7开启HugePage支持后, 会把自身的text段, 以及内存分配中的huge都采用大内存页来保存, 减少TLB miss, 从而提高性能.相关实现可参考Opcache实现中的accel_move_code_to_huge_pages()函数.


## 开启方法

以CentOS 6.5为例, 通过命令: 

sudo sysctl vm.nr_hugepages=128
分配128个预留的大页内存.

$ cat /proc/meminfo | grep Huge 
AnonHugePages:    444416 kB 
HugePages_Total:     128 
HugePages_Free:      128 
HugePages_Rsvd:        0 
HugePages_Surp:        0 
Hugepagesize:       2048 kB
然后在PHP.ini中加入

opcache.huge_code_pages=1

## 关于负载过高, 系统CPU使用占比过高的问题

当我们升级完第一个服务池时, 感觉整个升级过程还是比较顺利, 当灰度Page池, 低峰时一切正常, 但到了流量高峰, 系统CPU占用非常高, 如图: 


系统CPU的使用远超用户程序CPU的使用, 正常情况下, 系统CPU与用户程序CPU占比应该在1/3左右.但我们的实际情况则是, 系统CPU是用户CPU的2~3倍, 很不正常.

对比了一下两个服务池的流量, 发现Page池的流量正常比Home池高不少, 在升级Home池时, 没发现该问题, 主要原因是流量没有达到一定级别, 所以未触发该问题.当单机流量超过一定阈值, 系统CPU的使用会出现一个直线的上升, 此时系统性能会严重下降.

这个问题其实困扰了我们有一段时间, 通过各种搜索资料, 均未发现任何升级PHP7会引起系统CPU过高的线索.但我们发现了另外一个比较重要的线索, 很多软件官方文档里非常明确的提出了可以通过关闭Transparent HugePages(透明大页)来解决系统负载过高的问题.后来我们也尝试对其进行了关闭, 经过几天的观察, 该问题得到解决, 如图: 



## 什么是Transparent HugePages(透明大页)

简单的讲, 对于内存占用较大的程序, 可以通过开启HugePage来提升系统性能.但这里会有个要求, 就是在编写程序时, 代码里需要显示的对HugePage进行支持.

而红帽企业版Linux为了减少程序开发的复杂性, 并对HugePage进行支持, 部署了Transparent HugePages.Transparent HugePages是一个使管理Huge Pages自动化的抽象层, 实现方案为操作系统后台有一个叫做khugepaged的进程, 它会一直扫描所有进程占用的内存, 在可能的情况下会把4kPage交换为Huge Pages.

## 为什么Transparent HugePages(透明大页)对系统的性能会产生影响

在khugepaged进行扫描进程占用内存, 并将4kPage交换为Huge Pages的这个过程中, 对于操作的内存的各种分配活动都需要各种内存锁, 直接影响程序的内存访问性能.并且, 这个过程对于应用是透明的, 在应用层面不可控制,对于专门为4k page优化的程序来说, 可能会造成随机的性能下降现象.

## 怎么关闭Transparent HugePages(透明大页)

(1)查看是否启用透明大页.

[root@venus153 ~]# cat  /sys/kernel/mm/transparent_hugepage/enabled 
[always] madvise never
使用命令查看时, 如果输出结果为[always]表示透明大页启用了, [never]表示透明大页禁用.

(2)关闭透明大页.

echo never > /sys/kernel/mm/transparent_hugepage/enabled 
echo never > /sys/kernel/mm/transparent_hugepage/defrag
(3)启用透明大页.

echo always >  /sys/kernel/mm/transparent_hugepage/enabled 
echo always > /sys/kernel/mm/transparent_hugepage/defrag
(4)设置开机关闭.

修改/etc/rc.local文件, 添加如下行: 

if test -f /sys/kernel/mm/redhat_transparent_hugepage/enabled; then     
     echo never > /sys/kernel/mm/transparent_hugepage/enabled     
    echo never > /sys/kernel/mm/transparent_hugepage/defrag 
fi