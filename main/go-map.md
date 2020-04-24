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
var m := map[string]int // nil map; panic if written to
m := map[string]int{} // empty map
m := make(map[string]int, cap) // empty map
```

## 返回值 key是否存在

```go
m := map[string]int{"a":1, "b":2}
v1, exists1 := m["a"] // 1, true
v2, exists2 := m["xxx"] // 0, false
```

## len(), cap(), delete()

## 并发安全

```go
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
var sm sync.Map
sm.Store(1, "a")
if v,ok:=sm.Load(1);ok{
    fmt.Println(v)
}
```

## 拷贝

```go
originalMap := make(map[string]int)
originalMap["one"] = 1
originalMap["two"] = 2

targetMap := make(map[string]int)

for k, v := range originalMap {
    targetMap[k] = v
}
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
mf := map[int]func() int{
    1: func() int {return 10},
    2: func() int {return 20},
}

// map_func := make(map[KEY_TYPE]func() RETURN_TYPE)
mf := make(map[int]func() string)
mf[1] = func() string{ return "10" }
mf[2] = func() string{ return "20" }
```

## ref

[剖析golang map的实现](https://www.jianshu.com/p/092d4a746620)
