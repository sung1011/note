# go compare for | range

```go
// benchmark 287 ns/op
func ForSlice(s []string) {
	len := len(s)
	for i := 0; i < len; i++ {
		_, _ = i, s[i]
	}
}

// benchmark 509 ns/op
func RangeForSlice1(s []string) {
	for i, v := range s { // copy
		_, _ = i, v
	}
}

// benchmark 287 ns/op
func RangeForSlice2(s []string) {
	for i, _ := range s { // 避免copy
		_, _ = i, s[i]
	}
}

// benchmark 287 ns/op
func ForMap(s []int, m map[int]string) {
    for k, _ := range s { // 已知所有key的遍历map, 才能快
        _, _ := k, m[k]
    }
}

// benchmark 14531 ns/op
func RangeForMap1(m map[int]string) {
	for k, _ := range m { // 遍历map是很慢的行为
		_, _ = k, m[k]
	}
}

// benchmark 23199 ns/op
func RangeForMap2(m map[int]string) {
	for range m { // 遍历map是很慢的行为, 无关是否copy出来
	}
}
```

> 飞雪无情 for range <https://www.flysnow.org/2018/10/20/golang-for-range-slice-map.html>