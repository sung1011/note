# go sync.Map

    并发安全的map操作

## 场景

1. 只会增长的缓存系统中，一个 key 只写入一次而被读很多次
2. 多个 goroutine 为不相交的键集读、写和重写键值对

> 不常用, 场景比较笼统, 官方建议使用后自行评估是否适合

## 数据结构

```go
type Map struct {
	mu Mutex
	// 基本上你可以把它看成一个安全的只读的map
	// 它包含的元素其实也是通过原子操作更新的，但是已删除的entry就需要加锁操作了
	read atomic.Value // readOnly
	// 包含需要加锁才能访问的元素
	// 包括所有在read字段中但未被expunged（删除）的元素以及新加的元素
	dirty map[interface{}]*entry
	// 记录从read中读取miss的次数，一旦miss数和dirty长度一样了，就会把dirty提升为read，
	misses int
}

type readOnly struct {
	m       map[interface{}]*entry
	amended bool // 当dirty中包含read没有的数据时为true，比如新增一条数据
}

// expunged是用来标识此项已经删掉的指针
// 当map中的一个项目被删除了，只是把它的值标记为expunged，以后才有机会真正删除此项
var expunged = unsafe.Pointer(new(interface{}))

// entry代表一个值
type entry struct {
	p unsafe.Pointer // *interface{}
}
```

## usage

```go
func main() {
    var m sync.Map

    m.Store("hello", "world")

    v, ok := m.Load("hello")
    if ok {
        fmt.Println("The value of 'hello' is:", v)
    }

    // LoadOrStore returns the existing value for the key if present.
    // Otherwise, it stores and returns the given value.
    v, loaded := m.LoadOrStore("hello", "world!")
    if loaded {
        fmt.Println("The value of 'hello' was loaded:", v)
    } else {
        fmt.Println("The value of 'hello' was stored:", v)
    }

    // Range calls f sequentially for each key and value present in the map.
    // If f returns false, range stops the iteration.
    m.Range(func(key, value interface{}) bool {
        fmt.Println("Key:", key, "Value:", value)
        return true
    })

    // Delete deletes the value for a key.
    m.Delete("hello")
    v, ok = m.Load("hello")
    if !ok {
        fmt.Println("'hello' was deleted")
    }
}
```