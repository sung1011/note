# nginx 哈希表

## 场景

    静态不变的内容.即通常不会插入, 删除

## 源码

    src/core/ngx_hash

## 数据结构

```c
typedef struct {
  void value;   // 指向用户自定义元素数据的指针, 如果当前ngx_hash_elt_t槽为空, 则value的值为0
  u_short len;   // 元素关键字的长度
  u_char name[1];   // 元素关键字的首地址
} ngx_hash_elt_t;


typedef struct {
  ngx_hash_elt_t *buckets;   // 指向散列表的首地址, 也是第1个槽的地址
  ngx_uint_t size;   // 散列表中槽的总数
} ngx_hash_t;
```

## hash_max_size

    最大hash bucket个数

## hash_bucket_size

    与cpu cache向上对齐(64b)

## 实例

```bash
variables(steam/http)  
```

```bash
map(stream/http)  
```

```bash
http proxy # 反向代理  

    - ngx_http_uwsgi_module
    - ngx_http_scgi_module
    - ngx_http_fastcgi_module

http module  

    - ngx_http_referer_module
    - ngx_http_ssi_module
    - ngx_http_srcache_filter_module

http core  

    - server name
    - MIME types
```
