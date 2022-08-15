# MongoDB capped collection

    固定大小的集合, 支持高吞吐的插入操作和根据插入顺序的查询操作

## 注意
  
- 如果写比读多, 最好不要在上面创建索引
- 使用natual ordering可以有效地检索最近插入的元素, 因为capped collection能够保证自然排序就是插入的顺序.  
- capped collection不能被shard.  
- 可以在创建capped collection时指定collection中能够存放的最大文档数.  
