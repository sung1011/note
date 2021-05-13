# mysql sql

## setting

```sql
# binlog
show variables like '%log_bin%';
show master status;

# 复制
SELECT PLUGIN_NAME, PLUGIN_STATUS FROM INFORMATION_SCHEMA.PLUGINS WHERE PLUGIN_NAME LIKE '%semi%';
```
