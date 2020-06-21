# mongodb 索引

## 原理  

- [b-tree](ds-b-tree.md)
- [hash](ds-hash.md)

## 场景  

读多写少  

## 类型  

1. **Single Field Index 单字段索引**

2. **Compound Index 复合索引** 多字段联合索引，注意顺序（a,b,c创建复合索引，查a,b时索引有效，查b，c时索引无效）  

   - 顺序: ESR原则
     - E equal 精确放前面
     - S sort 排序放中间
       - R range 范围放最后

   > 如 `db.coll.find({gender: F, age:{$gte: 18}}).sort(joinTime: 1)` 建立 `{gender:1, joinTime:1, age:1}` 的联合索引最佳。

3. Multikey Index 多key索引 数组值索引  

4. hash 哈希索引 是指按照某个字段的hash值来建立索引，目前主要用于 MongoDB Sharded Cluster 的Hash分片，hash索引只能满足字段完全匹配的查询，不能满足范围查询等。  

5. Text Index 全文索引 关键词tag查找。  

   - 创建 `db.coll.ensureIndex({article:'text'})` // 限制 每个集合只能创建一个全文索引
   - 查找 `db.coll.find({$text: {$search: 'aa'}})`

6. Geospatial Index 地理索引 能很好的解决O2O的应用场景，比如『查找附近的美食』、『查找某个区域内的车站』等。  

7. 额外索引属性  

   - unique index 唯一索引  
   - TTL索引： 指定数据失效时间  
   - partial index 部分索引： 只针对符合某个特定条件的文档建立索引  
   - sparse index 稀疏索引： 只针对存在索引字段的文档建立索引，可看做是部分索引的一种特殊情况  

## 执行计划 -- 选择合适的索引

```bash
     评选查询计划------------------------------------------/
    /                                                   /
匹配查询计划---生成查询计划---评选候选计划---创建查询计划---执行查询计划---返回结果
```

> 评选候选计划: 多个候选计划后台多个线程同时尝试，谁执行快就选谁。

## 优化  

开启profiling慢查询，找到需要优化的请求  

- db.setProfilingLevel(1, 100); // 0 不开启 1 开启慢查询 2 记录所有  
  
## explain分析索引

```js
db.coll.find({name:"xx"}).explain(true)
```

## 关键词

- index
- key
- datapage
- covered query 覆盖查询。需要查询的字段都在索引中。
- ixscan 索引扫描
- collscan 集合扫描
- big O notation 时间复杂度
- selectivity 过滤性。
  