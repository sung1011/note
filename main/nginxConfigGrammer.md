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

## 模块
`event module`	搭建了独立于操作系统的事件处理机制的框架，及提供了各具体事件的处理。包括ngx_events_module， ngx_event_core_module和ngx_epoll_module等。nginx具体使用何种事件处理模块，这依赖于具体的操作系统和编译选项。  
`phase handler`	此类型的模块也被直接称为handler模块。主要负责处理客户端请求并产生待响应内容，比如ngx_http_static_module模块，负责客户端的静态页面请求处理并将对应的磁盘文件准备为响应内容输出。  
`output filter`	也称为filter模块，主要是负责对输出的内容进行处理，可以对输出进行修改。例如，可以实现对输出的所有html页面增加预定义的footbar一类的工作，或者对输出的图片的URL进行替换之类的工作。  
`upstream`	upstream模块实现反向代理的功能，将真正的请求转发到后端服务器上，并从后端服务器上读取响应，发回客户端。upstream模块是一种特殊的handler，只不过响应内容不是真正由自己产生的，而是从后端服务器上读取的。  
`load-balancer`	负载均衡模块，实现特定的算法，在众多的后端服务器中，选择一个服务器出来作为某个请求的转发服务器。  

## 作用域/指令上下文
`main`	nginx在运行时与具体业务功能（比如http服务或者email服务代理）无关的一些参数，比如工作进程数，运行的身份等。  
`http`  与提供http服务相关的一些配置参数。例如：是否使用keepalive啊，是否使用gzip进行压缩等。  
`server`	http服务上支持若干虚拟主机。每个虚拟主机一个对应的server配置项，配置项里面包含该虚拟主机相关的配置。在提供mail服务的代理时，也可以建立若干server.每个server通过监听的地址来区分。  
`location`	http服务中，某些特定的URL对应的一系列配置项。  
`mail`	实现email相关的SMTP/IMAP/POP3代理时，共享的一些配置项（因为可能实现多个代理，工作在多个监听地址上）。  

## ref
[ doc ](http://nginx.org/en/docs/beginners_guide.html#conf_structure)