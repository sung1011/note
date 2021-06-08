package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_bridge(t *testing.T) {
	Convey("", t, func() {
		// email/phone/msg/dingding/qq...
		sender := NewEmailMsgSender([]string{"test@test.com"})
		// error/warn/notice
		notify := NewErrorNotification(sender)
		err := notify.Notify("abc")
		So(err, ShouldBeNil)
	})
}
