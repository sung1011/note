# makefile

    一种常用的构建工具

> 代码变成可执行文件, 叫做编译(compile)；先编译这个, 还是先编译那个(即编译的安排), 叫做构建(build).

## usage

```bash
make # 执行第一个target

make clean # 伪目标
make a.txt # 文件

make -f rules.txt
make --file=rules.txt
```

## 格式

```makefile
<target> : <prerequistes>
[tab] <commands>

```

### 1. arget 目标

```makefile
# `make clean`会寻找当前目录名为clean的文件, 文件存在则不执行逻辑, 可通过声明`.PHPONY`以避免
.PHONY: *

# 可以定义多个目标, 空格分隔
clean clear del:
    rm *.o

dir:
    pwd
```

> 手册 关于target <https://www.gnu.org/software/make/manual/html_node/Special-Targets.html#Special-Targets>

### 2. prerequistes 前置条件

```makefile
.PHONY: result
# 当s1.txt s2.txt不存在, 会先执行其对应的target
result: s1.txt s2.txt
	cat s1.txt > result.txt
	cat s2.txt >> result.txt

s1.txt:
	echo 123 >> s1.txt # 用`>>`证明s1.txt只要存在就不会重复生成

s2.txt:
	echo 456 >> s2.txt
```

### 3. tab

```makefile
# 以 > 替代[tab]
.RECIPEPREFIX = >
all:
>echo Hello, world
```

### 4. commands 命令

```makefile
# 每行cmd在一个单独的shell进程执行
varlost:
	FOO=123;
	echo "$$FOO";       # err; 取不到foo的值
    
# 解决方法1: 换行转移
varlost:
	FOO=123; \
	echo "$$FOO";       # 123

# 解决方法2: 加上.ONESHELL命令
.ONESHELL:
varlost:
	FOO=123;
	echo "$$FOO";       # 123   (mac实测没成功 -.-)
```

## 语法

### 回声 echoing

    make会打印每条命令, 然后再执行, 即回声(echoing)

```makefile
# `@`关闭回声; 不显示pwd和ls命令字符, 只显示命令结果
test:
    pwd;
    # 注释123        # 该注释会有回声
	@# 注释456       # 一般只@注释和@echo命令
    @echo TODO;     # 一般只@注释和@echo命令
```

### 匹配模式

```makefile
# 类似正则匹配; 只作用于 < target > < prerequistes >
%.go: %.mod
    go build
```

### 变量和赋值

```makefile
# make的赋值和调用 $()
txt = hello
test1:
	@echo $(txt)

#shell的赋值和调用; $$ 因为make会对$转义
test2:
	@export foo=hello; echo $$foo   

#make的赋值运算符

    VARIABLE = value    # 在执行时扩展, 允许递归扩展.

    VARIABLE := value   # 在定义时扩展.

    VARIABLE ?= value   # 只有在该变量为空时才设置值.

    VARIABLE += value   # 将值追加到变量的尾端.
```

### 内置变量 Implicit-Variables

```makefile
# $(CC)当前编译器; 内置变量大多C语言相关的, 不常用
output:
    $(CC) -o output input.c
```

> 手册 关于implicit-variables <https://www.gnu.org/software/make/manual/html_node/Implicit-Variables.html>

### 自动变量 Automatic-Variables

```makefile
# $@ 代指< target >; 如make foo的$@就是foo
a.txt:
    touch $@
```

```makefile
# $< 代指第一个< prerequistes >
dest/%.txt: src/%.txt
    @[ -d dest ] || mkdir dest  # 不存在就新建
    cp $< $@        # cp src/%.txt dest/%.txt
```

```makefile
# $? 指代比目标更新的所有前置条件, 之间以空格分隔.比如, 规则为 t: p1 p2, 其中 p2 的时间戳比 t 新, $?就指代p2.
```

```makefile
# $^ 指代所有前置条件, 之间以空格分隔.比如, 规则为 t: p1 p2, 那么 $^ 就指代 p1 p2 .
```

```makefile
# $* 指代匹配符 % 匹配的部分,  比如% 匹配 f1.txt 中的f1 , $* 就表示 f1.
```

```makefile
# $(@D)和$(@F) 分别指向 $@ 的目录名和文件名.比如, $@是 src/input.c, 那么$(@D) 的值为 src , $(@F) 的值为 input.c.
```

```makefile
# $(<D) 和 $(<F) 分别指向 $< 的目录名和文件名
```

> 手册 关于automatic variables <https://www.gnu.org/software/make/manual/html_node/Automatic-Variables.html>

### 函数 Functions

```makefile
# 格式
$(function args)
# 或
${function args}
```

```makefile
# shell 执行shell命令
srcfiles := $(shell echo src/{00..99}.txt)
```

```makefile
# wildcard 替换bash通配符
srcfiles := $(wildcard src/*.txt)
```

```makefile
# subst 文本替换
# $(subst from,to,text)

# 1
$(subst ee,EE,feet on the street)

# 2
comma:= ,
empty:=
space:= $(empty) $(empty) # space变量用两个空变量作为标识符, 当中是一个空格
foo:= a b c
bar:= $(subst $(space),$(comma),$(foo)) # bar is now `a,b,c'.
```

```makefile
# patsubst 匹配替换
# $(patsubst pattern,replacement,text)
$(patsubst %.c,%.o,x.c.c bar.c)

# 简写 变量名 + 冒号 + 后缀名替换规则
min: $(OUTPUT:.js=.min.js) # 变量output中的.js全部替换为.min.js
```

> 手册 关于内置函数 <https://www.gnu.org/software/make/manual/html_node/Functions.html>

## ref

- GNU Make手册 <https://www.gnu.org/software/make/manual/make.html>
- Makefile文件教程 <https://gist.github.com/isaacs/62a2d1825d04437c6f08>
- 阮一峰make命令教程 <http://www.ruanyifeng.com/blog/2015/02/make.html>