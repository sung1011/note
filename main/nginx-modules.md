# nginx modules

## 类型 - 结构划分

![img](res/nginx-modules.jpeg)

- 核心module: HTTP、EVENT、MAIL、STEAM
- 基础module: HTTP Access、HTTP FastCGI、HTTP Proxy和HTTP Rewrite
- 第三方module: HTTP Upstream Request Hash、Notice和HTTP Access Keymodule.  

## 类型 - 功能划分

- Core(核心module): 构建nginx基础服务、管理其他module.  
- Handlers(处理器module): 此类module直接处理请求, 并进行输出内容和修改headers信息等操作.  
- Filters (过滤器module): 此类module主要对其他处理器module输出的内容进行修改操作, 最后由Nginx输出.  
- Proxies (代理类module): 此类module是Nginx的HTTP Upstream之类的module, 这些module主要与后端一些服务比如FastCGI等进行交互, 实现服务代理和负载均衡等功能.  

## 目录

```tree
# nginx/src/
|-- core    # 框架核心
|-- event
    |-- modules     # 基础module + 第三方module
    |   |-- ngx_devpoll_module.c
    |   |-- ngx_epoll_module.c
    |   |-- ngx_eventport_module.c
    |   |-- ngx_kqueue_module.c
    |   |-- ngx_poll_module.c
    |   |-- ngx_select_module.c
    |   |-- ngx_win32_poll_module.c
    |   `-- ngx_win32_select_module.c
    |-- ngx_event.c     # 核心module 优先执行
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

## 数据结构

```nginx
ngx_module_t module;
ngx_directives_t directives;
```

## 方法

ngx_xxx_merge_conf 合并配置  

## 实战

### 压缩

```nginx
# module: ngx_http_gzip_module  

gzip            on;  
gzip_min_length 1000;                       #小于1k的文件不压缩  
gzip_types      text/plain application/xml;  
gzip_static     on;                         #读取预先压缩的.gz文件, 存在则直接返回, 免去临时压缩返回的资源消耗.
```

### 浏览文件

```nginx
# module: ngx_http_autoindex_module

autoindex   on;
```

### 限制流量

```nginx
# module: ngx_http_core_module.Embedded Variables

set $limit_rate 1k; #限制BPS
```

### 防DDOS

```nginx
# module: ngx_http_limit_req_module 限制每秒请求数  
# module: ngx_http_limit_conn_module 限制ip连接数  
# geo, map 给上游(如lvs, haproxy)设置白名单  
```

### proxy && cache

```nginx
# module: ngx_http_proxy_module

proxy_pass                  <http://localhost:8000;>  
proxy_set_header Host       $host;  
proxy_set_header X-Real-IP  $remote_addr;  

proxy_cache_path            /tmp/nginxcache;  
proxy_cache                 mykey;  
proxy_cache_key             $host$url$is_args$args;  
proxy_cache_valid           200 302 302 1d;  
```

### 获取客户端ip

```nginx
# module: ngx_realip_module  
# stage: postread  
# 功能: 修改客户端地址$remote_addr  
# 变量: realip_remote_addr, realip_remote_port  
# directives: set_real_ip_from, real_ip_header, real_ip_recursive  

http.header.x-forwarded-for: 经过的ip的集合 如[115.204.33.1, 1.1.1.1]  
http.header.x-real-ip: 用户ip 如115.204.33.1  
```

### rewrite 重写

```nginx
# module: ngx_rewrite_module  
# stage: server-rewrite, rewrite  

# return
return 444 "body msg";

# error_page 重定向错误码处理的地址
error_page 444 /err.html

# if 条件
条件表达:

- 检查变量是否空或为0
- 将变量与字符串做比较 = !=
- 支持正则
- 检查文件是否存在 -f !-f
- 检查目录是否存在 -d !-d
- 检查文件 目录 软连是否存在 -e !-e
- 检查是否可执行文件 -x !-x
```

### limit 限流

```nginx
# 限制并发请求数
# module: http_limit_req_module  
# stage: pre_access  
# 算法: leaky bucket  突发流量限定为恒定流量, 故响应可能变慢, 超流量返回错误.  
# directives: limit_req_zone, limit_req, limit_req_log_level, limit_req_status  

- all worker (基于共享内存)  
- 进入pre_access前不生效  


# 限制并发连接数
# module: http_limit_conn_module  
# stage: pre_access  
# directives: limit_conn_zone, limit_conn, limit_conn_log_level, limit_conn_status  

- all worker (基于共享内存)  
- 进入pre_access前不生效  
- 限制的有效性取决于key的设定,  key一般用客户端ip (取真实客户端ip依赖realipmodule)  


    # limit_req_zone命令及limit_req命令限制单个IP的请求处理频率
	# 定义限流维度，一个用户一分钟一个请求进来，多余的全部漏掉
    #  1r/s代表1秒一个请求，1r/m一分钟接收一个请求， 如果Nginx这时还有别人的请求没有处理完，Nginx就会拒绝处理该用户请求

	limit_req_zone $binary_remote_addr zone=one:10m rate=1r/m;
 
	#绑定限流维度
	server{
		location/seckill.html{
			limit_req zone=zone;	
			proxy_pass http://lj_seckill;
		}
	}


    # Nginx提供burst参数结合nodelay参数可以解决流量突发的问题
  	# 定义限流维度，一个用户一分钟一个请求进来，多余的全部漏掉
    # 为什么就多了一个 burst=5 nodelay; 呢，多了这个可以代表Nginx对于一个用户的请求会立即处理前五个，多余的就慢慢来落，没有其他用户的请求我就处理你的，有其他的请求的话我Nginx就漏掉不接受你的请求
	limit_req_zone $binary_remote_addr zone=one:10m rate=1r/m;
 
	#绑定限流维度
	server{
		location/seckill.html{
			limit_req zone=zone burst=5 nodelay;
			proxy_pass http://lj_seckill;
		}
	}



    # limit_conn_zone + limit_conn限制并发连接数
    # 单个IP10并发, 整个服务最大100并发
  	http {
		limit_conn_zone $binary_remote_addr zone=myip:10m;
		limit_conn_zone $server_name zone=myServerName:10m;
	}
 
    server {
        location / {
            limit_conn myip 10;
            limit_conn myServerName 100;
            rewrite / http://www.lijie.net permanent;
        }
    }
```

