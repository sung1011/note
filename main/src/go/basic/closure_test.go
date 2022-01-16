package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// 闭包
// 官方译文: 闭包是一个函数值，它引用了函数体之外的变量。 这个函数可以对这个引用的变量进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。
// 主要理解 闭包函数 <--> 函数内变量, 闭包可修改函数内变量, 再次调用函数时变量是修改后的; 类似于 函数 <--> 全局变量 的关系

func calc() func() int {
	r := 1              // 函数内变量
	return func() int { // 闭包函数
		r = r + 4
		return r
	}
}

func TestCalc(t *testing.T) {
	Convey("", t, func() {
		funcA := calc() // 附带函数(calc)内的作用域, 即当前r=1
		a1 := funcA()   // r = 1 + 4; 附带函数(calc)内的作用域, 即calc中r也变为5
		So(a1, ShouldEqual, 5)

		a2 := funcA() // r = 5 + 4; 附带函数(calc)内的作用域, 即calc中r也变为9
		So(a2, ShouldEqual, 9)

		funcB := calc() // 新函数作用域, 即当前r=1
		b1 := funcB()

		So(b1, ShouldEqual, 5)
	})
}
