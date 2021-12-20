# php.ini

## base

request_slowlog_timeout = 1 慢查询条件  
slowlog = "" 慢查询log目录  
max_execution_time; 报出Fatal Error; 不包含system(), sleep()等系统调用, 数据库处理时间, 比较鸡肋.
max_execution_time = 30  
memory_limit = 8388608 (8M)  
disable_functions = "" 禁用函数, 多个由逗号隔开  
error_reporting  

- Deprecated最低级别错误, 程序继续执行  
- Notice 通知级别的错误 如直接使用未声明变量, 程序继续执行  
- Warning 警告级别的错误, 可能得不到想要的结果, 程序继续执行  
- Fatal error  致命级别错误致命级别错误, 程序不往下执行  
- Parse error 语法解析错误, 最高级别错误, 连其他错误信息也不呈现出来  
- E_USER_相关错误 用户设置的相关错误  

## extension  

- opcache  
  - opcache.enable=1  
  - opcache.memory_consumption=128; OPcache共享内存存储大小.用于存储预编译的opcode  
  - opcache.interned_strings_buffer=8; 字符串驻留的内存量(相同字符串指向同一地址 以节约内存)  
  - opcache.max_accelerated_files=4000; 利用opcache缓存的文件数(需大于你的项目中的所有PHP文件的总和)(find . -type f -print | grep php | wc -l)  
  - opcache.revalidate_freq=600 设置缓存的过期时间(单位是秒),为0的话每次都要检查  
  - opcache.validate_timestamps=0 opcache检测php代码变更, 并重新编译生成opcode的时间间隔(dev=0 online=default)  
  - opcache.fast_shutdown=1 php7.2已删除(推荐=1)  
  - opcache.file_cache=/tmp  
- yac  
  - yac.enable=1  
  - yac.keys_memory_size = 32M  
  - yac.values_memory_size = 128M  
  - yac.compress_threshold = -1  
- snappy 字符串压缩(压缩率约50%)  
- xdebug  
