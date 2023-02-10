# base64

    数据压缩算法, 是二进制文件的64进制显示

## 用途

    传输二进制文件时, 3Byte就需要24字符
    使用base64只需要传递3个字符
    不能使用ascii码是因为ascii码很多是无法输出的控制字符(可打印字符95个) 

## base64可打印字符

    字母A-Z、a-z    (26+26=52)
    数字0-9         (10)
    +、/            (2)
    =               (标记补全)

## 编码规则

    字符对应的二进制 切成 每6位一个字符, 与码表的64个字符对应

```js
    # 24bit
    Man的二进制 M（01001101）a（01100001）n（01101110）
    即01001101 01100001 01101110

    按每6个比特位一组来分组 010011 010110 000101 101110

    将6进制编译成Base64对应的字符：T W F u

    # 16bit
    Ma的二进制 01001101 01100001
    分组结果 010011 010110 0001
    补全 010011 010110 000100, 即TWE
    用1个=标记最后2位为补全, 即TWE=

    # 8bit
    M的二进制 01001101
    分组结果 010011 01
    补全 010011 010000, 即TQ
    用2个=标记最后4位为补全, 即TQ==
```

## go.base64

```go
	data := []byte("Man")
	str := base64.StdEncoding.EncodeToString(data) // TWFu

	data := []byte("ni hao!")
	str := base64.StdEncoding.EncodeToString(data) // bmkgaGFvIQ==

    data := []byte(`-_.~!*'();:@&=+$,/?#[]`)
	str := base64.StdEncoding.EncodeToString(data) // LV8ufiEqJygpOzpAJj0rJCwvPyNbXQ
```

## base64码表

                      Table 1: The Base 64 Alphabet

     Value Encoding  Value Encoding  Value Encoding  Value Encoding
         0 A            17 R            34 i            51 z
         1 B            18 S            35 j            52 0
         2 C            19 T            36 k            53 1
         3 D            20 U            37 l            54 2
         4 E            21 V            38 m            55 3
         5 F            22 W            39 n            56 4
         6 G            23 X            40 o            57 5
         7 H            24 Y            41 p            58 6
         8 I            25 Z            42 q            59 7
         9 J            26 a            43 r            60 8
        10 K            27 b            44 s            61 9
        11 L            28 c            45 t            62 +
        12 M            29 d            46 u            63 /
        13 N            30 e            47 v
        14 O            31 f            48 w         (pad) =
        15 P            32 g            49 x
        16 Q            33 h            50 y

## ref

- RFC <https://www.rfc-editor.org/rfc/rfc4648.html>