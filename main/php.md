# php  
  
## 基础  
  
### 浮点数  
- 浮点数运算 `bcadd, bcdiv, bcmod, bcmul, bcpow, bcsub, ...`  
  
## 面向对象 OOP  
  
### [ 概念 ](oop.md)  
### 对象  
生成: 类的实例化  
组成: 属性 方法  
### 特性  
继承  
封装  
多态  
### 继承, final  
### 访问控制  
public  
protected  
private  
### $this, self, parent,  
$this: 当前对象  
self: 当前类  
parent: 父类  
### 引用  
对象的引用相当于别名  
- $obj1 = &$obj; unset($obj); $obj1也会被删除.  
- $obj1 = $obj; unset($obj); $obj1依然存在.  
### 命名空间 namespace  
意义: 解决项目中类, 函数, 常量冲突问题; 别名提高可读性.  
实例  
- 定义 `namespace my\space`  
- 快捷导入 `use my\space\classA as ca`  
- 调用  
    - 类 `new \my\space\classA()`  
    - 别名类 `new ca()`  
    - 方法 `my\space\funcA()`  
    - 常量 `my\space\CONSTA`  
    - 全局方法 `\funcG()`  
- 自动加载  
    - __autoload()  
    - spl_autoload_register()  
### 后期静态绑定  
[ ex1 ](src/php/late_static_bindings1.php)  
[ ex2 ](src/php/late_static_bindings2.php)  
[ ex3 ](src/php/late_static_bindings3.php)  
### 接口  
接口是类的模板, 实现某接口就必须实现接口中定义的所有方法  
接口中所有方法的访问控制必须public  
当参数传递，以松耦合  
实现多个接口接口， 达到组合的效果， 以松耦合  
可以被实现`implements` 和 继承 `extends`  
### 抽象  
接口中有具体实现的方法， 就是抽象  
继承抽象的子类，其访问控制必须更宽松  
只能用来被继承`extends`  
### 魔术方法  
__get()  
__set()  
__isset()  
__unset()  
__call()  
- 调用: 调用不存在的方法  
__callStatic()  
- 调用: 调用不存在的静态方法  
__construct()  
__destruct()  
__toString()  
- 调用: 输出对象  
__invoke()  
- 调用: 把对象当函数执行  
__clone()  
- 调用: clone时  
- 说明: `$obj2 = clone $obj1`会进行浅拷贝（即$obj2是$obj1的拷贝, 但$obj2中的属性若是一个对象$objHang, 其保存的是指针地址, 即$objHang是浅拷贝）， 而深拷贝需要__clone()魔术方法。用以$obj1调用clone时， 内部的handler。  
__sleep()  
- 调用: 序列化 `serialize()`  
__wakeup()  
- 调用: 反序列化 `unserialize()`  
__set_state()  
- 调用: 导出时`var_export()`  

### [ 异常处理 ](phpException.md)
  
  
## 配置  
### php配置  
max_execution_time = 30  
memory_limit = 8388608 (8M)  
disable_functions = "" 禁用函数，多个由逗号隔开  
error_reporting  
- Deprecated最低级别错误，程序继续执行  
- Notice 通知级别的错误 如直接使用未声明变量，程序继续执行  
- Warning 警告级别的错误，可能得不到想要的结果, 程序继续执行  
- Fatal error  致命级别错误致命级别错误，程序不往下执行  
- Parse error 语法解析错误，最高级别错误，连其他错误信息也不呈现出来  
- E_USER_相关错误 用户设置的相关错误  
### php-fpm配置  
- master  
    - pm = static; 静态进程 (2G 50;4G 100;8G 200)  
        - pm.max_children = 300; 静态方式下开启的php-fpm进程数量  
    - pm = dynamic; 动态进程(有额外内存开销)  
        - pm.start_servers = 20; 动态方式下的起始php-fpm进程数量  
        - pm.min_spare_servers = 5; 动态方式下的最小php-fpm进程数量  
        - pm.max_spare_servers = 35; 动态方式下的最大php-fpm进程数量  
    - pm.max_requests = 10240; 每个worker处理多少个请求后会重启该线程 // 由于内存泄漏，泄漏的内存会累计，重启以归还内存  
    - 内存消耗 = max_children * max_requests; 静态进程内存消耗  
    - 内存消耗 = max_spare_servers * max_requests; 动态进程内存消耗  
    - rlimit_files = 1024; 文件打开描述符的rlimit限制, 默认系统值（ulimit -n）(一般要跟系统的同步更改)  
    - 覆盖ini: php_admin_value 如 php_admin_value[ memory_limit ] = 128M; php_admin_value[ date.timezone ] = Asia/Shanghai  
  
## 习题  
  
### 初级  
array_map, array_column, array_filter  
合并数组  
- array_merge(),  
- array_merge_recursive(),  
- array_combine(),  
- $arr1 + $arr2  
数组去重 array_unique(), array_flip()  
private, protect, public  
require, include  
static, global  
$this, self, parent  
常用的str, array操作函数  
  
### 中级  
魔术方法  
-  __construct(), __destruct(), __call(), __get(), __set(), __isset(), __unset(), __toString(), __clone(), __autoload(), __sleep(), __wakeup(), __set_state(), __invoke()  
abstruct, interface  
cgi, fastcgi, php-fpm, swoole  
反射  
迭代器原理  
ioc (DI)  
- 慢查询  
    - request_slowlog_timeout = 1 慢查询条件  
    - slowlog = "" 慢查询log目录  
- 执行时间  
    - request_terminate_timeout = 30s  
extension  
- opcache  
    - opcache.enable=1  
    - opcache.memory_consumption=128; OPcache共享内存存储大小。用于存储预编译的opcode  
    - opcache.interned_strings_buffer=8; 字符串驻留的内存量（相同字符串指向同一地址 以节约内存）  
    - opcache.max_accelerated_files=4000; 利用opcache缓存的文件数（需大于你的项目中的所有PHP文件的总和）(find . -type f -print | grep php | wc -l)  
    - opcache.revalidate_freq=600 设置缓存的过期时间（单位是秒）,为0的话每次都要检查  
    - opcache.validate_timestamps=0 opcache检测php代码变更，并重新编译生成opcode的时间间隔(dev=0 online=default)  
    - opcache.fast_shutdown=1 php7.2已删除（推荐=1）  
    - opcache.file_cache=/tmp  
- yac  
    - yac.enable=1  
    - yac.keys_memory_size = 32M  
    - yac.values_memory_size = 128M  
    - yac.compress_threshold = -1  
- snappy 字符串压缩（压缩率约50%）  
- xdebug  
传值, 引用传值  
  
### 高级  
字符串在手册中介绍，「PHP的字符串是二进制安全的」，这句话怎么理解，为什么是二进制安全？  
字符串连接符.，在内核中有哪些操作？多次.连接，是否会造成内存碎片过多？  
PHP中使用多线程和多进程分别有哪些优缺点？  
线上环境中，PHP进程偶尔会卡死（死锁），请问如何检测本质问题？  
PHP中创建多线程、多进程有哪些方式？互斥信号该如何实现？  
使用cUrl下载大文件时，占用内存太大，有没比较优化的方式？  
写代码来解决多进程/线程同时读写一个文件  
PHP的的这种弱类型变量是怎么实现的？  
垃圾回收    
  
### exp  
[ 锁 ](src/php/php_redis_lock.php)  
