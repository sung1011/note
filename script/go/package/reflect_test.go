package main

import (
	"reflect"
	"testing"
	"time"
	"unsafe"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Reflect_Type(t *testing.T) {
	Convey("Type", t, func() {
		date, _ := time.Parse("2006-01-02 15:04:05", "1989-10-11 04:00:00") // str -> time.Time
		stPTR := &struct {
			Name  string
			Age   int
			Birth time.Time `json:"birth" foo:"bar"`
		}{
			"sung",
			32,
			date,
		}
		// Ptr's Elem()
		typ := reflect.TypeOf(stPTR)
		So(typ.Kind(), ShouldEqual, reflect.Ptr)           // 指针类型
		So(typ.Elem().Kind(), ShouldEqual, reflect.Struct) // 取类型的元素(类似语言层*操作)

		// struct field & tag
		sf, _ := typ.Elem().FieldByName("Birth")
		So(sf.Type.NumField(), ShouldEqual, 3)
		So(sf.Name, ShouldEqual, "Birth")
		So(sf.Tag.Get("foo"), ShouldEqual, "bar")
	})
}

func Test_Reflect_Value(t *testing.T) {
	Convey("Value", t, func() {
		// 变量 -> reflect.Value -> Interface{} -> 变量
		var a int = 1024                                       // 变量
		valueOfA := reflect.ValueOf(a)                         // reflect.Value
		itf := valueOfA.Interface()                            // interface{}
		i := itf.(int)                                         // 类型断言 变回变量
		So(reflect.TypeOf(i).Kind(), ShouldEqual, reflect.Int) // int

		So(reflect.TypeOf(valueOfA.Int()).Kind(), ShouldEqual, reflect.Int64) // reflect.Value 强转int64
		So(valueOfA.CanAddr(), ShouldBeFalse)                                 // 不可被寻址 TODO why?
		So(valueOfA.CanSet(), ShouldBeFalse)                                  // 不可赋值 (需要可寻址 + 可导出的字段)
		So(valueOfA.CanInterface(), ShouldBeTrue)

		valueOfAPtr := reflect.ValueOf(&a)
		So(valueOfAPtr.Elem().CanAddr(), ShouldBeTrue) // 可寻址 a的值
		So(valueOfAPtr.CanSet(), ShouldBeFalse)        // 不可赋值 (需要可寻址 + 可导出的字段)
		x := valueOfAPtr.Elem()
		x.SetInt(1234)
		So(x.Int(), ShouldEqual, 1234)
		So(reflect.TypeOf(x).Kind(), ShouldEqual, reflect.Struct)
	})
}

func Test_Reflect_IsZero(t *testing.T) {
	Convey("Is Zero", t, func() {
		tests := []struct {
			isZero bool
			val    interface{}
		}{
			// bool
			{true, false},
			// int
			{true, 0},
			// string
			{true, ""},
			//
			{true, uintptr(0)},
			// array
			{true, [3]string{}},
			{true, [3]string{"", "", ""}},
			// chan
			{true, (chan string)(nil)},
			{false, make(chan string)},
			{false, reflect.New},
			// map
			{true, (map[string]string)(nil)},
			{false, map[string]string{}},
			{false, make(map[string]string)},
			// ptr
			{true, (*func())(nil)},
			{true, (*int)(nil)},
			{false, new(int)},
			// slice
			{true, ([]string)(nil)},
			{false, make([]string, 0)},
			{false, []string{}},
			// unsafe ptr
			{true, (unsafe.Pointer)(nil)},
			{false, (unsafe.Pointer)(new(int))},
		}
		for _, row := range tests {
			So(reflect.ValueOf(row.val).IsZero(), ShouldEqual, row.isZero)
		}
	})
}
