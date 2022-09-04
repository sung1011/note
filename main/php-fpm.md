# php-fpm

## config

    pm = static; 静态进程 (固定分配, 减少额外资源消耗, 但有进程切换消耗; memory_limit * max_children < 物理机内存; 可尽量多以满足高并发)  

- pm.max_children = 300; 静态方式下开启的php-fpm进程数量  

---

    pm = dynamic; 动态进程(一定范围内控制空闲worker数量; 有额外内存开销)  

- pm.start_servers = 20; 动态方式下的起始php-fpm进程数量  
- pm.max_children = 300; 动态方式下开启的最大php-fpm进程数量  
- pm.min_spare_servers = 5; 闲置子进程最小数量,  再小就fork  
- pm.max_spare_servers = 35; 闲置子进程最大数量,  再大就kill  

---

    pm = ondemand; 按需分配进程(一定范围内, 定时删除空闲worker, 直到只剩master; 内存友好, 不适合流量突发情况)  

- pm.max_children = 300;
- pm.process_idle_timeout = 10s;

---

    pm.max_requests = 10240; 每个worker处理多少个请求后会重启该线程 // 由于内存泄漏, 泄漏的内存会累计, 重启以归还内存  

- 内存消耗 = max_children * memory_limit; 静态进程内存消耗  
- 内存消耗 = max_spare_servers * memory_limit; 动态进程内存消耗  

---

    rlimit_files = 1024; 文件打开描述符的rlimit限制, 默认系统值(ulimit -n)(一般要跟系统的同步更改)  
    request_terminate_timeout (phpfpm.conf) 超时会报502Bad Gateway; 包含请求的一切时间; 会与 `max_execution_time` 同时生效, 谁先到达谁起作用.  
    覆盖ini: php_admin_value 如 php_admin_value[memory_limit] = 128M; php_admin_value[date.timezone] = Asia/Shanghai  

- process_control_timeout (phpfpm.conf) --- quit信号的超时时间, 超过该时间会在 `process_control_timeout+1` 后terminat.设置不合理, 则reload会导致terminat.建议值同 `request_terminate_timeout`  

## worker执行阶段

`fpm_request_stage_e` 源码中的结构体

| stage                       | 备注               |
| --------------------------- | ------------------ |
| FPM_REQUEST_ACCEPTING       | 空闲状态(等待请求) |
| FPM_REQUEST_READING_HEADERS | 读取头信息         |
| FPM_REQUEST_INFO            | 获取请求信息       |
| FPM_REQUEST_EXECUTING       | 执行状态           |
| FPM_REQUEST_END             | 请求结束状态       |

## scoreboard

```c
struct fpm_scoreboard_s {
    union {
        atomic_t lock;
        char dummy[16];
    };//锁状态
    char pool[32];//实例名称 例如: [www]
    int pm; //PM运行模式
    time_t start_epoch; //开始时间
    int idle;//procs的空闲数
    int active;//procs的使用数
    int active_max; //最大procs使用数
    unsigned long int requests;
    unsigned int max_children_reached; //到达最大进程数限制的次数
    int lq; //当前listen queue的请求数(accept操作, 可以过tcpi_unacked或getsocketopt获取)
    int lq_max;//listen queue大小
    unsigned int lq_len;
    unsigned int nprocs; //procs总数
    int free_proc; //从procs列表遍历下一个空闲对象的开始下标
    struct fpm_scoreboard_proc_s *procs[]; //列表
};
struct fpm_scoreboard_proc_s {
    union {
        atomic_t lock;
        char dummy[16];
    };//锁状态
    int used; //使用标识 0=未使用 1=正在使用
    time_t start_epoch; //使用开始时间
    pid_t pid; //进程id
    unsigned long requests; //处理请求次数
    enum fpm_request_stage_e request_stage; //处理请求阶段
    struct timeval accepted; //accept请求时间
    struct timeval duration; //脚本总执行时间
    time_t accepted_epoch;//accept请求时间戳(秒)
    struct timeval tv; //活跃时间
    char request_uri[128]; //请求路径
    char query_string[512]; //请求参数
    char request_method[16]; //请求方式
    size_t content_length; //请求内容长度 /* used with POST only */
    char script_filename[256];//脚本名称
    char auth_user[32];
#ifdef HAVE_TIMES
    struct tms cpu_accepted;
    struct timeval cpu_duration;
    struct tms last_request_cpu;
    struct timeval last_request_cpu_duration;
#endif
    size_t memory;//脚本占用的内存大小
};
```
