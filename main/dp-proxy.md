# 代理模式

    用代理类来给原始类附加访问控制和边缘功能

## 特点

- `高扩展性`
- `不改变原有接口`

## 实现

- `静态代理` 代理类重新实现原始类中的所有方法

> 新建proxy类包了一下原始类, 重新实现其功能

- `动态代理` 运行时动态创建原始类对应的代理类, 一般采用反射实现, 性能差

## code

- [go 静态代理](src/go/dp/proxy-static.go)
- [go 动态代理](src/go/dp/proxy-dynamic.go)

- [php](src/php_design_patterns/proxy/proxy.php)

## 场景

- RPC
- cache
- 非功能性需求
  - 监控
  - 统计
  - 鉴权
  - 限流