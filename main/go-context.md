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
// emptyCtx 不能取消, 无超时时间, 不能存储信息, 只作为ctx树的根节点
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

// emptyCtx的实例化方法
// 通常用于主函数, 初始化, 测试中
func Background() Context {
    return background
}

// emptyCtx的实例化方法
// 通常在不确定使用什么时使用
func TODO() Context {
    return todo
}
```

### valueCtx

```go
type valueCtx struct {
    // 父节点
    Context
    // 当前节点kv
    key, val interface{}
}

// 获取ctx链路上key对应的value, 如果当前ctx没有, 就沿ctx向上查找直到根节点
func (c *valueCtx) Value(key interface{}) interface{} {
    if c.key == key {
        return c.val
    }
    return c.Context.Value(key)
}
```

### WithValue

```go
// 向ctx添加kv; 以当前ctx作为父节点, 创建新的ctx子节点并附加kv
func WithValue(parent Context, key, val interface{}) Context {
    if key == nil {
        panic("nil key")
    }
    if !reflect.TypeOf(key).Comparable() {
        panic("key is not comparable")
    }
    return &valueCtx{parent, key, val}
}
```

### cancelCtx

```go
type cancelCtx struct {
    // 父节点
    Context

    mu       sync.Mutex            // protects following fields
    // 通道传递关闭信号
    done     chan struct{}         // created lazily, closed by first cancel call
    // 子节点
    children map[canceler]struct{} // set to nil by the first cancel call
    // 取消原因
    err      error                 // set to non-nil by the first cancel call
}

type canceler interface {
    cancel(removeFromParent bool, err error)
    Done() <-chan struct{}
}

func (c *cancelCtx) Done() <-chan struct{} {
    c.mu.Lock()
    if c.done == nil {
        c.done = make(chan struct{})
    }
    d := c.done
    c.mu.Unlock()
    return d
}

func (c *cancelCtx) Err() error {
    c.mu.Lock()
    err := c.err
    c.mu.Unlock()
    return err
}

// 取消 
// 设置取消原因, 发送取消信号, 遍历取消子节点, 从父节点移除自己
func (c *cancelCtx) cancel(removeFromParent bool, err error) {
    if err == nil {
        panic("context: internal error: missing cancel error")
    }
    c.mu.Lock()
    if c.err != nil {
        c.mu.Unlock()
        return // already canceled
    }
    // 设置取消原因
    c.err = err
    // 设置一个关闭的channel或者将done channel关闭，用以发送关闭信号
    if c.done == nil {
        c.done = closedchan
    } else {
        close(c.done)
    }
    // 将所有子节点context依次取消
    for child := range c.children {
        // NOTE: acquiring the child's lock while holding parent's lock.
        child.cancel(false, err)
    }
    c.children = nil
    c.mu.Unlock()

    // 将当前context节点从父节点上移除
    if removeFromParent {
        removeChild(c.Context, c)
    }
}
```

### WithCancel

```go
type CancelFunc func()

func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
    c := newCancelCtx(parent)
    propagateCancel(parent, &c)
    return &c, func() { c.cancel(true, Canceled) }
}

// newCancelCtx returns an initialized cancelCtx.
func newCancelCtx(parent Context) cancelCtx {
    // 将parent作为父节点context生成一个新的子节点
    return cancelCtx{Context: parent}
}

func propagateCancel(parent Context, child canceler) {
    if parent.Done() == nil {
        // parent.Done()返回nil表明父节点以上的路径上没有可取消的context
        return // parent is never canceled
    }
    // 获取最近的类型为cancelCtx的祖先节点
    if p, ok := parentCancelCtx(parent); ok {
        p.mu.Lock()
        if p.err != nil {
            // parent has already been canceled
            child.cancel(false, p.err)
        } else {
            if p.children == nil {
                p.children = make(map[canceler]struct{})
            }
            // 将当前子节点加入最近cancelCtx祖先节点的children中
            p.children[child] = struct{}{}
        }
        p.mu.Unlock()
    } else {
        go func() {
            select {
            case <-parent.Done():
                child.cancel(false, parent.Err())
            case <-child.Done():
            }
        }()
    }
}

func parentCancelCtx(parent Context) (*cancelCtx, bool) {
    for {
        switch c := parent.(type) {
        case *cancelCtx:
            return c, true
        case *timerCtx:
            return &c.cancelCtx, true
        case *valueCtx:
            parent = c.Context
        default:
            return nil, false
        }
    }
}
```

## 注意

- 不要把ctx放在结构体中,要以参数的方式进行传递
- 以ctx作为参数的函数方法,应该把ctx作为第一个参数,放在第一位
- 给一个函数方法传递ctx的时候,不要传递nil,如果不知道传递什么,就使用ctx.TODO
- ctx的Value相关方法应该传递必须的数据,不要滥用,什么数据都使用这个传递

## ref

- <https://zhuanlan.zhihu.com/p/110085652>
- <https://studygolang.com/articles/23247?fr=sidebar>
