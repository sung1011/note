# nginx 反向代理 负载均衡

## 反向代理流程

1. pass (处理content阶段)
   1. 缓存命中 [to 11. 发送响应header](_)
2. 根据指令生成发往上游的http header和body
3. request_buffering
   1. on; 读取完整请求body
   2. off; 边收边转发请求body
4. 根据负载均衡策略选择上游服务
5. 根据参数连接上游
6. 发送请求body(边读边发)
7. 处理并返回...
8. 接收和处理响应header
9. buffering
    1. on; 接收完整响应body
    2. off; 边收边处理上游响应headerbody
10. 返回响应header
11. 返回响应body(边读边发)
12. cache
    1. on; 响应body加入缓存
13. 关闭连接 或 复用连接

## 缓存

### 缓存流程 --- 客户端请求的缓存处理流程

1. 是否开启cache
    1. 否: 向上游发起请求
2. 是否匹配cache_methods
    1. 否: 向上游发起请求
3. 根据cache_convert_head进行HEAD->GET方法转换
4. 根据key生成关键字并md5操作
5. 检查cache_pass是否指明不缓存
    1. 是: 向上游发起请求
6. 在共享内存中查询缓存是否存在
    1. 不存在: [to 10](_)
7. 更新LRU链表及节点计数
8. 是否错误类响应且过期
    1. 是: 向上游发起请求
9. 文件存在且使用次数超出cache_min_users
    1. 是: 根据cache_background_update生成子请求; 向下游发送缓存的响应
    2. 否: 向上游发起请求
10. 共享内存中分配节点
11. 淘汰已过期的缓存, 再次分配

### 缓存流程 --- 接收上游响应的缓存处理流程

1. 是否匹配no_cache
   1. 否: 不更新缓存, 转发上游响应
2. 方法是否匹配cache_valid
   1. 否: 不更新缓存, 转发上游响应
3. 判断响应码是否为200或206
   1. 是: 更新缓存中的etag和last_modified
4. 处理缓存相关的响应header, 并检查是否有响应头不使用缓存
   1. 是: 不更新缓存, 转发上游响应
5. 读取、转发上游响应
6. 将临时文件mv至缓存目录
7. 更新共享内存状态

## 缓存肖峰

### 并发时由第1个请求激活cache, 其他请求利用该cache响应

- cache_lock 并发时仅1个请求发向上游, 其他请求等待那个请求返回并生成缓存, 使用该缓存响应客户端.
- cache_lock_timeout 并发时等待第1个请求返回的超时时间, 超时则自己向上游请求.
- cache_lock_age 并发时等待第1个请求返回的超时时间, 超时则再放行一个向上游请求.

### 减少回源请求, 使用stale陈旧的缓存; 并发时用旧缓存响应

- cache_use_stale
  - updating 并发时仅1个请求发向上游, 其他请求使用已经过期的旧缓存
  - error 与上游发生错误时, 使用缓存
  - timeout 与上游发生超时时, 使用缓存
  - http_(500|502|503|504|403|404|429) 缓存以上错误码的内容
- cache_background_update 并发时多个请求发向上游, 其他请求使用已经过期的旧缓存

### 缓存清除模块 ngx_cache_purge(第三方模块)

## 负载均衡

### AKF分类

- x轴: 基于算法分发请求  
- y轴: 基于URL对功能分发到不同的服务  
- z轴: 基于用户ip或者其他信息映射到某个特定集群或服务  

### 方式

- `权重,随机,加权随机 round-robin`
- `IP哈希 ip-hash`
- `最小连接数 least-connected`
- `哈希 hash` URL哈希或自定义的哈希
- `随机 random`
- `fair(第三方)` 根据响应时间  

> hash算法可以配置一致性hash, 以防扩缩容时缓存大量失效, 带来压力. `hash key consistent`  

> 使用upstream_zone模块, 以共享内存的方式将负载均衡算法共享到所有worker生效.

## Health-Checks

    module: ngx_http_upstream_module  
    command: upstream, server

```nginx
resolver 10.0.0.1;

upstream dynamic {
    zone upstream_dynamic 64k;

    server backend1.example.com      weight=5;
    server backend2.example.com:8080 fail_timeout=5s slow_start=30s;
    server 192.0.2.1                 max_fails=3;
    server backend3.example.com      resolve;
    server backend4.example.com      service=http resolve;

    server backup1.example.com:8080  backup;
    server backup2.example.com:8080  backup;
}

server {
    location / {
        proxy_pass http://dynamic;
        health_check match=welcome; # 每隔5秒向每个server发健康检查心跳
    }
}

match welcome {     # 自定义满足以下全部条件, 则认为是健康的.
    status 200;
    header Content-Type = text/html;
    body ~ "Welcome to nginx!";
}
```

