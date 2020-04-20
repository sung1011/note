# mongodb cmd

## 全家桶

- mongod 应用软件
- mongo 命令行管理工具
- mongos 路由进程，分片环境使用
- mongodump / mongostrore 备份和恢复
- mongoexport / mongoimport CSV/JSON等导入导出，主要用于不同系统间数据迁移
- compass 官方GUI
- Ops Manager（企业） 集群管理软件
- BI Connector（企业） SQL解释器、BI套接件
- MongoDB Charts（企业） 可视化工具
- Atlas（付费/免费） 云服务

## stat

```js
db.serverStatus() // 主要信息
  connections 连接数
  locks 锁
  network 网络
  opcounters CRUD统计
  repl 复制集信息
  wiredTiger
    block-manager 数据库读写情况
    session
    concurrentTransactions Ticket令牌读写量使用情况
  mem 内存
  metrics 性能指标
db.isMaster() // 次要信息 节点情况
mongostats // 命令行工具 简略信息

db.oplog.rs.find().sort({$natrual:-1}).limit(1).next.ts - db.oplog.rs.find().sort({$natrual:1}).limit(1).next.ts // 可容纳多久的写操作

//查询专注度，搜索扫描了多少个文档
var status = db.serverStatus()
status.metrics.queryExecutor.scanned / status.metrics.document.returned // 扫描文档
status.metrics.queryExcutor.scannedObjects / status.metrics.document.returned // 返回文档

// 内存排序情况
var status = db.serverStatus()
status.metrics.operation.scanAndOrder / status.opcounters.query

db.stats() // 内存数据大小 实例数据总量（压缩前）
storageSize // 落盘后占用磁盘大小
```

## insert

```js
db.coll.insert({name: "foo"})
db.coll.insertOne({name: "foo"})
db.coll.insertMany([{name: "foo"}, {name: "bar"}, {name: "baz"}]) // 比循环insert效率高很多
```

## find

```js
db.coll.count({year: 1949}) // 效率低 会遍历
db.coll.find({year: 1949})
db.coll.find({year: 1949, title:"batman"})
db.coll.find({$and: [{year: 1949}, {title: "batman"}]})
db.coll.find({$and: [{year: 1949}, {title: "batman"}]})
db.coll.find({title: /^bat/i})
db.coll.find({title: {$regex: /^bat/i})
db.coll.find({"from.country": "China"})
db.coll.find({"category": "action"}, {_id: 0, title: 1}) // projection 投影(字段)
```

## update

```js
db.coll.updateOne()
db.coll.updateMany()
db.coll.update({_id:123}, {$inc: {"foo.bar.baz":3}})
db.coll.update({_id:123}, {$set: {"address.uid5":"bj"}}) // 修改内嵌文档
db.coll.update({_id:123}, {$set: {"address.1":"bj"}}) // 修改数组元素值
db.coll.update({_id:123}, {$set: {"books.1.name":"haha"}}) // 修改数组元素内嵌文档的值
```

## remove

```js
db.coll.remove({name: "foo"})
db.coll.remove({time: {$lt: 2019}})
db.coll.remove({})
```

## drop

```js
db.coll.drop()
db.dropDatabase()
```

## aggregation 聚合  

阶段操作符  

- $count, $project, $match, $group, $sort, $limit, $unwind  
  
`db.mycol.aggregate([{group: {_id: 'sex', personCount: {$sum: 1}}}])`  

## lock

```js
db.fsyncLock() // 锁住db写入 rs锁住不会报错，解锁后自动同步。
db.fsyncUnlock() // 解锁db写入
```
