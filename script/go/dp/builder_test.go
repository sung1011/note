package main

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Builder_Simple(t *testing.T) {
	Convey("", t, func() {
		builder := &ResourcePoolConfigBuilder{}
		Convey("err empty name", func() {
			_, err := builder.Build()
			So(err != nil, ShouldBeTrue)
			So(err, ShouldBeError, errors.New("name can not be empty"))
		})

		Convey("err range", func() {
			builder.SetName("abc")
			max := 1
			min := 100
			builder.SetMaxIdle(max)
			builder.SetMinIdle(min)
			_, err := builder.Build()
			So(err != nil, ShouldBeTrue)
			So(err, ShouldBeError, fmt.Errorf("max idle(%d) cannot < min idle(%d)", max, min))
		})

		Convey("success", func() {
			builder.SetName("foo")
			builder.SetMinIdle(1)
			builder.SetMaxIdle(5)
			obj, err := builder.Build()
			So(err == nil, ShouldBeTrue)
			So(obj.maxTotal, ShouldEqual, defaultMaxTotal)
			So(obj.maxIdle, ShouldEqual, 5)
		})
	})
}

func Test_Builder_Opt(t *testing.T) {
	Convey("", t, func() {
		Convey("err empty name", func() {
			_, err := NewResourcePoolConfig("")
			So(err != nil, ShouldBeTrue)
			So(err, ShouldBeError, errors.New("name can not be empty"))
		})

		Convey("err range", func() {
			min := 100
			max := 1
			opts := []ResourcePoolConfigOptFunc{
				func(option *ResourcePoolConfigOption) {
					option.minIdle = min
					option.maxIdle = max
				},
			}
			_, err := NewResourcePoolConfig("foo", opts...)
			So(err != nil, ShouldBeTrue)
		})

		Convey("success", func() {
			min := 1
			max := 5
			name := "foo"
			opts := []ResourcePoolConfigOptFunc{
				func(option *ResourcePoolConfigOption) {
					option.minIdle = min
					option.maxIdle = max
				},
			}
			obj, err := NewResourcePoolConfig(name, opts...)

			So(err == nil, ShouldBeTrue)
			So(obj.name, ShouldEqual, name)
			So(obj.maxIdle, ShouldEqual, max)
		})
	})
}
