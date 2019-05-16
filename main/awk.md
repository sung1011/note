# awk  
  
## 基本用法  
awk 动作 文件名  
  
`awk '{print $1}' demo.txt`  打印第一列  
  
## 变量  
FILENAME：当前文件名  
FS：字段分隔符，默认是空格和制表符。  
RS：行分隔符，用于分割每一行，默认是换行符。  
OFS：输出字段的分隔符，用于打印时分隔字段，默认为空格。  
ORS：输出记录的分隔符，用于打印时分隔记录，默认为换行符。  
OFMT：数字输出的格式，默认为％.6g。  
  
`awk -F ':' '{print $(NF-1)}' demo.txt`  倒数第二个字段  
  
## 函数  
tolower()：字符转为小写。  
length()：返回字符串长度。  
substr()：返回子字符串。  
sin()：正弦。  
cos()：余弦。  
sqrt()：平方根。  
rand()：随机数。  
  
`awk -F ':' '{ print toupper($1) }' demo.txt`  转化大写  
  
## 条件  
awk '条件 动作' 文件名  

`awk -F ':' '/usr/ {print $1}' demo.txt`  只输出包含usr的行  
`awk -F ':' 'NR >3 {print $1}' demo.txt`  输出第三行以后的行  
`awk -F ':' '$1 == "root" || $1 == "bin" {print $1}' demo.txt`  第一个字段等于指定值的行  
  
## if语句  
`awk -F ':' '{if ($1 > "m") print $1; else print "---"}' demo.txt`  

[更多](related/awk.md)