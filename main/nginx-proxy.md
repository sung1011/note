# nginx 反向代理 负载均衡

## 反向代理

### 流程

1. proxy_pass (处理content阶段)
   1. proxy_cache命中 [to 11. 发送响应header](_)
2. 根据指令生成发往上游的http header和body
3. proxy_request_buffering
   1. on; 读取完整请求body
   2. off; 边收边转发请求body
4. 根据负载均衡策略选择上游服务
5. 根据参数连接上游
6. 发送请求body（边读边发）
7. 处理并返回...
8. 接收响应header
9. 处理响应header
10. proxy_buffering
    1. on; 接收完整响应body
    2. off; 边收边转发响应body
11. 发送响应header
12. 发送响应body（边读边发）
13. proxy_cache
    1. on; 响应body加入缓存
14. 关闭连接 或 复用连接

### 指令

代理修改请求

- proxy_pass url 代理上游url
- proxy_http_version 修改version
- proxy_method 修改method
- proxy_set_header 修改header
- proxy_set_body 修改body
- proxy_pass_request_headers on|off 是否转发http头部
- proxy_pass_request_body on|off 是否转发http包体

接收请求body

- proxy_request_buffering on|off
  - on: nginx接收完毕再转发
    1. 客户端网速慢时
    2. 上游并发处理能力低时
    3. 高吞吐量时
  - off: nginx边收边转发
    1. 上游及时处理时
    2. 能降低nginx读写磁盘
    3. 会令proxy_next_upstream功能失效
- client_body_buffer_size
- client_body_in_single_buffer
- client_max_body_size 检查header.content-length是否超限，超限制返回413
- client_body_temp_path 请求body存储路径
- client_body_in_file_only 请求的body是否存文件
- client_header_timeout 读取header时间限制，超时返回408
- client_body_timeout 读取body时间限制，超时返回408

与上游建立连接

- proxy_connect_timeout 与上游握手超时时间，超时返回502
- proxy_next_upstream 遇到特定状态码时，将请求分发到下一台上游机器
- proxy_socket_keepalive 与上游连接是否启动TCP keepalive
- proxy_bind 修改与上游的TCP连接中的Source IP Address
- proxy_ignore_client_abort 代理是否忽略客户端连接
- proxy_send_timeout 向上游发送请求的超时时间(两次write期间计时)

接收上游响应header (缓存)

- proxy_buffer_size 响应头大小限制，超出则`upstream sent too big header`
- proxy_buffers 内存存放响应body的大小，超过则写入磁盘
- proxy_buffering 是否接收完上游响应，再返回给客户端
- proxy_max_temp_file_size 缓存文件大小
- proxy_temp_file_write_size 缓存每次写入大小
- proxy_temp_path 缓存路径
- proxy_busy_buffers_size 及时响应客户端的大小（刚接收到n字节时就转发给客户端n字节）
- proxy_read_timeout 接收上游响应的超时时间（两次write期间计时）
- proxy_limit_rate 读取上游的速率
- proxy_store_access 缓存文件权限，并持久化到指定位置
- proxy_store 缓存文件root，并持久化到root位置

处理上游响应header

- proxy_ignore_headers 禁用某些上游响应头部
- proxy_hide_header 隐藏某些上游响应头部 (以下为默认隐藏的头部)
  - Date nginx发送响应头的时间，由ngx_http_header_filter_module填写
  - Server nginx版本，由ngx_http_header_filter_module填写
  - X-Pad apache为避免浏览器bug生成的头部，默认忽略
  - X-Accel- 用于控制nginx行为的响应，不需要向客户端转发
- proxy_pass_header 恢复proxy_hide_header隐藏的头部
- proxy_cookie_domain 修改上游返回的cookie
- proxy_cookie_path 替换上游返回的cookie
- proxy_redirect 修改返回的location header
- expires 缓存过期时间

### cache

### cache流程

TODO

### cache指令

- proxy_cache 缓存名
- proxy_cache_path 缓存路径，淘汰策略等
- proxy_cache_key 缓存生效的url关键词
- proxy_cache_vailid 遇到特定状态码，缓存不同的时长
- proxy_no_cache 参数为真时，不存入缓存
- proxy_cache_bypass 参数为真时，有缓存但不使用该缓存内容
- proxy_cache_convert_head 变更head方法为get方法
- $upstream_cache_status 缓存状态变量

## 负载均衡

### AKF分类

x轴： 基于算法分发请求  
y轴： 基于URL对功能分发到不同的服务  
z轴： 基于用户ip或者其他信息映射到某个特定集群或服务  

### method

权重,随机,加权随机 round-robin  
IP哈希 ip-hash  
最小连接数 least-connected  
哈希 hash: URL哈希或自定义的哈希
随机 random
fair(第三方): 根据响应时间  

- hash算法可以配置一致性hash, 以防扩缩容时缓存大量失效, 带来压力。 `hash key consistent`  
- 使用upstream_zone模块，以共享内存的方式将负载均衡算法共享到所有worker生效。

### Health checks

module: ngx_http_upstream_module  
command: upstream, server

- server状态
  - weight: 权重
  - max_conns: 最大并发的连接数
  - max_fails: 允许请求失败次数
  - fail_timeout: 经过max_fails失败后,服务暂停的时间(熔断)
  - backup: 备份服务器（当其余服务器异常时启用）  
  - down: 标记已下线  
  - resolve: 指定dns解析域名
  - route: set route name
  - service: set service name
  - slow_start: 防止新添加/恢复的主机被突然增加的请求所压垮，令weight从0开始慢慢增加到设定值。 （不能用于hash, ip_hash)
  - drain: 优雅关闭。即不再接受请求。

## 实战

### 配置

```conf
http {
    upstream myapp1 { # 负载均衡
        server srv1.example.com;
        server srv2.example.com;
        server srv3.example.com;
    }

    server {
        listen 80;

        location / {
            proxy_pass http://myapp1; # 代理
        }
    }
}
```
