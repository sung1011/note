// 类型保护

enum Type { Strong, Weak }

class java {
    javaV: any
    RunJava() { }
}

class js {
    jsV: any
    RunJs() { }
}

function getLang(type: Type, x: string | number) {
    let lang = type === Type.Strong ? new java() : new js()
    // asset (不推荐)
    if ((lang as java).RunJava) {
        (lang as java).RunJava()
    }
    // instanceof
    if (lang instanceof java) {
        lang.RunJava()
    }
    // in
    if ('javaV' in lang) {
        lang.RunJava()
    }

    // typeof
    if (typeof x === 'string') {
        x.length
    }
    // 类型尾词
    if (isJava(lang)) {
        lang.RunJava()
    }
}

// 类型尾词
function isJava(lang: java | js): lang is java {
    return (lang as java).RunJava !== undefined;
}