# go sync.WaitGroup

```go
var mu sync.Mutex

func main() {
	var wg sync.WaitGroup
	sl := make([]int, 0, 100) //结果集
	// wg.Add(100) // 一般不这样写               
	for i := 0; i < 100; i++ {
        wg.Add(1)           // counter原子自增n; 一定要在wg.Wait()之前调用, 即不能写在go func内
		go foo(&sl, i, &wg) // ??? 引用类型slice为啥还要用指针 ???
	}
	wg.Wait()                // 阻塞等待变回counter==0
	fmt.Println(len(sl), sl) // 100, [93, 5, 12, ...]
}

func foo(sl *[]int, i int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second) // mock逻辑消耗

	mu.Lock()
	*sl = append(*sl, i)
	mu.Unlock()

	wg.Done() // counter原子自减1
}
```
