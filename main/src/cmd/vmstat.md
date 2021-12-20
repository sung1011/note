# vmstat

- Procs(进程):
r: 运行队列中进程数量  
b: 等待IO的进程数量  
- Memory(内存):
swpd: 使用虚拟内存大小  
free: 可用内存大小  
buff: 用作缓冲的内存大小  
cache: 用作缓存的内存大小  
- Swap:
si: 每秒从交换区写到内存的大小  
so: 每秒写入交换区的内存大小  
- IO: (现在的Linux版本块的大小为1024bytes)  
bi: 每秒读取的块数  
bo: 每秒写入的块数  
- system: 
in: 每秒中断数, 包括时钟中断  
cs: 每秒上下文切换数  
- CPU(以百分比表示)
us: 用户进程执行时间(user time)  
sy: 系统进程执行时间(system time)  
id: 空闲时间(包括IO等待时间)  
wa: 等待IO时间  