- server状态
  - weight: 权重
  - max_conns: 最大并发的连接数
  - max_fails: 允许请求失败次数 (熔断)
  - fail_timeout: 经过max_fails失败后, 服务暂停的时间(熔断)
  - backup: 备份服务器(当其余服务器异常时启用)  
  - down: 标记已下线  
  - resolve: 指定dns解析域名
  - route: set route name
  - service: set service name
  - slow_start: 防止新添加/恢复的主机被突然增加的请求所压垮, 令weight从0开始慢慢增加到设定值. (不能用于hash, ip_hash)
  - drain: 优雅关闭, 即不再接受请求.

## 反向代理指令

| key                      | meaning                                | class              | uwsgi | fastcgi | scgi | http |
| ------------------------ | -------------------------------------- | ------------------ | ----- | ------- | ---- | ---- |
| pass                     | 指定上游                               | 构造请求内容       | o     | o       | o    | o    |
| pass_request_head        | 是否传递请求头部                       | 构造请求内容       | o     | o       | o    | o    |
| pass_request_body        | 是否传递请求包体                       | 构造请求内容       | o     | o       | o    | o    |
| method                   | 指定请求方法名                         | 构造请求内容       |       |         |      | o    |
| http_version             | 指定请求协议                           | 构造请求内容       |       |         |      | o    |
| set_header               | 设置请求头                             | 构造请求内容       |       |         |      | o    |
| set_body                 | 设置请求包体                           | 构造请求内容       |       |         |      | o    |
| pass_request_headers     | 是否转发请求头                         | 构造请求内容       |       |         |      | o    |
| pass_request_body        | 是否转发请求体                         | 构造请求内容       |       |         |      | o    |
| request_buffering   | 是否缓存请求包体                       | 接收请求body       | o     | o       | o    | o    |
| client_max_body_size     | 检查请求头content-length, 超返413      | 接收请求body       |       |         |      | o    |
| client_body_buffer_size  | 分配请求体buffer大小, 小则内存大则文件 | 接收请求body       |       |         |      | o    |
| client_body_temp_path    | 请求body临时文件存储路径               | 接收请求body       |       |         |      | o    |
| client_body_in_file_only | 请求的body是否存文件                   | 接收请求body       |       |         |      | o    |
| client_header_timeout    | 读取header时间限制, 超返408            | 接收请求body       |       |         |      | o    |
| client_body_timeout      | 读取body时间限制, 超返408              | 接收请求body       |       |         |      | o    |
| connect_timeout          | 连接上游超时时间, 超反502              | 建立连接并发送请求 | o     | o       | o    | o    |
| bind                     | TCP连接绑定地址(Source IP Address)     | 建立连接并发送请求 | o     | o       | o    | o    |
| socket_keepalive         | 使用TCPkeepalive                       | 建立连接并发送请求 | o     | o       | o    | o    |
| ignore_client_abort      | 忽略客户端关连接                       | 建立连接并发送请求 | o     | o       | o    | o    |
| headers_hash_bucket_size | 设置HTTP头部用到的hash buket size      | 建立连接并发送请求 |       |         |      | o    |
| headers_hash_max_size    | 设置HTTP头部用到的hash max size        | 建立连接并发送请求 |       |         |      | o    |
| send_timeout             | 发送请求超时时间                       | 建立连接并发送请求 | o     | o       | o    | o    |
| buffering                | 是否缓冲上游响应                       | 接收上游响应       | o     | o       | o    | o    |
| buffer_size              | 接收响应每次写文件的大小               | 接收上游响应       | o     | o       | o    | o    |
| buffers                  | 缓冲区的数量和大小                     | 接收上游响应       | o     | o       | o    | o    |
| store                    | 是否持久化缓冲包体文件                 | 接收上游响应       | o     | o       | o    | o    |
| store_access             | 设置持久化缓冲包体文件权限             | 接收上游响应       | o     | o       | o    | o    |
| temp_path                | 存放上游响应缓冲的目录                 | 接收上游响应       | o     | o       | o    | o    |
| temp_file_write_size     | 缓冲文件每次写入大小                   | 接收上游响应       | o     | o       | o    | o    |
| max_temp_file_size       | 临时文件最大大小                       | 接收上游响应       | o     | o       | o    | o    |
| busy_buffers_size        | 缓冲完成前转发包体的大小               | 接收上游响应       | o     | o       | o    | o    |
| read_timeout             | 读取响应超时时间(两次write期间计时)  | 接收上游响应       | o     | o       | o    | o    |
| limit_rate               | 读取响应限速率                         | 接收上游响应       | o     | o       | o    | o    |
| hide_header         | 隐藏某些响应头部                       | 处理上游响应header | o     | o       | o    | o    |
| pass_header              | 恢复hide_header隐藏的头部              | 处理上游响应header | o     | o       | o    | o    |
| ignore_headers           | 禁止处理某些响应头部                   | 处理上游响应header | o     | o       | o    | o    |
| cookie_domain            | 替换set-cookie头部中的域名             | 处理上游响应header |       |         |      | o    |
| cookie_path              | 替换set-cookie头部中的URL              | 处理上游响应header |       |         |      | o    |
| redirect                 | 修改重定向相应中location的值           | 处理上游响应header |       |         |      | o    |
| next_upstream            | 出指定状态码时更换上游                 | 处理上游响应header | o     | o       | o    | o    |
| next_upstream_timeout    | 更换上游超时                           | 处理上游响应header | o     | o       | o    | o    |
| next_upstream_tries      | 更换上游重试次数                       | 处理上游响应header | o     | o       | o    | o    |
| intercept_errors         | 拦截上游错误响应                       | 处理上游响应header | o     | o       | o    | o    |
| ssl_certificate          | 配置用于上游通讯的证书                 | SSL                | o     |         |      | o    |
| ssl_certificate_key      | 配置用于上游通讯的私钥                 | SSL                | o     |         |      | o    |
| ssl_ciphers              | 指定安全套件                           | SSL                | o     |         |      | o    |
| ssl_crl                  | 指定吊销证书链CRL文件验证上游的证书    | SSL                | o     |         |      | o    |
| ssl_name                 | 指定域名验证上游证书中域名             | SSL                | o     |         |      | o    |
| ssl_password_file        | 当私钥有密码时指定密码文件             | SSL                | o     |         |      | o    |
| ssl_protocols            | 指定具体某个版本的协议                 | SSL                | o     |         |      | o    |
| ssl_server_name          | 传递SNI信息至上游                      | SSL                | o     |         |      | o    |
| ssl_session_reuse        | 是否重用SSL连接                        | SSL                | o     |         |      | o    |
| ssl_trusted_certificate  | 验证上游服务的证书                     | SSL                | o     |         |      | o    |
| ssl_verify               | 是否验证上游服务的证书                 | SSL                | o     |         |      | o    |
| ssl_verify_depth         | 设置验证证书链的深度                   | SSL                | o     |         |      | o    |
| cache                    | 指定共享内存名                         | 缓存               | o     | o       | o    | o    |
| cache_path               | 缓存文件存放位置, 淘汰策略等           | 缓存               | o     | o       | o    | o    |
| cache_bypass             | 指定哪些请求不使用缓存                 | 缓存               | o     | o       | o    | o    |
| cache_background_update  | 开启子请求更新陈旧缓存                 | 缓存               | o     | o       | o    | o    |
| cache_key                | 定义缓存关键字                         | 缓存               | o     | o       | o    | o    |
| cache_max_range_offset   | 使用range协议的偏移                    | 缓存               | o     | o       | o    | o    |
| cache_methods            | 缓存哪些请求方法                       | 缓存               | o     | o       | o    | o    |
| cache_min_uses           | 多少请求后再缓存                       | 缓存               | o     | o       | o    | o    |
| cache_valid              | 缓存哪些响应及时长                     | 缓存               | o     | o       | o    | o    |
| force_ranges             | 强制使用range协议                      | 缓存               | o     | o       | o    | o    |
| cache_revalidate         | 有陈旧内容使用304                      | 缓存               | o     | o       | o    | o    |
| cache_use_stale          | 返回陈旧的缓存内容                     | 缓存               | o     | o       | o    | o    |
| no_cache                 | 指定哪些响应不会写入缓存               | 缓存               | o     | o       | o    | o    |
| cache_convert_head       | 将head方法转换为get方法                | 缓存               |       |         |      | o    |
| cache_lock               | 并发时加锁减少回源请求                 | 缓存               | o     | o       | o    | o    |
| cache_lock_age           | 回源请求达到该超时时间后再放行         | 缓存               | o     | o       | o    | o    |
| cache_lock_timeout       | 等待请求的最长等待时间                 | 缓存               | o     | o       | o    | o    |
| modifier1                |                                        | 独有配置           | o     |
| modifier2                |                                        | 独有配置           | o     |
| param                    |                                        | 独有配置           | o     | o       | o    |
| index                    |                                        | 独有配置           |       | o       |
| catch_stderr             |                                        | 独有配置           |       | o       |      |

- `request_buffering` on|off
  - on: nginx接收完毕再转发
    1. 客户端网速慢时
    2. 上游并发处理能力低时
    3. 高吞吐量时
  - off: nginx边收边转发
    1. 上游及时处理时
    2. 能降低nginx读写磁盘
    3. 会令next_upstream功能失效

- `hide_header` 隐藏某些上游响应头部 (以下为默认隐藏的头部)
  - Date nginx发送响应头的时间, 由ngx_http_header_filter_module填写
  - Server nginx版本, 由ngx_http_header_filter_module填写
  - X-Pad apache为避免浏览器bug生成的头部, 默认忽略
  - X-Accel- 用于控制nginx行为的响应, 不需要向客户端转发
    - X-Accel-Expire 控制缓存过期时间

## 实战

- [反向代理websocket](nginx-modules.md)

- [打开文件的缓存](nginx-modules.md)
