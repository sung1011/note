# todo

- stdout stdin stderr 重定向

## check

- crontab每个task的执行时间
- slow SQL

## issues

- 某DB实例压力大
- 某玩家请求多
- trace一次请求的所有log(不跨服务、跨服务)
- 慢请求
- pay记录过多如何处理
- 支付校验都有啥
- 如何管道标准错误(而非标准输出) `http://www.dovov.com/854.html`
- git ref `https://www.php.cn/manual/view/35099.html`
- why the design `https://draveness.me/whys-the-design/`
- ssh `http://www.ruanyifeng.com/blog/2011/12/ssh_port_forwarding.html`

## code

- etcd  
- mq  
- CURL  
- gRPC  
- protobuff  
- nginx源码  
- redis源码  

## tool

- tk  
- 批量文件改名 前缀 后缀 匹配  
- dotfile  
- ssh 秘钥管理 IP列表管理  

## other

```go
// matcher为副本
matcher, exists := matchers[feed.Type]

// v为副本
for _, v := range feeds {

}

// chan传递副本

// 如果一个接口类型只包含一个方法, 那类型名需要以er结尾 ex: Reader, Writer, Matcher...
type Matcher interface {
  Search(feed *Feed, searchTerm string) ([]*Result, error)
}

```