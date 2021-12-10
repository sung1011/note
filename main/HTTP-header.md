# HTTP header

## HTTP请求头

| Header              | 解释                                                     | 示例                                               |
| ------------------- | -------------------------------------------------------- | -------------------------------------------------- |
| Accept              | 指定C端可接受的内容类型。                                | Accept: text/plain, text/html                      |
| Accept-Charset      | 指定C端可接受的字符编码集。                              | Accept-Charset: iso-8859-5                         |
| Accept-Encoding     | 指定浏览器可以支持的web服务器返回内容压缩编码类型。      | Accept-Encoding: compress, gzip                    |
| Accept-Language     | 浏览器可接受的自然语言                                   | Accept-Language: en,zh                             |
| Accept-Ranges       | 可以请求实体的一部分，指定单位                           | Accept-Ranges: bytes                               |
| ~~Authorization~~   | HTTP鉴权，值为模式与base64(usr:pwd)，错误返401           | Authorization: Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==  |
| Cache-Control       | 指定请求和响应遵循的缓存机制                             | Cache-Control: no-cache                            |
| Connection          | 表示是否需要持久连接。（HTTP 1.1默认进行持久连接）       | Connection: close                                  |
| Cookie              | HTTP请求发送时，会把保存在该请求域名下的所有cookie值     | Cookie: $Version=1; Skin=new;                      |
| Content-Length      | 请求的内容长度                                           | Content-Length: 348                                |
| Content-Type        | 请求的与实体对应的MIME信息                               | Content-Type: application/x-www-form-urlencoded    |
| Date                | 请求发送的日期和时间                                     | Date: Tue, 15 Nov 2010 08:12:31 GMT                |
| Expect              | 请求的特定的服务器行为                                   | Expect: 100-continue                               |
| From                | 发出请求的用户的Email                                    | From: user@email.com                               |
| Host                | 指定请求的服务器的域名和端口号                           | Host: www.baidu.com                                |
| If-Match            | 只有请求内容与实体相匹配才有效                           | If-Match: “737060cd8c284d8af7ad3082f209582d”       |
| If-Modified-Since   | 指定时间之后S端有修改则200，无修改则304                  | If-Modified-Since: Sat, 29 Oct 2010 19:43:31 GMT   |
| If-None-Match       | 值为ETag，S端无修改则返回304                             | If-None-Match: “737060cd8c284d8af7ad3082f209582d”  |
| If-Range            | 如果实体未改变，服务器发送客户端丢失的部分，否则发送整   | If-Range: “737060cd8c284d8af7ad3082f209582d”       |
| If-Unmodified-Since | 只在实体在指定时间之后未被修改才请求成功                 | If-Unmodified-Since: Sat, 29 Oct 2010 19:43:31 GMT |
| Max-Forwards        | 限制信息通过代理和网关传送的时间                         | Max-Forwards: 10                                   |
| ~~Pragma~~          | 用来包含实现特定的指令                                   | Pragma: no-cache                                   |
| Range               | 只请求实体的一部分，指定范围                             | Range: bytes=500-999                               |
| Referer             | 先前网页的地址，当前请求网页紧随其后,即来路              | Referer: http://www.baidu.com/arch/71.html         |
| TE                  | 客户端愿意接受的传输编码，并通知服务器接受接受尾加头信息 | TE: trailers,deflate;q=0.5                         |
| Upgrade             | 向服务器指定某种传输协议以便服务器进行转换（如果支持）   | Upgrade: HTTP/2.0, SHTTP/1.3, IRC/6.9, RTA/x11     |
| User-Agent          | User-Agent的内容包含发出请求的用户信息                   | User-Agent: Mozilla/5.0 (Linux; X11)               |
| Via                 | 通知中间网关或代理服务器地址，通信协议                   | Via: 1.0 fred, 1.1 nowhere.com (Apache/1.1)        |
| Warning             | 关于消息实体的警告信息                                   | Warn: 199 Miscellaneous warning                    |

## HTTP响应头

