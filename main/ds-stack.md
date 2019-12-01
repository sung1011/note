# 栈

## 实战

### 最小栈

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
