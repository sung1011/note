# go sync.Mutex

## 场景

    共享资源: 避免并发的读写共享资源时, 造成数据竞争, 从而导致数据不一致的问题

## feature

### `临界区 critical section`

    Lock()和Unlock()之间的代码段称为资源的临界区, 临界区的代码同时间只有一个goroutine会执行.

### `不关联goroutine`

    可以在goroutine加锁, 另一个goroutine解锁; 不可重入

## Mutex

```go
// 数据结构; 源码
type Mutex struct {
	state int32
	sema uint32
}
```

```go
// 常规使用
type mytype struct {
	m sync.Mutex // 内置; 放在保护的字段(i)上面
	i int
}

func main() {
    ...
    mytype.Lock()
    mytype.i++
    mytype.Unlock()
    ...
}
```

## 易错场景

### 死锁

```go
// 因争夺共享资源而处于一种互相等待的状态
```

### copy已使用的mutex

```go
type Counter struct {
	sync.Mutex
	Count int
}

func main() {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo(c) // 复制锁
}

// 这里Counter的参数是通过复制的方式传入的, 会复制锁状态, 造成死锁
func foo(c Counter) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}
```

> go vet main.go可以检查copy问题

### 重入

```go
// mutex 是不可重入的 (数据结构中没记录gorountine id, 所以并不知道是不是同一个goroutine重入)
func main() {
    var mu sync.Mutex
    mu.Lock()
    mu.Lock() // 重入则报错(deadlock)
    mu.Unlock()
    mu.Unlock()
}
```

#### 改写为可重入

```go
// 方案1. 通过记录goroutine id, 判断是否是同一个goroutine重入
// RecursiveMutex 包装一个Mutex,实现可重入
type RecursiveMutex struct {
	sync.Mutex
	owner     int64 // 当前持有锁的goroutine id
	recursion int32 // 这个goroutine 重入的次数
}

func (m *RecursiveMutex) Lock() {
	gid := goid.Get()
	// 如果当前持有锁的goroutine就是这次调用的goroutine,说明是重入
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	// 获得锁的goroutine第一次调用，记录下它的goroutine id,调用次数加1
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1
}
func (m *RecursiveMutex) Unlock() {
	gid := goid.Get()
	// 非持有锁的goroutine尝试释放锁，错误的使用
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.owner, gid))
	}
	// 调用次数减1
	m.recursion--
	if m.recursion != 0 { // 如果这个goroutine还没有完全释放，则直接返回
		return
	}
	// 此goroutine最后一次调用，需要释放锁
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}
```

```go
// 方案2. 通过记录token, 判断是否是同一个goroutine重入
// Token方式的递归锁
type TokenRecursiveMutex struct {
sync.Mutex
token int64
recursion int32
}
// 请求锁，需要传入token
type TokenRecursiveMutex struct {
	sync.Mutex
	token     int64
	recursion int32
}

// 请求锁，需要传入token
func (m *TokenRecursiveMutex) Lock(token int64) {
	if atomic.LoadInt64(&m.token) == token { //如果传入的token和持有锁的token一致，
		m.recursion++
		return
	}
	m.Mutex.Lock() // 传入的token不一致，说明不是递归调用
	// 抢到锁之后记录这个token
	atomic.StoreInt64(&m.token, token)
	m.recursion = 1
}

// 释放锁
func (m *TokenRecursiveMutex) Unlock(token int64) {
	if atomic.LoadInt64(&m.token) != token { // 释放其它token持有的锁
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.token, token))
	}
	m.recursion--         // 当前持有这个锁的token释放锁
	if m.recursion != 0 { // 还没有回退到最初的递归调用
		return
	}
	atomic.StoreInt64(&m.token, 0) // 没有递归调用了，释放锁
	m.Mutex.Unlock()
}
```
