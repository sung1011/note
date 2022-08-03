package main

import (
	"log"
	"time"
)

type IUserStatic interface {
	Login(uname, pwd string) error
}

// 原始类
type UserStatic struct{}

// 原始功能
func (u *UserStatic) Login(uname, pwd string) error {
	return nil
}

// proxy 包了一下原始类, 重新实现其方法
type UserProxy struct {
	u *UserStatic
}

func NewUserProxy(u *UserStatic) *UserProxy {
	return &UserProxy{u: u}
}

func (p *UserProxy) Login(uname, pwd string) error {
	// before 这里可能会有一些统计的逻辑
	start := time.Now()

	r := p.u.Login(uname, pwd)

	// after 这里可能也有一些监控统计的逻辑
	log.Printf("user login cost time: %s", time.Now().Sub(start))

	return r
}
