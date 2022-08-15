# database

## 重复创建?

      issue: 创建数据时, 前端重复请求, 前端重试请求

- mysql

    1. 从`GenGUID`的服务中生成全局唯一ID
    2. 创建时GUID作为参数, 指定其GUID作为表的主键值(具有唯一性)

> 幂等

> 重复创建的请求建议忽略, 不抛出异常, 以第一次创建的返回值为准.

- redis

    1. 从`GenGUID`的服务中生成全局唯一ID
    2. 将GUID存入redis (KV)
    3. 创建时GUID作为参数, 判断其是否存在于redis
    4. 若存在, 则redis删除该key, 并继续逻辑
    5. 若不存在, 则视为重复创建

> 幂等

> 大规模项目需要关注redis单点问题

---

## 请求无序?

      请求顺序不固定

- redis

    1. 请求头中带有seqID参数
    2. S的redis中保存seqID
    3. S只接收seqID(C) == seqID(S)的请求
    4. S处理请求后,将S的seqID+=1
    5. seqID(C) > seqID(S)的请求被暂存起来,等待后续处理
    6. seqID(C) < seqID(S)的请求直接抛弃

---

## ABA修改?

        issue: ABA即先将某字段设置为a, 然后设置为b, 由于网络原因a没有正常返回导致客户端重试了a的请求, 导致最终字段值为a (期望b)

- 版本号

    1. 给数据表新增ver(版本号)字段
    2. 在update时将版本号作为条件
    3. 并且update时将ver+=1

> 幂等

> 对应SQL: `update <table> set <field>=<value>, ver=ver+1 where ver = 6`  

---

## schema不固定?

        issue: 不同的商品有不同的属性, 比如电脑: 内存 硬盘 CPU; 汽车: 车型, 发动机; 衣服: 颜色, 款式... 但需要抽象出一个商品属性表

- schema

    1. [MongoDB](mongodb.md)
    2. [Mysql-json](mysql.md)
    3. elasticSearch dynamic field mappings

---

## 浏览器存储?

- 存储

  - [Cookie](cookie.md)
  - LocalStorage
  - SessionStorage
  - IndexedDB
