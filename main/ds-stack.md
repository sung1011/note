# 栈

    先进后出的数组或链表

## 场景

    先进后出

## 操作

- push()
- pop()

## 实战

- `浏览器前进后退` 后退栈 前进栈

```bash
浏览网站abcde后, 点击两次后退的过程

# 后退栈
e d c b a  -> d c b a -> c b a
```

- `运算表达式` 操作数栈+运算符栈

```bash
# 1+2*3 出栈过程

# 操作数 栈
3 2 1  ->  6 1  ->  7
# 运算符 栈
 * +   ->   +   ->  _
```

- `函数调用栈` 函数代码栈

```bash
# 函数
main() {
    a = 1, b = 2
    rs = add(a, b)
}

add(x, y) {
    var sum int
    sum = x + y
    return sum
}

# 代码栈
x = 1
y = 2
sum = 3
...
rs = 0
a = 1
b = 2
```

- `最小栈` O(1)获取栈的最小元素

```bash
# push()时额外保存当前最小栈, 在pop()后不需要重新计算；队列同理.

stack = [10, 12, 9, 14]; # min = stack[2] = 9
minsStackIndex = [0, 2];

# push()
stack = [10, 12, 9, 14, 4]; # min = stack[4] = 4
minsStackIndex = [0, 2, 4]; # push()时发现4是最小栈, 则将其索引值同步push()到额外栈中.

# pop() pop() pop()
stack = [10, 12]; # min = stack[0] = 10
minsStackIndex = [0]; # pop()时同步pop()
```
