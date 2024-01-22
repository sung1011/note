# go sync.Once

`sync.Once` 用于执行只需要执行一次的操作，比如只需要初始化一次的操作。

## 方法

```go
func (o *Once) Do(f func()) // 只有第一次调用Do时才会执行f
```

## 用法

```go
// once方式 只执行一次
var defaultCacheOnce sync.Once
var defaultCache string

func main() {
	defaultCacheOnce.Do(func() {
		defaultCache = "hello"
	})
	fmt.Println("defaultCache: ", defaultCache) // hello
	defaultCacheOnce.Do(func() {
		defaultCache = "1234" // 不会执行
	})
	fmt.Println("defaultCache: ", defaultCache) // hello
    // 重置 defaultCacheOnce
	defaultCacheOnce = sync.Once{}
	defaultCacheOnce.Do(func() {
		defaultCache = "1234"
	})
	fmt.Println("defaultCache: ", defaultCache) // 1234
}
```

```go
// mutex方式 每次都竞争锁, 浪费资源
var connMu sync.Mutex
var conn net.Conn

func getConn() net.Conn {
	connMu.Lock()
	defer connMu.Unlock()
	if conn != nil {
		return conn
	}
	conn, _ = net.DialTimeout("tcp", "baidu.com:80", 10*time.Second)
	return conn
}
```

## 实现

```go
// 简单实现
// issue: 如果f()执行很慢, 会导致其他goroutine虽然看到done, 但f()还没执行完, 拿不到对应的结果
type Once struct {
    done uint32
}
func (o *Once) Do(f func()) {
    if !atomic.CompareAndSwapUint32(&o.done, 0, 1) {
        return
    }
    f()
}
```

```go
// 首次上锁, 并双检查
type Once struct {
	done uint32
	m    Mutex
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 0 {
		o.doSlow(f)
	}
}
func (o *Once) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 { // 双检查
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
```