# nginx modules

## 类型 - 结构划分
![ img ](res/nginx-modules.jpeg)
核心模块：HTTP模块、EVENT模块和MAIL模块  
基础模块：HTTP Access模块、HTTP FastCGI模块、HTTP Proxy模块和HTTP Rewrite模块，  
第三方模块：HTTP Upstream Request Hash模块、Notice模块和HTTP Access Key模块。  

## 类型 - 功能划分
Core(核心模块)：构建nginx基础服务、管理其他模块。  
Handlers（处理器模块）：此类模块直接处理请求，并进行输出内容和修改headers信息等操作。  
Filters （过滤器模块）：此类模块主要对其他处理器模块输出的内容进行修改操作，最后由Nginx输出。  
Proxies （代理类模块）：此类模块是Nginx的HTTP Upstream之类的模块，这些模块主要与后端一些服务比如FastCGI等进行交互，实现服务代理和负载均衡等功能。  

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
