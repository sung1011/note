# go start

## install

1. `wget https://golang.org/dl/go1.16.4.linux-amd64.tar.gz`
2. `rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.4.linux-amd64.tar.gz`
3. `export PATH=$PATH:/usr/local/go/bin`
4. `go version`

> doc `https://golang.org/doc/install`

## go env

| name        | describe                                 |
| ----------- | ---------------------------------------- |
| CGO_ENABLED | 指明cgo工具是否可用的标识。              |
| GOARCH      | 程序构建环境的目标计算架构。             |
| GOBIN       | 存放可执行文件的目录的绝对路径。         |
| GOCHAR      | 程序构建环境的目标计算架构的单字符标识。 |
| GOEXE       | 可执行文件的后缀。                       |
| GOHOSTARCH  | 程序运行环境的目标计算架构。             |
| GOOS        | 程序构建环境的目标操作系统。             |
| GOHOSTOS    | 程序运行环境的目标操作系统。             |
| GOPATH      | 工作区目录的绝对路径。                   |
| GORACE      | 用于数据竞争检测的相关选项。             |
| GOROOT      | Go语言的安装目录的绝对路径。             |
| GOTOOLDIR   | Go工具目录的绝对路径。                   |

## mod

### env

- `GO111MODULE` #111是对 > go1.11 有效的意思
  - `on` 所有目录使用`go module` (>= 1.16默认值)
  - `off` 禁用`go module`
  - `auto` 在GOPATH目录下使用传统`GOPATH`和`vendor`查找依赖, 非GOPATH使用`go module` (>= 1.11默认值)

> go module 是go1.11引入的概念, 解决对GOPATH的依赖

### go.mod file

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

### cmd

- `go mod`
  - `graph` print module requirement graph
  - `init` initialize new module in current directory
  - `vendor` make vendored copy of dependencies
  - `verify` verify dependencies have expected content
  - `why` why packages or modules are needed
  - `download` download modules to local cache ($GOPATH/pkg/mod)
  - `tidy` add missing and remove unused modules
  - `edit` edit go.mod from tools or scripts
    - `go mod edit -go=1.16`
    - `go mod edit -require tickles.cn/tk@v1.2.0` 增改指定require
    - `go mod edit -droprequire tickles.cn/tk` 删除指定的require, 不需要version
    - `go mod edit -replace example.com/a@v1.0.0=some.com/a@v1.0.0`
    - `go mod edit -replace example.com/a@v1.0.0=./a`
    - `go mod edit -dropreplace example.com/a@v1.0.0`
    - `go mod edit -json`
    - `go mod edit -fmt`

- `go list`
  - `go list -m all` list modules and dependencies instead of packages.
  - `go list -m -u all` list modules, dependencies and upgrade them.

- `go get`
  - `go get -u` update modules providing dependencies of packages

> go build/run 时有一个hook, 会自动调用go mod tidy, 该行为控制通过`go build -mod=readonly`或`go build -mod=vendor`
>
> go mod下载的包`按版本`保存在`$GOPATH/pkg/mod/`
