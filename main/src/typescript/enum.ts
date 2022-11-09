// 反向映射 (值是数字时)
enum Err {
    Notice,
    Info,
    Warn,
    Err,
}

Err["1"] // Info
Err.Info // 1

// 字符串枚举
enum Msg {
    Success = '成功',
    Fail = '失败'
}
Msg.Success

// 枚举成员
enum Char {
    a,           // 无初始值
    b = Char.a,  // 引用
    c = 1+3,     // 表达式
    d = Math.random(),
    e = '123'.length
}

// 常量枚举
const enum Month {
    Jan,
    Feb,
    Mar,
}
Month.Feb

// 枚举类型
enum E {a, b}
enum F {a = 0, b = 1}
enum G {a = 'apple', b = 'banana'}