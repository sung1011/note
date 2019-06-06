# 僵尸进程 不可中断进程

## 进程状态
R Running/Runnalbe: 正在运行，或在队列中的进程  
D Disk Sleep: 不可中断 Uninterruptible sleep (usually IO)  
- 进程正在跟硬件交互，交互过程不允许其他进程或中断打断
S Interruptible Sleep: 处于可中断休眠状态  
- 因等待某事件而被挂起，事件发生时会被唤醒
Z Zombie: 僵尸进程  
- 进程已经结束，但父进程尚未回收它。
I Idle: 空闲状态 不可中断睡眠睡眠的内核线程。  
- 对于某些内核线程来说，有可能并没有任何负载。主要与D区分。
- D状态进程会导致平均负载（loadAverage）升高，而I状态进程不会。
T 停止或被追踪  
- SIGSTOP信号使进程暂停，SIGCONT信号使进程恢复运行。 如:fg
s 包含子进程  
l 多线程，克隆线程 multi-threaded (using CLONE_THREAD, like NPTL pthreads do)  
< 高优先级  
N 低优先级  
L 有些页被锁进内存  
\+ 位于后台的进程组  
W 进入内存交换（从内核2.6开始无效）  
X 死掉的进程  

## 僵尸进程
