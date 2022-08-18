# interview-questions

## quest

### requirement

- 每月更新
- 10w人排行榜
- 每人分数<=10000的整数
- 相同分数, 先达到的靠前
- 从高到低排名
- 按uid查score
- 自己前后10名的score和rank

## answer

### 流程

1. 每个玩家身上记录记录自己的分数

2. 月底排名之前的5min 游戏显示`活动已经结束, 正在统计排行榜`. 以防止极限条件下 统计结束后才加分的bug
   - 若需求不允许预留这5min, 就与策划分情况讨论上述极限条件时如何处理 ( 加分 / 静默丢弃 / 提示丢弃 / 不允许开可能过期的局 ... )

3. 计算排名
   - 后台脚本执行 & 输出log

### 伪代码

```go
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

const tmpNum int = 9999999999

// value = score 拼接 时间参数
//   时间参数 = 9999999999 - 当前时间戳; 用来实现相同排名, 先达到的排在前面
// * 优化: 也可以用24bit表达194天内的秒数, 类似redis的LRU设计, 程序里简单的用了int
type value int
type score int

type IRankList interface {
	GetScoreByUid(uid int) score
	GetRange(uid, n int) []user
	Incr(uid int, sc score)
}

type user struct {
	uid int
	val value
}

func valToSc(val value) score {
	if val == 0 {
		return 0
	}
	s := strconv.FormatInt(int64(val), 10)
	sc, _ := strconv.Atoi(s[:len(s)-10])
	return score(sc)
}

func scToVal(sc score) value {
	valTm := int64(tmpNum) - time.Now().Unix()
	s := strconv.FormatInt(int64(sc), 10) + strconv.FormatInt(int64(valTm), 10)
	i, _ := strconv.Atoi(s)
	return value(i)
}

type RankList struct {
	Users []*user
}

func (rl *RankList) sort() {
	sort.Slice(rl.Users, func(i, j int) bool {
		return rl.Users[i].val > rl.Users[j].val
	})
}

func (rl *RankList) Incr(uid int, sc score) {
	for _, usr := range rl.Users {
		if uid == usr.uid {
			usr.val = scToVal(valToSc(usr.val) + sc)
			return
		}
	}
	u := &user{
		uid: uid,
		val: scToVal(sc),
	}
	rl.Users = append(rl.Users, u)
}

func (rl *RankList) GetScoreByUid(uid int) score {
	for _, usr := range rl.Users {
		if uid == usr.uid {
			return valToSc(usr.val)
		}
	}
	return score(0)
}

func (rl *RankList) GetRange(uid, n int) []user {
	rl.sort()
	var myRank int
	for rank, usr := range rl.Users {
		if uid == usr.uid {
			myRank = rank
		}
	}
	if myRank == 0 {
		return nil
	}
	rangeMin := myRank - n
	rangeMax := myRank + n
	var users []user
	for rank, usr := range rl.Users {
		if rank >= rangeMin && rank <= rangeMax {
			fmt.Printf("rank: %v, uid: %v, score: %v\n", rank, usr.uid, valToSc(usr.val))
			users = append(users, *usr)
		}
	}
	return users
}

func MockAndTest(rl IRankList) {
	// mock
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 50; i++ {
		rl.Incr(i, score(rand.Intn(10000)))
	}
	uid := 14

	// test; 简单输出结果, 为保持单文件没使用单元测试 _test.go
	fmt.Println("指定uid获取分数", rl.GetScoreByUid(uid))

	fmt.Println("指定uid获取前后10名玩家数据")
	rl.GetRange(uid, 10)
}

func main() {
	rl := new(RankList)
	MockAndTest(rl)
}
```