# go struct frequently

## 链表 与 双向链表

![img](res/go-struct-list.jpg)

```go
type Node struct { // 链表
    data    float64
    su      *Node // 下一节点指针
}

type Node struct { // 双向链表
    data float64
    su *Node // 下一节点指针
    pr *Node // 上一节点指针
}
```

## 二叉树

![img](res/go-struct-btree.jpg)

```go
type Tree strcut {
    le      *Tree
    data    float64
    ri      *Tree
}
```

## 重写String()

```go
type person struct {
	name string
	age  int
}

func (p *person) String() string {
	return p.name + ": " + strconv.Itoa(p.age)
}

func main() {
	p := new(person)
	p.name = "sun"
	p.age = 23
	fmt.Println(p) // sun: 23
}

```

