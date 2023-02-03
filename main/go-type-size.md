# go type size

| type        | size (byte) | value range                                |
| ----------- | ----------- | ------------------------------------------ |
| bool        | 1           | true/false                                 |
| ptr         | 8           |                                            |
| channel     | 8           | ptr                                        |
| map         | 8           | ptr                                        |
| func        | 8           | ptr                                        |
| slice       | 24          | ptr + len + cap                            |
| array       | ~           | T * len                                    |
| struct      | ~           |                                            |
| string      | 16          | ptr + len                                  |
| uint8       | 1           | 0~255                                      |
| uint16      | 2           | 0~65535                                    |
| uint32/rune | 4           | 0~4294967295                               |
| int64/int   | 8           | 0~18446744073709551615                     |
| int8        | 1           | -128~127                                   |
| int16       | 2           | -32768 ~ 32767                             |
| int32/rune  | 4           | -2147483648 ~ 2147483647                   |
| uint64/int  | 8           | -9223372036854775808 ~ 9223372036854775807 |
| float32     | 4           | -2147483648 ~ 2147483647                   |
| float64     | 8           | -9223372036854775808 ~ 9223372036854775807 |
| int         | 4/8         |

> `unsafe.Sizeof(T)` 占用字节数

> `int` 4byte in 32bit is equivalent to int32; 8byte in 64bit is equivalent to int64

> `timestamp` 如 1604556795 # int32: 2^32=4294967296 能容纳 10位数

> `timestamp millisecond` 如 1604556795000 # int64: 2^64=1.844674407370955e19 能容纳 20位数
