# mongodb cmd

## 全家桶

- mongod 应用软件
- mongo 命令行管理工具
- mongos 路由进程, 分片环境使用
- mongodump / mongostrore 备份和恢复
- mongoexport / mongoimport CSV/JSON等导入导出, 主要用于不同系统间数据迁移
- compass 官方GUI
- Ops Manager(企业) 集群管理软件
- BI Connector(企业) SQL解释器、BI套接件
- MongoDB Charts(企业) 可视化工具
- Atlas(付费/免费) 云服务

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

//查询专注度, 搜索扫描了多少个文档
var status = db.serverStatus()
status.metrics.queryExecutor.scanned / status.metrics.document.returned // 扫描文档
status.metrics.queryExcutor.scannedObjects / status.metrics.document.returned // 返回文档

// 内存排序情况
var status = db.serverStatus()
status.metrics.operation.scanAndOrder / status.opcounters.query

db.stats() // 内存数据大小 实例数据总量(压缩前)
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
db.coll.find({title: /^bat/i})
db.coll.find({title: {$regex: /^bat/i})
db.coll.find({"from.country": "China"})
db.coll.find({"category": "action"}, {_id: 0, title: 1}) // projection 投影(字段)
```

## update

```js
db.coll.update({}, {$inc: {"foo.bar.baz":3}}, false, true) // 修改多个;另外变种方法 updateOne, updateMany
db.coll.update({_id:123}, {$set: {"address.uid5":"bj"}}) // 修改内嵌文档
db.coll.update({_id:123}, {$set: {"address.1":"bj"}}) // 修改数组元素值
db.coll.update({_id:123}, {$set: {"books.1.name":"haha"}}) // 修改数组元素内嵌文档的值
db.coll.update({_id:123}, {$unset: {"foo.bar.baz":1, "foo.bar.quz":1}}) // 删除多个内嵌
```

## remove

```js
db.coll.remove({name: "foo"})
db.coll.remove({time: {$lt: 2019}})
db.coll.remove({})
```

## index

```js
db.coll.createIndex({name: 1}) // 1 升序 2 降序; ensureIndex是别名
db.coll.createIndex({name: 1}, {unique: true}) // 唯一索引
db.coll.createIndex({gender: 1, joinTime: 1, age: 1}) // 复合索引 创建遵循ESR原则
db.coll.createIndex({age: 1}, {partialFilterExpression: {age: {$get: 5}}}) // 索引部分创建 age大于5的才创建索引
db.coll.createIndex({name: 1}, {background: 1}) // 后台创建 推荐创建索引时必须加此选项
db.coll.getIndexes() // 查看集合索引; getIndexKeys()简化
db.coll.dropIndex({joinTime: 1}) // 删除索引
db.coll.find({age: {$gt: 20}}).explain(true) // 索引执行细节
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
db.fsyncLock() // 锁住db写入 rs锁住不会报错, 解锁后自动同步.
db.fsyncUnlock() // 解锁db写入
```
  
## 实战  

### 执行脚本

```bash
mongo --quiet < a.js | grep abc
```

### 输出全部结果

```js  
// dump.js  
var c = db.coll.find({ status:1 }, {}).limit(100)
while(c.hasNext()) {  
    printjson(c.next());  
}  
//mongo 127.0.0.1:27017/db1 dump.js> result.js  

var c = db.coll.find({ status:1 }, {}).limit(100); while(c.hasNext()) {  printjson(c.next()); }  
```

> 实质是拿到全部数据, 然后遍历
  
### 查询最后插入的数据

```js  
db.coll.find().skip(db.coll.count()-1).forEach(printjson)  
db.coll.find().limit(1).sort({$natural:-1})  
```  
  
### 查询文档的field

```js  
for(var key in db.coll.findOne({_id:"xxx"})) {print (key)}  
```  
  
### 查询内嵌embedded文档的keys

TODO
  
### doc size

```js  
Object.bsonsize(db.coll.findOne({type:"auto"}))  
```

### 单个doc 16M限制

TODO

### update vs findAndModify

1. update 可更新多个doc, 但只保证单个doc原子性
2. findAndModify 可以保证修改 与 返回结果(改前, 改后都可以)这两个步骤是原子的
3. findAndModify 若 upsert: true 并 无查询结果时, 并发状态下可能插入多个doc
4. findAndModify 在分片集群中, 查询必须包含分片key

### 导表 导列 导数据

```bash
mongoexport -d dbname -c collectionname -o file --type json/csv -f field

# -d : 数据库名
# -c : collection名
# -o : 输出的文件名
# --type :  输出的格式, 默认为json
# -f : 输出的字段, 如果-type为csv, 则需要加上-f "字段名"

mongoimport -d dbname -c collectionname --file filename --headerline --type json/csv -f field

# -d : 数据库名
# -c : collection名
# --type : 导入的格式默认json
# -f : 导入的字段名
# --headerline : 如果导入的格式是csv, 则可以使用第一行的标题作为导入的字段
# --file : 要导入的文件
```

> 文件为可读的 json 或 csv

## 备份 还原

```bash
mongodump -h IP --port 端口 -u 用户名 -p 密码 -d 数据库 -o 文件存在路径

mongorestore -h IP --port 端口 -u 用户名 -p 密码 -d 数据库 --drop 备份文件夹下的db路径

# --drop 替换
# 备份文件夹下的db路径
```

> 文件为不可读的bson

### 解析id

```bash
# 一个
ObjectId("5e0c4e087deadd1a6f8b4b57").getTimestamp()

# 遍历
db.{coll}.find().forEach(function(row){print("time: " + row._id.getTimestamp())})
```

> 截取_id的前8位, hex->dec.
