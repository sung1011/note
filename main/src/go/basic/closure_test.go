package main

import (
	"testing"
)

// 闭包
// 主要理解 闭包函数 <--> 函数内变量
// 类似于 函数 <--> 全局变量 的关系

func calc() func() int {
	r := 1              // 函数内变量
	return func() int { // 闭包函数
		r = r + 4
		return r
	}
}

func TestCalc(t *testing.T) {
	funcA := calc() // 附带函数(calc)内的作用域, 即当前r=1
	a1 := funcA()   // r = 1 + 4; 附带函数(calc)内的作用域, 即calc中r也变为5
	a2 := funcA()   // r = 5 + 4; 附带函数(calc)内的作用域, 即calc中r也变为9

	funcB := calc() // 新函数作用域, 即当前r=1
	b1 := funcB()

	if a1 != 5 {
		t.Errorf("err calc %v", a1)
	}
	if a2 != 9 {
		t.Errorf("err calc %v", a2)
	}
	if b1 != 5 {
		t.Errorf("err calc %v", b1)
	}

}
