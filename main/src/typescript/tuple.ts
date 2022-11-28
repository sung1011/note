// 声明
let alice:[string, number]
alice[0] = 'alice'
alice[1] = 18
alice[0].slice(1)

// 赋值
let bob:[string, number] = ["bob", 18];
// let bob: [string, number] = ["bob"]; // err: 必须赋值全部

// 越界
let cindy: [string, number] = ['cindy', 18]
cindy.push('female') // 越界后继续添加元素, 元素类型为变量类型的联合类型 string|number
// cindy.push(true) // err; 非string|number