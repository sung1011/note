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
```
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

## ref
[doc](http://nginx.org/en/docs/beginners_guide.html#conf_structure)
[measurement](http://nginx.org/en/docs/syntax.html)