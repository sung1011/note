package main

import (
	"fmt"
	"math"
)

func outputFEEE754float32(f float32) string {
	// 将该浮点数内存布局当做 uint32 看待（因为都占用 4 字节）
	// 这里实际上是做强制转换，内部实现是：return *(*uint32)(unsafe.Pointer(&f))
	buf := math.Float32bits(f)

	// 加上两处 -，结果一共 34 byte
	var result [34]byte

	// 从低字节开始
	for i := 33; i >= 0; i-- {
		if i == 1 || i == 10 {
			result[i] = '-'
		} else {
			if buf%2 == 1 {
				result[i] = '1'
			} else {
				result[i] = '0'
			}
			buf /= 2
		}
	}
	return fmt.Sprintf("%s", result)
}

func outputFEEE754float64(f float64) string {
	buf := math.Float64bits(f)

	var result [66]byte

	for i := 65; i >= 0; i-- {
		if i == 1 || i == 13 {
			result[i] = '-'
		} else {
			if buf%2 == 1 {
				result[i] = '1'
			} else {
				result[i] = '0'
			}
			buf /= 2
		}
	}
	return fmt.Sprintf("%s", result)
}
