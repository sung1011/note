# master feature

## kakura

aes-256-cbc

## rabbitmq

流量肖峰

## global

## backend

crontab方案
dynamicTask

## battle

战斗服方案流程 复盘与计算战力

## redis

分片aof导致损坏 -> rdb还原 -- 数据一致性  
集群变单例，master做aof，rewrite失败 -> 改slave做aof -- 备份都在slave做

## mongodb

写入慢 wj -> w1  
增量回档失败 -> 全量回档 -- 全量备份周期优化  

## other

技术选型
DDOS
sdk_source

## bug

福利号数据清空  
道具流向追踪  

## exp

提前准备必要脚本 - 发道具，修改删除数据，充值玩法  
监控报警 - 对redis持久化数据监控(尤其有赛程的玩法)并钉钉报警  
灾难预案 - 断电，物理机损坏，备份还原失败，突发bug协作。。。
单元测试 - 开发效率与单元测试取舍
代码互审 - merge request
错误监控 - 日常监控错误

## method

网
复

## github

`https://github.com/sung1011/note`
