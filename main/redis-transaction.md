# redis 事务

        隔离：事务中的所有命令都会序列化、按顺序地执行。事务在执行的过程中，不会被其他客户端发送来的命令请求所打断。
        原子：事务中的命令要么全部被执行，要么全部都不执行。
        但不会报错回滚

> redis脚本可代替redis事务, 可以理解为事务是脚本的简便写法.

## 相关命令

- MULTI
- EXEC
- DISCARD
- WATCH

## 用法

    WATCH 事务提供 CAS(check-and-set) 行为(乐观锁), 当watch的内容被其他请求修改, 则整个事务会被取消
    MULTI 开启事务后, 后续命令被放到一个队列中, 当EXEC时才依次执行和返回
    DISCARD 来清空队列 结束事务

    事务中若错误: 不会中断 和 回滚事务, 而是继续执行完事务
        - 执行前无语法检查

```bash
  WATCH                           # 乐观锁
      curr_point = GET user1_point
      curr_gold = GET user1_gold
      curr_point -= x
      curr_gold += x*100
  MULTI                           # 开启事务
      SET user1_point curr_point      # 命令1
      SET user1_gold curr_gold        # 命令2
  EXEC                            # 提交执行

  1) OK                           # 命令1返回
  2) OK                           # 命令2返回
```

## ref

- <http://www.redis.cn/topics/transactions.html>