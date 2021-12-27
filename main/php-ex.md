# php 实战

## 初级  

- map
  - array_map
  - array_column
  - array_filter  

- 合并数组  
  - array_merge(),  // [1,2] [a,b] -> [1,2,a,b]
  - array_merge_recursive(),  //
  - array_combine(),  // [1,2] [a,b] -> [1=>a,2=>b]
  - $arr1 + $arr2  // [1,2], [a,b,c] -> [1,2,c]
  - $arr1 + $arr2  // [a => 6, z => 1], [a => 1, b => 2, c => 3] -> [a => 6, z => 1, b => 2, c => 3]

- 数组去重
  - array_unique()
  - array_flip()  

- 访问控制符
  - private
  - protect
  - public  

- 包含
  - require
  - include  

- 静态
  - static
  - global  

- 访问
  - $this
  - self
  - parent  
  
## 中级  

- 魔术方法  
- php7为何比php5性能高
  - 变量存储字节减小, 减少内存占用, 提升变量操作速度
  - 改善数组结构, 数组元素和hash映射表被分配在同一块内存里, 降低了内存占用、提升了 cpu 缓存命中率
  - 改进了函数的调用机制, 通过优化参数传递的环节, 减少了一些指令, 提高执行效率
- abstruct, interface  
- cgi, fastcgi, php-fpm, swoole  
- 反射  
- 迭代器原理  
- ioc (DI)  

## 高级  

- 字符串在手册中介绍, 「PHP的字符串是二进制安全的」, 这句话怎么理解, 为什么是二进制安全？  --- php zval str数据结构`len + *val`; 不会将"abc \0 def"的`\0`错误的认为是结尾符而忽略后边的 def(而C语言会)  
- 字符串连接符., 在内核中有哪些操作？多次.连接, 是否会造成内存碎片过多？  --- 会
- PHP中使用多线程和多进程分别有哪些优缺点？  
- 线上环境中, PHP进程偶尔会卡死(死锁), 请问如何检测本质问题？  --- `https://blog.csdn.net/u010412301/article/details/52776584`
- PHP中创建多线程、多进程有哪些方式？互斥信号该如何实现？  --- `https://blog.csdn.net/ZHANG_TIMI/article/details/78342722`
- 使用cUrl下载大文件时, 占用内存太大, 有没比较优化的方式？--- `curl_setopt($ch, CURLOPT_WRITEFUNCTION, function($ch ,$str) use (&$flag){})`  
- 写代码来解决多进程/线程同时读写一个文件  --- `flock`
- PHP的的这种弱类型变量是怎么实现的？  
- 垃圾回收  

## exp  

[锁](src/php/php_redis_lock.php)  

