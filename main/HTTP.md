# HTTP

## 概念

基于TCP协议，属于OSI应用层的面向对象的超文本传输协议。

### Request

#### 请求行 start line

- **METHOD**
  - GET: 获取数据
  - HEAD: 只解析 header
  - POST: 新增数据（也改数据）
  - PUT: 修改数据
  - DELETE: 删除数据
  - CONNECT: 预留
  - OPTIONS: 查看可用的方法
  - TRACE: debug
  - 扩展
    - MKCOL
    - COPY
    - MOVE
    - LOCK
    - UNLOCK
    - PATCH: 局部PUT
- **URI**
  - scheme
  - host
  - port
  - path
  - query
- **VERSION**
  - HTTP/0.9
  - HTTP/1.0
    - 新增HEAD、POST等方法
    - 新增响应状态码，标记状态
    - 新增协议版本号概念
    - 新增header概念
    - 传输数据不限于文本
  - HTTP/1.1
    - 新增PUT、DELETE等方法
    - 新增缓存管理和控制
    - 允许持久连接，默认keep-alive，减少握手开销
    - 允许响应数据分块（chunked）
    - 支持host域参数
  - HTTP/2.0
    - 多路复用(多重请求)，并废除了1.1的管道
    - 二进制分帧（不再纯文本）
    - 流优先级
    - 服务器主动推送
    - 头部压缩(只传diff)
  - HTTP/3.0
  - HTTPS
    - ssl原理
    - CA证书
    - 签名
    - 公钥
    - 私钥
    - 对称秘钥

#### 请求头 header

- Accept-Encoding
- Authorization
- Cookie
- Content-Length
- Content-Type
- User-Agent
- If-None-Match: 客户端的ETag值，发给服务器以判定是否匹配服务器ETag
- If-Modify-Since: 从某时间数据是否修改。(以时间维度判定cache有效性)

#### 请求空行 CRLF

#### 请求体 entity/body

### Response

#### 响应行

- **Resp VERSION**
  - HTTP/1.0
  - HTTP/1.1
  - HTTP/2.0
  - HTTP/3.0

- **[status code](HTTP-statuscode.md)**

- **status message**

#### 响应头

- Allow: 支持哪些request method
- Content-Encoding  
- Content-Length  
- Content-Type
- Expires  
- Set-Cookie  
- Date  
- Expires
- Location
- Refresh
- Last-Modified: 最后一次改动时间。无改动时返回304。 (以时间维度判定cache有效性)
- ETag: 缓存的令牌

#### 响应体  

### 通讯流程

#### 三次握手

流程

- C -->> S: SYN seq=j // 在吗
- C <<-- S: ACK j+1, SYN seq=k // ok, 在
  - S: syn_recv
- C -->> S: ACK k+1 // ok, 建立连接吧
  - C: syn_send
  - C: established
  - S: established

意义

- 当第一次握手出现网络延迟， 第二次握手依然会正常响应， 这时第三次握手就会被抛弃，服务端没收到第三次握手就不会建立连接。反之，若无第三次握手，服务端在第二次握手就开始建立客户端已经认为过期抛弃的连接， 浪费资源。

#### 数据传输

- C -->> S: GET /foo HTTP/1.1
- C <<-- S: ACK
- C <<-- S: HTTP/1.1 200 OK
- C -->> S: ACK

#### 四次挥手

流程:

- C -->> S: fin seq=i // 我要关了
- C <<-- S: ACK i+1 // ok收到, 稍等, 我还有返回没发完呢
  - S: close_wait
  - C: fin_wait1
- C <<-- S: fin seq=j // (发完了)我关了, 你也关吧
  - S: last_wait
  - C: fin_wait2
- C -->> S: ACK j+1 // ok, 我也关了, 并且轮询几下检测是否真断了。
  - C: time_wait
  - C: closed
  - S: closed

意义

- 第二，三次挥手的意义在于，服务端可能不会立即关闭socket，故只能先回复一个ACK确认。  
- 第四次挥手会进入time_wait状态， 重发几次以确认服务端是否已经断开。
