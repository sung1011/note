# 栈

FILO

## 操作

- push
- pop

## 实战

- **浏览器前进后退** 后退栈+前进栈

```code
浏览网站abcde后, 点击两次后退的过程

# 后退栈
e d c b a  -> d c b a -> c b a  -> d c b a

# 前进栈
_  ->  e  ->  d e  -> e
```

- **运算表达式** 操作数栈+运算符栈

```code
1+2*3 出栈过程

# 操作数 栈
3 2 1  ->  6 1  ->  7

# 运算符 栈
* +  ->  +  ->  _
```

- **函数调用栈** 函数代码栈

```code
main() {
    a = 1, b = 2
    rs = add(a, b)
}

add(x, y) {
    var sum int
    sum = x + y
    return sum
}
```

```stack
x = 1
y = 2
sum = 3
...
rs = 0
a = 1
b = 2
```

- **最小栈**

- Q: O(1)获取栈的最小元素。
- A: 额外保存最小栈。
- E: 队列同理。

```c
stack = [10, 12, 9, 14]; // min = stack[2] = 9
minsStackIndex = [0, 2];

# push
stack = [10, 12, 9, 14, 4]; // min = stack[4] = 4
minsStackIndex = [0, 2, 4]; // push时发现4是最小栈，则将其索引值同步push到额外栈中。

# pop pop pop
stack = [10, 12]; // min = stack[0] = 10
minsStackIndex = [0]; // pop时同步pop
```
