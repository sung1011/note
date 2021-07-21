# go context

    多个协程交互时, 方便传递数据, 取消协程, 定时取消协程
    协程x和其子协程都会被取消, x的父协程不会被取消

## context 接口

```go
type Context interface {

    // 获取设置的超时自动Cancel的时间点
    Deadline() (deadline time.Time, ok bool)

    // 只读的struct{}, parent context发起了取消请求, 信号量可读, 应进行return退出goroutine
    Done() <-chan struct{}

    // 取消请求的原因error
    Err() error

    // 获取ctx上绑定的值
    Value(key interface{}) interface{}
}

```

## context方法

```go
// parent衍生ctx, 和CancelFunc
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
// parent衍生ctx, 和CancelFunc; 携带超时自动Cancel
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
// parent衍生ctx, 和CancelFunc; 携带超时自动Cancel
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
// parent衍生ctx; 携带k-v
func WithValue(parent Context, key, val interface{}) Context
```

## 源码

### emptyCtx

```go
// An emptyCtx is never canceled, has no values, and has no deadline. It is not
// struct{}, since vars of this type must have distinct addresses.
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
    return
}

func (*emptyCtx) Done() <-chan struct{} {
    return nil
}

func (*emptyCtx) Err() error {
    return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
    return nil
}

func (e *emptyCtx) String() string {
    switch e {
    case background:
        return "context.Background"
    case todo:
        return "context.TODO"
    }
    return "unknown empty Context"
}

var (
    background = new(emptyCtx)
    todo       = new(emptyCtx)
)

// Background returns a non-nil, empty Context. It is never canceled, has no
// values, and has no deadline. It is typically used by the main function,
// initialization, and tests, and as the top-level Context for incoming
// requests.
func Background() Context {
    return background
}

// TODO returns a non-nil, empty Context. Code should use context.TODO when
// it's unclear which Context to use or it is not yet available (because the
// surrounding function has not yet been extended to accept a Context
// parameter).
func TODO() Context {
    return todo
}
```

### valueCtx

TODO

### cancelCtx

TODO

### timerCtx

TODO

## 注意

- 不要把ctx放在结构体中,要以参数的方式进行传递
- 以ctx作为参数的函数方法,应该把ctx作为第一个参数,放在第一位
- 给一个函数方法传递ctx的时候,不要传递nil,如果不知道传递什么,就使用ctx.TODO
- ctx的Value相关方法应该传递必须的数据,不要滥用,什么数据都使用这个传递

## ref

- <https://studygolang.com/articles/23247?fr=sidebar>
- <https://zhuanlan.zhihu.com/p/110085652>
