# 二叉树

    每个节点拥有0~2个子节点.

```bash
                a                 Level0
              /    \
            b        c            Level1
           /  \     / \
         d     e   f   g          Level2
       /  \   / \ / \  / \
      h    i j                    Level3

# Root: a
# Sub-tree: d h i ; 
#   - Parent-Node: d
#   - Child-Node: h i
#   - Left-Node: h
#   - Right-Node: i
# Siblings: b&c / d&e / f&g
```


## 关联

    `双向链表`中有prev和next指针, 如果每个节点有两个next,就成了`二叉树`.
    `二叉树`如果节点间互相连接, 就是`图 graph`.

## ref

> [双向链表](ds-linkedlist.md)  

> [二叉树](ds-binary-tree.md)  
