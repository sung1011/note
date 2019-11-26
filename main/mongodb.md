# mongodb  
  
## 日常命令 cmd  
  
## [index 索引](mongodb-index.md)

## aggregation 聚合  

阶段操作符  

- $count, $project, $match, $group, $sort, $limit, $unwind  
  
`db.mycol.aggregate([{group: {_id: 'sex', personCount: {$sum: 1}}}])`  
  
## sync 同步

initial sync 全量同步  
oplog 增量同步  

- Primary上的写操作完成后，会向特殊的local.oplog.rs特殊集合写入一条oplog，Secondary不断的从Primary取新的oplog并应用  
  
## [replica set 副本集](mongodb-relicset.md)

## shard 分片  

### 组成  

- **config**  

用来保存数据，保证数据的高可用性和一致性。可以是一个单独的mongod实例，也可以是一个副本集。在生产环境下Shard一般是一个Replica Set，以防止该数据片的单点故障。可以将所有shard的副本集放在一个服务器多个mongodb实例中。  
  
- **router**  

路由就是mongos的实例，客户端直接连接mongos，由mongos把读写请求路由到指定的Shard上去。  
一个Sharding集群，可以有一个mongos，也可以为每个App Server配置一个mongos以减轻路由压力。  
注意这里的mongos并不要配置为rs，因为只是个路由，并不存储数据，配置多个mongos的意思是配置多个单独的mongos实例。  

- **shards**  

用来保存数据，保证数据的高可用性和一致性。可以是一个单独的mongod实例，也可以是一个副本集。在生产环境下Shard一般是一个Replica Set，以防止该数据片的单点故障。可以将所有shard的副本集放在一个服务器多个mongodb实例中。  
  
### 分片键 Shard keys  

### 优点  

读写方面： sharding将读写负载均匀到各个shard，且workload上限可以通过水平扩展来增加。  
扩容方面： 每个shard保存一部分数据，可以通过增加shards来扩容。  
高可用方面： 即便某个shard不可用了，整个集群也可以对外提供服务，只不过访问down掉的shard会报"Connection refused"的错误。而且MongoDB3.2以后可以为每个shard都配置副本集（replica set），这样保证最大程度的高可用性。  

### 缺点  

数据量较少时不建议使用sharding，毕竟读写都要经过一层路由会有性能损耗，直接表现就是ips和qps会降低。  

### 额外  

sharding集群不支持一些常规的单实例方法，如group()，可以使用mapReduce()或者aggregate()中的group来替代，因此建议从一开始学习就直接使用aggregate(),这种写法较为简单明了，且统一化易于识别。  
对于没有用到shard key的查询，路由进行全集群广播（broadcast operation），对每个shard都查一遍进行scatter/gather，此时效率会很低。  

### 策略  

hash  
ranged  
  
## capped collection

Capped Collection是性能出色的有着固定大小的集合，以LRU（least Recently Used，最近最少使用）规则和插入顺序执行age-out（老化移出）处理，自动维护集合中对象的插入顺序。多用以日志归档。  

### 注意
  
- 如果写比读多，最好不要在上面创建索引
- 使用natual ordering可以有效地检索最近插入的元素，因为capped collection能够保证自然排序就是插入的顺序。  
- capped collection不能被shard.  
- 可以在创建capped collection时指定collection中能够存放的最大文档数。  
  
### 实例  

oplog
  
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

## ref

[mongodb是如何实现ACID](https://blog.csdn.net/czq7511/article/details/77531903)
