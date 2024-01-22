package main

import (
	"sync"
	"testing"
)

type Counter struct {
	sync.Mutex
	n int32
}

// sync.Mutex 的单元测试用例
func Test_Mutex(t *testing.T) {
	var cntr Counter
	cntr.Lock()
	defer cntr.Unlock()
	cntr.n++
}
