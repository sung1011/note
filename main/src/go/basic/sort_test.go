package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

/*

- sort <https://blog.csdn.net/K346K346/article/details/118314382>
*/

func Test_Sort_Func(t *testing.T) {
	// sort内部实现4种排序, 根据情况自动选用
	//   插入排序（Shell 排序）、归并排序、堆排序和快速排序。
}

func Test_Sort_BIGFILE(t *testing.T) {
	/*
		// 大文件排序 5e行 每行一个数字
			归并 太耗内存OOM
			sort 太慢24min
			bitmap 耗时190s, 使用内存 5e的内存/32
			外部排序 切成小文件 + 归并; 13min, 耗内存最小, 读写IO较高
	*/
}
