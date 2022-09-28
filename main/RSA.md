# RSA

    RSA 数学家Rivest、Shamir 和 Adleman提出的非对称加密算法

## 加密解密

    使用公钥加密的数据,利用私钥进行解密  
    使用私钥加密的数据,利用公钥进行解密  

## 应用

### HTTPS

1. C 请求 S, 返回 S.pubKey

2. C 用 S.pubKey 加密 (C.pubKey + C对称加密key) 发送给 S

    - 双向验证: 数据hash后生成摘要, 用C.priKey对摘要加密 生成数字签名, 发给S以防篡改
      - 摘要 = hash( data )
      - 签名 = C.priKey加密( 摘要 )

3. 双方后续通信基于 C对称加密key 作为key进行对称加密

## 数学原理

    互质、欧拉函数、欧拉定理、模反元素