| Header             | 解释                                                      | 示例                                                |
| ------------------ | --------------------------------------------------------- | --------------------------------------------------- |
| Accept-Ranges      | 表明服务器是否支持指定范围请求及哪种类型的分段请求        | Accept-Ranges: bytes                                |
| Age                | 从原始服务器到代理缓存形成的估算时间（以秒计，非负）      | Age: 12                                             |
| Allow              | 允许的HTTP方法，不允许则返回405                           | Allow: GET, HEAD                                    |
| Cache-Control      | 告诉所有的缓存机制是否可以缓存及哪种类型                  | Cache-Control: no-cache                             |
| Content-Encoding   | web服务器支持的返回内容压缩编码类型。                     | Content-Encoding: gzip                              |
| Content-Language   | 响应体的语言                                              | Content-Language: en,zh                             |
| Content-Length     | 响应体的长度                                              | Content-Length: 348                                 |
| Content-Location   | 请求资源可替代的备用的另一地址                            | Content-Location: /index.htm                        |
| Content-MD5        | 返回资源的MD5校验值                                       | Content-MD5: Q2hlY2sgSW50ZWdyaXR5IQ==               |
| Content-Range      | 在整个返回体中本部分的字节位置                            | Content-Range: bytes 21010-47021/47022              |
| Content-Type       | 返回内容的MIME类型                                        | Content-Type: text/html; charset=utf-8              |
| Date               | 原始服务器消息发出的时间                                  | Date: Tue, 15 Nov 2010 08:12:31 GMT                 |
| ETag               | 内容实体的令牌(哈希值)                                    | ETag: “737060cd8c284d8af7ad3082f209582d”            |
| Expires            | 响应过期的日期和时间                                      | Expires: Thu, 01 Dec 2010 16:00:00 GMT              |
| Last-Modified      | 请求资源的最后修改时间                                    | Last-Modified: Tue, 15 Nov 2010 12:45:26 GMT        |
| Location           | 用来重定向接收方到非请求URL的位置来完成请求或标识新的资源 | Location: http://www.baidu.com/arch/94.html         |
| Pragma             | 包括实现特定的指令，它可应用到响应链上的任何接收方        | Pragma: no-cache                                    |
| Proxy-Authenticate | 它指出认证方案和可应用到代理的该URL上的参数               | Proxy-Authenticate: Basic                           |
| refresh            | 应用于重定向或一个新的资源被创造，在5秒之后重定向         | Refresh: 5; url=http://www.baidu.com/arch/94.html   |
| Retry-After        | 如果实体暂时不可取，通知客户端在指定时间之后再次尝试      | Retry-After: 120                                    |
| Server             | web服务器软件名称                                         | Server: Apache/1.3.27 (Unix) (Red-Hat/Linux)        |
| Set-Cookie         | 设置Http Cookie                                           | Set-Cookie: UserID=JohnDoe; Max-Age=3600; Version=1 |
| Trailer            | 指出头域在分块传输编码的尾部存在                          | Trailer: Max-Forwards                               |
| Transfer-Encoding  | 文件传输编码                                              | Transfer-Encoding:chunked                           |
| Vary               | 告诉下游代理是使用缓存响应还是从原始服务器请求            | Vary: *                                             |
| Via                | 告知代理客户端响应是通过哪里发送的                        | Via: 1.0 fred, 1.1 nowhere.com (Apache/1.1)         |
| Warning            | 警告实体可能存在的问题                                    | Warning: 199 Miscellaneous warning                  |
| WWW-Authenticate   | 表明客户端请求实体应该使用的授权方案                      | WWW-Authenticate: Basic                             |

## 关联

- 缓存
  - pragma
  - cache-control
  - expire
- 分段
  - accept-ranges
  - range

- GZIP (、ZLIB、DEFLATE)

  - 场景1：获取S端临时压缩的数据(如: nginx压缩并返回)，在C端解压。

    1. req header `Accept-Encoding: gzip` 表达C端想要gzip
    2. resp header `Content-Encoding: gzip` S端压缩数据、返回此响应头与压缩后的数据
       - resp header `Content-Length: 1234` 为压缩后的大小
    3. 客户端（浏览器）收到该响应头，进行解压（流式）

  - 场景2：获取S端已压缩的数据(如：上传时已经压缩过的文件)，在C端解压。
    1. 不要带req header `Accept-Encoding: gzip` 发起普通请求
    2. resp header `Content-Encoding: gzip` 返回此响应头与压缩后的数据，S端需要预先设定规则，划定哪些文件是gzip格式，返回此header。
    3. C端（如：浏览器）收到该响应头，进行解压（流式）

  - 场景3：压缩请求数据，发起请求
    1. req header `Content-Encoding: gzip` 在C端代码进行压缩，带此header告知S端请求格式为gzip

> 优点 压缩

> 缺点 压缩与解压耗时耗CPU

> 原理 找到类似的字符串, 并替换. 所以对很多标签的代码文件压缩率高（xml, html）

## ref

`https://www.cnblogs.com/LO-ME/p/7377082.html`
