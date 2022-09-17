# go map

## 数据结构

```go
type hmap struct {
    count     int // 元素的个数
    flags     uint8 // 标记读写状态，主要是做竞态检测，避免并发读写
    B         uint8  // 可以容纳 2 ^ N 个bucket
    noverflow uint16 // 溢出的bucket个数
    hash0     uint32 // hash 因子

    buckets    unsafe.Pointer // 指向数组buckets的指针
    oldbuckets unsafe.Pointer // growing 时保存原buckets的指针
    nevacuate  uintptr        // growing 时已迁移的个数

    extra *mapextra
}

type mapextra struct {
    overflow    *[]*bmap
    oldoverflow *[]*bmap

    nextOverflow *bmap
}

// 桶结构 A bucket for a Go map.
type bmap struct {
    // tophash generally contains the top byte of the hash value
    // for each key in this bucket. If tophash[0] < minTopHash,
    // tophash[0] is a bucket evacuation state instead.
    tophash [bucketCnt]uint8    // 记录着每个key的高8个bits
    // Followed by bucketCnt keys and then bucketCnt elems.
    // NOTE: packing all the keys together and then all the elems together makes the
    // code a bit more complicated than alternating key/elem/key/elem/... but it allows
    // us to eliminate padding which would be needed for, e.g., map[int64]int8.
    // Followed by an overflow pointer.
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
m := make(map[string]int, 100) // empty map; len=0; cap=100 常用, 出于性能考虑

m := map[string]int{} // empty map; prt=xxx len=0 cap=0;

// nil ptr map
// panic: assignment to entry in nil map
// "panic if written to" -- 没意义; 不能赋值; 可以查找 删除 循环 统计; 需要用上述方法初始化内存再用
var m map[string]int 
```

> `len()` 对map而言, len()是key的个数; map有cap的概念, 但无法使用cap()

> `cap()` map不能使用cap(), 但make(map[T]T, cap)时可以初始化容量

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

- <https://www.jianshu.com/p/092d4a746620>
- <https://learnku.com/articles/33919>