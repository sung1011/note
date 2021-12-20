# go map

## 数据结构

```go
type hmap struct {
    count     int // # live cells == size of map.  Must be first (used by len() builtin)
    flags     uint8
    B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
    noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
    hash0     uint32 // hash seed

    buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
    oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
    nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

    extra *mapextra // optional fields
}

type mapextra struct { // mapextra holds fields that are not present on all maps.
    overflow    *[]*bmap
    oldoverflow *[]*bmap

    nextOverflow *bmap
}

type bmap struct { // A bucket for a Go map.
    tophash [bucketCnt]uint8
}
```

## 创建 初始化 访问

```go
m := map[string]int{"red": 1, "green": 2, "blue": 3} // len=3; 常用

m := make(map[string]int, 10) // empty map; len=0; cap=10 常用

m := map[string]int{} // empty map; prt=xxx len=0 cap=0; 可append

var m = map[string]int // nil ptr map; panic if written to -- 没啥意义, 不可用, 不能赋值和append

```

> `len()` 对map而言, len就是cap

> `cap()` map有cap的概念, 但无法使用cap()

> `key限制` 不能使用 == 做比较的类型不能做map的key; 如`切片slice`, `函数func`, `包含切片和函数的结构体struct{slice, func}`; chan可以做key

> `val限制` 任何类型

## 返回值 key是否存在

```go
m := map[string]int{"a":99, "b":22}

v1, exists1 := m["a"] // 99, true

v2, exists2 := m["xxx"] // 0, false
```

## 作为参数

直接传递

> map的副本值(即:指针)作为参数传递给函数.

## 并发安全

```go
// 同步锁
type RWMap struct {
    m map[string]int
    sync.RWMutex
}
func (r RWMap) Get(key string) int {
    r.RLock()
    defer r.RUnlock()
    return r.m[key]
}
func (r RWMap) Set(key string, val int) {
    r.Lock()
    defer r.Unlock()
    r.m[key] = val
}
```

```go
// 原子锁
var sm sync.Map
sm.Store(1, "a")
if v,ok:=sm.Load(1);ok{
    fmt.Println(v)
}
```

## 赋值 (保持引用)

```go
// map赋值, 会同步修改底层.
originalMap := make(map[string]int, 10)
originalMap["a"] = 1
originalMap["b"] = 2
originalMap["c"] = 3
originalMap["d"] = 4
originalMap["e"] = 5
targetMap := make(map[string]int, 2)
targetMap = originalMap
targetMap["d"] = 4444

fmt.Println(originalMap) // map[a:1 b:2 c:3 d:4444 e:5]
```

## 拷贝 (解引用)

```go
originalMap := make(map[string]int)
originalMap["a"] = 1
originalMap["b"] = 2

targetMap := make(map[string]int)
for k, v := range originalMap {
    targetMap[k] = v
}
```

## 删除key

```go
originalMap := make(map[string]int, 2)
originalMap["a"] = 1
originalMap["b"] = 2
delete(originalMap, "a") // map[a:1]
```

## 集合

```go
// 推荐结构 map[Type]struct{}
sets := map[string]struct{} (
    "x": {},
    "y": {},
)
```

## 函数作为map的值

```go
// map_func := map[KEY_TYPE]func() RETURN_TYPE {......}
mf := map[int]func()int{
    1: func() int { return 10 },
    2: func() int { return 20 },
}

// map_func := make(map[KEY_TYPE]func() RETURN_TYPE)
mf := make(map[int]func()string)
mf[1] = func() string{ return "aaa" }
mf[2] = func() string{ return "bbb" }
```

## 排序

```go
// 根据key排序
func sortKey(mp map[string]int) {
   var newMp = make([]string, 0)
   for k, _ := range mp {
      newMp = append(newMp, k)
   }
   sort.Strings(newMp)
   for _, v := range newMp {
      fmt.Println("根据key排序后的新集合》》   key:", v, "    value:", mp[v])
   }
}

// 根据value排序
func sortValue(mp map[string]int) {
   var newMp = make([]int, 0)
   var newMpKey = make([]string, 0)
   for oldk, v := range mp {
      newMp = append(newMp, v)
      newMpKey = append(newMpKey, oldk)
   }
   sort.Ints(newMp)
   for k, v := range newMp {
      fmt.Println("根据value排序后的新集合》》  key:", newMpKey[k], "    value:", v)
   }
}
```

## ref

[剖析golang map的实现](https://www.jianshu.com/p/092d4a746620)
