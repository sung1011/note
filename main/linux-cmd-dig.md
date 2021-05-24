# dig

```go
$ dig redhat.com

// 头信息 版本和输入参数
; <<>> DiG 9.7.3-RedHat-9.7.3-2.el6 <<>> redhat.com
;; global options: +cmd
// 概述
    // status: NOERROR 无报错 / NXDOMAIN 没查到
    // ANSWER: 4 查询结果个数
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 62863
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 4, ADDITIONAL: 3
// 查询段 查询条件
    // 即查redhat.com.的A记录 (默认A, 经常需要查ns)
;; QUESTION SECTION:
;redhat.com.                    IN      A
// 结果段 查询结果
    // 即redhat.com.的A记录为209.132.183.81
;; ANSWER SECTION:
redhat.com.             267     IN     CNAME    redhat.amazoneaws.com.
redhat.amazoneaws.com.     37      IN      A       209.132.183.81
// 来源段 由哪些dns解析出来的
    // 即redhat.com.的解析是由一组二级dns解析出来的
;; AUTHORITY SECTION:
redhat.com.             73      IN      NS      ns4.redhat.com.
redhat.com.             73      IN      NS      ns3.redhat.com.
redhat.com.             73      IN      NS      ns2.redhat.com.
redhat.com.             73      IN      NS      ns1.redhat.com.
// 附加段 来源段S的IP
;; ADDITIONAL SECTION:
ns1.redhat.com.         73      IN      A       209.132.186.218
ns2.redhat.com.         73      IN      A       209.132.183.2
ns3.redhat.com.         73      IN      A       209.132.176.100
// 统计段
;; Query time: 13 msec
;; SERVER: 209.144.50.138#53(209.144.50.138)
;; WHEN: Thu Jan 12 10:09:49 2012
;; MSG SIZE  rcvd: 164
```

## 参数

- `@` 指定DNS服务器; 如 dig @ns1.redhat.com redhat.com
- `+trace` 跟踪完整的解析过程
- `+short` 简短输出
- `+noall` 关闭所有段落; 如 dig -noall +answer 只显示answer
- `+nocomments` 关闭注释行
- `-t` 指定类型, 一般可省略; 如 dig ns, dig AAAA, dig ANY
- `-x` 通过IP反向查询域名; 如 dig -x 209.132.186.218; 需要S额外配置
- `-f` 批量域名查询, 指定一个写了很多域名(换行隔开)的文件

> 可配置该命令默认携带的参数 `$HOME/.digrc`
