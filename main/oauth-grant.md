# oauth 授权 (获取token)

## 流程

```seq
    第三方应用 app --- auth req ---> res owner      // app要求用户授权
    第三方应用 app <--- auth grant --- res owner    // 用户同意授权
app
    第三方应用 app --- auth grant ---> 微信验证 auth server   // app使用上一步的授权，向认证服申请令牌
    第三方应用 app <--- access token --- 微信验证 auth server // 认证服认证，并发放令牌
app
    第三方应用 app --- access token ---> 微信 res server // app使用令牌向资源服请求资源
    第三方应用 app <--- protected resource --- 微信 res server // 资源服验证令牌，并提供资源
```

## 类型

- `授权码 authorization-code`
- `隐藏式 implicit`
- `密码式 password`
- `客户端凭证 client credentials`

### 授权码（authorization-code）

app先申请授权码，user登录web-b使app获得授权码，app的server端再通过授权码获取令牌。

```bash

# 1. app跳转到web-b，user在web-b登录后，选择授权or不授权，授权则返回授权码
https://b.com/oauth/authorize         # web-a req
 ?response_type=code          # 授权方式: 授权码
 &client_id={CLIENT_ID}       # app是谁
 &redirect_uri={CALLBACK-WEB-A}  # 授权成功（or失败）后的回调
 &scope=read                  # 授权范围 如只读

# 2. 选择授权，web-b 回调 web-a，并附带授权码
https://a.com/{callback-web-a}       # 回调 web-a
 ?code={AUTHORIZATION_CODE}  # 附带授权码

# 3. web-a的后端用授权码向web-b请求令牌
https://b.com/oauth/token?   # web-a req
 client_id=CLIENT_ID&        # appId
 client_secret={CLIENT_SECRET}& # appSecret
 grant_type=authorization_code& # 授权方式
 code={AUTHORIZATION_CODE}&  # 授权码
 redirect_uri={CALLBACK-WEB-A}  # 令牌颁发后的回调

# 4. web-b发送令牌json数据给web-a

a.com/{CALLBACK-WEB-A}

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

> 步骤1 是重定向，必须get，敏感信息可能会被中间人攻击  
> 步骤3 由后端发起请求，保障了安全

### 隐藏式（implicit）

app是纯前端应用，跟授权码方式项目，没有授权码这个中间步骤。

```bash

# 1. app跳转到web-b，user在web-b登录后，选择授权or不授权，授权则直接返回返回token
https://b.com/oauth/authorize?
  response_type=token&          # 授权方式: token
  client_id={CLIENT_ID}&
  redirect_uri={CALLBACK-WEB-A}&
  scope=read

# 2. 选择授权，web-b 回调 web-a，并附带授权码
https://a.com/{CALLBACK-WEB-A}#code={AUTHORIZATION_CODE} # 回调web-a 附带授权码 (注意是锚点 而非query)
```

> 注意，令牌的位置是 URL 锚点（fragment），而不是查询字符串（querystring），这是因为 OAuth 2.0 允许跳转网址是 HTTP 协议，因此存在"中间人攻击"的风险，而浏览器跳转时，锚点不会发到服务器，就减少了泄漏令牌的风险。  
> 隐藏式非常不安全，用于安全需求不高的场景，并且有效期务必很短。

### 密码式（password）

将web-b的账号密码直接交给web-a进行授权。

```bash

# 1. 传递web-b的用户名，密码
https://oauth.b.com/token?
  grant_type=password&      # 授权方式: 密码
  username={USERNAME}&
  password={PASSWORD}&
  client_id={CLIENT_ID}
```

> 只适用于其他方式都无法采用的情况

### 凭证式（credentials）

命令行下请求令牌

```bash

# 1. web-a 的 server 进行请求
https://oauth.b.com/token?
  grant_type=client_credentials&    # 授权方式: 凭证
  client_id={CLIENT_ID}&            # appID
  client_secret={CLIENT_SECRET}     # appSecret
```
