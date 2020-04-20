# shell 基础

## 参数

```shell
ls -l // 短格式 手动输入

ls --list // 长格式 脚本中使用
```

## echo

```shell
echo -n a;echo b; // -n 输出出现在同一行
ab

echo -e "hello\nword" // -e 解释特殊字符
hello
word
```

## && ||

```shell
cat a && ls a // cat成功才会ls
cat a || ls a // cat失败才会ls
```
