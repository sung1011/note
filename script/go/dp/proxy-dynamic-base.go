package main

type IUser interface {
	Login(username, password string) error
}

// User 用户
// @proxy IUser
type User struct{}
