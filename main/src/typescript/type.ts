// 声明
// { 变量 / 函数 }: type

// bool
var v1:boolean = true
v1 = false

// number  int+float
var v2:number=123
v2 = 3.14

// string
var v3:string = "bar"
v3= "qux"

// array
var v4:number[] = [3, 22, 123]
var v5:string[] = ["aa", "bb"]
var v6:any[] = [2123, "bb"]

var v7:Array<number> = [3, 22, 123]
var v8:Array<any> = [true, "ad", 312]
var v9:Array<number|string> = [3, 22, 123, '4']

// tuple 元组
var v10:[string, number, boolean] = ["abc", 123, true]

// enum 枚举
enum foo { success = 1, error = 2 }
var v11:foo = foo.success // 1
var v12:string = foo[1] // "success"

enum Color{red, blue, orange = 5, yellow}
var b:Color=Color.blue // 1; 索引值
var o:Color=Color.orange // 5; 索引值
var y:Color=Color.yellow // 6; 索引值继续

// any 任意
var v13:any = 123
v13 = "xxx"
v13 = [2, "d", true]
v13 = () => {}

// func 函数
var add = (x: number, y: number) => x+y
var compute: (x: number, y:number) => number
compute = (a, b) => a+b

// obj 对象
var obj1:object = {x: 1, y:2}
// obj.x = 1 // 报错
var obj2:{x: number, y:number} = {x: 1, y:2}
obj2.x = 123

// void 方法没有返回值, 返回undefined
function v14():void {
}

// symbol 具有唯一的值
var sy1:symbol = Symbol()
var sy2 = Symbol()
// sy1 != sy2

// undefined, null 是所有类型的子类型
var un: undefined = undefined
var nu: null = null
var num = 123
num = undefined  // 需要strictNullChecks: fale
num = null  // 需要strictNullChecks: fale

// nerver 永远不会有返回值的类型 null+undefined
var v15 = () => {
    throw new Error('error');
}
var v16 = () => {
    while(true){}
}