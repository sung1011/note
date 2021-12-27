# linux命令

## openssl

```bash
# sha256校验文件签名
openssl digest -sha256 go1.17.5.darwin-amd64.pkg
# crc32校验文件
crc32 go1.17.5.darwin-amd64.pkg
```

## 重定向 >

| 文件描述符FD | 名称         | 缩写   | 默认值 |
| ------------ | ------------ | ------ | ------ |
| 0            | 标准输入     | stdin  | 键盘   |
| 1            | 标准输出     | stdout | 屏幕   |
| 2            | 标准错误输出 | stderr | 屏幕   |

> 黑洞 /dev/null

### 简单重定向

```bash
# cmd > file; 清空file, 输出重定向到file; bash的

# cmd >> file; 不清空file, 输出重定向到file

# cmd < file; 使cmd从file读入
yum install -y $(<package.list)

# cmd << file; 
cat <<EOF
abc
def
EOF

# cmd <<< word

# cmd <> file

# cmd >| file

```

### FD重定向

```bash
# cmd >&n	把输出送到文件描述符n

# cmd m>&n	把输出 到文件符m的信息重定向到文件描述符n

# cmd >&-	关闭标准输出

# cmd <&n	输入来自文件描述符n

# cmd m<&n	m来自文件描述符n

# cmd <&-	关闭标准输入

# cmd <&n-	移动输入文件描述符n而非复制它.(需要解释)

# cmd >&n-	移动输出文件描述符 n而非复制它.(需要解释)

```

### 组合重定向

```bash
# cmd 2>file	把文件描述符2重定向到file, 即把错误输出存到file中.
find ~ -type f 2> /dev/null

# cmd > file 2>&1	把标准错误重定向到标准输出, 再重定向到file, 即stderr和stdout都被输出到file中

# cmd &> file	功能与上一个相同, 更为简便的写法.

# cmd >& file	功能仍与上一个相同.

# cmd > f1 2>f2	把stdout重定向到f1, 而把stderr重定向到f2

# tee files	把stdout原样输出的同时, 复制一份到files中.

# tee files	把stderr和stdout都输出到files中, 同时输出到屏幕.

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

综艺和电影我经常是看了二创才去爱优腾和影院的
实际蹭热度吸流量 不等于 损害平台利益
之前看到爱奇艺老总访谈时说: 一场球赛用户从二创得知最终比分和哪个球员进的 就不会来平台看整场球赛了.
觉得好扯, 那体育新闻更剧透.
营收一开始下滑就立刻找茬和涨价.
