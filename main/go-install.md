# go install

## install

```js
wget https://golang.google.cn/dl/go1.20.5.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

// golang.org 国外官网
```

## env

| name          | describe                                        | detail                                                    |
| ------------- | ----------------------------------------------- | --------------------------------------------------------- |
| CGO_ENABLED   | 指明cgo工具是否可用的标识.                      |
| GOARCH        | 程序构建环境的目标计算架构.                     |
| GOBIN         | 存放可执行文件的目录的绝对路径.                 |
| GOCHAR        | 程序构建环境的目标计算架构的单字符标识.         |
| GOEXE         | 可执行文件的后缀.                               |
| GOHOSTARCH    | 程序运行环境的目标计算架构.                     |
| GOOS          | 程序构建环境的目标操作系统.                     |
| GOHOSTOS      | 程序运行环境的目标操作系统.                     |
| `GOPATH`      | 工作区目录的绝对路径.                           | 目录: bin, pkg, src                                       |
| GORACE        | 用于数据竞争检测的相关选项.                     |
| `GOROOT`      | Go语言的安装目录的绝对路径.                     |
| GOTOOLDIR     | Go工具目录的绝对路径.                           | cmd: go tool                                              |
| `GOPROXY`     | 指明拉取和下载依赖包的代理地址.                 | default: proxy.golang.org,direct; 大陆: goproxy.cn,direct |
| `GO111MODULE` | 构建模式是GO-Module还是GOPATH                   | 值: on, off, auto(在gopath目录下则用gopath)               |
| `GOPRIVATE`   | 设置私有库, 跳过goproxy和校验检查; 多个逗号隔开 |
| GOINSECURE    | 私有库不支持https时设置                         |

### env cmd

```bash
    go env -w GO111MODULE=auto
    go env -w GOPRIVATE=github.com/org_name
```

## ref

- doc <https://golang.org/doc/install>

