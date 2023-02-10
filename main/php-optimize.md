# php optimize

## 写时复制 cow

```php
$a = range(0, 1000);
var_dump(memory_get_usage()); // n
$b = $a;
var_dump(memory_get_usage()); // n; 内存没涨, 共用同一块内存空间
$a = range(10, 2000);
var_dump(memory_get_usage()); // n++; 内存涨了, 开辟了新内存; 写时复制cow


$a = range(0, 1000);
var_dump(memory_get_usage()); // n
$b = $a;
var_dump(memory_get_usage()); // n; 内存没涨, 共用同一块内存空间
$a = range(10, 2000);
var_dump(memory_get_usage()); // n; 内存没涨
```

## JIT Just-In-Time Compilation

```js
# php with opcache

                    ---- `Zend Compiler`  
                    |       Opcodes: #1 ADD $a, $b, ~0 #2 ASSIGN $c, ~0
                    |         `Optimizer`
                    |            Optimized Opcodes: #1 ADD $a, $b, $c
                    |                        |
    PHP ----> `Zend Opcache` ---------> Opcodes Cache ----> `Zend VM` ----> Opcode Handlers ----> X86 CPU
  # $c=$a+$b                                                              #1 ZEND_ADD_HANDLER

# php with opcache plus JIT
                    ---- `Zend Compiler`  
                    |       Opcodes: #1 ADD $a, $b, ~0 #2 ASSIGN $c, ~0
                    |         `Optimizer`
                    |            Optimized Opcodes: #1 ADD $a, $b, $c
                    |              `JIT Compiler`
                    |                Machine Code: #0:41 01 c8 add %rcx %rax
                    |                        |
    PHP ----> `Zend Opcache` ------------> X86 CPU
  # $c=$a+$b                            


# Opcache会做opcode层面的优化，比如图中的俩条opcode合并为一条

# JIT在Opcache优化之后的基础上，再次优化，直接生成机器码

# PHP8的JIT是在Opcache之中提供的, 目前只支持x86架构的CPU

# JIT是在Opcache优化的优化基础之上进行优化的，不是替代
```