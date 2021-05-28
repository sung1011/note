# php  
  
## 基础  
  
### 浮点数  

- 浮点数运算 `bcadd, bcdiv, bcmod, bcmul, bcpow, bcsub, ...`  

```php
$value = "2.1";
if (is_numeric($value) && intval($value) == $value) {
    echo 'int' . PHP_EOL;
    $value += 0;
} elseif (is_numeric($value) && strpos($value, '.') !== false) {
    echo 'float' . PHP_EOL;
    $value = bcadd($value, 0, 2);
}
var_dump($value);
```
  
## 面向对象 OOP  
  
### [设计模式](dp.md)  

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

### $this, self, parent

$this: 当前对象  
self: 当前类  
parent: 父类  

### 引用  

对象的赋值，本身就是引用。

```php
$obj1 = &$obj; // 改为 `$obj1 = $obj;` 以下结果相同
$obj->foo = 123;
echo $obj1->foo; // 123
unset($obj);
echo $obj1->foo; // 123
```

变量的赋值

```php
$a = 1;
$b = $a; // a b 指向同一个值
$a = 2; // zend写时赋值 相当于新声明了$a
echo $b; // 1
```

变量的引用

```php
$a = 1;
$b = &$a; // a b 指向同一个地址
$a = 2;
echo $b; // 2
```

### 命名空间 namespace  

意义: 解决项目中类, 函数, 常量冲突问题; 别名提高可读性.  
实例  

- 定义 `namespace my\space`  
- 快捷导入 `use my\space\classA as ca`  
- 调用  
  - 类 `new \my\space\classA()`  
  - 别名类 `new ca()`  
  - 方法 `my\space\funcA()`  
  - 常量 `my\space\CONSTA`  
  - 全局方法 `\funcG()`  
- 自动加载  
  - __autoload()  
  - spl_autoload_register()  

### 后期静态绑定  

[ex1](src/php/late_static_bindings1.php)  
[ex2](src/php/late_static_bindings2.php)  
[ex3](src/php/late_static_bindings3.php)  

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

- __get()  
- __set()  
- __isset()  
- __unset()  
- __call() 调用不存在的方法  
- __callStatic() 调用不存在的静态方法  
- __construct()  
- __destruct()  
- __toString() 输出对象  
- __invoke() 把对象当函数执行  
- __clone()  
  - 调用: clone时  
  - 说明: `$obj2 = clone $obj1`会进行浅拷贝（即$obj2是$obj1的拷贝, 但$obj2中的属性若是一个对象$objHang, 其保存的是指针地址, 即`$objHang`是浅拷贝）， 而深拷贝需要__clone()魔术方法。用以`$obj1`调用clone时， 内部的handler。  
- __sleep() 序列化 `serialize()`  
- __wakeup() 反序列化 `unserialize()`  
- __set_state() 导出时`var_export()`  

### [异常处理](phpException.md)
  
## 配置  

### [php.ini](php-ini.md)

### [php-fpm](php-fpm.md)

## 实战  
  
### 初级  

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
  
### 中级  

- 魔术方法  
- php7为何比php5性能高
  - 变量存储字节减小，减少内存占用，提升变量操作速度
  - 改善数组结构，数组元素和hash映射表被分配在同一块内存里，降低了内存占用、提升了 cpu 缓存命中率
  - 改进了函数的调用机制，通过优化参数传递的环节，减少了一些指令，提高执行效率
- abstruct, interface  
- cgi, fastcgi, php-fpm, swoole  
- 反射  
- 迭代器原理  
- ioc (DI)  

### 高级  

- 字符串在手册中介绍，「PHP的字符串是二进制安全的」，这句话怎么理解，为什么是二进制安全？  --- php zval str数据结构`len + *val`; 不会将"abc \0 def"的`\0`错误的认为是结尾符而忽略后边的 def（而C语言会）  
- 字符串连接符.，在内核中有哪些操作？多次.连接，是否会造成内存碎片过多？  --- 会
- PHP中使用多线程和多进程分别有哪些优缺点？  
- 线上环境中，PHP进程偶尔会卡死（死锁），请问如何检测本质问题？  --- `https://blog.csdn.net/u010412301/article/details/52776584`
- PHP中创建多线程、多进程有哪些方式？互斥信号该如何实现？  --- `https://blog.csdn.net/ZHANG_TIMI/article/details/78342722`
- 使用cUrl下载大文件时，占用内存太大，有没比较优化的方式？--- `curl_setopt($ch, CURLOPT_WRITEFUNCTION, function($ch ,$str) use (&$flag){})`  
- 写代码来解决多进程/线程同时读写一个文件  --- `flock`
- PHP的的这种弱类型变量是怎么实现的？  
- 垃圾回收  

### exp  

[锁](src/php/php_redis_lock.php)  

### ref

- [Fpm启动机制及流程分析———详细](http://www.mamicode.com/info-detail-2625546.html)
- [PHP内核探索之变量（7）- 不平凡的字符串](https://blog.csdn.net/fvjuerh/article/details/68946281)
- [php下载大文件的方法](https://blog.csdn.net/dengjiexian123/article/details/53057593)
- [亿级用户 PC 主站的 PHP7 升级实践](https://www.infoq.cn/article/practice-of-PHP7-upgrade-for-the-PC-master-station)
- [phpfpm运行原理](https://blog.csdn.net/sinat_38804294/article/details/94393621)
- [PHP传值和传引用区别](https://blog.csdn.net/chengjianghao/article/details/81507752)
