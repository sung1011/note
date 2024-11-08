# trie tree 字典树, 前缀树

## 场景

- 统计和排序大量的字符串(但不限于字符串)出现频率
- 路由
- 敏感词匹配

> 敏感词匹配: 适合中等规模的敏感词数量, 大规模的敏感词可以用AC自动机

## 特性

- 根节点不包含字符, 除根节点外每个节点只包含一个字符
- 从根节点到某一节点, 路径上经过的字符连接起来, 为该节点对应的字符串
- 每个节点的所有子节点包含的字符都不相同

## 结构

```js
// 查找come
                 .
              /    \
            b        c
           /        / \
         a         d   o
        / \             \
       d   c             m
            \             \
             k             e  
    (bad)   (back) (cd)    (come)

// 每个节点值可以是字母(统计), 也可以是单词(路由)
// 节点的信息可以存 come出现次数(统计), 位置(统计), 对应的handler(路由)
```

## 实例

### 路由

```js
// 路由查找 /v1/foo/xxx/abc 的api
                .
              /    \
            v1       v2
           /        / \
         foo     login logout
        / \             \
      bar :name          abc
            \
            abc
// 实现
type node struct {
    segment string
    children []*node
    handler  func()
}
// 递归node进行匹配
// 指定路由的优先级大于通配符(:name)
// 插入路由时, 检查是否已有
```

### gin路由 (radix tree)

```js
// 路由查找 /v2/logout/abc 的api
                .
              /    \
    v1/foo/bar     v2/log
                    / \
                   in out
                        \
                         abc
// 压缩树高
// 创建慢, 查找快
```

### 敏感词匹配

```go
package main

import (
	"fmt"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) Search(text string) bool {
	for i := 0; i < len(text); i++ {
		node := t.root
		for j := i; j < len(text); j++ {
			ch := rune(text[j])
			if _, ok := node.children[ch]; !ok {
				break
			}
			node = node.children[ch]
			if node.isEnd {
				return true
			}
		}
	}
	return false
}

func main() {
	trie := NewTrie()
	sensitiveWords := []string{"badword1", "badword2", "badword3"}
	for _, word := range sensitiveWords {
		trie.Insert(word)
	}

	text := "this is a text with ==badword1== in it"
	if trie.Search(text) {
		fmt.Println("Sensitive word found !")
	} else {
		fmt.Println("Not found")
	}
}
```