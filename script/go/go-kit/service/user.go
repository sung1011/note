package service

type IUser interface {
	GetName(uid int64) string
}

type User struct{}

func (u *User) GetName(uid int64) string {
	return "tickles"
}
