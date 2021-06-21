# 命令模式

    将命令请求封装成一个对象, 可以将不同请求来进行参数化

## 特点

## 实现

- 接口 函数封装成对象, 实现相同接口, 使其可以通过参数传递
- 参数方法 函数可以直接当做参数传递

## code

- [go 接口](../script/go/dp/command.go)
- [go 参数方法](../script/go/dp/command-func.go)

- [php](src/php_design_patterns/command/command.php)

## 场景

- 控制命令的执行. 如: 异步,延迟,排队,撤销,重做,记录