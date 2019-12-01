# HTTP

## 概念

基于TCP协议，属于OSI应用层的面向对象的超文本传输协议。

## 版本

[HTTP/0.9](HTTP-0.9.md)

[HTTP/1.0](HTTP-1.0.md)

[HTTP/1.1](HTTP-1.1.md)

[HTTP/2](HTTP-2.md)

[HTTP/3](HTTP-3.md)

## Request

### 请求行 line

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
  - HTTP/1.1
  - HTTP/2
  - HTTP/3
  - HTTPS
    - ssl原理
    - CA证书
    - 签名
    - 公钥
    - 私钥
    - 对称秘钥

### 请求头 header

- Accept-Encoding
- Authorization
- Cookie
- Content-Length
- Content-Type
- User-Agent
- If-None-Match: 客户端的ETag值，发给服务器以判定是否匹配服务器ETag
- If-Modify-Since: 从某时间数据是否修改。(以时间维度判定cache有效性)

### 请求空行 CRLF

### 请求体 entity/body

## Response

### 响应行

- **Resp VERSION**
  - HTTP/1.0
  - HTTP/1.1
  - HTTP/2
  - HTTP/3

- **[status code](HTTP-statuscode.md)**

- **status message**

### 响应头

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

### 响应体  

## [通讯流程 TCP](TCP.md)

## ref

[队头阻塞](https://cloud.tencent.com/developer/article/1509279)
[HTTP历史演变](https://www.cnblogs.com/imstudy/p/9234124.html)
