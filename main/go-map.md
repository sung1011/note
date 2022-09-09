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

```go
	t := make(map[int]int)
	fmt.Printf("%p\n", t)   // 0xc000090180
	println(t)              // 0xc000090180

```

## 创建 初始化 访问

```go
m := map[string]int{"red": 1, "green": 2, "blue": 3} // len=3; 常用

m := make(map[string]int) // empty map; len=0; cap=0 常用; 超过len会自动扩容
m := make(map[string]int, 100) // empty map; len=0; cap=10 常用, 出于性能考虑

m := map[string]int{} // empty map; prt=xxx len=0 cap=0;

// nil ptr map
// panic: assignment to entry in nil map
// "panic if written to" -- 没意义; 不能赋值; 可以查找 删除 循环 统计; 需要用上述方法初始化内存再用
var m map[string]int 
```

> `len()` 对map而言, len就是cap; map有cap的概念, 但无法使用cap()

## key限制

```bash
    无法 == 做比较的类型不能做map的key
    如:
        `切片 slice`
        `函数 func`
        `map`
        `(包含不可比较T的)结构体 struct{slice, func}`

# 可以做key: interface{}, chan, ptr, array
	make(map[chan int]bool)
	make(map[interface{}]bool)
    make(map[*interface{}]bool)
	make(map[[3]int]bool)
```

```go
type foo struct {
}
type bar struct {
	sl []int // 不可比较类型
}

func main() {
	var m1 map[foo]bool // ok
	var m1 map[bar]bool // 编译报错: invalid map key type bar
}
```

> `val限制` 无限制; 任何类型

## [比较](go-type-compare.md#map)

## 返回值 key是否存在

```go
m := map[string]int{"a":100, "b":200}

v := m["a"] // 100
v, exists1 := m["a"] // 100, true
v, exists2 := m["xxx"] // 0, false
```

## 不可获取元素地址

```go
    // 原因: map的增长可能会导致元素地址变更
    m := map[string]int{"sun": 1001}
    _ := &m["sun"]      // 编译错误: cannot take the address of m1["hello"]
```


## 作为参数

        直接传递, 指针的值 (8byte)

## 迭代

```go
for key,value := range my_map {
    println("key:",key," value:",value)
}
for key := range my_map {
    println("key:",key," value:",value)
}
```

> [迭代](src/go/basic/range_test.go) range时增删有可能增删循环次数; range改立刻生效; range只创建一次i, v

## 并发安全

    同一个变量在多个goroutine中访问需要保证并发安全

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

// ---------------------------------
func foo() {
    var lock sync.Mutex
    lock.Lock()
    mapList["c"] = 3
    lock.Unlock()
}
```

```go
// 原子锁
var lock sync.Map
lock.Store("a", 1)
lock.Store("b", 2)
if v,ok:=lock.Load("a");ok{
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
targetMap := make(map[string]int, 2) // 即使这里len是2, 后面也会全赋值过来
targetMap = originalMap
targetMap["d"] = 4444 // map[a:1 b:2 c:3 d:4444 e:5]
```

## 深拷贝 (解引用)

```go
originalMap := make(map[string]int)
originalMap["a"] = 1
originalMap["b"] = 2

targetMap := make(map[string]int)
for k, v := range originalMap {
    targetMap[k] = v
}
```

## delete(), len()

    删除指定key delete(map, key)
    key的个数 len(map)

```go
m := make(map[string]int)
m["a"] = 1
m["b"] = 2
fmt.Println(len(m)) // 2
delete(m, "a") // map[a:1]
fmt.Println(len(m)) // 1
```

## 集合

```go
// 推荐结构 map[Type]struct{}
sets := map[string]struct{} (
    "x": {},
    "y": {},
)

// 加锁
type inter interface{}

type Set struct {
	m map[inter]bool
	sync.RWMutex
}

func New() *Set {
	return &Set{
		m: map[inter]bool{},
	}
}

func (s *Set) Add(item inter) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}
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
// 根据key排序后迭代
func sortKey(mp map[string]int) {
   var newMpKey = make([]string, 0)
   for k, _ := range mp {
      newMpKey = append(newMpKey, k)
   }
   sort.Strings(newMpKey)
   for _, v := range newMpKey {
      fmt.Println("key:", v, "value:", mp[v])
   }
}

// 根据value排序后迭代
func sortValue(mp map[string]int) {
   var newMp = make([]int, 0)
   var newMpKey = make([]string, 0)
   for oldk, v := range mp {
      newMp = append(newMp, v)
      newMpKey = append(newMpKey, oldk)
   }
   sort.Ints(newMp)
   for k, v := range newMp {
      fmt.Println("key:", newMpKey[k], "value:", v)
   }
}
```

## ref

- [剖析golang map的实现](https://www.jianshu.com/p/092d4a746620)
