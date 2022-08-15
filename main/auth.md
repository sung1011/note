# auth

## token

## [Cookie](cookie.md)

    S端通过Cookie在C端记录会话状态

## session

## [OAuth](oauth.md)

    分离`第三方应用(app)`和`资源所有者(user)`
    user同意以后
    app向资源鉴权服务器(res-auth-server)申请授权码
    app-server向`资源服务器(res-server)`申请令牌.
    app通过令牌请求数据.

## [JWT](JWT.md)

    S认证后, 返回一段密文(JWT)给C, JWT包含账号非敏感信息, 只有S能解开, 改过程不依赖额外数据 只依赖一个S端的secretKey.
    作用: 认证登陆过, 传递payload信息

## SSO