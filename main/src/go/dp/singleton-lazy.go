package main

import "sync"

var sltLazy *Instance
var once = &sync.Once{}

func GetInstanceLazy() *Instance {
	if sltLazy == nil {
		once.Do(func() { // 类级别锁
			sltLazy = &Instance{}
		})
	}
	return sltLazy
}
