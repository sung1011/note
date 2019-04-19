# linux

## 内存分配和多线程原理

## 进程间通信方式
- 管道
- 信号量
- 信号
- 消息队列
- 套接字

## io模型
- 阻塞
- 非阻塞
- 多路复用
- 信号驱动
- 异步

## crontab

## related
[文件描述符](related/文件描述符.md)  
[关于网络IO中的同步、异步、阻塞、非阻塞](related/关于网络IO中的同步、异步、阻塞、非阻塞.md)

## Makefile

## cmd
- [awk](linux/awk.md)
    - 某列求和 `awk -F "\t" '{sum+=$1}END{print sum}'`
- [grep](linux/grep.md)
- du
    - 文件大小 `du -sh * | sort -n`
- lsof
    - 端口占用排序排重 `lsof -n | awk '{print $2}'|sort|uniq -c |sort -nr`
- ulimit
    - 每个进程可打开的文件数 `ulimit -n`
- fuser
    - 查看占用端口的进程 `fuser -n tcp 9000`

### HugePage 透明大页
- [hugepage](linux/hugepage.md)

### 其他
- grep, awk, top, ps, tar, ln
- telnet, lsof, scp, ssh, vim
