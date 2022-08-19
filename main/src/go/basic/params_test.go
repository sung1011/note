package main

import (
	"fmt"
	"testing"
)

type user struct {
	Name string
}

func Test_Params(t *testing.T) {
	ValParamsStruct() // 值类型 struct; 传ref会改变
	ValParamsInt()    // 值类型 int; 传ref会改变

	RefParamsSlice() // 引用类型 slice; 不论传val还是ref都会改变

}

func ValParamsStruct() {
	fmt.Println("---------val int---------")
	var u user
	u.Name = "sun"
	fmt.Printf("s: %v, u: %v, ptr: %p\n", "origin", u, &u)
	structModifyVal(u)
	fmt.Printf("s: %v, u: %v, ptr: %p\n", "after val", u, &u) // u.Name没改
	structModifyRef(&u)
	fmt.Printf("s: %v, u: %v, ptr: %p\n", "after ref", u, &u) // u.Name改了
}

func structModifyVal(u user) { // 值拷贝 地址变了
	fmt.Printf("s: %v, u: %v, ptr: %p\n", "val inner 1", u, &u)
	u.Name = "xxx"
	fmt.Printf("s: %v, u: %v, ptr: %p\n", "val inner 2", u, &u)
}

func structModifyRef(u *user) { // 引用指针的拷贝 地址变了
	fmt.Printf("s: %v, u: %v, ptr: %p\n", "ref inner 1", u, &u)
	u.Name = "xxx"
	fmt.Printf("s: %v, u: %v, ptr: %p\n", "ref inner 2", u, &u)
}

//--------

func ValParamsInt() {
	fmt.Println("---------val struct---------")
	var i int
	i = 123
	fmt.Printf("s: %v, i: %v, ptr: %p\n", "origin", i, &i)
	intModifyVal(i)
	fmt.Printf("s: %v, i: %v, ptr: %p\n", "after val", i, &i) // i 没改
	intModifyRef(&i)
	fmt.Printf("s: %v, i: %v, ptr: %p\n", "after ref", i, &i) // i 改了
}

func intModifyVal(i int) { // 值拷贝 地址变了
	fmt.Printf("s: %v, i: %v, ptr: %p\n", "val inner 1", i, &i)
	i = 444
	fmt.Printf("s: %v, i: %v, ptr: %p\n", "val inner 2", i, &i)
}

func intModifyRef(ip *int) { // 引用指针的拷贝 地址变了
	fmt.Printf("s: %v, i: %v, ptr: %p\n", "ref inner 1", *ip, &ip)
	*ip = 444
	fmt.Printf("s: %v, i: %v, ptr: %p\n", "ref inner 2", *ip, &ip)
}

//--------

func RefParamsSlice() {
	fmt.Println("---------ref slice---------")
	var sl []int
	sl = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("s: %v, i: %v, ptr: %p\n", "origin", sl, &sl)
	slModifyVal(sl)
	fmt.Printf("s: %v, sl: %v, ptr: %p\n", "after val", sl, &sl) // sl 变了
	slModifyRef(&sl)
	fmt.Printf("s: %v, sl: %v, ptr: %p\n", "after ref", sl, &sl) // sl 改了
}

func slModifyVal(sl []int) { // 值拷贝 地址变了
	fmt.Printf("s: %v, sl: %v, ptr: %p\n", "val inner 1", sl, &sl)
	sl[3] = 444
	fmt.Printf("s: %v, sl: %v, ptr: %p\n", "val inner 2", sl, &sl)
}

func slModifyRef(sl *[]int) { // 引用指针的拷贝 地址变了
	fmt.Printf("s: %v, sl: %v, ptr: %p\n", "ref inner 1", *sl, &sl)
	(*sl)[3] = 888
	fmt.Printf("s: %v, sl: %v, ptr: %p\n", "ref inner 2", *sl, &sl)
}
