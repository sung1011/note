# 平衡二叉树

    基于二叉搜索树, 自平衡(左右层级差不会大于1)
    比rbtree更快的查询速度

## 结构

```bash
# AVL ; left=3L right=4L
                3
              /    \
            2        5
           /        / \
         1         4   6
                        \
                         7

# 非AVL ; left=2L right=4L
                3
              /    \
            2        5
                    / \
                   4   6
                        \
                         7
```

## ref

> [二叉树](ds-binary-tree.md)  

> [二叉搜索树](ds-binary-search-tree.md)  

> `https://zhuanlan.zhihu.com/p/56066942`  

