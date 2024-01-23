# go map

## 数据结构

```js
// bucket0
----------
k[0] k[1] k[2] k[3] k[4] k[5] k[6] k[7]
v[0] v[1] v[2] v[3] v[4] v[5] v[6] v[7]
overflow  // 8个key都填满 且尚未达到扩容条件, 会新建立bucket挂在旧bucket的overflow上, bucket之间成链表状态
---------

// bucket1
----------
k[0] k[1] k[2] k[3] k[4] k[5] k[6] k[7]
v[0] v[1] v[2] v[3] v[4] v[5] v[6] v[7]
overflow  
-------|--
       |
        // bucket1-1
        --------
        k[0] k[1] k[2] k[3] k[4] k[5] k[6] k[7]
        v[0] v[1] v[2] v[3] v[4] v[5] v[6] v[7]
        overflow  
        --------


// key
v, ok := m[key]

hashcode = Hash(key) = XXXXYYYY   (高位区:XXXX 低位区:YYYY)

低位区 决定bucket的位置, 避免遍历bucket (如bucket1)
高位区 决定bucket中key的位置, 避免遍历key (如bucket1.k4)

// val
如: map[int8]int64

    go方案 节约空间
    8+64 = 72
    kkkkkkkk vvvvvvvv vvvvvvvv vvvvvvvv ...
    k占用8 v占用8 v占用8 v占用8


    kv紧邻方案 padding浪费空间
    8*16 = 128
    k_______ vvvvvvvv k_______ vvvvvvvv k _______vvvvvvvv ...
    k占用1 pading占用7 v占用8

// 扩容
`count > loadFactor*2^B`扩容:
新建一个两倍规模的map, assign和delete时逐步做迁移和排空
原map会挂在hmap的oldbuckets指针, 全部迁移结束才会被释放

`overflow-bucket 过多` 扩容:
重建一个规模一样的map, assign和delete时逐步做迁移和排空

```

```go
type hmap struct {
    count     int    // map的元素的个数; 即len()
    flags     uint8  // 标记读写状态，主要是做竞态检测，避免并发读写; 定义了4种状态 iterator, oldIterator, HashWriting, sameSizeGrow
    B         uint8  // 2 ^ B = bucket; bucket数量的以2为底的对数
    noverflow uint16 // overflow(溢出)的 bucket 个数
    hash0     uint32 // hash函数种子值

    buckets    unsafe.Pointer // 指向数组buckets的指针, 即bmap
    oldbuckets unsafe.Pointer // growing(扩容阶段) 原指向原buckets的指针
    nevacuate  uintptr        // growing 已迁移的个数

    extra *mapextra // 可选字段，用于存储 overflow buckets
}

type mapextra struct {
    overflow    *[]*bmap
    oldoverflow *[]*bmap

    nextOverflow *bmap
}

// 桶结构
type bmap struct {
    tophash [bucketCnt]uint8    // 8个元素 k:桶数 v:记录着每个key的高8个bits
}
```

## 创建 初始化 访问

```go
m := map[string]int{"red": 1, "green": 2, "blue": 3} // len=3; 常用

m := make(map[string]int) // empty map; len=0; cap=0 常用; 超过len会自动扩容
m := make(map[string]int, 100) // empty map; len=0; cap=100 常用, 出于性能考虑

m := map[string]int{} // empty map; prt=xxx len=0 cap=0;

// nil ptr map
// panic: assignment to entry in nil map
// "panic if written to" -- 不能赋值会报panic; 但可以查找 删除 循环 统计
var m map[string]int 
```

> `len()` 对map而言, len()是key的个数; map有cap的概念, 但无法使用cap()

> `cap()` map不能使用cap(), 但make(map[T]T, cap)时可以初始化容量

## key限制

```js
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

## key是否存在

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

```go
func main() {
	m := map[string]int{
		"k1": 1,
		"k2": 2,
	}
	foo(m)
	fmt.Println("", m) // map[k1:1 k2:222]; 引用传参, 会修改原值
}

func foo(m map[string]int) {
	m["k2"] = 222
}
```

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
// delete(), foreach(), len()
```

```go
// 分片同步锁 (性能)
var SHARD_COUNT = 32

// 分成SHARD_COUNT个分片的map
type ConcurrentMap []*ConcurrentMapShared

// 通过RWMutex保护的线程安全的分片，包含一个map
type ConcurrentMapShared struct {
	items        map[string]interface{}
	sync.RWMutex // Read Write mutex, guards access to internal map.
}

// 创建并发map
func New() ConcurrentMap {
	m := make(ConcurrentMap, SHARD_COUNT)
	for i := 0; i < SHARD_COUNT; i++ {
		m[i] = &ConcurrentMapShared{items: make(map[string]interface{})}
	}
	return m
}

// 根据key计算分片索引
func (m ConcurrentMap) GetShard(key string) *ConcurrentMapShared {
	return m[uint(fnv32(key))%uint(SHARD_COUNT)]
}

func (m ConcurrentMap) Set(key string, value interface{}) {
	// 根据key计算出对应的分片
	shard := m.GetShard(key)
	shard.Lock() //对这个分片加锁，执行业务操作
	shard.items[key] = value
	shard.Unlock()
}

func (m ConcurrentMap) Get(key string) (interface{}, bool) {
	// 根据key计算出对应的分片
	shard := m.GetShard(key)
	shard.RLock()
	// 从这个分片读取key的值
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}
```

```go
// 原子锁
var m sync.Map
m.Store("a", 1)
m.Store("b", 2)
if v,ok := m.Load("a");ok{
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