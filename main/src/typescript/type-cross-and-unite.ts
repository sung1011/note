// 交叉接口
interface iDog {
    run(): void
}
interface iCat {
    jump(): void
}

let pet: iDog & iCat = {
    run() { },
    jump() { }
}

// 联合类型
let au1: string | number = 1
let au2: 2 | 3 | 4 = 3
let au3: 'x' | 'y' | 'z' = 'x'