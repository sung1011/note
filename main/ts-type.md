# ts type

```ts
// bool
var foo:boolean = true
foo = false

// number  int+float
var foo:number=123
foo = 3.14

// string
var foo:string = "bar"
foo = "qux"

// array
var foo:number[] = [3, 22, 123]
var foo:string[] = ["aa", "bb"]
var foo:any[] = [2123, "bb"]
var foo:Array<number> = [3, 22, 123]
var foo:Array<any> = [true, "ad", 312]

// tuple 元组
var foo:[string, number, boolean] = ["abc", 123, true]

// enum 枚举
enum foo { success = 1, error = 2 }
var f:foo = foo.success // 1
var g:string = foo[1] // "success"

enum Color{red, blue, orange = 5, yellow}
var b:Color=Color.blue // 1; 索引值
var o:Color=Color.orange // 5; 索引值
var y:Color=Color.yellow // 6; 索引值继续

// any 任意
var foo:any = 123
foo = "xxx"
foo = [2, "d", true]


// void 方法没有返回值
function foo():void {
}

// nerver 声明 null+undefined
var foo:number | undefined
var bar:number | undefined | null

var a:never // any类型的子类型; 不常用, 一般用any替代
a = (() => {
    throw new Error('abc');
})()
```


