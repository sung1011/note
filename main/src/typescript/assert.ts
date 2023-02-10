// <类型>值     (不推荐)
// 值 as 类型   (推荐)

interface ApiError extends Error {
    code: number
}

interface HttpError extends Error {
    statusCode: number
}

function isApiError(err:Error) {
    if (typeof (err as ApiError).code == 'number') {
        return true
    }
    return false
}

// 联合类型可以被断言为其中一个类型
// 父类可以被断言为子类
// 任何类型都可以被断言为 any
// any 可以被断言为任何类型