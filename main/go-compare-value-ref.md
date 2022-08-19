# go compare value | ref

        本质上传递的都是指针的值 (方法内指针地址都变了)
        值类型新指针指向了copy后的新值
        引用类型新指针指向旧值

## code

- [params](src/go/basic/params_test.go)

## ref

- <https://www.flysnow.org/2018/02/24/golang-function-parameters-passed-by-value.html>