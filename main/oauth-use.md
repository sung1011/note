# oauth 操作

## 使用token

```bash
# 将token加在header中
curl -H "Authorization: Bearer ACCESS_TOKEN" \
"https://api.b.com"
```

## 更新token

res-server 可以在授权时一次性颁发两个令牌,  REFRESH_TOKEN 用于更新令牌.

```bash
https://b.com/oauth/token?
  grant_type=refresh_token&         # 授权方式: 更新
  client_id={ CLIENT_ID }&
  client_secret={ CLIENT_SECRET }&
  refresh_token={ REFRESH_TOKEN }

```

> 也可以重走一遍授权流程, 但体验不佳
