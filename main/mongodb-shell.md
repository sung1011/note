# mongodb shell

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
