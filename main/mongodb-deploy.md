# MongoDB 部署

## 数量

    部署奇数个成员
    副本集最多50个成员, 可投票成员最多7个

## 容错

| num-of-members | majority-required-to-elect | fault-tolerance |
| -------------- | -------------------------- | --------------- |
| 3              | 2                          | 1               |
| 4              | 3                          | 1               |
| 5              | 3                          | 2               |
| 6              | 4                          | 2               |

> 单个副本集不应部署多个Arbiter

## Hidden和Delay属性

    备份, 统计, 减缓压力

## 读写分离

## 异地多活

    每个IDC至少1个成员, 最好奇数个IDC

## 使用tag操作成员

## 使用journaling方式停电

## 主机名

    使用逻辑DNS主机名替换IP
