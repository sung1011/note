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
var m := map[string]int
m := map[string]int{}
m := make(map[string]int, cap)
```

## nil map 与 空 map

nil map: `var nil_map map[string]int`  
空 map: `empty_map := map[string]int{}`, `empty_map := make(map[string]int)`

## 返回值 key是否存在

```go
m := map[string]int{"a":1, "b":2}
v1, exists1 := m["a"] // 1, true
v2, exists2 := m["z"] // 0, false
```

## len(), cap(), delete()

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
