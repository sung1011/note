# redis过期

## 依赖

依赖计算机时钟，调整时钟会令key立刻过期。  

## 被动删除

当尝试访问它时，key会被发现并主动的过期。  

## 主动删除(定期删除)

1. 检测随机的20个keys是否过期, 过期则删除
2. 如果有多于25%的keys过期, 立刻重复步奏1

> 后台线程周期性执行

> 源码 src/expire.c activeExpireCycle()

## 复制AOF过期

slaves不会独立处理过期, 会等到master执行DEL命令。
