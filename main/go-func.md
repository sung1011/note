# go func

## 内置函数

```go
$ go doc builtin | grep func

// 关闭channel
func close(c chan<- Type)

// 删除map中的元素
func delete(m map[Type]Type1, key Type)

// 处理错误
func panic(v interface{})
func recover() interface{}

// 底层的输出函数，用来调试用。在实际程序中，应该使用fmt中的print类函数
func print(args ...Type)
func println(args ...Type)

// 操作复数(虚数)
func complex(r, i FloatType) ComplexType
func imag(c ComplexType) FloatType
func real(c ComplexType) FloatType

// 追加slice
func append(slice []Type, elems ...Type) []Type

// 分配内存并初始化
func make(t Type, size ...IntegerType) Type     // slice、map、channel
func new(Type) *Type

// 获取slice的容量
func cap(v Type) int

// 拷贝slice
func copy(dst, src []Type) int

// 获取
// - slice的长度
// - map的元素个数
// - array的元素个数
// - 指向array的指针时获取array的长度
// - string的字节数
// - channel的channel buffer中的未读队列长度
func len(v Type) int
```