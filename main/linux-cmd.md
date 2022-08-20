# linux命令

## stdbuf

```bash
tail -f access.log | stdbuf -oL cut -d' ' -f1 | uniq # 输出不缓冲, 直接显示; | (管道) 会将内容read到kernel, 具有缓冲区, 未写满的缓冲无法传递给后续程序
```

> 管道与缓冲 <https://www.cnblogs.com/outsrkem/p/11200697.html>

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
# 清空file, 输出重定向到file; bash的
cmd > file

# 不清空file, 输出重定向到file
cmd >> file

# cmd读入file
cmd < file
# 如 yum install -y $(<package.list)

# cmd << file; 
cat <<EOF
abc
def
EOF

cmd <<< word

cmd <> file

cmd >| file

```

### echo

```bash
echo a; echo b
a
b

# -n 输出出现在同一行
echo -n a;echo b 
ab

# -e 解释特殊字符
echo -e "hello\nword" 
hello
word
```

### && ||

```bash
cat log && ls log # cat成功才会ls
cat log || ls log # cat失败才会ls
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

```bash
# 扫描IP段 -v 可视化; -z 扫描时不发数据; -w 超时时间
nc -vzw 2 47.93.191.198 8080-8088
# Connection to 47.93.191.198 port 8080 [tcp/http] succeeded!
# nc: connectx to 47.93.191.198 port 8081 (tcp) failed: Connection refused
# ...
# nc: connectx to 47.93.191.198 port 8088 (tcp) failed: Connection refused
```

## [awk](linux-cmd-awk.md)  

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

    协议: ICMP（Internet Control Message Protocol 因特网报文控制协议）

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

## top

```bash
    go build -o foo && ./foo
    top -pid $(pidof foo)       
```


## ref

- linux重定向 <http://linux-wiki.cn/wiki/Bash%E7%9A%84%E8%BE%93%E5%85%A5%E8%BE%93%E5%87%BA%E9%87%8D%E5%AE%9A%E5%90%91>