# go compare struct.field.func | type-interface

```go
// struct.field.func
package main

type Foo struct {
  	say func() error
}

func main() {
  	f := &Foo{
      say: func() error { // 每次实例化时, 实现一次成员方法say(); 可以不实现
      	return nil
      },
	}
	f.say()
}

```

```go
// type-interface
package main

type IFoo interface {
	say() error
}

type Foo struct { // 定义结构体时, 实现一次接口方法say(); 必须实现
}

// 
func (this *Foo) say() error {
    return nil
}
```