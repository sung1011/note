# linux进程

        进程是具有一定独立功能的程序关于某个数据集合上的一次运行活动,进程是系统进行资源分配和调度的一个独立单位.每个进程都有自己的独立内存空间, 不同进程通过进程间通信来通信.由于进程比较重量, 占据独立的内存, 所以上下文进程间的切换开销(栈、寄存器、虚拟内存、文件句柄等)比较大, 但相对比较稳定安全.  

## 内容: 进程中拥有的数据  

        那个程序的可执行机器代码的一个在存储器的映像.  
        分配到的存储器(通常是虚拟的一个存储器区域).
        存储器的内容包括可执行代码、特定于进程的数据(输入、输出)、调用堆栈、堆栈(用于保存运行时运输中途产生的数据).  
        分配给该进程的资源的操作系统描述符, 诸如文件描述符(Unix术语)或文件句柄(Windows)、数据源和数据终端.  
        安全特性, 诸如进程拥有者和进程的权限集(可以容许的操作).  
        处理器状态(内文), 诸如寄存器内容、物理存储器定址等.当进程正在运行时, 状态通常存储在寄存器, 其他情况在存储器.  

## 状态: 进程当前的动作  

- `新生 new` 进程新产生中.  
- `运行 running` 正在运行.  
- `等待 waiting` 等待某事发生, 例如等待用户输入完成.亦称“阻塞”(blocked)  
- `就绪 ready` 排班中, 等待CPU.  
- `结束 terminated` 完成运行.  

## 调度: 将任务分配至资源的过程  

- 依序循环调度(RR, Round-robin scheduling)  
- 最短先做排班(SJN, Shortest job next)  
- Shortest remaining time  
- Weighted round-robin scheduling  
- 速率单调 (RMS)  
- Deadline-monotonic scheduling (DMS)  
- Earliest deadline first scheduling (EDF)  
- Two-level scheduling  
- 先进先出  
- LIFO  
- Fair-share scheduling  
- 完全公平调度 (CFS)  
- Least slack time scheduling (LST)  
- Multilevel Feedback Queue  
- Take scheduling  
- Gang scheduling  
- Least-connection scheduling  
- Weighted least-connection scheduling  
- Shortest expected delay scheduling  
- Never queue scheduling  
- List scheduling  
- Genetic Anticipatory  
- Lottery Scheduling  
- 关键路径  
  
## 进程间通信 IPC Inter-Process Communication  
  
### 目的  

- `数据传输` 一个进程需要将它的数据发送给另一个进程, 发送的数据量在一个字节到几M字节之间  
- `共享数据` 多个进程想要操作共享数据, 一个进程对共享数据  
- `通知事件` 一个进程需要向另一个或一组进程发送消息, 通知它(它们)发生了某种事件(如进程终止时要通知父进程).  
- `资源共享` 多个进程之间共享同样的资源.为了作到这一点, 需要内核提供锁和同步机制.  
- `进程控制` 有些进程希望完全控制另一个进程的执行(如Debug进程), 此时控制进程希望能够拦截另一个进程的所有陷入和异常, 并能够及时知道它的状态改变.  

### 途径  

- `信号 singal` 无需知道进程状态; 不能传递复杂消息, 只能用来同步  
- `消息队列 Message queue` 只能传递无格式的字节流, 缓冲区大小受限
- `匿名管道 pipe / 流管道 s_pipe / 文件` 单向的, 速度慢, 容量有限, 只有父子进程能通讯  
- `命名管道` 任何进程间都能通讯, 但速度慢
- `共享内存 shared-memory` 能够很容易控制容量, 速度最快, 但需要保证同步通信.
- `信号量 Semaphore` 是一个原子计数器, 用来做互斥锁控制进程的并发访问.
- [套接字 socket](linux-socket.md) 可用于机器间通讯

## 孤儿进程

    父进程意外退出, 联系不上孩子
    一个父进程退出，而它的一个或多个子进程还在运行，那么那些子进程将成为孤儿进程。

## 僵尸进程

    子进程意外退出, 连不上父亲
    子进程退出时，父进程并未对其发出的SIGCHLD信号进行适当处理，导致子进程停留在僵死状态等待其父进程为其收尸，这个状态下的子进程就是僵死进程。


  