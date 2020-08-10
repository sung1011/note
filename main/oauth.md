# OAuth

OAuth 引入了一个授权层，用来分离两种不同的角色：`第三方应用(app)`和`资源所有者(user)`。user同意以后，`资源服务器(res server)`可以向第三方应用颁发令牌。第三方应用通过令牌，去请求数据。

## key

- `Third-party application` 第三方应用程序(app、web-a)，又称"客户端"（client）但也可能是app的server。 --- 第三方app
- `User Agent` 用户代理。 --- 浏览器
- `Authorization server` 认证服务器(web-b-auth)，即服务提供商专门用来处理认证的服务器。 --- 微信认证服
- `Resource server` 资源服务器(web-b)，即服务提供商存放用户生成的资源的服务器。它与认证服务器，可以是同一台服务器，也可以是不同的服务器。 --- 微信
- `Resource Owner` 资源所有者，又称"用户"（user）。 --- 自己
- `HTTP service` HTTP服务提供商，简称"服务提供商"。 --- 腾讯

## [token授权](oauth-grant.md)

## [token操作](oauth-use.md)
