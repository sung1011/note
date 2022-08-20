# go debug

## 协程

```go
# trace.go
package main

import (
    "os"
    "fmt"
    "runtime/trace"
)

func main() {

    //创建trace文件
    f, err := os.Create("trace.out")
    if err != nil {
        panic(err)
    }

    defer f.Close()

    //启动trace goroutine
    err = trace.Start(f)
    if err != nil {
        panic(err)
    }
    defer trace.Stop()

    //main
    fmt.Println("Hello World")
}
```

```bash
go run trace.go // 生成trace.out
go tool trace trace.out
# 点击输出的链接 跳转到web界面查看
```

## time 查看耗时

```bash
    time go run test.go

    real	0m0.843s
    user	0m0.216s
    sys	    0m0.389s

    # real：从程序开始到结束，实际度过的时间；
    # user：程序在用户态度过的时间；
    # sys：程序在内核态度过的时间。
```

## GCTrace 查看内存

```bash
    go build -o foo && GODEBUG='gctrace=1' ./foo

    # 格式
    # gc # @#s #%: #+#+# ms clock, #+#/#/#+# ms cpu, #->#-># MB, # MB goal, # P

```

```txt
格式
    如 gc # @#s #%: #+#+# ms clock, #+#/#/#+# ms cpu, #->#-># MB, # MB goal, # P

  - `gc #`        GC次数的编号，每次GC时递增  
  - `@#s`         距离程序开始执行时的时间  
  - `#%`          GC占用的执行时间百分比  
  - `#+...+#`     GC使用的时间  
  - `#->#-># MB`  GC开始，结束，以及当前活跃堆内存的大小，单位M  
  - `# MB goal`   全局堆内存大小  
  - `# P`         使用processor的数量  
```

## ReadMemStats() 查看内存

```go
// 调用即打印内存情况
func readMemStats() {   
    var ms runtime.MemStats
    runtime.ReadMemStats(&ms)
    log.Printf(" ===> Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}
```

## net/http/pprof包 web查看内存,CPU

## go tool pprof 查看内存

## ref

- <https://www.yuque.com/aceld/golang/ga6pb1>