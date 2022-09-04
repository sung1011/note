# php oop

## [设计模式](dp.md)  

## 对象  

    生成: 类的实例化  
    组成: 属性 方法  

## 特性  

    继承  
    封装  
    多态  

## 继承, final  

## 访问控制  

    public  
    protected  
    private  

## $this, self, parent

    $this: 当前对象  
    self: 当前类  
    parent: 父类  

## 引用  

    对象的赋值, 本身就是引用.

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

## 命名空间 namespace  

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

## 后期静态绑定  

- [ex1](src/php/late_static_bindings1.php)  
- [ex2](src/php/late_static_bindings2.php)  
- [ex3](src/php/late_static_bindings3.php)  

## 接口  

    接口是类的模板, 实现某接口就必须实现接口中定义的所有方法  
    接口中所有方法的访问控制必须public  
    当参数传递, 以松耦合  
    实现多个接口接口,  达到组合的效果,  以松耦合  
    可以被实现`implements` 和 继承 `extends`  

## 抽象  

    接口中有具体实现的方法,  就是抽象  
    继承抽象的子类, 其访问控制必须更宽松  
    只能用来被继承`extends`  

## 魔术方法  

```bash
__get()  
__set()  
__isset()  
__unset()  
__call() 调用不存在的方法  
__callStatic() 调用不存在的静态方法  
__construct()  
__destruct()  
__toString() 输出对象  
__invoke() 把对象当函数执行  
__clone()  
__sleep() 序列化 `serialize()`  
__wakeup() 反序列化 `unserialize()`  
__set_state() 导出时`var_export()`  
```

> clone: `$obj2 = clone $obj1`会进行浅拷贝(即$obj2是$obj1的拷贝, 但$obj2中的属性若是一个对象$objHang, 其保存的是指针地址, 即`$objHang`是浅拷贝), 而深拷贝需要__clone()魔术方法.用以`$obj1`调用clone时, 内部的handler

