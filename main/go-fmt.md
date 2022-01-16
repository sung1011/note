# go fmt

## code

- [fmt](src/go/package/fmt_test.go)

## 运算符 operand

| operand | type      | desc                                     | ex                                    |
| ------- | --------- | ---------------------------------------- | ------------------------------------- |
| %v      | general   | 万能默认                                 |
| %T      | general   | 类型                                     |
| %%      | general   | 百分号%                                  |
| %d      | int       | 十进制dec(不带符号)                      |
| %+d     | int       | 十进制dec(带符号)                        |
| %o      | int       | 八进制oct(不带0)                         |
| %#o     | int       | 八进制oct(带0)                           |
| %x      | int       | 十六进制hex(小写a~f)                     |
| %X      | int       | 十六进制hex(大写A~F)                     |
| %#x     | int       | 十六进制hex(带0x)                        |
| %U      | int       | Unicode                                  |
| %#U     | int       | Unicode(带字符)                          |
| %b      | int       | 二进制bin                                |
| %q      | int       | 带单引号                                 |
| %5d     | int-width | 右对齐,最少宽度5                         | __123                                 |
| %-5d    | int-width | 左对齐,最少宽度5                         | 123__                                 |
| %05d    | int-width | 0占位,最少宽度5                          | 00123                                 |
| %f      | float     | 保留6位小数点(=%.6f)                     | {"%+.3f", -1.0, "-1.000"},            |
| %e      | float     | 科学计数法                               | {"% .3e", 1.0, " 1.000e+00"},         |
| %g      | float     | 最少数字                                 |
| %.3g    | float     | 截取3位数字                              |
| %.3f    | float     | 截取3位小数                              |
| %s      | str       | 字符串                                   |
| %q      | str       | 字符串(带引号)                           |
| %#q     | str       | 字符串(带反引号, 若存在反引号则用双引号) |
| %x      | str       | 十六进制hex(小写a~f)                     |
| %X      | str       | 十六进制hex(大写A~F)                     |
| % x     | str       | 十六进制hex(带空格)                      | {"% x", "xyz", "78 79 7a"},           |
| %5s     | str-width | 右对齐,最少宽度5                         |
| %-5s    | str-width | 左对齐,最少宽度5                         |
| %.5s    | str-width | 截取5字符                                |
| %5.7s   | str-width | 最少宽度5, 截取7字符                     |
| %v      | struct    | 值                                       | {tickles 32}                          |
| %+v     | struct    | 键值                                     | {Name:tickles Age:32}                 |
| %#v     | struct    | 类型和键值                               | main.sample{Name:\"tickles\", Age:32} |
| %t      | bool      | true/false                               |
| %p      | ptr       | 指针(带有0x的十六进制前缀)               |
| %#p     | ptr       | 指针(不带0x的十六进制前缀)               |

## ref

<https://www.cnblogs.com/forever521Lee/p/10700549.html>