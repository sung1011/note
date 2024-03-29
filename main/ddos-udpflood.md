# ddos-udp-flood UDP反射放大攻击

    攻击者向肉鸡发送伪造的数据包, 其源IP伪造为了被攻击服务器的IP
    肉鸡会将UDP返回到该IP, 并且返回内容会比请求内容增大非常多, 形成攻击.

## 影响

- CPU, 带宽占满, 无法正常提供服务
- 隐蔽性强

## 防御

- 增加集群节点
- 扩容带宽
- 地理位置过滤
- 服务白/黑名单
- 源端口波动限制 对反射端口进行统计, 当其请求数激增则临时封禁
- 偏移字节数学习 学习攻击的包文中相同特征, 进行筛选丢弃
- 包文长度学习 统计正常包文长度, 对过大过小包文进行筛选丢弃
- 限速
- 云服务高防

## 放大

| protoc       | port  | multiple    |
| ------------ | ----- | ----------- |
| DNS          | 53    | 28~54       |
| NTP          | 123   | 557         |
| SNMP         | 161   | 6           |
| SSDP         | 1900  | 31          |
| PORTMAP      | 111   | 7~28        |
| QOTD         | 17    | 140         |
| CHARGEN      | 19    | 359         |
| TFTP         | 69    | 60          |
| NETBIOS      | 138   | 4           |
| MEMCACHED    | 11211 | 10000~50000 |
| WS_DISCOVERY | 3702  | 70~500      |
| CLDAP        | 389   | 56~70       |

> 黑客通过快速扫描全球开放memcached端口的服务器, 将其作为肉鸡.

## ref

- aliyun高防设置说明 <https://help.aliyun.com/document_detail/215582.html?spm=5176.22414175.sslink.1.6a154e38ogZD95>
- 常见UDP反射放大攻击 <https://zhuanlan.zhihu.com/p/83793355>