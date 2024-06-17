# cookie

    S端通过http响应头Set-Cookie使客户端更新cookie
    C端通过Cookie请求头携带cookie给S

## 用途

- `用户标识` 登录后C端保存一串唯一字符xxx, 并在每次请求时携带, S端找到xxx对应的uid

- `填充信息, 用户喜好` 自动填充账号 邮箱

- `浏览器` 同一个域, 不同页签, 共享数据

## 方式

    S端通过http响应头Set-Cookie使客户端更新cookie
    C端通过Cookie请求头携带cookie给S

## 持久化

### Session Cookie

    临时Cookie, 浏览器关闭就消失

### persistent cookie

    持久化Cookie, 保存在磁盘, 浏览器关闭也不会消失

## 传输

    C对S的每个请求都会带有全部cookie数据

## 不可跨域

    浏览器保证了google只能操作google的cookie, baidu只能操作baidu的cookie, 不同域名间cookie是隔离的

## 编码

    UTF-8中文
    base64图片

## 最大限制

    每个kv 4kb
    (而非一个域名下所有cookie)

## 安全性

    TODO

## go http.cookie

```go
go doc http.cookie
package http // import "net/http"

type Cookie struct {
	Name  string        // key
	Value string        // value

	Path       string    // optional
	Domain     string    // optional
	Expires    time.Time // optional; 老版本浏览器
	RawExpires string    // for reading cookies only

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int         // 新版本浏览器
	Secure   bool
	HttpOnly bool
	SameSite SameSite
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}
    A Cookie represents an HTTP cookie as sent in the Set-Cookie header of an
    HTTP response or the Cookie header of an HTTP request.

    See https://tools.ietf.org/html/rfc6265 for details.

func (c *Cookie) String() string
func (c *Cookie) Valid() error


cookie.Path("/WEB16");
    代表访问WEB16应用中的任何资源都携带cookie
cookie.Path("/WEB16/cookietest");
    代表访问WEB16中的cookietest时才携带cookie信息
cookie.Domain(".foo.com");
    这对foo.com域下的所有主机都生效(如www.foo.com)，但不包括子域www.abc.foo.com

```

## 属性

| name    | type   | describe                                                                                                                                          |
| ------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| name    | str    | Cookie的名称.Cookie一旦创建,名称便不可更改                                                                                                        |
| value   | object | Cookie的值.如为Unicode字符,需要为字符编码.如为二进制数据,则需用BASE64编码.                                                                         |
| maxAge  | int    | Cookie失效的时间,单位秒.>0,则在maxAge秒之后失效.<0,为临时Cookie,关闭浏览器即失效,浏览器也不会以任何形式保存.=0,删除该Cookie.默认为–1              |
| domain  | string | 可以访问该Cookie的域名.如果设置为“.google.com”,则所有以“google.com”结尾的域名都可以访问该Cookie.注意第一个字符必须为“.”                           |
| path    | string | Cookie使用路径.如为/sessionWeb/,则只有contextPath为/sessionWeb的程序可以访问该Cookie.如果为/,则本域名下都可以访问该Cookie.注意最后一个字符必须为/ |
| secure  | bool   | Cookie是否仅被使用安全协议传输.安全协议有HTTPS,SSL等,在网络上传输数据之前先将数据加密.默认为false                                        |
| comment | string | Cookie的用处说明.浏览器显示Cookie信息的时候显示该说明                                                                                             |
| version | int    | Cookie使用的版本号.0表示遵循Netscape的Cookie规范,1表示遵循W3C的RFC 2109规范                                                                       |

## ref

- <https://www.zhihu.com/question/19786827>
