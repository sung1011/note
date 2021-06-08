# 单例模式

      确保一个类只有一个实例, 提供一个全局访问点

## 特点

- 节省内存
- 线程安全

## 实现

- `饿汉`
  - 类加载时创建实例, 创建过程是线程安全
    - 服务启动时慢, 执行时快
    - fail-fast 若有问题启动时能尽早暴露
  - 不支持延迟加载

- `懒汉`
  - getInstance()时创建, 线程不安全需要加锁, 并发能力比饿汉弱
  - 延迟加载

## 细节

- 构造函数是private访问权限
- 考虑对象创建时是否线程安全
- 是否支持延迟加载
- getInstance()性能问题(是否加锁)

## issue

- 不支持有参数的构造函数
- 可测性不友好
  - 不利于mock, 不同用例不易修改单例

## code

- [go 饿汉](../script/go/dp/singleton-hungry.go)
- [go 懒汉](../script/go/dp/singleton-lazy.go)

- [php](src/php_design_patterns/singleton/mysql_singleton.php)

## 场景

- DB连接数据库 (饿汉)
- 加载model实例 (懒汉)
