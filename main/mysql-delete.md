# mysql delete

## 删除大数据

- 直接删

```sql
delete from orders where time < SUBDATE(CURDATE(),INTERVAL 3 month);
```

- 分批删

```sql
delete from orders where time < SUBDATE(CURDATE(),INTERVAL 3 month) order by id limit 1000;
```

- 转化为按主键删

```sql
select max(id) from orders where time < SUBDATE(CURDATE(),INTERVAL 3 month);

delete from orders where id < ? order by id limit 1000;
```

- 重建新表

```sql
# 新建同构表
create table orders_tmp like orders;
# 将需要保留的数据重建到新表中
insert into orders_tmp (select * from order where time >= SUBDATE(CURDATE(),INTERVAL 3 month));
# 保守的抛弃旧表
rename table orders to orders_droppd;
rename table orders_tmp to orders;

drop table orders_droppd;
```

> `order by`: 由于B+tree的主键有序性, id相近的记录在磁盘的物理文件大致也在一起.

> `标记删除`: 数据库数据存储在二进制文件,类似数组,难以删除中间部分.所以删除的行为只是标记(即不减硬盘空间),可用`OPTIMIZE TABLE`真正删除,但会锁表很慢.
