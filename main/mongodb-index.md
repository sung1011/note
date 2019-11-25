# mongodb 索引

## 原理  

btree和hash  
  
## 场景  

读多写少  
  
## 类型  

- **Single Field Index 单字段索引**
- **Compound Index 复合索引** 多字段联合索引，注意顺序（a,b,c创建复合索引，查a,b时索引有效，查b，c时索引无效）  
- Multikey Index 多key索引 数组值索引  
- hash 哈希索引 是指按照某个字段的hash值来建立索引，目前主要用于 MongoDB Sharded Cluster 的Hash分片，hash索引只能满足字段完全匹配的查询，不能满足范围查询等。  
- Text Index 文本索引 能解决快速文本查找的需求，比如有一个博客文章集合，需要根据博客的内容来快速查找，则可以针对博客内容建立文本索引。  
- Geospatial Index 地理索引 能很好的解决O2O的应用场景，比如『查找附近的美食』、『查找某个区域内的车站』等。  

额外索引属性  

- unique index 唯一索引  
- TTL索引： 指定数据失效时间  
- partial index 部分索引： 只针对符合某个特定条件的文档建立索引  
- sparse index 稀疏索引： 只针对存在索引字段的文档建立索引，可看做是部分索引的一种特殊情况  

## 优化  

开启profiling慢查询，找到需要优化的请求  

- db.setProfilingLevel(1, 100); // 0 不开启 1 开启慢查询 2 记录所有  
  
## explain分析索引  
