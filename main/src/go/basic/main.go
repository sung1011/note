package main

import (
	"fmt"
	"sync"
)

func main() {
	bPool := sync.Pool{
		New: func() interface{} {
			return make([]int, 5)
		},
	}

	get := bPool.Get().([]int)
	get[0] = 1
	bPool.Put(get)
	// pool  中存在刚刚更改的对象直接获取服用
	get0 := bPool.Get()
	fmt.Println("get0: ", get0)
	bPool.Put([]int{1, 23, 4, 6, 6})
	// pool  中存在刚刚put的对象不需要创建
	get1 := bPool.Get()
	fmt.Println("get1: ", get1)
	// pool  中不存在对象 需要new
	get2 := bPool.Get()
	fmt.Println("get2: ", get2)
	get3 := bPool.Get()
	fmt.Println("get3: ", get3)
}
