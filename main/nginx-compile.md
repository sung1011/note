# nginx编译

## 下载解压

```bash
wget http://nginx.org/download/nginx-1.16.0.tar.gz 

tar -xzf nginx-1.16.0.tar.gz
```

## 源码目录结构

```bash
.
├── auto            # 自动检测系统环境以及编译相关的脚本
│   ├── cc          # 关于编译器相关的编译选项的检测脚本
│   ├── lib         # nginx编译所需要的一些库的检测脚本
│   ├── os          # 与平台相关的一些系统参数与系统调用相关的检测
│   └── types       # 与数据类型相关的一些辅助脚本
├── conf            # 存放默认配置文件, 在make install后, 会拷贝到安装目录中去
├── contrib         # 存放一些实用工具, 如geo配置生成工具(geo2nginx.pl), nginx-conf高亮工具
├── html            # 存放默认的网页文件, 在make install后, 会拷贝到安装目录中去
├── man             # nginx的man手册
└── src             # 存放nginx的源代码
    ├── core        # nginx的核心源代码, 包括常用数据结构的定义, 以及nginx初始化运行的核心代码如main函数
    ├── event       # 对系统事件处理机制的封装, 以及定时器的实现相关代码
    │   └── modules # 不同事件处理方式的模块化, 如select、poll、epoll、kqueue等
    ├── http        # nginx作为http服务器相关的代码
    │   └── modules # 包含http的各种功能模块
    ├── mail        # nginx作为邮件代理服务器相关的代码
    ├── misc        # 一些辅助代码, 测试c++头的兼容性, 以及对google_perftools的支持
    └── os          # 主要是对各种不同体系统结构所提供的系统函数的封装, 对外提供统一的系统调用接口
```

## 编译安装

### 步骤

1. 依赖

   ```bash
   yum install gcc gcc-c++ automake pcre pcre-devel zlip zlib-devel openssl openssl-devel  
   ```

   ```bash
       gcc为GNU Compiler Collection的缩写, 可以编译C和C++源代码等, 它是GNU开发的C和C++以及其他很多种语言 的编译器(最早的时候只能编译C, 后来很快进化成一个编译多种语言的集合, 如Fortran、Pascal、Objective-C、Java、Ada、 Go等.)
       gcc 在编译C++源代码的阶段, 只能编译 C++ 源文件, 而不能自动和 C++ 程序使用的库链接(编译过程分为编译、链接两个阶段, 注意不要和可执行文件这个概念搞混, 相对可执行文件来说有三个重要的概念: 编译(compile)、链接(link)、加载(load).源程序文件被编译成目标文件, 多个目标文件连同库被链接成一个最终的可执行文件, 可执行文件被加载到内存中运行).因此, 通常使用 g++ 命令来完成 C++ 程序的编译和连接, 该程序会自动调用 gcc 实现编译.
       gcc-c++也能编译C源代码, 只不过把会把它当成C++源代码, 后缀为.c的, gcc把它当作是C程序, 而g++当作是c++程序；后缀为.cpp的, 两者都会认为是c++程序, 注意, 虽然c++是c的超集, 但是两者对语法的要求是有区别的.
       automake是一个从Makefile.am文件自动生成Makefile.in的工具.为了生成Makefile.in, automake还需用到perl, 由于automake创建的发布完全遵循GNU标准, 所以在创建中不需要perl.libtool是一款方便生成各种程序库的工具.
       pcre pcre-devel: 在Nginx编译需要 PCRE(Perl Compatible Regular Expression), 因为Nginx 的Rewrite模块和HTTP 核心模块会使用到PCRE正则表达式语法.
       zlip zlib-devel: nginx启用压缩功能的时候, 需要此模块的支持.
     openssl openssl-devel: 开启SSL的时候需要此模块的支持.
   ```

2. 配置 `./configure --prefix=<path>`

   - objs目录存放临时文件
   - ./configure --help  查看配置参数

   ```bash
   --prefix 安装路径  
   --with-xxx 指定添加模块 (默认不添加)  
   --without-xxx 指定去除模块 (默认会添加)  
   --with-xxx=dynamic 指定添加动态模块 (默认不添加)(使用时需在nginx.conf中引用`load_module modules/xxx.so`)
   ```

3. 编译 `make`

   - objs目录存放临时生成的中间文件, 动态模块(.so)文件

4. 安装 `make install`  

## 实战

### https

- `./configure ... --with-http_ssl_module --with-http_v2_module` 必要SSL模块
- `certbot --nginx --nginx-ctl /opt/nginx/sbin/nginx --nginx-server-root /opt/nginx/conf -d "tickles.cn" -d "*.tickles.cn"` certbot配置泛域名证书; 在服务商新增TXT; 在服务商新增 * 的二级域名
- `certbot renew` 更新过期的证书 (可加进crontab)

## ref

- doc <http://nginx.org/en/docs/install.html>
- Nginx安装 <http://www.nginx.cn/install>
