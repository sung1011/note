# grep

简介

--

　　grep (global search regular expression(RE) and print out the line,全面搜索正则表达式并把行打印出来)是一种强大的文本搜索工具，它能使用正则表达式搜索文本，并把匹配的行打印出来。

　　Unix的grep家族包括grep、egrep和fgrep。egrep和fgrep的命令只跟grep有很小不同。egrep是grep的扩展，支持更多的re元字符， fgrep就是fixed grep或fast grep，它们把所有的字母都看作单词，也就是说，正则表达式中的元字符表示回其自身的字面意义，不再特殊。linux使用GNU版本的grep。它功能更强，可以通过-G、-E、-F命令行选项来使用egrep和fgrep的功能。

　　grep的工作方式是这样的，它在一个或多个文件中搜索字符串模板。如果模板包括空格，则必须被引用，模板后的所有字符串被看作文件名。搜索的结果被送到标准输出，不影响原文件内容。grep可用于shell脚本，因为grep通过返回一个状态值来说明搜索的状态，如果模板搜索成功，则返回0，如果搜索不成功，则返回1，如果搜索的文件不存在，则返回2。我们利用这些返回值就可进行一些自动化的文本处理工作。

命令格式：grep \[option\] pattern file

**grep的常用选项：** -V： 打印grep的版本号 -E： 解释PATTERN作为扩展正则表达式，也就相当于使用egrep。 或操作 -F :   解释PATTERN作为固定字符串的列表，由换行符分隔，其中任何一个都要匹配。也就相当于使用fgrep。 -G:   将范本样式视为普通的表示法来使用。这是默认值。加不加都是使用grep。

**匹配控制选项：** -e :  使用PATTERN作为模式。这可以用于指定多个搜索模式，或保护以连字符（ - ）开头的图案。指定字符串做为查找文件内容的样式。 -f :  指定规则文件，其内容含有一个或多个规则样式，让grep查找符合规则条件的文件内容，格式为每行一个规则样式。 -i :  搜索时候忽略大小写 -v:  反转匹配，选择没有被匹配到的内容。 -w：匹配整词，精确地单词,单词的两边必须是非字符符号(即不能是字母数字或下划线) -x：仅选择与整行完全匹配的匹配项。精确匹配整行内容(包括行首行尾那些看不到的空格内容都要完全匹配) -y：此参数的效果和指定“-i”参数相同。

**一般输出控制选项：** -c： 抑制正常输出;而是为每个输入文件打印匹配线的计数。 --color \[= WHEN\]：让关键字高亮显示，如--color=auto -L：列出文件内容不符合指定的范本样式的文件名称 -l : 列出文件内容符合指定的范本样式的文件名称。 -m num：当匹配内容的行数达到num行后,grep停止搜索,并输出停止前搜索到的匹配内容 -o: 只输出匹配的具体字符串,匹配行中其他内容不会输出 -q：安静模式,不会有任何输出内容,查找到匹配内容会返回0,未查找到匹配内容就返回非0 -s：不会输出查找过程中出现的任何错误消息，-q和-s选项因为与其他系统的grep有兼容问题，shell脚本应该避免使用-q和-s，并且应该将标准和错误输出重定向到/dev/null 代替。

**输出线前缀控制：** -b：输出每一个匹配行(或匹配的字符串)时在其前附加上偏移量(从文件第一个字符到该匹配内容之间的字节数) -H：在每一个匹配行之前加上文件名一起输出(针对于查找单个文件),当查找多个文件时默认就会输出文件名 -h：禁止输出上的文件名的前缀。无论查找几个文件都不会在匹配内容前输出文件名 --label = LABEL：显示实际来自标准输入的输入作为来自文件LABEL的输入。这是特别在实现zgrep等工具时非常有用，例如gzip -cd foo.gz | grep --label = foo -H的东西。看到 也是-H选项。 -n：输出匹配内容的同时输出其所在行号。 -T：初始标签确保实际行内容的第一个字符位于制表位上，以便对齐标签看起来很正常。在匹配信息和其前的附加信息之间加入tab以使格式整齐。

