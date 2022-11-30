// 映射类型

interface objMapping {
    a: string
    b: number
    c: boolean
}

// 映射为readonly属性
type ReadonlyObj = Readonly<objMapping>

// 映射为可选属性
type PartialObj = Partial<objMapping>

// 映射为某些属性
type PickObj = Pick<objMapping, 'a' | 'b'>

// 映射为新属性
type RecordObj = Record<'x' | 'y', objMapping>