package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Adapter(t *testing.T) {
	Convey("", t, func() {
		Convey("aliyun implement interface", func() {
			var a ICreateServer = &AliyunClientAdapter{
				Client: AliyunClient{},
			}
			a.CreateServer(1.0, 2.0)

			So(&AliyunClientAdapter{}, ShouldImplement, (*ICreateServer)(nil))
		})
		Convey("aws implement interface", func() {
			var a ICreateServer = &AwsClientAdapter{
				Client: AWSClient{},
			}
			a.CreateServer(1.0, 2.0)

			So(&AwsClientAdapter{}, ShouldImplement, (*ICreateServer)(nil))
		})

	})
}
