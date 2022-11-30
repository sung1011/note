// 条件类型

type TypeName<T> =
    T extends string ? "string" :
    T extends number ? "num" :
    T extends boolean ? "bool" :
    T extends undefined ? "undefined" :
    T extends Function ? "func" :
    "obj"

type T1 = TypeName<string>
type T2 = TypeName<string[]>
// (A|B) extends U ? X : Y
// (A extends U ? X : Y) | (B extends U ? X : Y)

type T3 = TypeName<string | string[]>

// Exclude<T, U>
type Diff<T, U> = T extends U ? never : T
type T4 = Diff<"a" | "b" | "c", "a" | "e">
type T4_1 = Exclude<"a" | "b" | "c", "a" | "e">

// NonNullable<T>
type NotNull<T> = Diff<T, undefined | null>
type T5 = NotNull<string | number | undefined | null>
type T5_1 = NonNullable<string | number | undefined | null>

// Extract<T, U>
type T6 = Extract<"a" | "b" | "c", "a" | "e">

// ReturnType<T>
type T7 = ReturnType<() => string>