/**
 * 类实现接口
 */
interface Animals {
    // new (name: string): void     // 只能实现公有成员, 构造函数无法被实现, 所以不要定义构造函数
    eat(): void;
}
interface Human extends Animals {    // 接口继承接口
    foo: string;
}
interface Speaker {
    lang: string
}
class Asian implements Human, Speaker { // 实现多个接口
    foo: string;      // 实现Human
    lang: string;      // 实现speaker
    constructor(foo: string) {
        this.lang = ""
        this.foo = ""
    }
    eat() { }         // 实现Animal
    sleep() { }       // 可以额外实现方法
}

/**
 * 接口实现类 
 * (几乎用不到)
 */
class Point {
    x: number;
    y: number;
    constructor(x: number, y: number) {
        this.x = x;
        this.y = y;
    }
}

interface Point3d extends Point {
    z: number;
}

let point3d: Point3d = { x: 1, y: 2, z: 3 };