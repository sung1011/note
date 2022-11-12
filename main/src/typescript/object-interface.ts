/**
 * 定义接口对象
 */
interface Person {
    readonly id: string;    // 必须实现, 只读属性赋值对象后不可修改
    name: string;           // 必须实现
    age: number;            // 必须实现
    gender?: string;        // 可选属性, 不必实现
    [xx: string]: any;      // 任意属性
}

/**
 * 实现接口 需要严格实现接口定义 id, name, age
 * 可选属性gender可以实现 也可不实现
 */
let tom: Person = {
    id: 'abc',          // 必须实现
    name: 'Tom',        // 必须实现
    age: 25,            // 必须实现
    // gender: 'male',  // 可以不实现
    xyz: {},            // 任意属性
};

/**
 * err: 只读属性赋值后不可修改
 */
// tom.id = 'a'

/**
 * err: 少实现age 
 */
// let tom2: Person = {
//     id: 'abc',
//     name: 'Tom',
// };

/**
 * err: 多实现addr 
 */
// let tom3: Person = {
//     id: 'abc',
//     name: 'Tom',
//     age: 25,
//     addr: 'xxx',
// };
