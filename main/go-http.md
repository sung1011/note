# go http

## ResponseWriter接口

    构造HTTP response

```go
type ResponseWriter interface {
	Header() Header             // 构造Header
	Write([]byte) (int, error)  // 写入响应数据
	WriteHeader(statusCode int) // 写入状态码
}
// Header结构
type Header map[string][]string
```

## ListenAndServe

    监听服务

```go
func ListenAndServe(addr string, handler Handler) error
```

## 对比 Handle, HandleFun, Handler, HandlerFunc

### Handle

    是函数, 用来给url绑定handler

```go
// Handler用于响应一个HTTP request
// 接口方法ServerHTTP应该用来将response header和需要响应的数据写入到ResponseWriter中，然后返回。返回意味着这个请求已经处理结束，不能再使用这个ResponseWriter、不能再从Request.Body中读取数据，不能并发调用已完成的ServerHTTP方法
// handler应该先读取Request.Body，然后再写ResponseWriter。只要开始向ResponseWriter写数据后，就不能再从Request.Body中读取数据
// handler只能用来读取request的body，不能修改已取得的Request(因为它的参数Request是指针类型的)
func Handle(pattern string, handler Handler) {}
```

### HandleFunc

    是函数, 用来给url绑定handler

```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {}
```

### Handler, HandlerFunc

    是类型, 用来处理请求

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
```

## ref

- <https://www.cnblogs.com/f-ck-need-u/p/10020951.html>