class Foo {
     // 访问修饰符 public, protected, private
    private name: string
    protected age: number | undefined;   
    public sex: number;

    constructor(name:string, age?:number) {
        this.name = name
        this.age = age
        this.sex = 2
    }

    run():void{
    }
}

class Bar extends Foo {
    constructor(name: string, age: number) {
        super(name, age) // 初始化父类的constructor
    }
}

var w = new Bar("sun", 33)
w.run()
