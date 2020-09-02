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

```bash
# 编码
base64 {file}
# 解码
base64 -d {code} > {file}
echo {code} | base64 -d > {file}
```

## curl

```bash
# application/x-www-form-urlencoded
curl localhost:3000/api/basic -X POST -d 'hello=world'

# application/json
curl localhost:3000/api/json -X POST -d '{"hello": "world"}' --header "Content-Type: application/json"

# 文件内容作为要提交的数据
curl localhost:3000/api/json -X POST -d @data.json --header "Content-Type: application/json" # json
curl localhost:3000/api/basic -X POST -d @data.txt # x-www-from-urlencoded

# multipart/form-data 上传文件
curl localhost:3000/api/multipart -F raw=@raw.data -F hello=world

```
