# nginx 红黑树

## 场景

    高度不超过2倍log(n)  
    增删改查算法复杂度O(log(n))  
    遍历O(n)

## 源码

src/core/ngx_rbtree

## 实例

ngx_conf_module

ngx_event_timer_rbtree

ngx_http_file_cache

ngx_http_geo_module

ngx_http_limit_conn_module

ngx_http_limit_req_module

ngx_http_lua_shdict:ngx.shared.DICT LRU链表性质

resolver ngx_resolver_t

ngx_stream_geo_module

ngx_stream_limit_conn_module
