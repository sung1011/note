# cookie

    Cookie通过在客户端记录会话状态, 确定用户身份
    S端可以修改C端的cookie(响应头`Set-Cookie`), C对S的每个请求都会带有全部cookie数据

## 不可跨域

    浏览器保证了google只能操作google的cookie, baidu只能操作baidu的cookie, 不同域名间cookie是隔离的

## 编码

    UTF-8中文
    base64图片

## 最大限制

    4kb

> 是指一个name=value的value的限制, 而非一个域名下所有cookie

> 每个请求都会带着全部cookie

## 安全性

TODO

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
