# nginx配置语法

## 语法

配置文件由`指令`与`指令块`构成  
每条指令以;分号结尾，指令与参数以空格分隔  
指令块以{}大括号将多条指令组织在一起  
include用来引入配置  
以#符号作为注释  
以$符号作为变量  
部分指令的参数支持正则  

## 示例

```bash
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

### 数据结构

```c
variables_hash_bucket_size 变量名大小64B
variables_hash_max_size 变量数量最大值1024个
```

### http_头部名字

得到header中对应key的val 如 `curl localhost -H 'foo=bar'; $http_foo = bar`  
预定义除外 如 `http_host, http_cookie, http_user_agent, http_referer, http_via, http_x_forwarded_for ...`  

## ref

[doc](http://nginx.org/en/docs/beginners_guide.html#conf_structure)
[measurement](http://nginx.org/en/docs/syntax.html)
