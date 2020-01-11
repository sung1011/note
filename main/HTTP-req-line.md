# HTTP请求行

```HTTP
GET /hello?foo=bar HTTP/1.1
```

## **METHOD**

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

## **URI**

- scheme
- host
- port
- path
- query

## **VERSION**

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
