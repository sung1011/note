# scale

## 数据单位

- 1B(byte, 字节) = 8 bit
- 1KB(Kibibyte, 千字节) = 1024B = 2^10 B;
- 1MB(Mebibyte, 兆字节, 百万字节, 简称“兆”) = 1024KB = 2^20 B;
- 1GB(Gigabyte, 吉字节, 十亿字节, 又称“千兆”) = 1024MB = 2^30 B;
- 1TB(Terabyte, 万亿字节, 太字节) = 1024GB = 2^40 B;

> int16范围的bitmap, 2^16/8/1024 = 8k空间

> int32范围的bitmap, 2^32/8/1024/1024 = 512M空间

> int64范围的bitmap, 2^64/8/1024/1024/1024/1024 = 2097152T空间

## 数据实例

- 1英文字符(word)           = 2byte [ASCII, Unicode, UTF-8]
- 1中文字符                 = 2~4byte(B) [ASCII, Unicode]
- 1中文字符                 = 3byte(B) [UTF-8]

## 进制

- B(Binary)二进制
  - 1011B = 11D
- O(Octal)八进制
- D(Decimal)十进制
- H(Hexadecimal)十六进制

## 时间

- 秒s
- 毫秒ms
  - 1000ms = 1s
- 微妙μs
  - 1000μs = 1ms
- 纳秒ns
  - 1000ns = 1μs

- 念 = 0.018s = 1刹那
- 弹指 = 7.2s = 20刹那
- 须臾 = 48min = 160000刹那
- 昼夜 = 24h = 4800000刹那

## 时区

- `UTC` 世界协调时间 Coordinated Universal Time

      调节时钟时间, 世界标准时间, 以原子时秒为基础, 与GMT近似(误差不超过0.9s).

> 北京时区 UTC+8; 北京时间=UTC时间+8小时

> GMT=UTC+0

- `GMT` 格林威治标准时间 Greenwich Mean Time;

      英国伦敦格林尼治标准时间, 太阳横穿格林尼治子午线(本初子午线)时的时间.

- `DST` 夏季节约时间 Daylight Saving Time

      夏季节约时间，即夏令时；是为了利用夏天充足的光照而将时间调早一个小时，北美、欧洲的许多国家实行夏令时；

- `CST` 四个不同时区的缩写

      Central Standard Time (USA) UT-6:00 美国标准时间
      Central Standard Time (Australia) UT+9:30 澳大利亚标准时间
      China Standard Time UT+8:00 中国标准时间
      Cuba Standard Time UT-4:00 古巴标准时间

- `tz` 时区 TimeZone

      以GMT/UTC为起点, 将世界划分24个时区; 英国格林尼治为零时区(GMT+00), 东1-12区, 西1-12区, 北京为东8区(GMT+08).

- `UNIX` 时间戳

      以GMT/UTC时间 [ 1970-01-01T00:00:00 ] 为起点, 到具体时间的秒数; 无关时区.

## 其他

- 256 = 2^8
- 65536 = 2^16
- 4294967296 = 2^32
- 1.84467441e19 = 2^64
- 3.40282367e38 = 2^128
- 1.15792089e77 = 2^256
