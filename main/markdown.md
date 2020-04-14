# markdown 不常用格式

## 细节折叠

<details>
<summary>Click here to see terminal history + debug info</summary>
<pre>
488 cd /opt/LLL/controller/laser/
489 vi LLLSDLaserControl.c
490 make
491 make install
492 ./sanity_check
493 ./configure -o test.cfg
494 vi test.cfg
495 vi ~/last_will_and_testament.txt
</pre>
</details>

## 删除线

被~~狠心删除~~的内容

## 嵌套列表

1. 第一项：
    - 第一项嵌套的第一个元素
    - 第一项嵌套的第二个元素
2. 第二项：
    - 第二项嵌套的第一个元素
    - 第二项嵌套的第二个元素

## 区块引用

> 第一层
>> 第二层
>>> 第三层

## 多行引用

> 第1行(转移)\
第2行(双空格)  
第3行(无)
第4行

## html元素

使用 <kbd>Ctrl</kbd>+<kbd>Alt</kbd>+<kbd>Delete</kbd> 重启电脑

## 注释

[comment]: <> (哈哈我是注释，不会在浏览器中显示。)
[comment]: <> (哈哈我是注释，不会在浏览器中显示。)
[comment]: <> (哈哈我是注释，不会在浏览器中显示。)
[//]: <> (哈哈我是注释，不会在浏览器中显示。)
[//]: # (哈哈我是注释，不会在浏览器中显示。)

```comment
[comment]: <> (哈哈我是注释，不会在浏览器中显示。)
[comment]: <> (哈哈我是注释，不会在浏览器中显示。)
[comment]: <> (哈哈我是注释，不会在浏览器中显示。)
[//]: <> (哈哈我是注释，不会在浏览器中显示。)
[//]: # (哈哈我是注释，不会在浏览器中显示。)
```

## 注脚

这是一个markdown [^1] 的注脚示例。

> vscode的md不支持

[^1]: "我是Markdown注脚1"

## 连接引用

网页链接[Google][1]  
文件连接[自身文件][self]  
图片连接[图片][p]

   [1]: http://www.google.com/
   [self]: markdown.md
   [p]: res/markdownlogo.png

## 缩放图片

<img src="res/markdownlogo.png" width=30%>

![altaltalt](res/markdownlogo.png "鼠标锚点")
