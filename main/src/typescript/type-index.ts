// 索引类型

// keyof T
interface iObjIndex{
    a:number,
    b:string,
}
let key: keyof iObjIndex

// T[K]
let value: iObjIndex['a']

// T extends U

// 索引约束+泛型
let objIndex = {
    a: 1,
    b: 2,
    c: 3,
}

function getValues(obj: any, keys: string[]) {
    return keys.map(key => obj[key])
}
function getValuesNew<T, K extends keyof T>(obj: T, keys: K[]): T[K][] {
    return keys.map(key => obj[key])
}
