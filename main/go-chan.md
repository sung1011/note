# go channel

- [chan](src/go/basic/chan_test.go)

## 场景

    消息传递: 不同goroutine之间线程安全的数据交流
    任务编排: goroutine之间有互相等待or依赖的顺序关系, 按一定规律执行

## 特征

    读empty阻塞, 读close空零, 写close异常, 读写nil阻塞


## 数据结构

```js
```

```go
	t := make(chan int)
	fmt.Printf("%p\n", t) // 0xc00008c060
	println(t) // 0xc00008c060
```

## 实例

- [实例](src/go/basic/chan_test.go)

## unbuffered chan

    sender ch<- 持续阻塞直到 receiver <-ch
    receiver <-ch 持续阻塞直到 sender ch<-

```go
    // 初始化
    // make(chan TYPE,CAP)
    ch := make(chan int)    // ch := make(chan int, 0)

    // 写入
    //   写close异常
    //   写nil阻塞
    ch<-123

    // 读取
    //   读空阻塞
    //   读close空零; 控制goroutine顺序
    //   读nil阻塞
    <-ch                // 读取赋值 val = <-ch

    // 关闭chan
    //   只能发送方(ch<-)关闭; 表达不再生产; 接收方(<-ch)关闭会报错
    //   最好defer
    //   close(ch)并非必须, 某些时候可以自动关闭
    close(ch)  

    // chan是否关闭 
    val, ok := <-ch     // ok == false 表达chan已关闭

    // nil chan
    //   常用来禁用select的某分支
    ch = nil

    // 基本用法
    func main() {
        ch := make(chan int)
        go func() {
            ch <- 1
            defer close(ch)   // 发送方close()
        }()
        <-ch
    }
```

## buffered chan

    sender ch<- 不断发送直到len=cap发生阻塞
    receiver 按照FIFO顺序出队<-ch, 直到队空阻塞

```go
    // 初始化
    ch := make(chan int, 10)        
    fmt.Println(len(ch), cap(ch))   // len=0; cap=10

    // 其他同 unbuffered chan

    // 基本用法
    counter := make(chan int, 10)
	go func() {
		counter <- 31
		counter <- 32
		counter <- 33
		defer close(counter)
	}()
	fmt.Println(<-counter)  // 31
	fmt.Println(<-counter)  // 32
	fmt.Println(<-counter)  // 33

```

## nil chan 阻塞

```go
func send(c chan int) {
	for {
		c <- rand.Intn(10)
	}
}

func add(c chan int) {
	sum := 0
	t := time.NewTimer(1 * time.Second) // 或 t := time.After(1 * time.Second); <-t
	for {
		select {
		case input := <-c:      // 1s内不断执行
			sum = sum + input
		case <-t.C:             // 1s后执行
			c = nil             // 阻塞另一分支
			fmt.Println(sum)
		}
	}
}

func main() {
	c := make(chan int)
	go add(c)
	go send(c)
	time.Sleep(3 * time.Second)
}
```

## chan chan 双重

    外层chan作为chan的加工厂

```go
var chch1 chan chan int
// 发送通道给外层通道
chch1 <-ch1
chch1 <-ch2

// 从外层通道取出内层通道
c <-chch1

// 操作取出的内层通道
c <-123
val := <-c
```

## deadlock 死锁

    所有goroutine都被阻塞, 就会出现死锁

```go
	ch := make(chan int)
	ch <- 1                 // 阻塞线程; fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
```

## select 多路监听

    所有case都阻塞, 则阻塞等待
    case同时满足, 则随机一个case处理
    如果case都不满足, 并存在default, 处理default
    select中有发送ch<-很可能会出现死锁, 常见解决办法
        - select加上default
        - select加上超时时间
        - 不放在for无限循环

```go
select {
	// ch1有数据时, 读取到v1变量中
	case v1 := <-ch1:
		...
	// ch2有数据时, 读取到v2变量中
	case v2 := <-ch2:
		...
	// 所有case都不满足条件时, 执行default
	default:
		...
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
