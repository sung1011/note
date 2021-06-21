package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_State(t *testing.T) {
	Convey("leader state", t, func() {
		m := &Machine{GetLeaderApproveState()}
		So(m.GetStateName(), ShouldEqual, "LeaderState")
		Convey("leader approve -> finance", func() {
			m.Approval()
			// leader approve 后, 对象主体已经从leader变为finance
			So(m.GetStateName(), ShouldEqual, "FinanceState")
			Convey("finance reject -> leader", func() {
				m.Reject()
				// finance拒绝后, 对象主体从finance回到leader
				// 即 状态变化后, 对象主体变化, approval, reject 行为同步变化
				So(m.GetStateName(), ShouldEqual, "LeaderState")
				Convey("leader approve -> finance (again)", func() {
					m.Approval()
					So(m.GetStateName(), ShouldEqual, "FinanceState")
					Convey("finance approve", func() {
						m.Approval()
					})
				})
			})
		})
	})
}
