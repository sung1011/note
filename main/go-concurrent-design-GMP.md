# go GMP

## 组成

```bash
    G: Goroutine 协程

        go关键字加函数调用的代码就是创建了一个G对象, 是对一个要并发执行的任务的封装, 也可以称作用户态线程.
        属于用户级资源, 对OS透明, 具备轻量级, 可以大量创建, 上下文切换成本低等特点.  

    P: Processor 逻辑处理器
      
        逻辑处理器, 主要作用是管理G对象(每个P都有一个G队列), 并为G在M上的运行提供本地化资源.  
        最多个数GOMAXPROCS可配置

        # 配置: runtime.GOMAXPROCS(); 默认为CPU核数

    M: Machine (Thread) 内核线程

        M的作用就是执行G中包装的并发任务 (调度器的主要职责就是将G公平合理的安排到多个M上去执行)
        阻塞会创建新的, 空闲时也会回收

        # 配置: runtime/debug包SetMaxThreads()
    
    本地队列

        (优先)存放G
        若满了会放在全局队列

    全局队列

        存放G
```