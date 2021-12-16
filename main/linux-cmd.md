# linux命令

## 重定向 >

```bash
TODO
```

## |

```bash
# 上个命令的stdout接入到下个命令的stdin
echo 123 | base64
```

## tee

```bash
# stderr, stdout输出到屏幕和文件
echo 123 | tee file1
```

## [dig](linux-cmd-dig.md)

## nsloopup

## cloc

## nc

## [awk](src/cmd/awk.md)  

## [grep](ref/grep.md)  
  
## du  

```bash
du -sh * | sort -n # 文件大小
```

## sort

```bash
sort -nr -k4 # 以第四列进行数值倒序排序 (不加-n是按ascii)
```
  
## lsof  

```bash
lsof -i :{port} #  查看占用端口的进程
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

## sar

```bash
sar -n {sock}
```

## tcpdump

```bash
tcpdump -iany tcp port 9000
```

## netstat

```bash
netstat -nat
```

## ping

## traceroute

## pathping

## mtr

## base64

```bash
# 编码
base64 {file}
echo {txt} | base64

# 解码
base64 -d {code} > {file}
echo {code} | base64 -d
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

## ref

- linux重定向 <http://linux-wiki.cn/wiki/Bash%E7%9A%84%E8%BE%93%E5%85%A5%E8%BE%93%E5%87%BA%E9%87%8D%E5%AE%9A%E5%90%91>