# go compare

## [os | bufio | ioutil](go-compare-rw.md)

        read: 在G级别以上 os, bufio更优; 
        read oneline: 在G级别以上 ReadSlice() ReadLine() 优于 ReadBytes() ReadString() 

## [gzip | zip | tar](go-compare-zip.md)

        zip是win的, gzip压缩, tar打包

## [struct.field.func | type-interface](go-compare-fieldfunc-interface.md.md)

        field.func 每次实例化都要实现一遍

## [for | range](go-compare-for-range.md)

        出现copy会慢一些
        map的遍历无论是否copy都慢 (毕竟key是分散的)

## [new | make](go-compare-new-make.md)

        make只适用于chan map slice;
        make可以指定len cap
        make返回类型本身

## [参数传递值 | 引用](go-compare-value-ref.md)

        todo