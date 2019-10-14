# redis

1. server.c
    1. redisServer
    2. redisCommand
    3. ht implementation
    4. main()
        1. initServerConfig()
        2. loadServerConfig()
        3. supervised || daemonize()
        4. initServer()
           1. setupSignalHandlers()
           2. server init
           3. aeCreateEventLoop()
           4. listenToPort()
           5. db init
           6. aeCreateTimeEvent()
              1. LRU
              2. RDB BGSAVE
              3. AOF rewrite
              4. 打印统计信息
              5. rehash
              6. close timeout conn
              7. slave request master sync
           7. aeCreateFileEvent()
              1. aeApiAddEvent()
              2. epoll_ctl()
                 1. acceptTcpHandler()
                 2. acceptCommonHandler()
                 3. createClient() # 创建redisClient对象 && 连接fd注册看到事件循环。
        5. aeMain()
              1. aeProcessEvents()
              2. aeApiPoll()
              3. evport_wait() / epoll_wait() || kqueue_wait() / select_wait()
2. 数据结构
   1. db
   2. object
   3. encoding
