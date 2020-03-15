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
db.coll.insertMany([{name: "foo"}, {name: "bar"}, {name: "baz"}])
```

## find

```js
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
db.coll.update()
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
