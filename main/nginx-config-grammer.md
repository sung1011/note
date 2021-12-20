# nginx配置语法

## 语法

```txt
配置文件由`指令`与`指令块`构成  
每条指令以;分号结尾, 指令与参数以空格分隔  
指令块以{}大括号将多条指令组织在一起  
include用来引入配置  
以#符号作为注释  
以$符号作为变量  
部分指令的参数支持正则  
```

## 示例

```conf
http {
    include       mime.types;
    #access_log  logs/access.log  main;

    server {
        listen      80;
        server_name example.org www.example.org;
        root        /data/www;

        location / {
            index   index.html index.php;
        }

        location ~* \.(gif|jpg|png)$ {
            expires 30d;
        }

        location ~ \.php$ {
            fastcgi_pass  localhost:9000;
            fastcgi_param SCRIPT_FILENAME
                        $document_root$fastcgi_script_name;
            include       fastcgi_params;
        }
    }
}
```

## 变量

### 预定义变量

|                     |                                                                                                                         |
| ------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| $arg_PARAMETER      | GET请求中变量名PARAMETER参数的值.                                                                                      |
| $args               | 这个变量等于GET请求中的参数.例如, foo=123&bar=blahblah;这个变量只可以被修改                                            |
| $binary_remote_addr | 二进制码形式的客户端地址.                                                                                              |
| $body_bytes_sent    | 传送页面的字节数                                                                                                        |
| $content_length     | 请求头中的Content-length字段.                                                                                          |
| $content_type       | 请求头中的Content-Type字段.                                                                                            |
| $cookie_COOKIE      | cookie COOKIE的值.                                                                                                     |
| $document_root      | 当前请求在root指令中指定的值.                                                                                          |
| $document_uri       | 与$uri相同.                                                                                                            |
| $host               | 请求中的主机头(Host)字段, 如果请求中的主机头不可用或空, 则为处理请求的server名称.                                      |
| $hostname           | 机器名使用 gethostname系统调用的值                                                                                      |
| $http_HEADER        | HTTP请求头中的内容, HEADER为HTTP请求中的内容转为小写, -变为_(破折号变为下划线), 例如: $http_user_agent(Uaer-Agent的值); |
| $sent_http_HEADER   | HTTP响应头中的内容, HEADER为HTTP响应中的内容转为小写, -变为_(破折号变为下划线), 例如: $sent_http_cache_control.        |
| $is_args            | 如果$args设置, 值为"?", 否则为"".                                                                                      |
| $limit_rate         | 这个变量可以限制连接速率.                                                                                              |
| $nginx_version      | 当前运行的nginx版本号.                                                                                                 |
| $query_string       | 与$args相同.                                                                                                           |
| $remote_addr        | 客户端的IP地址.                                                                                                        |
| $remote_port        | 客户端的端口.                                                                                                          |
| $remote_user        | 已经经过Auth Basic Module验证的用户名.                                                                                 |
| $request_filename   | 当前连接请求的文件路径, 由root或alias指令与URI请求生成.                                                                |
| $request_body       | 这个变量(0.7.58+)包含请求的主要信息.在使用proxy_pass或fastcgi_pass指令的location中比较有意义.                       |
| $request_body_file  | 客户端请求主体信息的临时文件名.                                                                                        |
| $request_completion | 如果请求成功, 设为"OK"；如果请求未完成或者不是一系列请求中最后一部分则设为空.                                          |
| $request_method     | 这个变量是客户端请求的动作, 通常为GET或POST.                                                                           |
| $request_uri        | 这个变量等于包含一些客户端请求参数的原始URI, 它无法修改, 请查看$uri更改或重写URI.                                      |
| $scheme             | 所用的协议, 比如http或者是https, 比如rewrite ^(.+)$ $scheme://example.com$1 redirect;                                   |
| $server_addr        | 服务器地址, 在完成一次系统调用后可以确定这个值, 如果要绕开系统调用, 则必须在listen中指定地址并且使用bind参数.          |
| $server_name        | 服务器名称.                                                                                                            |
| $server_port        | 请求到达服务器的端口号.                                                                                                |
| $server_protocol    | 请求使用的协议, 通常是HTTP/1.0或HTTP/1.1.                                                                              |
| $uri                | 求中的当前URI(不带请求参数, 参数位于args)                                                                               |

### 数据结构

- variables_hash_bucket_size  变量名大小64B
- variables_hash_max_size     变量数量最大值1024个

### http_头部名字

```bash
直接获取header中对应key的val. 
如 `curl localhost -H 'foo=bar'; # 则 $http_foo = bar`  

# 预定义除外 如 `http_host, http_cookie, http_user_agent, http_referer, http_via, http_x_forwarded_for ...`  
```

## ref

- [tool](https://www.digitalocean.com/community/tools/nginx)
- [doc](http://nginx.org/en/docs/beginners_guide.html#conf_structure)
- [measurement](http://nginx.org/en/docs/syntax.html)
