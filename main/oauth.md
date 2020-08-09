# OAuth

OAuth 引入了一个授权层，用来分离两种不同的角色：`第三方应用(app)`和`资源所有者(user)`。user同意以后，`资源服务器(res server)`可以向第三方应用颁发令牌。第三方应用通过令牌，去请求数据。

## key

- `Third-party application` 第三方应用程序(app、web-a)，又称"客户端"（client）但也可能是app的server。 --- 第三方app
- `User Agent` 用户代理。 --- 浏览器
- `Authorization server` 认证服务器(web-b-auth)，即服务提供商专门用来处理认证的服务器。 --- 微信认证服
- `Resource server` 资源服务器(web-b)，即服务提供商存放用户生成的资源的服务器。它与认证服务器，可以是同一台服务器，也可以是不同的服务器。 --- 微信
- `Resource Owner` 资源所有者，又称"用户"（user）。 --- 自己
- `HTTP service` HTTP服务提供商，简称"服务提供商"。 --- 腾讯

## 流程

    第三方应用 app --- auth req ---> res owner      // app要求用户授权
    第三方应用 app <--- auth grant --- res owner    // 用户同意授权
app
    第三方应用 app --- auth grant ---> 微信验证 auth server   // app使用上一步的授权，向认证服申请令牌
    第三方应用 app <--- access token --- 微信验证 auth server // 认证服认证，并发放令牌
app
    第三方应用 app --- access token ---> 微信 res server // app使用令牌向资源服请求资源
    第三方应用 app <--- protected resource --- 微信 res server // 资源服验证令牌，并提供资源

## feature

## 类型

- `授权码 authorization-code`
- `隐藏式 implicit`
- `密码式 password`
- `客户端凭证 client credentials`

### 授权码（authorization-code）

app先申请授权码，user登录web-b使app获得授权码，app的server端再通过授权码获取令牌。

```bash
# 1. app向web-b申请授权码， user在web-b登录后，选择授权or不授权

b.com/oauth/authorize         # web-a req
 ?response_type=code          # 授权类型: 授权码
 &client_id={CLIENT_ID}       # app是谁
 &redirect_uri={callback-web-a}  # 授权成功（or失败）后的回调
 &scope=read                  # 授权范围 如只读

# 2. 选择授权，web-b 回调 web-a，并附带授权码

a.com/{callback-web-a}       # web-b req(callback)
 ?code={AUTHORIZATION_CODE}  # 附带授权码

# 3. web-a的后端用授权码向web-b请求令牌

https://b.com/oauth/token?   # web-a req
 client_id=CLIENT_ID&        # appId
 client_secret={CLIENT_SECRET}& # appSecret
 grant_type=authorization_code& # 授权类型
 code={AUTHORIZATION_CODE}&  # 授权码
 redirect_uri={callback-web-a}  # 令牌颁发后的回调

# 4. web-b发送令牌json数据给web-a

a.com/{callback-web-a}

{    
"access_token":"ACCESS_TOKEN",
"token_type":"bearer",
"expires_in":2592000,
"refresh_token":"REFRESH_TOKEN",
"scope":"read",
"uid":100101,
"info":{}
}

```

### 隐藏式（implicit）

app是纯前端应用，跟授权码方式项目，没有授权码这个中间步骤。

### 密码式（password）

将web-b的账号密码直接交给web-a进行授权。

### 凭证式（credentials）

命令行下请求令牌
