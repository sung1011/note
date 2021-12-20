package main

import "fmt"

type Cooker interface {
	open()
	fire()
	cooke()
	outfire()
	close()
}

// 父类(模板)
type CookMenu struct {
	Dish string
}

func (CookMenu) open() {
	fmt.Println("打开开关")
}

func (CookMenu) fire() {
	fmt.Println("开火")
}

// 做菜, 交给具体的子类实现
func (CookMenu) cooke() {
}

func (CookMenu) outfire() {
	fmt.Println("关火")
}

func (CookMenu) close() {
	fmt.Println("关闭开关")
}

// 封装具体步骤
func doCook(cook Cooker) {
	cook.open()
	cook.fire()
	cook.cooke()
	cook.outfire()
	cook.close()
}

// 子类
type XiHongShi struct {
	CookMenu // 父类
}

// 重写父类(模板)中的cooke()
func (p *XiHongShi) cooke() {
	fmt.Println("做西红柿")
	p.Dish = "我是一盘西红柿"

}

// 子类
type ChaoJiDan struct {
	CookMenu // 父类
}

func (p *ChaoJiDan) cooke() {
	fmt.Println("做炒鸡蛋")
	p.Dish = "我是一盘炒鸡蛋"
}
