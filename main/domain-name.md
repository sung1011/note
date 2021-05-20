# domain name

## 规则

- 不区分大小写
- 不能使用标点符号, 除了连字符(-)
- 字符数不超过255

## 层级

```bash
主机名.次级域名.顶级域名.根域名
即: host.sld.tld.root
如: app.tickles.cn. (root可省略, . 可以保留)
```

- `根域名 root` 对于所有域名都一样, 一般会省略. 也可以简写为`.`.如: tickles.cn.
- `顶级域名 TLD top-level domain` 如: com, cn, net, co
- `次级域名 SLD second-level domain` 自定义域名.如: tickles.cn
- `主机名 host` 可用户自定义, 定义为`*`即泛域名.如: *.tickles.cn, app.tickles.cn

## 查询

