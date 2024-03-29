# JWT json web token

@@json @@jwt @@跨语言 @@协议 @@加密 @@SSO @@非对称加密

    S认证后, 返回一段密文(JWT)给C, JWT包含账号非敏感信息, 只有S能解开, 该过程不依赖额外数据, 只依赖一个S端的secretKey.
    作用: 认证登陆过, 传递payload信息

> 跟普通token对比, JWT是加密计算出来的 而非 保存在server的

## 特征

- 跨语言
- 不需要保存在server, 对内存和分布式友好
- 传输数据量小

## 用途

- 认证
- 交换payload信息(无法篡改)

## 缺点

- 无法使用过程中废止某个token(因为S不保存会话信息)

## 流程

1. C发送账号密码
2. S验证, 并返回`JWT`
3. C通过每个请求的 HTTP header 将`JWT`传递给服务端
4. S对`JWT`验签, 并获取payload信息

> HTTP-header: `Authorization: Bearer {JWT}`  

## 构成

```js
JWT = `{header}.{payload}.{signature}`
```

### header 头部

声明类型和算法

```js
{
    'typ': 'JWT',  // 类型
    'alg': 'HS256' // 算法
}

header = base64(data) // eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9
```

### payload 载荷 (claims)

      CS共享一些非敏感信息(C端轻易可解密)

```js
{
   // registered claims 注册的声明
   iss: jwt签发者 // issuer
   sub: jwt所面向的用户 // subject
   aud: 接收jwt的一方 // audience
   exp: jwt的过期时间, 这个过期时间必须要大于签发时间 // expireation time
   nbf: 定义在什么时间之前, 该jwt都是不可用的. // not before
   iat: jwt的签发时间 // issued at
   jti: jwt的唯一身份标识, 主要用来作为一次性token,从而回避重放攻击. // JWT ID

   // public claims 公共的声明(自定义)
   uid: 1001,
   name: sun,

   // private claims 私有的声明(自定义) 用于在同意使用它们的各方之间共享信息, 并且不是注册的或公开的声明.
   foo: 123,
}

payload = base64(data) // eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9
```

> 实际就是预定义声明 和 自定义声明  

> JWT默认不加密, 因此payload中不要存敏感信息  

> 也可以将生成JWT再加密一层

### signature 签名 放篡改

```js
data = base64UrlEncode(header) + '.' + base64UrlEncode(payload);
secretKey = 'secretKey_in_server' // secretKey/salt/key/秘钥 保存在server中
hashMethod = header.alg // HMACSHA256

signature = base64UrlEncode(hashMethod(data, secretKey)); // TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ
```

> key 可以是对称加密秘钥, 也可以是非对称加密的公钥+私钥, 决定于加密算法

## code

```go
package main

import (
   "fmt"
   "log"
   "github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secret_key_in_server")

type JWTClaims struct {
   Foo string `json:"foo"`
   UID int `json:"uid"`

   jwt.StandardClaims
}

func genToken() (string, error) {
   claims := JWTClaims{
      Foo: "bar",
      UID: 654,

      StandardClaims: jwt.StandardClaims{
         Issuer:    "sunji",
         ExpiresAt: 3600,
      },
   }
   signMethod := jwt.GetSigningMethod("HS256")
   token := jwt.NewWithClaims(signMethod, claims)
   return token.SignedString(secretKey) // eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJ1aWQiOjY1NCwiZXhwIjozNjAwLCJpc3MiOiJzdW5qaSJ9._ckmHA0u6szAZqvij_hlJiSMP1O1fgYXxtfTEkFfp4U
}

func parseToken(tokenString string) {
   token, _ := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
      return secretKey, nil
   })
   if cl, ok := token.Claims.(*JWTClaims); ok && token.Valid {
      fmt.Printf("uid: %v, foo: %v", cl.UID, cl.Foo) // uid: 654, foo: bar
   } else {
      log.Fatal(err)
   }
}

```
