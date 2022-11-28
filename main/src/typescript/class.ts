// 继承; 访问修饰符
class Father {
    // 访问修饰符 public, protected, private, readonly
    public readonly identity: string;

    constructor(public readonly firstName:string) {
        this.identity = 'father'
    }

    introduce(): string {
        return this.identity
    }
}

class Son extends Father {
    public name: string
    private age: number | undefined;
    protected sex: number;
    constructor(name: string, age: number) {
        super('sun') // 初始化父类的constructor
        this.name = name
        this.age = age
        this.sex = 1
    }

    hello(): string {
        return 'hello, I\'m '+this.name+ ' my dad is '+super.introduce() // 只能访问super方法, 不能访问属性
    }
}

let w = new Son("sun", 33)
w.hello()

// 存取器
class chaosName {
    get name() {
        return 'XXX'
    }

    set name(v) {
    }
}

let chaosNameObj = new chaosName()
chaosNameObj.name = 'daf'


// 静态方法
class Animal {
    private name:string
    constructor(name: string) {
        this.name = name
    }
    static isAnimal(a:Animal) {
        return a instanceof Animal;
    }
}
let a = new Animal('lion');
Animal.isAnimal(a);

// 抽象方法
abstract class Animal2 {
    public name;
    public constructor(name:string) {
        this.name = name;
    }
    public abstract sayHi():void;
}

class Cat extends Animal2 {
    public sayHi() {
        console.log(`Meow, My name is ${this.name}`);
    }
}