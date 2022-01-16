package main

import (
	"log"
	"time"
)

type IOrder interface {
	OrderList() []Order
}

// 装饰器模式 用组合替代继承
type UserDecorator struct {
	u *UserStatic
	o []Order
}

type Order struct {
	Id int64
	t  time.Time
	v  int64
}

func (p *UserDecorator) Login(uname, pwd string) error {
	// before 这里可能会有一些统计的逻辑
	start := time.Now()

	r := p.u.Login(uname, pwd)

	// after 这里可能也有一些监控统计的逻辑
	log.Printf("user login cost time: %s", time.Now().Sub(start))

	return r
}

// 装饰器模式 扩展方法
func (p *UserDecorator) OrderList() []Order {
	return p.o
}
