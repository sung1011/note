# go string

    字符串是不可变的
    但可以重新赋值

## 数据结构

```go
    // $GOROOT/src/reflect/value.go; StringHeader是string的运行时表达
    // string为值类型, 固定占用16字节 = len8 + ptr8
    type StringHeader struct {
        Data uintptr // 指针; 指向底层array被切的第一个元素地址 (切口)
        Len  int
    }
```

```go
	var s = "hello"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s)) // 将string类型变量地址显式
	fmt.Printf("0x%x\n", hdr.Data)                     // 通过反射, 找到底层数组地址 0x10a30e0
	p := (*[5]byte)(unsafe.Pointer(hdr.Data))          // 获取Data字段所指向的数组的指针
	sl := (*p)[:]                                      // 赋值给切片
	for _, b := range sl {
		fmt.Printf("%c ", b) // [h e l l o ] // 输出底层数组的内容
	}

    // 运行时 
    len: 5
    data: x10a30e0  --------> [h, e, l, l, o]  # 指向底层array

```

## 操作

### 传参

    传参时, 会复制一份string 16byte
    不要传递string的指针, 会导致内存泄漏, 因为GC只会处理堆上的数据, 传递指针会导致底层数据逃逸到堆上, 极大增加GC的压力

### raw string

    所见即所得的原始字符串, 不转义
    `符号本身无法被写入raw

```go
	s := `123'"
456\n\r***\*^$
789`
```

### count

```go
	s := "你好啊"
	fmt.Println(len(s), utf8.RuneCountInString(s)) // 字节数9, 字符数3
```

### rune

    单引号表达rune(int32)类型, 单引号内是单个字符对应值为ASCII值

```go
	sl := []rune{'a', '中', '\u4e2d', '\n', ' ', '*'}
    fmt.Println(sl) // [97 20013 20013 10 32 42]

    # \u4e2d 或 \U00004e2d 是Unicode字符
```

### UTF-8

todo

### 迭代

```go
    // 下标
	s := "中国人"
	fmt.Printf("%v %T", s[0], s[0]) // 228 uint8; 得到字节, 而非字符

    s[0] = "e" // error: 字符串是不可变的

```

```go
    // for 迭代   字节视角
	s := "中国人"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v ", s[i]) // 228 184 173 229 155 189 228 186 186; 获得UTF-8的字节 
	}
```

```go
    // forrange迭代    字符视角
	s := "中国人"
	for _, v := range s {
		fmt.Println("", v) // 20013 22269 20154; 获得字符的Unicode码点值
	}
```

### 拼接

```go
    // s1+s2; 效率低
    func plusConcat(n int, str string) string {
        s := ""
        for i := 0; i < n; i++ {
            s += str
        }
        return s
    }
    // fmt.Sprintf(); 效率低
    func sprintfConcat(n int, str string) string {
        s := ""
        for i := 0; i < n; i++ {
            s = fmt.Sprintf("%s%s", s, str)
        }
        return s
    }
    // bytes.Buffer
    func bufferConcat(n int, s string) string {
        buf := new(bytes.Buffer)
        for i := 0; i < n; i++ {
            buf.WriteString(s)
        }
        return buf.String()
    }
    // []byte
    func byteConcat(n int, str string) string {
        buf := make([]byte, 0)
        for i := 0; i < n; i++ {
            buf = append(buf, str...)
        }
        return string(buf)
    }
    // []byte 已知str最终长度; 效率高
    func preByteConcat(n int, str string) string {
        buf := make([]byte, 0, n*len(str))
        for i := 0; i < n; i++ {
            buf = append(buf, str...)
        }
        return string(buf)
    }
    // strings.Builder
    func builderConcat(n int, str string) string {
        var builder strings.Builder
        for i := 0; i < n; i++ {
            builder.WriteString(str)
        }
        return builder.String()
    }
    // strings.Builder 已知str最终长度; 效率最好
    // BenchmarkBuilderConcat-8   16855    0.07 ns/op   0.1 MB/op       1 allocs/op
    // BenchmarkPreByteConcat-8   17379    0.07 ms/op   0.2 MB/op       2 allocs/op
    // 相比preByteConcat少1次内存分配, 少string转换
    func builderConcat(n int, str string) string {
        var builder strings.Builder
        builder.Grow(n * len(str))
        for i := 0; i < n; i++ {
            builder.WriteString(str)
        }
        return builder.String()
    }
    // strings.Join(str []string, sep string) string
```

- `+` vs `strings.Builder`

- `strings.Builder` vs `bytes.Buffer`

### 比较

    先len是否相同
    再ptr是否指向同一个array
    最后逐个遍历字符对比

```go
	var s1, s2 string
	s1 = "abc"
	s2 = "ab" + "c"
	fmt.Println(s1 == s2) // true

	s1 = "abc"
	s2 = "ab"
	fmt.Println(s1 == s2, s1 > s2, s1 < s2) // false true false
    // 不是比较长短, 而是比较第一个不同的字符
	s1 = "abc"
	s2 = "db"
	fmt.Println(s1 == s2, s1 > s2, s1 < s2) // false false true
```

### []byte优化互转

    不创建内存, 直接强制转换
    这是由于str与[]byte的数据结构相同(只相差一个cap), 故他们的内存布局上是对齐的, 可以直接指针替换
    制约: string是不可变的, 一旦修改[]byte会发生严重错误, defer+recover也无法捕获

```go
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func main() {
    s1 := "hello"
    b1 := StringToBytes(s1)
    fmt.Println(b1) // [104 101 108 108 111]

    b2 := []byte("hello")
    s2 := BytesToString(b2)
    fmt.Println(s2) // hello
    
    // 强转的弊端:
    // 由于string是不可变的, 所以strToBy后, 修改[]byte会发生严重错误, defer+recover也无法捕获
    s3 := "hello"
    b3 := StringToBytes(s3)
    b3[0] = 'H' // error
}
```

## ref

- 字符串拼接 <https://geektutu.com/post/hpg-string-concat.html>
- 字符串转换[]byte <https://segmentfault.com/a/1190000040289417>