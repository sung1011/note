# nginx modules

## 压缩
### ngx_http_gzip_module
gzip            on;
gzip_min_length 1000;   #小于1k的文件不压缩
gzip_types      text/plain application/xml;

## 浏览文件
### ngx_http_autoindex_module
autoindex   on;

## 限制流量
### ngx_http_core_module.Embedded Variables
set $limit_rate 1k; #限制BPS

## 防DDOS
ngx_http_limit_req_module 限制每秒请求数  
ngx_http_limit_conn_module 限制ip连接数  
geo, map 给上游（如lvs, haproxy）设置白名单  

## proxy && cache
### ngx_http_proxy_module
proxy_pass       http://localhost:8000;  
proxy_set_header Host      $host;  
proxy_set_header X-Real-IP $remote_addr;  

proxy_cache_path /tmp/nginxcache;  
proxy_cache mykey;  
proxy_cache_key $host$url$is_args$args;  
proxy_cache_valid 200 302 302 1d;  
