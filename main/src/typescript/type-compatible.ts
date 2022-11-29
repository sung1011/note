// 类型兼容
// 结构之间 成员少的兼容成员多的
// 函数之间 参数多的兼容参数少的

// string兼容null
// null是string的子类型
let tc:string = 'abc'
tc = null

// 接口兼容
interface tc1 {
    a:any
    b:any
}
interface tc2 {
    a:any
    b:any
    c:any
}
let tc1:tc1 = {a:1, b:1}
let tc2:tc2 = {a:1, b:1, c:1}
tc1 = tc2 // 兼容
// tc2 = tc1 // err; 不兼容

// 函数兼容
function hof(handler:(a:number, b:number) => void) {

}
let hofhandler1 = (a:number, b:number) => {}
hof(hofhandler1)
let hofhandler2 = (a:number, b:number, c:number) => {}
// hof(hofhandler2) // err; 不兼容
let hofhandler3 = (a:number) => {}
hof(hofhandler3)
let hofhandler7 = (a:string) => {}
// hof(hofhandler7) // err; 不兼容

// 可选参数和剩余参数兼容
let hofhandler4 = (a:number, b:number) => {}
let hofhandler5 = (a?:number, b?:number) => {}
let hofhandler6 = (...args: number[]) => {}

hofhandler4 = hofhandler5
hofhandler4 = hofhandler6

hofhandler5 = hofhandler6 // strictFunctionTypes: false
hofhandler5 = hofhandler4 // strictFunctionTypes: false

hofhandler6 = hofhandler5
hofhandler6 = hofhandler4

// 接口参数兼容
interface Point3D {
    x: number;
    y: number;
    z: number;
}
interface Point2D {
    x: number;
    y: number;
}
let p3d = (p: Point3D) => {};
let p2d = (p: Point2D) => {};
p3d = p2d
// p2d = p3d // err; 不兼容 参考配置 strictFunctionTypes: false

// 返回值兼容
let tcFunc1 = () => ({name: 'alice'})
let tcFunc2 = () => ({name: 'alice', addr: ''})
tcFunc1 = tcFunc2
// tcFunc2 = tcFunc1; err不兼容; 参考配置 strictFunctionTypes: false

// 枚举兼容性
enum Fruit {Apple, Banana}
enum Color {Red, Blue}
// let fruitApple: Fruit.Apple = 3; // err
// let fruitColor: Color.Red = Fruit.Apple // err

// 类兼容性
// 泛型兼容性

