# mongodb design

## 1对1 内嵌

## 1对n 数组 {address:['bj', 'tj', 'sh']}

## n对n 数组+关联 {address: [010, 022, 021]} {code: {110:bj, 022:tj, 021:sh}}

## 分桶

- 场景 时序数据，如每分钟股票数据
- 痛点 数据采集频繁，数据量小而多，索引占用空间大
- 方案 将一小时内的数据内嵌到同一个文档

## 列转行

- 场景 大量相似字段，如多语言属性(zh,us,fr,it...)，产品属性(color,size)
- 痛点 索引过多
- 方案 多个字段转化为一个数组字段的元素，建立一条联合索引

```js
{zh:2019, us:2011, fr:2016}

db.coll.createIndex({zh:1})
db.coll.createIndex({fr:1})
db.coll.createIndex({us:1})
//
{release:[ {contry:zh, date:2019}, {contry:us, date:2011}, {contry:fr, date:2016} ]}

db.coll.createIndex({"release.contry":1, "release.date":1})
```

## 子集

- 场景
- 痛点
- 方案

## 预聚合

- 场景 需要聚合的数据
- 痛点 统计计算耗时长
- 方案 添加冗余统计字段，更新数据时同时更新统计字段

```js
{num:5} -> {num:12}

{num:5, avg:5, daily:1, weekly:1} -> {num:12, avg:6, daily:2, weekly:2}
```

## 文档版本

- 场景 版本衍变
- 痛点 文档格式不同，难以管理
- 方案 增加版本号字段；快速过滤掉不需要升级的文档；不同版本做不同管理校验。

```js
{name: haha, company: lala}

{name: haha, company: [lala, xixi, dede], schema_version:"2.0"}
```

## 近似处理

- 场景 网页计数，结果不需要准确的排行
- 痛点 写入太频繁
- 方案 间隔写入，每隔10或100次写入一次
