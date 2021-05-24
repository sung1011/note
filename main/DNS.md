# DNS & domain name

## 域名规则

- 不区分大小写
- 不能使用标点符号, 除了连字符(-)
- 字符数不超过255
- TCP/UDP:53

> 早期只支持英文, IDNs推出后, 支持了多语言. 通过`punycode转码`, 如 中国.cn -> xn-fiqs8s.cn

## 域名结构

```bash
# 主机名.次级域名.顶级域名.根域名
# 三级域名.二级域名.一级域名.根域名
host.SLD.TLD.root 如: app.tickles.cn. (root可省略为 . 亦可忽略)
```

- `根域名 root domain` 对于所有域名都一样, 一般会省略. 也可以简写为`.`.如: tickles.cn.
- `顶级域名 TLD top-level domain` 如: com, cn, net, co
- `次级域名 SLD second-level domain (二级域名)` 自定义域名.如: tickles.cn; 申请需要备案.
- `主机名 host (三级域名)` 可用户自定义, 定义为`*`即泛域名.如: *.tickles.cn, app.tickles.cn, www.tickles.cn
- `四/五/六级域名` 同理延续

> `子域名` 相对概念, 三级是二级的子域, 二级是一级的子域

> www.sina.com.cn 中cn是顶级 com是二级, sina是三级, www是四级 (而非com.cn是顶级)

## DNS 域名解析流程

1. 检查C端(浏览器)缓存中是否有解析结果
2. C端请求`域名解析S`解析域名app.tickles.cn对应的IP
3. 检查`域名解析S`缓存中是否有解析结果
4. `域名解析S`请求`root domain` (根域名dns S), 返回`TLD`(cn)S地址
5. `域名解析S`请求`TLD` (顶级域名dns S), 返回`SLD`(tickles.cn)S地址
6. `域名解析S`请求`SLD` (二级域名dns S), 返回`host`(app.tickles.cn)地址IP
7. `域名解析S`返回C端 域名app.tickles.cn对应的实际IP

> 域名解析服务: 代理请求 和 缓存结果; 地址配置在 /etc/resolve.conf

> [dig](linux-cmd-dig.md)观察域名解析过程 `dig +trance <domain name>`; root->TLD->SLD->host; 查到`A`记录,即找到IP; Received 为每个阶段分界;

## DNS资源记录 RR

- `Domain` 域名 如: tickles.cn.
- `TTL` 生命周期 如: 600
- `class` 网络协议类型 如: IN
- `type` 资源记录类型 如: A,AAAA,NS,TXT,CNAME,MX
- `rdata` 资源记录数据 如: 47.93.191.198

### DNS资源记录类型 RR-type

- `A` 域名对应的IP IPv4

      【domain】 IN A 【IP地址】
      ns1.exmaple.com. IN A 198.51.100.2


- `AAAA` 域名对应的IP IPv6

      【domain】 IN AAAA 【IP地址】
      ns1.exmaple.com. IN AAAA 8fe0::8f61:ac8:30cd:a16e

- `NS` Name Server 域名服务器; 上级域名服务器地址

      【domain】 IN NS 【DNS服务器】
      example.com. IN NS ns1.example.com.

- `TXT` Text; 域名说明; 验证控制权

      【domain】 IN TXT 【任意字符串】
      ns1.exmaple.com. IN TXT "联系电话：XXXX"

- `CNAME` Canonical Name 规范名,别名

      【別名】 IN CNAME 【原名】
      sub.example.com. IN CNAME hoge.example.com.

- `MX` Mail Exchanger 邮件交换

      【domain】 IN MX 【优先度】 【邮件服务器】
      example.com. IN MX 10 mail.example.com.

- `SPF` 送信测邮件服务器确认规则

      【domain】 IN TXT 【送信侧邮件服务器确认规则】
      exmaple.com. IN TXT "v=spf1 ip4:198.51.100.1 ~all"

- `SOA` Start of Authority 起始授权机制

- `PTR` Pointer 指针

- `SRV` Service

## DNS域名状态

- `ACTIVE`
- `REGISTRY-LOCK`
- `REGISTRY-HOLD`
- `REGISRAR-LOCK`
- `REGISRAR-HOLD`
- `REDEMPTIONPERIOD`
- `PENDINGRESTORE`
- `PENDINGDELETE`

## ref

- `https://www.cnblogs.com/alummox/p/11173449.html`
- `https://zhuanlan.zhihu.com/p/88260838`
- `https://www.sohu.com/a/308514296_404292`