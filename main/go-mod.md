# go 包依赖

    GOPATH
    vendor
    go module

## GOPATH

```bash
    通过`go get {pkg}@vX.Y.Z`下载包和其依赖到GOPATH
    运行时go编译器从GOPATH下搜索依赖的第三方包
    默认: $HOME/go/

    包版本机制
        v{X}.{Y}.{Z}
        - v 固定的版本号前缀
        - X 主版本号 major; 不同X互 代码不兼容 如: v1.8.0 & v2.0.0 不兼容
        - Y 次版本号 minor; 相同X不同Y 代码兼容 如: v1.8.0 & v1.9.0 兼容
        - Z 补丁版本号 patch; 

    ver获取规则
        go get ...@v2 会更新到v2.x.x的最新版本

    MVS (整体)最小版本原则
        proj 依赖 a包 + b包
        a 依赖 c v1.1.0
        b 依赖 c v1.3.0
        c 最新版本为 v1.7.0

        go mod 会选择v1.3.0 (满足所有包a, b依赖的最小版本), 而非最新版本
        去掉b包后, 依然保持1.3.0不会降级到1.1.0
        若增加d包 (d 依赖 c v1.6.0), go mod 会选择更新到v1.6.0
```

> `go get {pkg}` e.g. go get -v github.com/sirupsen/logrus@1.9.0

> `go list -m -versions {pkg}` 查看所有版本

## vendor

    依赖包都存放在vendor文件夹下
    运行时go优先查找vendor下的依赖
    最佳实践是vendor也提交到git, 保障版本稳定

> `缺点` 项目依然必须放在GOPATH下; 手动管理vendor; 提交代码库影响体积, 速度, 评审等

## go mod

### 操作

```bash
    `初始化`
        go mod init {domain-name}/{user-name}/{repo-name}
        #  e.g. go mod init github.com/sung1011/note

    `更新依赖信息`
        go mod tidy 
        # 默认情况下 go mod下载的包`按版本`保存在`$GOPATH/pkg/mod/`
        # 可通过 GOMODCACHE 修改go mod下载的包的目录

    `新增`
        go get {pkg}
        go mod tidy # 需要先编辑go.mod指定版本

    `升级/降级`
        go get {pkg}@vX.Y.Z # 会同步变更go.mod
        go mod tidy # 需要先编辑go.mod指定版本
        源码中直接指定主版本 e.g. import _ "github.com/go-redis/redis/v7"

```

### ENV

#### GOPROXY

    设置代理 默认为GOPROXY="https://goproxy.io,direct"

#### GO111MODULE

    go mod发布于go1.11, 111 是对 > go1.11 有效的意思
    `on` 所有目录使用`go module` (>= 1.16默认值)
    `off` 禁用`go module`
    `auto` 在GOPATH目录下使用传统`GOPATH`和`vendor`查找依赖, 非GOPATH使用`go module` (>= 1.11默认值)

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

