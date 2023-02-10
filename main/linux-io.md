# linux IO模型

## 基础

### 同步  

    请求发起方对结果的获取是`主动问询`
  
### 异步  

    请求发起方对结果的获取是`被动通知`
  
### 阻塞  

    调用(针对io)后线程`挂起`等待结果  
  
### 非阻塞  

    调用(针对io)后线程`立刻返回`去忙别的
  
> 同步/异步针对结果; 阻塞/非阻塞针对线程

## 组合  

### 同步阻塞 BIO  

    请求方发起, 线程一直等待结果  
  
### 同步非阻塞 NIO  

    请求方发起, 线程不等待结果直接返回, 通过主动轮询不断查看结果  
  
### 异步阻塞 (没意义)  

    请求方发起后, 挂起等待服务方通知  
  
### 异步非阻塞 AIO  

    请求方发起后, 返回做其他任务, 等待服务方通知  
  
## 实例

### 多线程 [同步阻塞]

```js
    每个usr线程阻塞等待sys的返回.

    # issues
        阻塞
```

### select多路复用 [同步非阻塞]  

```js
    1个线程同时遍历多个io请求结果.

    # flow
        1. 将需要io操作的fd注册到events中
        2. select线程阻塞等待select()系统调用返回
        3. 当数据从sys就绪, 会 select() > 0
        4. select线程遍历events找到准备就绪的fd, 将其标记为就绪状态fd_set = 1.
        5. usr线程正式发起read/write请求, 从sys读取那些标记就绪的数据.

    # issue
        句柄存在上限
        重复初始化
        逐个遍历排查句柄状态

    # 结构
        fd_set rfds, wfds;

    # 伪代码
        {
            select(socket);
            while(1) {
                sockets = select();
                for(socket in sockets) {
                    if(can_read(socket)) {
                    read(socket, buffer);
                    process(buffer);
                }
                }
            }
        }
```

> 采用一个1024数组存储状态, 最多同时检查1024.

> 非线程安全

### poll多路复用 [同步非阻塞]  

```js
       链表结构的select

        # flow
           1. 主动轮询链表(采用链表避免数组长度限制), 其他同select(依然遍历全部fd, 看哪个fd有sys返回)  

        # issue
           逐个遍历排查句柄状态; 某个sys线程就绪将poll激活时, poll并不知道哪个fd有数据, 只能自己遍历.
```

> 非线程安全
  
### epoll多路复用 [异步非阻塞]  

```js
        当fb准备就绪, 会被异步回调通知得到哪些fb已经ready

        # flow
           1. 通过epoll_ctl注册fd(注册到RBTREE结构)
           2. 一旦该fd就绪, 内核就会采用类似callback的回调机制来激活该fd, epoll_wait便可以收到通知. (即不用再遍历而是监听回调进行io)
           3. usr,sys通过共享内存传递消息.
        
        # 结构
           int epfd;
           struct epoll_event *events;
```

> 线程安全

### 信号驱动 (SIGIO) [同步非阻塞]  

    主动轮询链表, 设备就绪后信号通知将就绪fd事件放入链表.  
  
### aio [异步非阻塞]  

    io完成后, 信号通知或者回调线程  
  
## ref

- [IO多路复用的三种机制Select, Poll, Epoll](https://www.jianshu.com/p/397449cadc9a)
- [聊聊IO多路复用之select、poll、epoll详解](https://my.oschina.net/xianggao/blog/663655)
- [IO多路复用—由Redis的IO多路复用](https://blog.csdn.net/happy_wu/article/details/80052617)
- [select、poll、epoll之间的区别(搜狗面试)](https://www.cnblogs.com/aspirant/p/9166944.html)
