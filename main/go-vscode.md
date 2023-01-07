# go vscode

## 填充空值到结构体

```go
type Foo struct{
    Name string
    Count int32
}

func main() {
    foo := &Foo{}
    fmt.Println(foo)
}
```

- 方法1
  1. 命令: Fill Struct

- 方法2
  1. 光标到 &Foo{}
  2. 点击 `quick fix` (cmd+.)
  3. 点击 `rewrite... fill Person{}`

```go
func main() {
    foo := &Foo{
        Name: "",
        Age:  0,
    }
    fmt.Println(foo)
}
```

## 实现接口

```go
package main

// User represent Person
type Person interface{
    // GetName return person's name
    getName() string
    // GetAge return person's age
    getAge() string
}

type Student struct{
    Name string
    Age int32
}
```

1. 命令:`Generate Interface Stubs`
2. 输入参数`s *Student main.Person`. 即生成如下接口的实现方法代码(with 注释)

```go
// GetName return person's name
func (stu *Student) getName() string {
	panic("not implemented") // TODO: Implement
}

// GetAge return person's age
func (stu *Student) getAge() string {
	panic("not implemented") // TODO: Implement
}
```

## 生成单元测试用例

> package: gotests

```go
func add(a, b int) int {
	return a + b
}
```

1. 命令:`Generate Unit Tests For File / Function / Package`

```go
func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
```

## 生成tag

```go
type Student struct{
    Name string
    Age int32
}
```

1. 命令:`Add Tags To Struct Fields`

```go
/**
vscode配置项:
   "go.addTags": {
     "tags": "json,xml",
     "options": "json=omitempty",
     "promptForTags": false,
     "transform": "snakecase", // snakecase下划线; camelcase小驼峰
     "template": ""
   }
*/
type Student struct{
    Name string `json:"name,omitempty" xml:"name"`
    MyAge int32   `json:"my_age,omitempty" xml:"my_age"`
}
```

## 跳转查看接口的实现

1. 命令:`go to implementations` (/右键)

## 同时重命名关联数据

1. 命令:`rename symbol` (/右键)

## 提取变量 / 方法

> package: godoctor

```go
func add(a, b int) int {
	if a < 0 || b < 0 {
		panic("negative")
	}
	return a + b
}
```

- 方法1
  1. 选中 `a < 0 || b < 0`
  2. 点击快速修复(cmd+.)
  3. 点击 `Extract variable`
- 方法2
  1. 选中 `a < 0 || b < 0`
  2. 命令 `Extract variable`

- 提取方法类似

## 保存的钩子 test / format / vet(是否可运行) / lint(严格性)

1. setting中搜索 `go:save`
2. 查看修改 build on build/test/cover/vet/lint

## 右键快捷方式

1. 修改vscode.setting; 命令:`preference: open user settings (JSON)`

```setting
    "go.editorContextMenuCommands": {
        "toggleTestFile": false,
        "addTags": true,
        "removeTags": false,
        "fillStruct": false,
        "testAtCursor": false,
        "testFile": false,
        "testPackage": false,
        "generateTestForFunction": true,
        "generateTestForFile": false,
        "generateTestForPackage": false,
        "addImport": true,
        "testCoverage": true,
        "playground": false,
        "debugTestAtCursor": false,
        "benchmarkAtCursor": false
    },
```

## 语言解析

> package: gopls

      语言服务解析, 不依赖于IDE, 而是独立的服务来后台解析实现解析, 提示, 跳转等
