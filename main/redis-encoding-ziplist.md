# redis 编码 ziplist

    OBJ_ENCODING_ZIPLIST  

## 作用于

    OBJ_ZSET  
    OBJ_LIST  
    OBJ_HASH  

## 特性

    entries的数据存储在连续内存, 数据可压缩
    时间换空间

## 数据结构

```js
{
    zlbytes: 123,       # 记录整个ziplist的大小
    zltail: 3,          # ziplist开始指针与最后一个entry之间的偏移量，通过该偏移量可以获得最后一个entry
    zllen: 4,           #  entry数量
    entries: []entry{
        {
            previous_entry_length: 54,      # 上一个entry的大小
            encoding: 3,                    # 记录content的类型以及长度
            content: 'xljsd',                    # 一个整形或者字节数组
        }
    }  # 存储具体数据的节点, 紧密的数据可压缩
    zlend: '\0',          # 结尾标识符
}

# `previous_entry_length` 当<263byte时只占1byte; 当>=263字节时占用5byte; 导致当对队头进行增删时, 最差可能链式的导致后面的previous_entry_length依次重新分配空间.
```

## API
