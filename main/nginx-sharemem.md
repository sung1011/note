# nginx 共享内存

    效率高: 管道和消息队列等需要4次拷贝, 共享内存2次.

## 锁

    自旋锁

## slab内存管理

### 原理

    Bestfit: 共享内存每个页面4k, 将其切成多个slot.

> 32, 64, 128.... 51字节的数据会被放入64的slot.

> 相同大小slot组成链表.

### 优点

- 适合小对象  
- 避免碎片, 使用率高.  
- 重复使用分配了的slot避免初始化

### 监控管理

    ngx_slab_stat

## 使用共享内存的数据结构和模块

### rbtree

```bash
- ngx_stream_limit_conn_module # 流控
- ngx_http_limit_conn_module
- ngx_stream_limit_req_module
- http_cache
  - ngx_http_file_cache_module
  - ...proxy...
  - ...scgi...
  - ...uwsgi...
  - ...fastcgi...
- ssl
  - ngx_http_ssl_module
  - ...mail...
  - ...stream...
```

### 单链表

- ngx_http_upstream_zone_module
- ngx_stream_upstream_zone_module

### ngx_http_lua_api

- lua_shared_dict 分配共享内存(rbtree, 链表)