**上下文线控制选项：** -A num：匹配到搜索到的行以及该行下面的num行 -B num：匹配到搜索到的行以及该行上面的num行 -C num：匹配到搜索到的行以及上下各num行

**文件和目录选择选项：** -a： 处理二进制文件，就像它是文本;这相当于--binary-files = text选项。不忽略二进制的数据。 --binary-files = TYPE：如果文件的前几个字节指示文件包含二进制数据，则假定该文件为类型TYPE。默认情况下，TYPE是二进制的，grep通常输出一行消息二进制文件匹配，或者如果没有匹配则没有消息。如果TYPE不匹配，grep假定二进制文件不匹配;这相当于-I选项。如果TYPE是文本，则grep处理a二进制文件，如果它是文本;这相当于-a选项。警告：grep --binary-files = text可能会输出二进制的垃圾，如果输出是一个终端和如果可能有讨厌的副作用终端驱动程序将其中的一些解释为命令。 -D：如果输入文件是设备，FIFO或套接字，请使用ACTION处理。默认情况下，读取ACTION，这意味着设备被读取，就像它们是普通文件一样。如果跳过ACTION，设备为 默默地跳过。 -d:  如果输入文件是目录，请使用ACTION处理它。默认情况下，ACTION是读的，这意味着目录被读取，就像它们是普通文件一样。如果跳过ACTION，目录将静默跳过。如果ACTION是recurse，grep将递归读取每个目录下的所有文件;这是相当于-r选项。 --exclude=GLOB：跳过基本名称与GLOB匹配的文件（使用通配符匹配）。文件名glob可以使用*，？和\[...\]作为通配符，和\\引用通配符或反斜杠字符。搜索其文件名和GLOB通配符相匹配的文件的内容来查找匹配使用方法:grep -H --exclude=c* "old" ./* c*是通配文件名的通配符./* 指定需要先通配文件名的文件的范围,必须要给*,不然就匹配不出内容,(如果不给*,带上-r选项也可以匹配)

--exclude-from = FILE：在文件中编写通配方案,grep将不会到匹配方案中文件名的文件去查找匹配内容

--exclude-dir = DIR：匹配一个目录下的很多内容同时还要让一些子目录不接受匹配,就使用此选项。

 --include = GLOB：仅搜索其基本名称与GLOB匹配的文件（使用--exclude下所述的通配符匹配）。

-R ,-r :以递归方式读取每个目录下的所有文件; 这相当于-d recurse选项。

**其他选项：**

--line-buffered： 在输出上使用行缓冲。这可能会导致性能损失。

--mmap：启用mmap系统调用代替read系统调用

-U：将文件视为二进制。

-z：将输入视为一组行，每一行由一个零字节（ASCII NUL字符）而不是a终止新队。与-Z或--null选项一样，此选项可以与排序-z等命令一起使用来处理任意文件名。

简述

-a   --text   #不要忽略二进制的数据。 将 binary 文件以 text 文件的方式搜寻数据 
-A<显示行数>   --after-context=<显示行数> #除了显示符合范本样式的那一列之外，并显示该行之后的内容。   
-b   --byte-offset   #在显示符合样式的那一行之前，标示出该行第一个字符的编号。   
-B<显示行数>   --before-context=<显示行数> #除了显示符合样式的那一行之外，并显示该行之前的内容。   
**-c    --****count   #计算符合样式的行数。**   
-C<显示行数>    --context=<显示行数>或-<显示行数> #除了显示符合样式的那一行之外，并显示该行之前后的内容。   
-d <动作>      --directories=<动作> #当指定要查找的是目录而非文件时，必须使用这项参数，否则grep指令将回报信息并停止动作。   
-e<范本样式>  --regexp=<范本样式> #指定字符串做为查找文件内容的样式。   
-E      --extended-regexp   #将样式为延伸的普通表示法来使用。   
-f<规则文件>  --file=<规则文件> #指定规则文件，其内容含有一个或多个规则样式，让grep查找符合规则条件的文件内容，格式为每行一个规则样式。   
-F   --fixed-regexp   #将样式视为固定字符串的列表。   
-G   --basic-regexp   #将样式视为普通的表示法来使用。   
-h   --no-filename   #在显示符合样式的那一行之前，不标示该行所属的文件名称。   
-H   --with-filename   #在显示符合样式的那一行之前，表示该行所属的文件名称。   
**-i    --ignore-case** **#忽略字符大小写的差别。**   
-l    --file-with-matches   #列出文件内容符合指定的样式的文件名称。   
-L   --files-without-match   #列出文件内容不符合指定的样式的文件名称。   
**-n   --line-****number   #在显示符合样式的那一行之前，标示出该行的列数编号。 **  
-q   --quiet或--silent   #不显示任何信息。   
-r   --recursive   #此参数的效果和指定“-d recurse”参数相同。   
**-s   --no-****messages   #不显示错误信息。**   
**-v   --revert-****match   #显示不包含匹配文本的所有行。**   
-V   --version   #显示版本信息。   
-w   --word-regexp   #只显示全字符合的列。   
-x    --line-regexp   #只显示全列符合的列。   
-y   #此参数的效果和指定“-i”参数相同。

--color=auto ：可以将找到的关键词部分加上颜色的显示

**使用实例：**
=========

**一、常用用法**
----------

grep -i pattern files ：不区分大小写地搜索。默认情况区分大小写，
grep -l pattern files ：只列出匹配的文件名，
grep -L pattern files ：列出不匹配的文件名，
grep -w pattern files ：只匹配整个单词，而不是字符串的一部分（如匹配‘magic’，而不是‘magical’），
grep -C number pattern files ：匹配的上下文分别显示\[number\]行，
grep pattern1 | pattern2 files ：显示匹配 pattern1 或 pattern2 的行，
grep pattern1 files | grep pattern2 ：显示既匹配 pattern1 又匹配 pattern2 的行。
  
这里还有些用于搜索的特殊符号： < 和 > 分别标注单词的开始与结尾。
例如：
grep man * 会匹配 ‘Batman’、‘manic’、‘man’等，
grep \'<man\\' * 匹配‘manic’和‘man’，但不是‘Batman’，
grep \'<man>\\' 只匹配‘man’，而不是‘Batman’或‘manic’等其他的字符串。
\'^\\'：指匹配的字符串在行首，
\'$\\'：指匹配的字符串在行尾，
如果您不习惯命令行参数，可以试试图形界面的‘grep’，如 reXgrep 。这个软件提供 AND、OR、NOT 等语法，还有漂亮的按钮 :-) 。如果您只是需要更清楚的输出，不妨试试 fungrep 。

.grep 搜索字符串
命令格式:
grep string filename
寻找字串的方法很多，比如说我想找所有以M开头的行.此时必须引进pattern的观
念.以下是一些简单的□例，以及说明： ^M 以M开头的行，^表示开始的意思
M$ 以M结尾的行，$表示结束的意思 ^\[0-9\] 以数字开始的行，\[\]内可列举字母 ^\[124ab\] 以1,2,4,a,或b开头的行 ^b.503 句点表示任一字母 * 星号表示0个以上的字母(可以没有) + 加号表示1个以上的字母
. 斜线可以去掉特殊意义 <eg> cat passwd | grep ^b 列出大学部有申请帐号者名单
cat passwd | grep ^s 列出交换学生申请帐号者名单
cat passwd | grep \'^b.503\\' 列出电机系各年级...
grep \'^.\\' myfile.txt 列出所有以句点开头的行

**1、查找指定进程**

命令：ps -ef|grep java

**2、查找指定进程个数**

命令：ps -ef|grep -c java

或ps -ef|grep java **-c**

**3、从文件中读取关键词进行搜索，默认是显示的是行 **

命令1：cat test.txt | grep -f test2.txt

命令2（**显示行号**）：cat test.txt | grep -**n**f test2.txt

作用：输出test.txt文件中含有从test2.txt文件中读取出的关键词的内容行，可用于按指定关键词（放到一个文件中）搜索日志文件。

**另一种用法：将多个文件之间相同的行输出来**

\# cd /etc/sysconfig/network-scripts/

\# grep  "IPADDR" ifcfg-eth0  ifcfg-lo      #默认不加参数指定过滤关键字，外加多个文件，只是将多个文件里面有匹配的行输出

ifcfg-eth0:IPADDR=192.168.1.108

ifcfg-lo:IPADDR=127.0.0.1

\# grep **-f ifcfg-eth0  ifcfg-lo**  #grep -f 文件1 文件2 ,会将多个文件之间相同的行输出出来

ONBOOT=yes

**-o:只显示被模式匹配到的字符串，而不是整个行**

命令：grep -o "you" ab.log 

\# grep "root"  /etc/passwd  #先看下正常的过滤，会将整个一行过滤出来

root:x:0:0:root:/root:/bin/bash

operator:x:11:0:operator:/root:/sbin/nologin

\# grep -o "root" /etc/passwd  #加o之后的操作，只过滤关键字出来

root

root

root

root

\# grep -o "root:.*0" /etc/passwd    #加上正则表达式，这样才是正确的用法，不用输出一整行，只是输出一小段

root:x:0:0

\# grep -o "root" -b   /etc/passwd  

-b和-o一般是配合使用的，一行中字符串的字符是从该行的第一个字符开始计算，起始值为0。这里左边的数字就是此关键字在此文件中的起始位置，第一个root出现在0位置，然后字符字母有一个算一个，你就一个个的向右数吧，下一个root出现在11位置以此类推。

0:root

11:root

17:root

414:root

**4、从文件中查找关键词，忽略大小写，默认情况区分大小写  
**

命令1：grep 'linux' test.txt

命令2（从多个文件中查找）：grep 'linux' test.txt test2.txt

　　多文件时，输出查询到的信息内容行时，会把文件的命名在行最前面输出并且加上":"作为标示符

命令3（忽略大小写）：grep -i  'linux' test.txt

命令：find . -name ".log" | grep -i error | grep -vi "info"

1）使用find -name 来列出所有log文件，重定向给grep  
2）使用grep -i 来查找包含error的行  
3）使用grep -vi 来查找不包含info的行

**5、grep不显示本身**

命令：

ps aux|grep \\\[s\]sh

ps aux | grep ssh | grep -v "grep"      #不包含grep ssh这条命令

grep -v root /etc/passwd | grep -v nologin  #将/etc/passwd，将没有出现 root 和nologin的行取出来

6、**-r 递归查找子目录**  
 查找当前目录及其子目录下面包含匹配字符的文件

\# grep ‘ab’ * #在当前目录搜索带'ab'行的文件  
\# grep -r ‘ab’ * #在当前目录及其子目录下搜索'ab'行的文件  
\# grep -l -r ‘ab’ * #在当前目录及其子目录下搜索'ab'行的文件，但是不显示匹配的行，只显示匹配的文件  
\# grep -nr BLOG* . # 查找子目录，匹配后输出行号，这里的点表示当前目录  
\# grep -lr BLOG* . #查找子目录，匹配后只输出文件名

查询不包含某个目录

#grep -R --exclude-dir=node_modules 'some pattern' /path/to/search   #不包含txt目录 

7、**列出关键字所在行的前几行与后几行也一起显示**

**  **-A -B -C****

很多时候，我们并关心匹配行而是关心匹配行的上下文。这时候-A -B -C就有用了  
-A n 后n行，A记忆为(After)  
-B n 前n行，B记忆为(Before)  
-C n 前n行，后n行，C记忆为(Center)

\[root@www ~\]# dmesg | grep -n -A3 -B2 --color=auto 'eth'
245-PCI: setting IRQ 10 as level-triggered 246-ACPI: PCI Interrupt 0000:00:0e.0\[A\] -> Link \[LNKB\] ... 247:eth0: RealTek RTL8139 at 0xee846000, 00:90:cc:a6:34:84, IRQ 10
248:eth0: Identified 8139 chip type 'RTL-8139C'
249-input: PC Speaker as /class/input/input2 250-ACPI: PCI Interrupt 0000:00:01.4\[B\] -> Link \[LNKB\] ... 251-hdb: ATAPI 48X DVD-ROM DVD-R-RAM CD-R/RW drive, 2048kB Cache, UDMA(66)
\# 如上所示，你会发现关键字 247 所在的前两行及 248 后三行也都被显示出来！

**8、--line-buffered 打开buffering 模式**

有一个文件是动态的，它不断地添加信息到文件的尾部，而你想要输出包含某些信息的行。即持续的grep一个动态的流

**9、e与E区别**

grep想同时过滤多个条件**或操作**

错误写法：

\# netstat -an|grep "ESTABLISHED|WAIT"      #默认grep不支持多条件匹配    

正确写法：

\# netstat -an|grep **-E "ESTABLISHED|WAIT"**     #加上-E 多条件用""包起来，然后多条件之间用|管道符分开

tcp        0     52 192.168.1.108:22            192.168.1.104:54127         ESTABLISHED 

\# ps -aux|grep **-e udevd -e master**|awk {'print $(NF-1)'}|sort|uniq    #而-e呢不用""包起来，-e 指定一个匹配条件

/sbin/udevd

/usr/bin/salt-master

  grep -E '123|abc' filename  // 找出文件（filename）中包含123或者包含abc的行  
  egrep '123|abc' filename    // 用egrep同样可以实现  
  awk '/123|abc/' filename   // awk 的实现方式

**与操作**

 grep pattern1 files | grep pattern2 ：显示既匹配 pattern1 又匹配 pattern2 的行。

**10、-c 统计行数**

\# grep -i "abc" test.txt|wc -l  #不分大小写。test.txt里面包含abc过滤条件的为2行

2

\# grep -yc "abc" test.txt  #-c呢，就是不显示行的内容，直接显示有几行

2

\# cat  /etc/passwd|wc -l 

55

\# grep  -c "^.*$" /etc/passwd  #那么我们除了wc -l用来统一一个文件有多少行以外，又多了一种统计文件多少行的方法

55

**11、 -m的使用**

\# cat test2.txt  #这是测试文件

abc 1

abc 2

abc 3

abc 4

abc 5

\# grep -m 3 "abc" test2.txt  #只匹配到了第三行就退出了

abc 1

abc 2

abc 3

**二、与正则表达式结合**
--------------

grep的规则表达式:
\     反义字符：如"\\"\\""表示匹配"" \[ \- \] 匹配一个范围，\[0-9a-zA-Z\]匹配所有数字和字母 * 所有字符，长度可为0 + 前面的字符出现了一次或者多次 ^  #匹配行的开始 如：'^grep'匹配所有以grep开头的行。    
$  #匹配行的结束 如：'grep$'匹配所有以grep结尾的行。    
.  #匹配一个非换行符的字符 如：'gr.p'匹配gr后接一个任意字符，然后是p。 \*  #匹配零个或多个先前字符 如：'*grep'匹配所有一个或多个空格后紧跟grep的行。    
.* #一起用代表任意字符。   
\[\]   #匹配一个指定范围内的字符，如'\[Gg\]rep'匹配Grep和grep。    
\[^\]  #匹配一个不在指定范围内的字符，如：'\[^A-FH-Z\]rep'匹配不包含A-R和T-Z的一个字母开头，紧跟rep的行。    
\\(..\\)  #标记匹配字符，如'\\(love\\)'，love被标记为1。    
\<      #到匹配正则表达式的行开始，如:'\\<grep'匹配包含以grep开头的单词的行。    
\\>      #到匹配正则表达式的行结束，如'grep\\>'匹配包含以grep结尾的单词的行。    
x\\{m\\}  #重复字符x，m次，如：'0\\{5\\}'匹配包含5个o的行。    
x\\{m,\\}  #重复字符x,至少m次，如：'o\\{5,\\}'匹配至少有5个o的行。    
x\\{m,n\\}  #重复字符x，至少m次，不多于n次，如：'o\\{5,10\\}'匹配5--10个o的行。   
\\w    #匹配文字和数字字符，也就是\[A-Za-z0-9\]，如：'G\\w*p'匹配以G后跟零个或多个文字或数字字符，然后是p。   
\\W    #\\w的反置形式，匹配一个或多个非单词字符，如点号句号等。   
\\b    #单词锁定符，如: '\\bgrep\\b'只匹配grep。  
  
POSIX字符:
为了在不同国家的字符编码中保持一至，POSIX(The Portable Operating System Interface)增加了特殊的字符类，如\[:alnum:\]是\[A-Za-z0-9\]的另一个写法。要把它们放到\[\]号内才能成为正则表达式，如\[A- Za-z0-9\]或\[\[:alnum:\]\]。在linux下的grep除fgrep外，都支持POSIX的字符类。
\[:alnum:\]    #文字数字字符   
\[:alpha:\]    #文字字符   
\[:digit:\]    #数字字符   
\[:graph:\]    #非空字符（非空格、控制字符）   
\[:lower:\]    #小写字符   
\[:cntrl:\]    #控制字符   
\[:print:\]    #非空字符（包括空格）   
\[:punct:\]    #标点符号   
\[:space:\]    #所有空白字符（新行，空格，制表符）   
\[:upper:\]    #大写字符   
\[:xdigit:\]   #十六进制数字（0-9，a-f，A-F） 

**例：通过管道过滤ls -l输出的内容，只显示以a开头的行。**

首与行尾字节 ^ $，^ 符号，在字符类符号(括号\[\])之内与之外是不同的！ 在 \[\] 内代表『反向选择』，在 \[\] 之外则代表定位在行首的意义！

$ ls -l | grep \\'^a\\'

$ ls -l | grep  ^a

$ ls -l | grep  ^\[^a\]      #输出非a开头的行，反向选择

$ grep -n '^$' express.txt      #找出空白行，因为只有行首跟行尾 (^$)

**例：显示所有以d开头的文件中包含test的行。**

$ grep \\'test\\' d*

****例：**输出以hat结尾的行内容**

$ cat test.txt |grep hat$

****例：**显示在aa，bb，cc文件中匹配test的行。**

$ grep \\'test\\' aa bb cc

****例：**显示所有包含每个字符串至少有5个连续小写字符的字符串的行。**

在一组集合字节中，如果该字节组是连续的，例如大写英文/小写英文/数字等等，就可以使用\[a-z\],\[A-Z\],\[0-9\]等方式来书写，那么如果我们的要求字串是数字与英文呢？就将他全部写在一起，变成：\[a-zA-Z0-9\]。

$ grep \\'\[a-z\]{5}\\' aa

$ grep -n '\[0-9\]' regular_express.txt  　　#取得有数字的那一行

$ grep -n '^\[a-z\]' regular_express.txt 　　 #只输出开头是小写字母的那一行

$ grep -n '^\[^a-zA-Z\]' regular_express.txt   #不输出开头是英文的

$ grep -n '\\.$' regular_express.txt   　　　　 #只输出行尾结束为小数点 (.) 的那一行

注意：小数点具有其他意义，所以必须要使用转义字符(\\)来加以解除其特殊意义！

******例：****显示包含ed或者at字符的内容行**

命令：cat test.txt |grep -E "ed|at"

**例：**如果west被匹配，则es就被存储到内存中，并标记为1，然后搜索任意个字符（.*），这些字符后面紧跟着另外一个es（1），找到就显示该行。如果用egrep或grep -E，就不用""号进行转义，直接写成\\'w(es)t.*1\\'就可以了。

$ grep \\'w(es)t.*1\\' aa

****例：**显示当前目录下面以.txt 结尾的文件中的所有包含每个字符串至少有7个连续小写字符的字符串的行**

命令：grep '\[a-z\]\\{7\\}' *.txt

**例：查询IP地址、邮箱、手机号**

这里用到了-o和-P命令  
man grep查看  
-o, --only-matching：  
              Show only the part of a matching line that matches PATTERN.  
-P, --perl-regexp：  
              Interpret PATTERN as a Perl regular expression.

也就是说-o，只显示匹配行中匹配正则表达式的那部分，-P，作为Perl正则匹配

192.168.0.1

abc@163.com

匹配ABC类IP地址即 1.0.0.1---223.255.255.254

命令（IP）：grep -oP "(\[0-9\]{1,3}\\.){3}\[0-9\]{1,3}" file.txt

或grep -E --color "\\<(\[1-9\]|\[1-9\]\[0-9\]|1\[0-9\]\[0-9\]|2\[0-1\]\[0-9\]|22\[0-3\])\\.(\[0-9\]|\[1-9\]\[0-9\]|1\[0-9\]\[0-9\]|2\[0-4\]\[0-9\]|25\[0-5\])\\.(\[0-9\]|\[1-9\]\[0-9\]|1\[0-9\]\[0-9\]|2\[0-4\]\[0-9\]|25\[0-5\])\\.(\[1-9\]|\[1-9\]\[0-9\]|1\[0-9\]\[0-9\]|2\[0-4\]\[0-9\]|25\[0-4\])\\>" file.txt

邮箱是任意长度数字字母@任意长度数字字母

命令（邮箱）：grep -oP "\[a-zA-Z0-9_-\]+@\[a-zA-Z0-9_-\]+(\\.\[a-zA-Z0-9_-\]+)+" file.txt

手机号码是1\[3|4|5|8\]后面接9位数字的

命令（手机号）：grep -E "\\<1\[3|4|5|8\]\[0-9\]{9}\\>"  file.txt

**例：任意一个字节 . 与重复字节 ***

. (小数点)：代表『一定有一个任意字节』的意思； \* (星号)：代表『重复前一个字符， 0 到无穷多次』的意思，为组合形态

$ grep -n '\[0-9\]\[0-9\]*' regular_express.txt     #找出『任意数字』的行

$ grep -n 'g.*g' regular_express.txt       #找出以g行首与行尾的行，当中的字符可有可无

这个 .\* 的 RE 表示任意字符是很常见的.

**例：限定连续 RE 字符范围 {}**

利用 . 与 RE 字符及 * 来配置 0 个到无限多个重复字节

打算找出两个到五个 o 的连续字串，该如何作？这时候就得要使用到限定范围的字符 {} 了。 但因为 { 与 } 的符号在 shell 是有特殊意义的，因此， 我们必须要使用字符   \ 来让他失去特殊意义才行。 

$ grep -n 'o\\{2\\}' regular_express.txt

$ grep -n 'go\\{2,5\\}g' regular_express.txt  #要找出 g 后面接 2 到 5 个 o ，然后再接一个 g 的字串

$ grep -n 'go\\{2,\\}g' regular_express.txt    #想要的是 2 个 o 以上的 goooo....g 呢？除了可以是 gooo*g 