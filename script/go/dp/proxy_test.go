package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Proxy_Static(t *testing.T) {
	Convey("proxy static", t, func() {
		proxy := NewUserProxy(&UserStatic{})
		err := proxy.Login("admin", "123456")
		So(err, ShouldBeNil)
	})
}

func Test_Proxy_Dynamic(t *testing.T) {
	want := `package main

type UserProxy struct {
	child *User
}

func NewUserProxy(child *User) *UserProxy {
	return &UserProxy{child: child}
}

func (p *UserProxy) Login(username, password string) (r0 error) {
	// before 这里可能会有一些统计的逻辑
	start := time.Now()

	r0 = p.child.Login(username, password)

	// after 这里可能也有一些监控统计的逻辑
	log.Printf("user login cost time: %s", time.Now().Sub(start))

	return r0
}
`
	Convey("proxy dynamic", t, func() {
		str, err := generate("./proxy-dynamic-base.go")
		So(err, ShouldBeNil)
		So(str, ShouldEqual, want)
	})
}