> [漏桶](algo-leaky-bucket.md)

> [令牌桶](algo-token-bucket.md)

### access 认证


```nginx
# 限制ip
# module: http_access_module  
# stage: access  
# directives: alow, deny  

# 限制用户名密码
# module: http_access_module  
# stage: access  
# directives: auth_basic, auth_basic_user_file  
# 工具: 密码文件生成依赖httpd-tools库, `htpasswd -c < file > -b < username > < password >`  

# 向上游服务验证用户名密码
# module: http_auth_request_module  
# stage: access  
# directives: auth_request, auth_request_set  
# 原理: 向上游服务转发请求, 若上游返回200则验证通过, 否则验证失败.  

# 配置条件
# module: ngx_http_core_module  
# directives: satisfy all|any  
# 原理: all全部放行才放行, any任一放行就放行  

location / {
    satisfy any; 任一满足即可.如访问以下ip 或 密码验证正确
    allow 192.168.1.0/32;
    deny  all;
    auth_basic           "closed site";
    auth_basic_user_file conf/htpasswd;
}
```  

### pre_content

```nginx
# 试图访问多个url路径, 若文件都不存在则返回最后一个url或者code
# module: ngx_http_try_file_module  
# stage: pre_content  
# directives: try_file  

# 流量拷贝, 处理请求时, 生成子请求访问其他服务, 但不处理其返回值
# module: ngx_http_mirror_module  
# stage: pre_content  
# directives: mirror, mirror_request_body  
```

### content

```nginx
# static (已并入ngx_http_core_module)

# index 返回主页

# autoindex 返回目录结构

# proxy_pass 反向代理

# concat (第三方) 同时请求/下载多个小文件

# module: ngx_http_concat  
# stage: content  
# usage: `https://localhost/??a.js,b.css,res/c.js`  
```

### log

```nginx
# module: ngx_http_log_module  
# stage: log  
# directives: access_log, log_format, open_log_file_cache  

- access_log 日志
  - buffer: 缓存大小超过设定  
  - flush: 缓存时间超过设定  
  - gzip: 缓存压缩比率

- log_format 日志格式

- open_log_file_cache 文件缓存 以减少含有变量路径的日志打开关闭的消耗
  - max: 设置缓存中的最大文件描述符数量, 如果缓存被占满, 采用LRU算法将描述符关闭.
  - inactive: 设置存活时间, 默认是10s
  - min_uses: 设置在inactive时间段内, 日志文件最少使用多少次后, 该日志文件描述符记入缓存中, 默认是1次
  - valid: 设置检查频率, 默认60s
  - off: 禁用缓存
```

### filter

```nginx
# 替换返回
# module: ngx_http_sub_filter_module
# directives: sub_filter, sub_filter_last_modified, sub_filter_once, sub_filter_types

# 前后添加返回 (添加的body内容为子请求的返回值)
# module: ngx_http_addition_filter_module
# directives: add_before_body, add_after_body, addition_types
```

### referer模块 携带client信息以防盗链

### map模块 按条件匹配设置某个新变量的值

### split_clients模块 按百分比匹配设置某个新变量的值以A/B方案测试

### geo模块 根据ip地址匹配设置新变量的值以区分不同地区的客户端请求

### geoip模块 是geo的增强版

### 即时清除cache

```nginx
# module: ngx_cache_purge
# directives: proxy_cache_purge
```

### 反向代理websocket

```nginx
# module: ngx_http_proxy_module  

请求:

- proxy_http_version 1.1; // HTTP/1.1
- proxy_set_header Connection "upgrade"; // Connection:keep-alive,Upgrade
- proxy_set_header Upgrade $http_upgrade; // Upgrade: websocket

响应行:

- HTTP/1.1 101 Web Socket Protocol Handshake
```

### slice 通过range协议, 分解并缓存大的数据块

```nginx
# module: http_slice_module  
# directives: slice
```

### openfilecache 打开文件的缓存

```nginx
# module: ngx_http_core_module
# directives: open_file_cache, open_file_cache_errors, open_file_cache_min_users, open_file_cache_valid

- open_file_cache max=N [inactive=time]
  - max 每个worker最多缓存多少个文件, 超出LRU淘汰
  - inactive 多少秒后没被访问, 则淘汰

```

### 防爬虫

```nginx
server{
	listen 80;
	server_name 127.0.0.1; 
	# 添加如下内容即可防止爬虫
	if ($http_user_agent ~* "qihoobot|Baiduspider|Googlebot|Googlebot-Mobile|Googlebot-Image|Mediapartners-Google|Adsbot-Google|Feedfetcher-Google|Yahoo! Slurp|Yahoo! Slurp China|YoudaoBot|Sosospider|Sogou spider|Sogou web spider|MSNBot|ia_archiver|Tomato Bot") 
	{
		return 403;
	}
}
```

### 限制浏览器

```nginx
if ($http_user_agent ~* "Firefox|MSIE")
{
     return 403;
}
```
