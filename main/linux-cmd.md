# linux命令

## nc

## [awk](src/cmd/awk.md)  

## [grep](ref/grep.md)  
  
## du  

```bash
du -sh * | sort -n # 文件大小
```
  
## lsof  

```bash
lsof -n | awk '{print $2}' | sort | uniq -c |sort -nr # 端口占用排序排重
```
  
## ulimit  

```bash
ulimit -n # 每个进程可打开的文件数
```
  
## fuser  

```bash
fuser -n tcp 9000
```

## sar -n {sock}

## lsof -i :{port}

```bash
lsof -i :{port} #  查看占用端口的进程
```

## tcpdump -iany tcp port 9000

## netstat -nat

## ping

## traceroute

## pathping

## mtr

## nsloopup

## base64

- 编码 `base64 {file}`
- 解码 `base64 -d {code} > {file}`, `echo {code} | base64 -d > {file}`
