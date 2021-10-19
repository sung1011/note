# go mod

## env

- `GO111MODULE` #111是对 > go1.11 有效的意思
  - `on` 所有目录使用`go module` (>= 1.16默认值)
  - `off` 禁用`go module`
  - `auto` 在GOPATH目录下使用传统`GOPATH`和`vendor`查找依赖, 非GOPATH使用`go module` (>= 1.11默认值)

> go module 是go1.11引入的概念, 解决对GOPATH的依赖

## go.mod file

```bash
module github.com/sung1011/testing

go 1.14

# 包含
#
# 细节: 不用git tag管理的代码库, 会生成伪版本 v0.0.0.xxx
require (
    example.com/other/thing v1.0.2
    example.com/new/thing/v2 v2.3.4
)

# 不包含,排除
#
# 角色: 代码库是应用 (main)
# 场景: 主动排除有bug的依赖版本
# 范围: 为main项目时有效, 被引用时无效
# 
exclude (
    example.com/old/thing v1.2.3
)

# 别名
# 
# 范围: 为main项目时有效, 被引用时无效
# 例子: 由于没有科学上网 `replace golang.org/x/image v0.0.0-xxx => github.com/golang/image v0.0.0-xxx`
replace (
    example.com/bad/thing v1.4.5 => example.com/good/thing v1.4.5
    example.com/bad/code v1.4.5 => ./../good/code
)

# 撤回
#
# 角色: 代码库是依赖库
# 场景: 不小心发布了包含错误代码的版本时, 可以使用retract(撤回)来提醒依赖了本库的开发者, 这是一个有错误的版本
# 效果: 开发者会在 go get 标记撤回的版本时, 看到如下警告, 警告的内容是retract的注释
# 
# $ go get github.com/sung1011/testing@v1.9.0
#   `retracted by module author: Mistake happened in the version DO NOT USE`
retract (
   // Failed to update the message DO NOT USE
   v1.9.5
   // Mistake happened in the version DO NOT USE
   v1.9.0, 
 )
```

> doc `https://golang.org/ref/mod`

