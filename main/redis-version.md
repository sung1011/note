# redis 版本

## Redis2.6

Redis2.6在2012年正是发布，经历了17个版本，到2.6.17版本，相对于Redis2.4，主要特性如下：

1. **服务端支持Lua脚本。**
2. 去掉虚拟内存相关功能。
3. 放开对客户端连接数的硬编码限制。
4. 键的过期时间支持毫秒。
5. 从节点支持只读功能。
6. 两个新的位图命令：bitcount和bitop。
7. 增强了redis-benchmark的功能：支持定制化的压测，CSV输出等功能。
8. 基于浮点数自增命令：incrbyfloat和hincrbyfloat。
9. redis-cli可以使用--eval参数实现Lua脚本执行。
10. shutdown命令增强。
11. 重构了大量的核心代码，所有集群相关的代码都去掉了，cluster功能将会是3.0版本最大的亮点。
12. info可以按照section输出，并且添加了一些统计项
13. sort命令优化

## Redis2.8

Redis2.8在2013年11月22日正式发布，经历了24个版本，到2.8.24版本，相比于Redis2.6，主要特性如下：

1. 添加部分主从复制的功能，在一定程度上降低了由于网络问题，造成频繁全量复制生成RDB对系统造成的压力。
2. 尝试性的支持IPv6.
3. **可以通过config set命令设置maxclients。**
4. 可以用bind命令绑定多个IP地址。
5. Redis设置了明显的进程名，方便使用ps命令查看系统进程。
6. config rewrite命令可以将config set持久化到Redis配置文件中。
7. 发布订阅添加了pubsub。
8. Redis Sentinel第二版，相比于Redis2.6的Redis Sentinel，此版本已经变成生产可用。

## Redis3.0（里程碑)

Redis3.0在2015年4月1日正式发布，相比于Redis2.8主要特性如下：

Redis最大的改动就是添加Redis的分布式实现Redis Cluster。

1. **Redis Cluster：Redis的官方分布式实现。**
2. 全新的embedded string对象编码结果，优化小对象内存访问，在特定的工作负载下载速度大幅提升。
3. LRU算法大幅提升。
4. migrate连接缓存，大幅提升键迁移的速度。
5. migrate命令两个新的参数copy和replace。
6. 新的client pause命令，在指定时间内停止处理客户端请求。
7. bitcount命令性能提升。
8. cinfig set设置maxmemory时候可以设置不同的单位（之前只能是字节）。
9. Redis日志小做调整：日志中会反应当前实例的角色（master或者slave）。
10. incr命令性能提升。

## Redis3.2

Redis3.2在2016年5月6日正式发布，相比于Redis3.0主要特征如下：

1. 添加GEO相关功能。
2. SDS在速度和节省空间上都做了优化。
3. 支持用upstart或者systemd管理Redis进程。
4. **新的List编码类型：quicklist。**
5. 从节点读取过期数据保证一致性。
6. 添加了hstrlen命令。
7. 增强了debug命令，支持了更多的参数。
8. Lua脚本功能增强。
9. 添加了Lua Debugger。
10. config set 支持更多的配置参数。
11. 优化了Redis崩溃后的相关报告。
12. 新的RDB格式，但是仍然兼容旧的RDB。
13. 加速RDB的加载速度。
14. spop命令支持个数参数。
15. cluster nodes命令得到加速。
16. Jemalloc更新到4.0.3版本。

## Redis4.0

可能出乎很多的意料，Redis3.2之后的版本是4.0，而不是3.4、3.6、3.8。

一般这种重大版本号的升级也意味着软件或者工具本身发生了重大改革。下面是Redis4.0的新特性：

1. 提供了模块系统，方便第三方开发者拓展Redis的功能。
2. **PSYNC2.0：优化了之前版本中，主从节点切换必然引起全量复制的问题。**
3. **提供了新的缓存剔除算法：LFU（Last Frequently Used） ，并对已有算法进行了优化。**
4. 提供了非阻塞del和flushall/flushdb功能，有效解决删除了bigkey可能造成的Redis阻塞。
5. 提供了memory命令，实现对内存更为全面的监控统计。
6. 提供了交互数据库功能，实现Redis内部数据库的数据置换。
7. 提供了RDB-AOF混合持久化格式，充分利用了AOF和RDB各自优势。
8. Redis Cluster 兼容NAT和Docker。

## Redis5.0

TODO
