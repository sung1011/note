# nginx信号

## master
`TERM`、`INT` 快速关闭  
`QUIT` 从容关闭  
`HUP` 平滑重启，重新加载配置文件  
`USR1` 重新打开日志文件  
`USR2` 平滑升级可执行程序  
`WINCH` 从容关闭工作进程  

## worker
`TERM`  
`QUIT`  
`USR1`  

## nginx -s
stop    SIGTERM
quit    SIGQUIT
reopen  SIGUSR1
reload  SIGHUP

## 实战
### 热部署
1. 移出老nginx二进制，移入新nginx二进制文件。
2. `kill -USR2 < old master pid >` 。 平滑升级: 新老master,work都在运行，老nginx不处理新请求，并不再监听80端口。
3. `kill -WINCH < old master pid >`。 优雅关闭老worker。但保留老master，以备版本回退（其实不必保留）。
4. 需要版本回退时重复热部署操作即可。

### 日志切割
1. mv < log > < bak log >
2. kill -USR1 < nginx master pid >
3. 可将以上放入crontab