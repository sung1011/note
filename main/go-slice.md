# go 切片

## 数据结构

```bash
    `ptr` 底层array被切的第一个元素地址 (切口)
    `len` 切出的元素个数
    `cap` 从切口到底层array结尾的个数

    # `占用` 固定24byte = ptr8 + len8 + cap8
```

## 创建 初始化 访问

```go
sl := make([]int, 5) // len=5 cap=5
sl := make([]int, 3, 5) // len=3 cap=5
sl := []int{30, 50, 20}
sl := []string{99: "foo"} // 100个元素
var sl []int

// sl := make([]int) // missing len argument to make([]int)
```

## 数组 与 切片

```go
// 数组
arr := [...]string{"a", "b", "c", "d", "e", "f"}
// 切片
sl1 := arr[2:4] // [c, d] len=2 cap=4 (即6-2)
// 切片的切片, cap进一步缩小
sl1_1 := sl1[1:3] // [d, e] len=2 cap=3 (即4-1)
// 索引3限制cap
sl2 := arr[2:4:5] // [c, d] len=2 cap=3 (即5-2)
```

## 赋值 (保持引用)

```go
// 切片赋值, 会同步修改底层数组的元素值.
arr := [...]string{"a", "b", "c", "d"}
sl := arr[2:] // [c d]
sl[0] = "ccc" // arr: [a b ccc, d]
```

## 拷贝 (解引用)

```go
// 仅覆盖值 && 不改变len, cap && 不报错
n := copy([]int{1, 2, 3}, []int{666, 777, 888, 999}) //sl1 = {666, 777, 888}; n = 3 (即copy了3个值)

n := copy([]int{1,2,3}, []int{666, 777}) //sl1 = {666, 777, 3}; n = 2 (即copy了2个值)
```

> `copy(dst, src)` go doc builtin copy

## append 增长/扩容

```go
sl := make([]string, 1, 4); //[] len=1; cap=4

// cap足够时 增长: 增加可用元素 & 同步修改底层数组的元素值.  len++
sl1 := append(sl, []string{"a", "b", "c"}...) // ["" a b c] len=4 cap=4; sl&sl1共享底层arr

// cap不足时 扩容: 新建底层数组(cap倍增).  len++ cap+++++
sl2 := append(sl1, "x") // ["" a b c x] len=5 cap=8; sl2新建了底层arr
```

> append的扩容机制: 更换底层array; cap < 1000时翻倍扩;  cap >= 1000 每次扩25%

## 删除

```go
// 删除开头
a = []int{1, 2, 3}
a = a[1:] // 删除开头1个元素
a = a[N:] // 删除开头N个元素

// 删除开头, 且不移动数据指针
a = []int{1, 2, 3}
a = append(a[:0], a[1:]...) // 删除开头1个元素
a = append(a[:0], a[N:]...) // 删除开头N个元素

// 删除开头, 且不移动数据指针
a = []int{1, 2, 3}
a = a[:copy(a, a[1:])] // 删除开头1个元素
a = a[:copy(a, a[N:])] // 删除开头N个元素


// 删除中间
a = []int{1, 2, 3, ...}
a = append(a[:i], a[i+1:]...) // 删除中间1个元素
a = append(a[:i], a[i+N:]...) // 删除中间N个元素
a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
a = a[:i+copy(a[i:], a[i+N:])] // 删除中间N个元素

// 删除尾部
a = []int{1, 2, 3}
a = a[:len(a)-1] // 删除尾部1个元素
a = a[:len(a)-N] // 删除尾部N个元素

// 删除指定位置
seq := []string{"a", "b", "c", "d", "e"}
index := 2
fmt.Println(seq[:index], seq[index+1:]) // [a, b] [d, e]
seq = append(seq[:index], seq[index+1:]...)
fmt.Println(seq) // [a, b, c, d]
```

## 并发安全

        与map一样, 也需要考虑并发安全
        同步锁 / chan

## nil切片 & 空切片

- nil 切片

  `var sl []int`  ptr为nil, len 0, cap 0  

- 空 切片

  `sl := make([]int, 0)` ptr有值(指向底层的空array), len 0, cap 0

## 作为参数

      直接传递, ptr+len+cap 的值 (24byte)

## 迭代

```go
// v是每个元素的副本, v实质是items的副本在迭代 (修改循环内v 不会同步修改sl)
items = []string{"a", "b", "c"} //或 items = [...]string{"a", "b", "c"}
for i, v := range items {
    v = "xxx"
    fmt.Println(items, v) // [a b c], xxx
}
```

> [迭代](src/go/basic/range_test.go) range时增删不影响循环次数; range改立刻生效; range只创建一次i, v

## 内存GC

      问题: 当切片对应的底层数组很大, 而GC不会回收正在被引用的对象, 造成内存浪费.  
      解决: 当函数的返回值是指向底层数组的数据结构(如slice), 应在函数内copy slice到新slice并返回新slice.
      效果: 函数退出时老slice对应的较大的底层数组会被回收, 保存在内存的是新的小slice.

```go
func foo() {
  old_sl := ... // 切片返回值
  new_sl := make([]T, len(old_sl)) // 空切片 指针指向新的底层数组
  copy(new_sl, old_sl) // 深拷贝到新切片
  return new_sl // 新切片返回值
}
```

## ref

- <`https://blog.csdn.net/lengyuezuixue/article/details/81197691`>
