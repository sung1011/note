# go panic

## [code](src/go/basic/panic_test.go)

## 数据结构

```go
// runtime/runtime2.go
type _panic struct {
    argp      unsafe.Pointer
    arg       interface{}    // panic 的参数
    link      *_panic        // 链接下一个panic结构体; 与defer类似, 是个单向链表
    recovered bool           // 是否恢复
    aborted   bool           // the panic was aborted
}

// goroutine
type g struct {
    // ...
    _panic         *_panic // panic 链表
    _defer         *_defer // defer 链表
    // ...
}

func gorecover(argp uintptr) interface{} {
    // 只处理 gp._panic 链表最新的这个 _panic；
    gp := getg()
    p := gp._panic
    if p != nil && !p.recovered && argp == uintptr(p.argp) {
        p.recovered = true
        return p.arg
    }
    return nil
}
```

## 触发

- 主动 调用panic()
- 被动 编译器报错. 如 除0操作
- 被动 内核发送给进程信号触发. 如 非法寻址

## 基础使用

    recover()必须放在defer中
    并在panic之前执行

```go
func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("error is", e)
		}
	}()
	panic("foo")
}
```

## 注意事项

- g1触发panic(), g2也会挂掉
- g1触发panic(), g2无法recover()
- 对于服务, 每个g都要有自己的recover()
- 不要将panic()作为错误返回给调用者
- 大多数错误, 当做断言在使用, 即触发断言, 代表可能有bug; 而error是预定义的错误

## ref

- 常见panic <https://www.jianshu.com/p/d7a596aba481>
- 底层 <https://blog.csdn.net/qcrao/article/details/120092891>
- Go多协程并发的异常处理 <https://segmentfault.com/a/1190000023691221>
