# cron

## config

- `/etc/crontab` # 系统级。一般不要动
- `/etc/cron.d/` # 系统级。全局计划任务时可改
- `/var/spool/cron/<user>` # 用户级。可用 crontab -l, -e等进行操作。

> 以上配置修改后，都不必重载服务即生效

## crontab

```bash
# 以下都是用户级操作, 对应/var/spool/cron/
-e　编辑该用户的计时器设置,如果不指定用户，则表示编辑当前用户的crontab文件
-l　列出该用户的计时器设置,如果不指定用户，则表示显示当前用户的crontab文件内容。
-r　删除某个用户的crontab文件，如果不指定用户，则默认删除当前用户的crontab文件。
-i  在删除用户的crontab文件时给确认提示。
-u user 用来设定某个用户的crontab服务，例如，“-u ixdba”表示设定ixdba用户的crontab服务，此参数一般有root用户来运行。
file file是命令文件的名字,表示将file做为crontab的任务列表文件并载入crontab (覆盖到/var/spool/cron/)。
```

## cmd

```bash

# 环境
SHELL=/bin/bash
PATH=/sbin:/bin:/usr/sbin:/usr/bin
HOME=/User/sunji # 执行脚本的主目录

# 示例
HOME=/User/sunji1
* * * * * a.sh
HOME=/User/sunji2
* * * * * b.sh

# For details see man 4 crontabs

# Example of job definition:
# .---------------- minute (0 - 59)
# |  .------------- hour (0 - 23)
# |  |  .---------- day of month (1 - 31)
# |  |  |  .------- month (1 - 12) OR jan,feb,mar,apr ...
# |  |  |  |  .---- day of week (0 - 6) (Sunday=0 or 7) OR sun,mon,tue,wed,thu,fri,sat
# |  |  |  |  |
# *  *  *  *  * user-name  command to be executed

# 示例
30 21 * * * /usr/local/etc/rc.d/lighttpd restart       #每晚的21:30重启apache。
45 4 1,10,22 * * /usr/local/etc/rc.d/lighttpd restart  #每月1、10、22日的4 : 45重启apache。
10 1 * * 6,0 /usr/local/etc/rc.d/lighttpd restart      #每周六、周日的1 : 10重启apache
0,30 18-23 * * * /usr/local/etc/rc.d/lighttpd restart  #每天18 : 00至23 : 00之间每隔30分钟重启apache。
0 23 * * 6 /usr/local/etc/rc.d/lighttpd restart        #每星期六的11 : 00 pm重启apache。
* */1 * * * /usr/local/etc/rc.d/lighttpd restart       #每一小时重启apache
* 23-7/1 * * * /usr/local/etc/rc.d/lighttpd restart    #晚上11点到早上7点之间，每隔一小时重启apache
0 11 4 * mon-wed /usr/local/etc/rc.d/lighttpd restart  #每月的4号与每周一到周三的11点重启apache
0 4 1 jan * /usr/local/etc/rc.d/lighttpd restart       #一月一号的4点重启apache
*/30 * * * * /usr/sbin/ntpdate 210.72.145.44           #每半小时同步一下时间

# 超时kill 与 锁
* * * * * timeout -s SIGINT 100 flock -xn /tmp/lock /path/to/php /path/to/file

# run-parts 执行文件夹内所有
01 * * * * root run-parts /etc/cron.hourly //每小时执行/etc/cron.hourly内的脚本
02 4 * * * root run-parts /etc/cron.daily //每天执行/etc/cron.daily内的脚本
22 4 * * 0 root run-parts /etc/cron.weekly //每星期执行/etc/cron.weekly内的脚本
42 4 1 * * root run-parts /etc/cron.monthly //每月去执行/etc/cron.monthly内的脚本
```

## crond

```bash
systemctl start   crond    #启动服务
systemctl stop    crond    #关闭服务
systemctl restart crond    #重启服务
systemctl reload  crond    #重新载入配置
systemctl status  crond    #查看crontab服务状态
systemctl enable  crond    #开启开机自动启动
systemctl disable crond    #禁止开机自动启动

# or

/sbin/service crond start    #启动服务
/sbin/service crond stop     #关闭服务
/sbin/service crond restart  #重启服务
/sbin/service crond reload   #重新载入配置
ntsysv                       #查看crontab服务是否已设置为开机启动（用方向键和tab操作）
chkconfig –level 35 crond on #开机自动启动
```
