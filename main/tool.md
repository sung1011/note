# tool

## unix2dos

    dos2unix - DOS/Mac - Unix文件格式, 编码转换器 [lf-crlf](ref/lf-crlf.md)

## shadowsocks 翻墙

### require  

    shadowsocks: 翻墙  
    polipo: socks5协议转http/https代理
  
### configuration  

    查看shadowsocks的http代理监听端口(1087), 并设置环境变量

```bash
export http_proxy="http://127.0.0.1:1087"  
export https_proxy=$http_proxy  
```  