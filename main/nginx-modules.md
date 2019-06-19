# nginx modules

## 类型 - 结构划分
![ img ](res/nginx-modules.jpeg)
核心模块：HTTP、EVENT、MAIL、STEAM  
基础模块：HTTP Access模块、HTTP FastCGI模块、HTTP Proxy模块和HTTP Rewrite模块，  
第三方模块：HTTP Upstream Request Hash模块、Notice模块和HTTP Access Key模块。  

## 类型 - 功能划分
Core(核心模块)：构建nginx基础服务、管理其他模块。  
Handlers（处理器模块）：此类模块直接处理请求，并进行输出内容和修改headers信息等操作。  
Filters （过滤器模块）：此类模块主要对其他处理器模块输出的内容进行修改操作，最后由Nginx输出。  
Proxies （代理类模块）：此类模块是Nginx的HTTP Upstream之类的模块，这些模块主要与后端一些服务比如FastCGI等进行交互，实现服务代理和负载均衡等功能。  

## 源码
```
# nginx/src/
|-- core    # 框架核心
|-- event
    |-- modules     # 基础模块 + 第三方模块
    |   |-- ngx_devpoll_module.c
    |   |-- ngx_epoll_module.c
    |   |-- ngx_eventport_module.c
    |   |-- ngx_kqueue_module.c
    |   |-- ngx_poll_module.c
    |   |-- ngx_select_module.c
    |   |-- ngx_win32_poll_module.c
    |   `-- ngx_win32_select_module.c
    |-- ngx_event.c     # 核心模块 优先执行
    |-- ngx_event.h
    |-- ngx_event_accept.c
    |-- ngx_event_connect.c
    |-- ngx_event_connect.h
    |-- ngx_event_openssl.c
    |-- ngx_event_openssl.h
    |-- ngx_event_openssl_stapling.c
    |-- ngx_event_pipe.c
    |-- ngx_event_pipe.h
    |-- ngx_event_posted.c
    |-- ngx_event_posted.h
    |-- ngx_event_timer.c
    |-- ngx_event_timer.h
    `-- ngx_event_udp.c
|-- http
|-- mail
|-- misc
|-- os
`-- stream
```

## 实战
### 压缩
#### ngx_http_gzip_module
gzip            on;
gzip_min_length 1000;   #小于1k的文件不压缩
gzip_types      text/plain application/xml;

### 浏览文件
#### ngx_http_autoindex_module
autoindex   on;

### 限制流量
#### ngx_http_core_module.Embedded Variables
set $limit_rate 1k; #限制BPS

### 防DDOS
ngx_http_limit_req_module 限制每秒请求数  
ngx_http_limit_conn_module 限制ip连接数  
geo, map 给上游（如lvs, haproxy）设置白名单  

### proxy && cache
#### ngx_http_proxy_module
proxy_pass       http://localhost:8000;  
proxy_set_header Host      $host;  
proxy_set_header X-Real-IP $remote_addr;  

proxy_cache_path /tmp/nginxcache;  
proxy_cache mykey;  
proxy_cache_key $host$url$is_args$args;  
proxy_cache_valid 200 302 302 1d;  

### 获取客户端ip
#### ngx_realip_module
阶段: postread  
功能: 修改客户端地址$remote_addr  
变量: realip_remote_addr, realip_remote_port  
指令: set_real_ip_from, real_ip_header, real_ip_recursive  
http.header.x-forwarded-for: 经过的ip的集合 如[115.204.33.1, 1.1.1.1]
http.header.x-real-ip: 用户ip 如115.204.33.1

### rewrite 重写
#### ngx_rewrite_module
阶段: server rewrite, rewrite
##### rewrite
##### return
return 444 "body msg";
##### error_page 重定向错误码处理的地址
error_page 444 /err.html

### if 条件
#### ngx_rewrite_module
阶段: server rewrite, rewrite
条件表达:
- 检查变量是否空或为0
- 将变量与字符串做比较 = !=
- 支持正则
- 检查文件是否存在 -f !-f
- 检查目录是否存在 -d !-d
- 检查文件 目录 软连是否存在 -e !-e
- 检查是否可执行文件 -x !-x

### limit 限流
#### http_limit_req_module 限制并发请求数
阶段: pre_access  
算法: leaky bucket  突发流量限定为恒定流量, 故响应可能变慢, 超流量返回错误。
指令: limit_req_zone, limit_req, limit_req_log_level, limit_req_status
范围:  
- all worker (基于共享内存)  
- 进入pre_access前不生效  
#### http_limit_conn_module 限制并发连接数
阶段: pre_access  
指令: limit_conn_zone, limit_conn, limit_conn_log_level, limit_conn_status  
范围:   
- all worker (基于共享内存)  
- 进入pre_access前不生效  
- 限制的有效性取决于key的设定， key一般用客户端ip (取真实客户端ip依赖realip模块)  