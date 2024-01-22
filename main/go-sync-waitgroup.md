# go sync.WaitGroup

## 场景

    任务编排: goroutine之间有互相等待or依赖的顺序关系, 按一定规律执行

## 方法

```go
func (wg *WaitGroup) Add(delta int) // 一般主goroutine调用
func (wg *WaitGroup) Done() // == Add(-1); 一般在其他goroutine中调用
func (wg *WaitGroup) Wait() // 调用这个方法的goroutine会阻塞, 直到counter==0; 一般主goroutine调用
```

## 使用

```go
// 屏蔽区为另一种写法
var mu sync.Mutex

func main() {
	var wg sync.WaitGroup
	sl := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
        wg.Add(1)           // counter原子自增n; 一定要在wg.Wait()之前调用, 因此不能写在go func内
		go worker(&sl, i, &wg) // 传sl指针是因为append
	}
    // for i := 0; i < 100; i++ {
    //     worker(&sl, i, &wg) 
	// }
	wg.Wait()                // 阻塞等待变回counter=0
	fmt.Println(len(sl), sl) // 100, [93, 5, 12, ...]
}

func worker(sl *[]int, i int, wg *sync.WaitGroup) {
    // wg.Add(1)
    // go func() {
	mu.Lock()
	*sl = append(*sl, i)
	mu.Unlock()
	wg.Done() // counter原子自减1
    // }()
}
```

## 易错场景

### 计数器add, done不一致

```go
// deadlock
wg.Add(3)
wg.Done()
wg.Done()

// panic
wg.Add(2)
wg.Done()
wg.Done()
wg.Done()
```

### add未调用就wait

```go
func main() {
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        go worker(&wg)
    }
    wg.Wait()
}

func worker(wg *sync.WaitGroup) {
    wg.Add(1)
    // ...
    wg.Done()
}
```

> vet可以帮助检测出这些问题