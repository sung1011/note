# HTTP

    @@HTTP

    基于TCP协议，属于OSI应用层的面向对象的超文本传输协议。

## 版本

[HTTPS](HTTPS.md)

[HTTP/0.9](HTTP-0.9.md)

[HTTP/1.0](HTTP-1.0.md)

[HTTP/1.1](HTTP-1.1.md)

[HTTP/2](HTTP-2.md)

[HTTP/3](HTTP-3.md)

## Request

```HTTP
GET /hello?foo=bar HTTP/1.1
Accept-Language: en-us,de;q=0.5
Accept-Encoding: gzip, deflate
Content-Type: application/x-www-form-urlencoded
Content-Length: 18
Connection: Keep-Alive

action=login&uid=4

```

### [请求行 line](HTTP-req-line.md)

### [请求头 header](HTTP-header.md)

### 请求空行 CRLF

### 请求体 entity/body

## Response

```HTTP
HTTP/1.1 200 OK
Date: Sun, 15 Nov 2015 11:02:04 GMT
Server: bfe/1.0.8.9
Content-Length: 2605
Content-Type: application/javascript
Cache-Control: max-age=315360000
Expires: Fri, 13 Jun 2025 09:54:00 GMT
Content-Encoding: gzip
Set-Cookie: H_PS_PSSID=2022_1438_1944_1788; path=/; domain=test.com
Connection: keep-alive

<html>
<head><title> Index.html </title></head>
</html>
```

### [响应行 line](HTTP-resp-line.md)

### [响应头 header](HTTP-header.md)

### 响应体 body

## 网络流程

1. `HTTP` GET google.com http1.1
2. `DNS` domain name: google.com; IP 39.156.69.79
3. `TCP` src port: 5678; dst port: 80
4. `IP` src IP: 120.244.152.147; dst IP 39.156.69.79
5. `MAC` src MAC: dc:a9:04:8f:98:aa; src GateWay MAC: 32:35:2f:dc:e4:8a
6. `Router` dst {router1 ip}, {router2 ip}, {router3 ip}, {router4 ip} ...;


## [TCP协议](TCP.md)

## [跨域 cors](HTTP-cors.md)

## ref

- [队头阻塞](https://cloud.tencent.com/developer/article/1509279)
- [HTTP历史演变](https://www.cnblogs.com/imstudy/p/9234124.html)
