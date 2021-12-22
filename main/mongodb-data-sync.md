# MongoDB 数据同步

    为了维护共享数据集的最新副本, 副本集中的从节点成员可以从其他成员同步或复制数据.

## 初始化

    成员间数据全量同步

### 流程

1. 克隆所有数据库(除local库)
2. 对collection应用所有的更改; 应用来自源库的oplog将数据更到最新; 状态从STARTUP2转为SECONDARY

### 选择源

    initialSyncSourceReadPreference

        primary             必须主
        primaryPreferred    优先主
        seconday            必须从
        secondayPreferred   优先从
        nearest             网络延迟最小

### 容错

    选取初始化同步源, 会遍历所有成员两次
    若两次遍历后依然无法选出, 则重试
    重试最多10次


## 复制

    从节点从源复制oplog, 并多线程批量写自己

### 选择源

- 配置 chining
  - true 副本集成员间选择源
  - false 主节点作为同步源

### [流程 oplog](mongodb-oplog.md)

### 容错

    选取初始化同步源, 会遍历所有成员两次
    若两次遍历后依然无法选出, 则重试

## ref

- docs <https://docs.mongodb.com/manual/reference/replica-configuration/#mongodb-rsconf-rsconf.settings.chainingAllowed>