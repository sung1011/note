# ssh

## 验证

1. 口令
2. 秘钥

## 文件

   id_rsa：保存私钥
   id_rsa.pub：保存公钥
   authorized_keys：保存已授权的客户端公钥
   known_hosts：保存已认证的远程主机ID

## 连接

```bash
ssh user@host

　　The authenticity of host 'host (12.18.429.21)' can't be established.
　　RSA key fingerprint is 98:2e:d7:e0:de:9f:ac:67:28:c2:42:2d:37:16:58:4d.
　　Are you sure you want to continue connecting (yes/no)?
```

1. 首次ssh时问询是否确定连接公钥指纹为 `98:2e:d7:e0:de:9f:ac:67:28:c2:42:2d:37:16:58:4d` 的机器，这是为了防止中间人攻击，公钥指纹需自行对比。
2. 若同意（yes），则将远程主机的公钥保存在 `.ssh/known_hosts` 中，今后的ssh连接从而可以跳过该问询。

> RSA算法的公钥长达1024位，很难对比，所以MD5将其转化为128位的公钥指纹。

## cmd

```bash
ssh -vt user@host # 调试

ssh -p 1234 user@host # 指定端口

ssh user@host 'mkdir -p .ssh && cat >> .ssh/authorized_keys' < ~/.ssh/id_rsa.pub # 保存公钥
```
