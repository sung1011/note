# mongodb  

## [概述](mongodb-overview.md)
  
## [命令 cmd](mongodb-shell.md)

## [索引index](mongodb-index.md)

## aggregation 聚合  

阶段操作符  

- $count, $project, $match, $group, $sort, $limit, $unwind  
  
`db.mycol.aggregate([{group: {_id: 'sex', personCount: {$sum: 1}}}])`  
  
## sync 同步

initial sync 全量同步  
oplog 增量同步  

- Primary上的写操作完成后，会向特殊的local.oplog.rs特殊集合写入一条oplog，Secondary不断的从Primary取新的oplog并应用  
  
## [replica set 副本集](mongodb-relicset.md)

## [mongos shard 分片](mongodb-mongos.md)

## capped collection

Capped Collection是性能出色的有着固定大小的集合，以LRU（least Recently Used，最近最少使用）规则和插入顺序执行age-out（老化移出）处理，自动维护集合中对象的插入顺序。多用以日志归档。  

### 注意
  
- 如果写比读多，最好不要在上面创建索引
- 使用natual ordering可以有效地检索最近插入的元素，因为capped collection能够保证自然排序就是插入的顺序。  
- capped collection不能被shard.  
- 可以在创建capped collection时指定collection中能够存放的最大文档数。  

### 设计模式

- 分桶
- 子集
- 预聚合
- 列转行
- 文档版本
- 近似处理
  
## 实战  

### 输出全部结果

```js  
// dump.js  
var c = db.coll.find({status:1}).limit(5)  
while(c.hasNext()) {  
    printjson(c.next());  
}  
//mongo 127.0.0.1:27017/db1 dump.js> result.js  
```  
  
### 查询最后插入的数据

```js  
db.coll.find().skip(db.coll.count()-1).forEach(printjson)  
db.coll.find().limit(1).sort({$natural:-1})  
```  
  
### 查询文档的keys

```js  
for(var key in db.coll.findOne({_id:"xxx"})) {print (key)}  
```  
  
### 查询内嵌embedded文档的keys TODO  
  
### doc size

```js  
Object.bsonsize(db.coll.findOne({type:"auto"}))  
```

### 单个doc 16M限制

TODO

### update vs findAndModify

1. update 可更新多个doc，但只保证单个doc原子性
2. findAndModify 可以保证修改 与 返回结果（改前，改后都可以）这两个步骤是原子的
3. findAndModify 若 upsert: true 并 无查询结果时, 并发状态下可能插入多个doc
4. findAndModify 在分片集群中，查询必须包含分片key

## ref

[mongodb是如何实现ACID](https://blog.csdn.net/czq7511/article/details/77531903)
