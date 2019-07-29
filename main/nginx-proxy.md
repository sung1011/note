# nginx 反向代理 缓存 负载均衡

## 反向代理

module: ngx_http_proxy_module  
directives:

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
- client_body_timeout 读取body时间是否超限，超时返回408

与上游建立连接

- proxy_connect_timeout 与上游握手超时时间，超时返回502
- proxy_next_upstream

## 缓存

### 指令

proxy_buffering, proxy_buffer_size, proxy_busy_buffers_size, proxy_temp_path, proxy_max_temp_file_size

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
