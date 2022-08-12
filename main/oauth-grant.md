# oauth 授权 (获取token)

## members

- `Third-party application` 第三方应用程序(app), 又称"客户端"(client)但也可能是app的server. --- 第三方app
- `User Agent` 用户代理. --- 浏览器
- `Authorization server` 认证服务器(res-server-auth), 即服务提供商专门用来处理认证的服务器. --- 百度账号认证服务
- `Resource server` 资源服务器(res-server), 即服务提供商存放用户生成的资源的服务器.它与认证服务器, 可以是同一台服务器, 也可以是不同的服务器. --- 百度云盘/地图/翻译服务
- `Resource Owner` 资源所有者, 又称"用户"(user). --- 自己
- `HTTP service` HTTP服务提供商, 简称"服务提供商". --- 百度

## 流程

```bash
                        auth-req
    app         ---------------------> res owner        # app要求用户授权
                        auth-code
    app         <--------------------- res owner        # 用户登录并同意授权

                        auth-code 
    app         ---------------------> auth server      # app使用上一步的授权码(区分权限), 向认证服申请令牌
                        access-token
    app-server  <--------------------- auth server      # 认证, 并发放令牌

                       access-token
    app-server  ---------------------> res server       # app-server使用令牌向资源服请求资源
                       protected-resource
    app-server  <--------------------- res server       # 验证令牌, 并提供资源
```

> `auth-code` 决定了是否安全, 尽量不给用户(防止被劫持); 细分了资源权限

> `auth-server, res-server` 是同一企业下的不同服务; auth-server可以做SSO, 多个app同一个auth-server

## 类型

- `授权码 authorization-code`
- `隐藏式 implicit`
- `密码式 password`
- `客户端凭证 client credentials`

### 授权码(authorization-code)

```bash
# 适用于: app
# user登录res-server获得auth-code 并回调给app-server, app-server再通过auth-code申请token.


# 1. app跳转到res-server, 登录后授权
app -> res-server

  https://b.com/oauth/authorize           # App req
   ?response_type=code                    # 授权方式: 授权码
   &client_id={CLIENT_ID}                 # app是谁
   &redirect_uri={CALLBACK-App}           # 授权成功(or失败)后的回调
   &scope=read                            # 授权范围 如只读

# 2. 授权后 res-server 携带授权码回调 App-server
res-server -> app-server

  https://a.com/{callback-App-server}     # 回调 App-server
   ?code={AUTHORIZATION_CODE}             # 附带授权码

# 3. App-server用授权码向res-server请求令牌
app-server -> res-server

  https://b.com/oauth/token?              # App-server req
   client_id=CLIENT_ID&                   # appId
   client_secret={CLIENT_SECRET}&         # appSecret
   grant_type=authorization_code&         # 授权方式
   code={AUTHORIZATION_CODE}&             # 授权码
   redirect_uri={CALLBACK-App}            # 令牌颁发后的回调

# 4. res-server发送令牌json数据给App
res-server -> app

  a.com/{CALLBACK-App}

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

> 步骤1 是重定向, 必须get, 敏感信息可能会被中间人攻击  

> 步骤3 由后端发起请求, 保障了安全

### 隐藏式(implicit)

```bash
# 适用于: app是纯前端应用, 无后端
# 直接向前端返回token, 没有授权码这个中间步骤. (易劫持不安全)


# 1. app跳转到 res-server 登录后, 授权并直接返回返回token
app -> res-server

  https://b.com/oauth/authorize?
    response_type=token&          # 授权方式: token
    client_id={CLIENT_ID}&
    redirect_uri={CALLBACK-App}&
    scope=read

# 2. 选择授权, res-server 回调 App, 并附带token
res-server -> app

  https://a.com/{CALLBACK-App}#code={TOKEN} # 回调App 附带授权码 (注意是锚点 而非query)
```

> `锚点` 注意, 令牌的位置是 URL 锚点(fragment), 而不是查询字符串(querystring). 这是因为 OAuth 2.0 允许跳转网址是 HTTP 协议, 因此存在"中间人攻击"的风险, 而浏览器跳转时, 锚点不会发到服务器, 就减少了泄漏令牌的风险.  

> 隐藏式非常不安全, 用于安全需求不高的场景, 并且有效期务必很短.

### 密码式(password)


```bash
# 适用于: app其他方式都无法采用的情况
# 直接发送res-server的账号密码以获取信任


# 1. 传递res-server的用户名, 密码
app -> res-server

  https://oauth.b.com/token?
    grant_type=password&      # 授权方式: 密码
    username={USERNAME}&
    password={PASSWORD}&
    client_id={CLIENT_ID}
```


### 凭证式(clint credentials)

```bash
# 适用于: 服务器终端
# 通过证书请求令牌


# 1. App-server 进行请求
https://oauth.b.com/token?
  grant_type=client_credentials&    # 授权方式: 凭证
  client_id={CLIENT_ID}&            # appID
  client_secret={CLIENT_SECRET}     # appSecret
```

## ref

- 阮一峰oauth2<https://www.ruanyifeng.com/blog/2019/04/oauth-grant-types.html>