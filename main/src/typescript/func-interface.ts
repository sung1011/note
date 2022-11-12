// 函数类型接口
interface Add {
    (x: number, y: number):number
}
let tmp: Add = (a, b) => {return a+b}