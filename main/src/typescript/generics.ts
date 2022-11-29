// any
function createArray1(length: number, value: any): Array<any> {
    let result = [];
    for (let i = 0; i < length; i++) {
        result[i] = value;
    }
    return result;
}

createArray1(3, 'x'); // ['x', 'x', 'x']

// 泛型
function createArray2<T>(length: number, value: T): Array<T> {
    let result: T[] = [];
    for (let i = 0; i < length; i++) {
        result[i] = value;
    }
    return result;
}
createArray2<string>(3, 'x'); // ['x', 'x', 'x'] // 指定泛型类型
createArray2(3, 'x'); // ['x', 'x', 'x'] // 不指定泛型类型(推断)

// 泛型 约束参数
function log1<T>(value: T): T {
    return value
}
log1<string[]>(['x', 'y'])
log1(['x', 'y'])

// 泛型 约束类型定义
type Log2 = <T>(value: T) => T
let log2: Log2 = log1

// 泛型 约束接口
interface iLog1 {
    <T>(value: T): T
}
interface iLog2<T> { // 内部所有方法都是泛型<T>
    (value: T): T
}
// 泛型 约束类
class Log1<T>{
    // static foo(value:T) {} // 静态成员不能被泛型约束
    run(value:T) {
        return value
    }
}
(new Log1<number>()).run(123);
(new Log1()).run({a:1}); // 推断

// 泛型 类型约束
interface Len{
    length: number
}
function Log2<T extends Len>(value:T):T {
    return value
}
Log2(["s"])
// Log2(123) // err; 没有length属性

// 多类型T, U
function swap<T, U>(tuple: [T, U]): [U, T] {
    return [tuple[1], tuple[0]];
}
swap([7, 'seven']); // ['seven', 7]

// 泛型约束 - 无约束
function loggingIdentity1<T>(arg: T): T {
    // console.log(arg.length); // err; 泛型不一定包含length方法
    return arg;
}

// 泛型约束 - 约束存在length方法
interface Lengthwise {
    length: number;
}

function loggingIdentity2<T extends Lengthwise>(arg: T): T {
    console.log(arg.length);
    return arg;
}

// 泛型约束 - 多个类型之间互相约束
function copyFields<T extends U, U>(target: T, source: U): T {
    for (let id in source) {
        target[id] = (<T>source)[id];
    }
    return target;
}
let x = { a: 1, b: 2, c: 3, d: 4 };
copyFields(x, { b: 10, d: 20 });