# go sync.Cond

    等待, 通知
    (难用, 一般用channel替代)

## 方法

```go
type Cond
    func NeWCond(l Locker) *Cond
    func (c *Cond) Broadcast() // 唤醒所有等待的goroutine
    func (c *Cond) Signal() // 随机唤醒一个等待的goroutine
    func (c *Cond) Wait() // 阻塞等待通知; 一般在Lock()和Unlock()之间调用
```

## code

- [basic](src/go/sync/cond_test.go)