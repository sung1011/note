# redis对象引用计数

## 对象回收

```c
void incrRefCount(robj *o); // 对象的引用计数+1
void decrRefCount(robj *o); // 对象的引用计数-1
robj resetRefCount(robj *o); // 对象的引用计数清零(但不释放对象)
```

## 对象共享

server初始化时会创建一些共享对象，供自定义的值对象共享。

```c
void createSharedObjects(void); // server初始化时创建共享对象
// 整数类: 0 1 2 ... 9999
// 字符类: \r\n; +PONG\r\n; +OK\r\n
// 报错类: -ERR no such key\r\n
// 命令类: DEL; RPOP
// 特别标记类: minstring maxstring
```

> 若某key得值对象是共享对象，则`object refcount <key>`返回2147483647(INT_MAX)。
