package service

type IUser interface {
	GetName(uid int64) string
	SetName(uid int64, name string) error
	AddExp(v int64)
}

type User struct {
	Id   int64
	Name string
	Exp  int64
}

func (u *User) GetName(uid int64) string {
	if uid == 99 {
		u.Name = "tickles"
	} else {
		u.Name = "nobody"
	}
	return u.Name
}

func (u *User) SetName(uid int64, name string) error {
	u.Name = name
	return nil
}

func (u *User) AddExp(v int64) {
	u.Exp += v
}
