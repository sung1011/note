# go sync.Mutex

## Mutex

```go
// 数据结构; 源码
type Mutex struct {
	state int32
	sema uint32
}

// 常规使用
//   lock与lock阻塞互斥
type mytype struct {
	m   sync.Mutex      // 内置
	var int
}
```

> `临界区` Lock()和Unlock()之间的代码段称为资源的临界区(critical section), 临界区的代码同时间只有一个goroutine会执行.

> `不关联goroutine` 可以在goroutine加锁, 另一个goroutine解锁

## RWMutex

```go
// 数据结构; 源码
//   r与r兼容
//   r与w互斥阻塞
//   w与w互斥阻塞
type RWMutex struct {
	w           Mutex  // held if there are pending writers
	writerSem   uint32 // 写锁需要等待读锁释放的信号量
	readerSem   uint32 // 读锁需要等待写锁释放的信号量
	readerCount int32  // 读锁后面挂起了多少个写锁申请
	readerWait  int32  // 已释放了多少个读锁
}
func (rw *RWMutex) RLocker() Locker
func (rw *RWMutex) RLock()      // 读锁; rr兼容 rw, ww互斥
func (rw *RWMutex) RUnlock()    // 读解锁
func (rw *RWMutex) Lock()       // 锁; lock与lock互斥
func (rw *RWMutex) Unlock()     // 解锁; 
```