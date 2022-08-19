# protocal  
  
## 网络七层模型 (OSI Open System Interconnection)  

        7 应用层  
        6 表示层  
        5 会话层  
        4 传输层  
        3 网络层  
        2 数据链路层  
        1 物理层  

## TCP/IP 四层模型与协议

| 模型               | 对应OSI七层               | 协议                                                                     |
| ------------------ | ---------------------- | ------------------------------------------------------------------------ |
| 应用层 application | 应用层, 表示层, 会话层 | HTTP协议(TCP)、TELNET协议、DNS协议(UDP)、FTP协议(UDP)、SMTP协议、POP协议 |
| 传输层 transport   | 传输层                 | TCP协议(传输控制协议)、UDP协议(用户数据报协议)                           |
| 网络层 internet    | 网络层                 | IP协议(网际协议)、ICMP协议(互联网控制报文协议)                           |
| 数据链路层 link    | 数据链路层, 物理层     | Ethernet ARP协议(地址解析协议)、Mac                                      |

## [协议列表](protocal-list.md)

## 协议详解
  
### 7 应用层  

- [HTTP](HTTP.md)  

      line, header, body
      持久连接
      pipeline

- [HTTPS](HTTPS.md)  

      http + tls

- [websocket](websocket.md)
- [SSH Secure Shell](SSH.md)
- [cookie](cookie.md)

      S端通过Cookie在C端记录会话状态

- [session](session.md)
- telnet

### 6 表示层

- [SSL/TLS](SSL-TLS.md)  

### 5 会话层

- RPC

### 4 传输层

- [TCP](TCP.md)

      三次握手, 四次挥手
      超时重传
      拥塞控制
      粘包拆包

- UDP  
- [socket套接字抽象层](socket.md)  

### 3 网络层

- [IP](IP.md)  
- [ICMP](ICMP.md)

      Traceroute by ping

- 路由器

      外网IP寻址, 在多个路由器间计算最优路由路径, 将数据送达指定主机.

### 2 数据链路层

- [ARP 地址解析协议 Address Resolution Protocol](ARP.md)

      IP-MAC映射, 查不到就广播

- RARP  
- MAC 媒体访问控制 Medium Access Control  
- 链路层地址

      MAC地址 48位主机的物理地址, 局域网内唯一.  

- 交换机

      内网MAC寻址, 从第一次经由交换机的数据包中截取MAC地址并记录MAC-LAN口映射.  

### 1 物理层

- 网线  
- 集线器hub  

## ref

- <https://mp.weixin.qq.com/s/jiPMUk6zUdOY6eKxAjNDbQ>