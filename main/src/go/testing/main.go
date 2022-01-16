package main

import "bytes"

// Square 平方
func Square(a int) int {
	return a * a
}

// StrAppend 字符串直接拼接
func StrAppend(s ...string) {
	ret := ""
	for _, elem := range s {
		ret += elem
	}
}

// StrAppendByBuffer buff方式字符串拼接
func StrAppendByBuffer(s ...string) {
	var buf bytes.Buffer
	for _, elem := range s {
		buf.WriteString(elem)
	}
}
