# go channel

    nil读写阻塞，写关闭异常，读关闭空零

## 数据结构

    todo

```go
	t := make(chan int)
	fmt.Printf("%p\n", t) // 0xc00008c060
	println(t) // 0xc00008c060
```

## 基础用法

```go
func main() {
    ch := make(chan int)
    go func() {
        ch <- 1
    }()
    <-ch
}
```

## 是否关闭

```go
r, ok <-ch
if !ok {
    fmt.Println("is close")
}
```

## 控制速率

```go
var wg = sync.WaitGroup{}

func busi(ch chan int) {
	for t := range ch {
		fmt.Println("go task = ", t, ", goroutine count = ", runtime.NumGoroutine())
		wg.Done()
	}
}

func sendTask(task int, ch chan int) {
	wg.Add(1)
	ch <- task
}

func main() {
	ch := make(chan int)
	goCnt := 3 //goroutine的数量
	for i := 0; i < goCnt; i++ {
		go busi(ch)
	}

	taskCnt := 100000 // math.MaxInt64 模拟用户需求业务的数量
	for t := 0; t < taskCnt; t++ {
		//发送任务
		sendTask(t, ch)
	}
	wg.Wait()
}
```


## ref

- <https://www.cnblogs.com/liang1101/p/7285955.html>
