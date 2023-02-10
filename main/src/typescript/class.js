// ES5
function Foo(name, age) {
    // 属性
    this.name = name;
    this.age = age;
    // 实例方法
    this.run = function () {
    }
}

// 原型链 属性
Foo.prototype.sex = "man"
// 原型链 方法
Foo.prototype.work = function () { }
// 静态方法
Foo.getInfo = function () { }

// 实例化
var f = new Foo("sun", 33);
f.name // sun
f.run()
f.sex // man
f.work()
Foo.getInfo()

// 继承 对象冒充实现继承;
// [x] 可以继承构造函数
// [x] 实例化时给父类传参
// [ ] 继承原型链
function Bar(name, age) {
    Foo.call(this, name, age) // 对象冒充继承 `构造 + 参数`
}
var b = new Bar("bbb", 123);
b.run() // ; 可以继承构造函数的方法
b.name // bbb; 可以给父类传参
// b.sex // undefined; 无法继承原型链上的属性和方法

// 继承 原型链实现继承;
// [x] 可以继承构造函数
// [ ] 实例化时给父类传参
// [x] 继承原型链
function Baz(name, age) { }
Baz.prototype = new Foo(); // 原型链继承 `构造 + 原型链`
var z = new Baz("zzz", 789)
z.run()
// z.name // undefined; 无法给父类传参
z.sex // man; 可以继承原型链方法

// 继承 原型链+对象冒充
// [x] 可以继承构造函数
// [x] 实例化时给父类传参
// [x] 继承原型链
function Quz(name, age) {
    Foo.call(this, name, age) // 对象冒充继承 `构造 + 参数`
}
Quz.prototype = Foo.prototype // 原型链继承 `原型链`
// 或 Quz.prototype = new Foo() // 原型链继承 `构造 + 原型链`
var q = new Quz("qqq", 9999)
q.run()
q.name // qqq
q.sex // man