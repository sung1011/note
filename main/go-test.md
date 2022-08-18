# GO unittest 单元测试 & benchmark 基准测试  

## 命令  

`go test` 测试  
`go test -v` 测试详情  
`go test -cover` 覆盖率  
`go test -bench=.` 基准测试  
`go test -bench=. -benchmem` 基准测试详情  

## 约定  

框架: go内置testing包  
文件: `*_test.go`; 与被测试代码放在同一个包  
函数: 名称格式`Test[^a-z]`; 参数格式`*testing.T`  

## 样例  

目录  

```bash
├── main.go  
└── somepackage  
    ├── foo.go          被测试文件  
    └── foo_test.go     单测文件  
```

文件  

```go
package somepackage  
  
import (  
    "testing"  
)  
  
func TestFoo(t *testing.T) {  
    T.log(123)  
}  
```  

## 方式  

- [表组测试](src/go/testing/main_test.go)  
- [accert 断言](src/go/testing/main_test.go)  
- [mock 模拟](src/go/testing/main_test.go)  
- [基准测试](src/go/testing/main_test.go)  
- BDD  
