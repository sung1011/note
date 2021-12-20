package main

// 采用原课程的示例, 并且做了一些裁剪
// 假设我们现在有一个较为复杂的对话框, 里面包括, 登录组件, 注册组件, 以及选择框
// 当选择框选择“登录”时, 展示登录相关组件
// 当选择框选择“注册”时, 展示注册相关组件

import (
	"fmt"
	"reflect"
)

// Input 假设这表示一个输入框
type Input string

// String String
func (i Input) String() string {
	return string(i)
}

// Selection 假设这表示一个选择框
type Selection string

// Selected 当前选中的对象
func (s Selection) Selected() string {
	return string(s)
}

// Button 假设这表示一个按钮
type Button struct {
	onClick func()
}

// SetOnClick 添加点击事件回调
func (b *Button) SetOnClick(f func()) {
	b.onClick = f
}

// IMediator 中介模式接口
type IMediator interface {
	HandleEvent(component interface{}) int
}

// Dialog 对话框组件 -- 中介对象
type Dialog struct {
	LoginButton         *Button
	RegButton           *Button
	Selection           *Selection
	UsernameInput       *Input
	PasswordInput       *Input
	RepeatPasswordInput *Input
}

// HandleEvent HandleEvent
func (d *Dialog) HandleEvent(component interface{}) int {
	switch {
	case reflect.DeepEqual(component, d.Selection):
		if d.Selection.Selected() == "登录" {
			fmt.Println("select login")
			fmt.Printf("show: %s\n", d.UsernameInput)
			fmt.Printf("show: %s\n", d.PasswordInput)
			return 3
		} else if d.Selection.Selected() == "注册" {
			fmt.Println("select register")
			fmt.Printf("show: %s\n", d.UsernameInput)
			fmt.Printf("show: %s\n", d.PasswordInput)
			fmt.Printf("show: %s\n", d.RepeatPasswordInput)
			return 2
		} else {
			// others, 如果点击了登录按钮, 注册按钮
			return 4
		}
	}
	return 0
}
