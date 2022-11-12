
// 函数
function foo1() { }

// 无返回
function foo2():void { }

// 匿名
var foo3 = function(){ }

// 必选参数
function foo4(name:string) { }

// 可选参数 可选参数必须放在后面
function foo5(name:string, age?:number) { }

// 默认可选参数
function foo6(name:string, age:number = 30) { }

// 剩余参数
function foo7(init:number, ...rest:number[]): number {
    return init + rest.reduce((pre, cur) => pre + cur)
}
foo7(1000, 1, 2, 3, 4, 5)


// 重载 同名函数不同参数
function foo8(name:string):string;
function foo8(age:number):number;
function foo8(name:string, age:number):number;
function foo8(p1:any, p2?:any):any { // 最宽泛的实现, 在函数内部区分参数类型
    if (p2 != undefined) {
        // name age
    } else if (typeof p1 == 'string') {
        // name
    } else {
        // age
    }
}
foo8(2)
foo8("zxc")
foo8("zxc", 123)
// foo8(true) ; error

// 箭头函数
function foo9(cb:() => void) {}
foo9(() => {
})