# go cmd

- `go`
  - `bug`         start a bug report
  - `build`       compile packages and dependencies

        编译出可执行文件

  - `clean`       remove object files and cached files
  - `doc`         show documentation for package or symbol
  - `env`         print Go environment information
  - `fix`         update packages to use new APIs
  - `fmt`         gofmt (reformat) package sources
  - `generate`    generate Go files by processing source
  - `get`         add dependencies to current module and install them
    - `go get -u` update modules providing dependencies of packages

  - `install`     compile and install packages and dependencies

        build后可执行文件安装到bin, 其依赖安装到pkg

  - `list`        list packages or modules
    - `go list -m all` list modules and dependencies instead of packages.
    - `go list -m -u all` list modules, dependencies and upgrade them.
  - `mod`         module maintenance
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

  - `run`         compile and run Go program
  - `test`        test packages
  - `tool`        run specified go tool
  - `version`     print Go version
  - `vet`         report likely mistakes in packages

> go build/run 时有一个hook, 会自动调用go mod tidy, 该行为控制通过`go build -mod=readonly`或`go build -mod=vendor`