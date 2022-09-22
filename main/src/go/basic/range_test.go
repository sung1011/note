package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// # ref - <https://learnku.com/articles/26861>

func Test_Range(t *testing.T) {
	Convey("", t, func() {
		Convey("issue1: range只创建1次i, v", func() {
			Convey("Q: range是i, v都只创建1次, 导致v的地址是同一个", func() {
				sl := []int{1, 2, 3}
				for _, v := range sl { // range的坑, v只在第一次循环时创建了1次, 开辟了一个地址&v
					t.Log(v, &v) // 3次v不同, 但是是同一个地址
				}
			})

			Convey("A: 重新赋值创建n次", func() {
				sl := []int{1, 2, 3}
				for _, v := range sl {
					tmp := v       // 创建了3次, 是3个不同的变量
					t.Log(v, &tmp) // 3次v不同, 是不同地址
				}
			})

			Convey("A: 用for代替range", func() {
				sl := []int{1, 2, 3}
				for i := 0; i < len(sl); i++ {
					t.Log(&sl[i])
				}
			})
		})

		Convey("issue2: slice循环中增删, 不影响其循环次数", func() {
			sl := []int{1, 2, 3}
			for i := range sl {
				if i == 1 {
					sl = append(sl, 4, 5) // 增; 底层array(值类型)已经是新array了, 所以不影响迭代中的旧sl
					sl[4] = 555           // 改值会正常改, 立刻生效
					// sl = sl[4:]           // 删; 删除 0 1 2 3
				}
				t.Logf("循环次数:%v, 值:%v", i, sl) // 还是输出3次
			}
		})

		Convey("issue3: map循环中增删, 由于其随机特性, 不确定是否/有可能 会增减循环次数", func() {
			m := map[int]int{
				1: 111,
				2: 222,
				3: 333,
			}
			for i := range m { // i 在1~3中随机
				// 当恰巧i == 最后一次, 会增加循环次数
				if i == 3 {
					m[3] = 250 // 改值会正常改, 立刻生效
					m[7] = 777 // 增; map(引用类型), 可能会被影响
					m[8] = 888
					m[9] = 999
				}
				t.Log(m)
			}
		})
	})
}
