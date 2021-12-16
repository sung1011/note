# 微服务 服务框架

连接 协议 序列化 压缩

## 外部建立连接

- [HTTP](HTTP.md)  
- [SOCKET](socket.md)

### IO模型

- BIO  
- NIO  
- AIO  

> [IO模型](linux-io.md)

## 内部通信协议

- [HTTP](HTTP.md)  
- Dubbo  

## 序列化压缩

- XML  
- JSON  
- protocol buffer  
- Thrift  

### 序列化压缩 - 评估标准

1. 支持数据结构丰富度
2. 跨语言
3. 性能
4. 可读性

## 方案

- Dubbo 阿里
- Motan 新浪
- Tars 腾讯
- Spring Cloud 国外pivotal
- gRPC 谷歌
- Thrift Facebook
