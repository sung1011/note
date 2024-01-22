# go sync.RWMutex

## data structure

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