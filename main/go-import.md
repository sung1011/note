# go import

## 导入

```go
// 单行导入
import "path1"   // 注意: 这里是路径, 而非包名
import "path2"

// 多行导入
import (
    "path1"
    "path2"
)
```

## 自定义包名

```go
import (
    "crypto/rand"
    mrand "math/rand" // 将名称替换为mrand避免冲突
)
```

## 匿名导入包

```go
import (
    _ "path/to/package" // 触发导入包的init()
)
```

## 包的初始化函数 init()

```go
func init() {

}
```

> 初始化顺序: 以深度优先的顺序. 如包的引用关系为main→A→B→C, 则`init()`的调用顺序为C.init→B.init→A.init→main

> 同一个包的多个`init()`顺序随机

> `init()`不能被其他函数调用

