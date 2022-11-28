// arr 初始化
let arr: Array<number> = new Array<number>(); // 泛型
// let arr = [1, 2, 3];                       // 任意类型数组
// let arr: number[] = [1, 2, 3];             // number数组
// interface NumberArray {                    // 接口表示数组(不常用)
//     [index: number]: number;
// }
// let fibonacci: NumberArray = [1, 1, 2, 3, 5];

arr = [1, 2, 3, 7, 5, 9, 1]; // 给数组直接赋值
console.log("数组", arr, "长度为:", arr.length);

// 向数组中增加元素
console.log("向数组 尾 添加一个元素4, 修改后的数组长度为: " + arr.push(4));
console.log("向数组 头 添加一个元素0, 修改后的数组长度为: " + arr.unshift(0));
console.log(arr);

console.log("向数组下标为2插入 一个 元素 99: " + arr.splice(2, 0, 99));
console.log(arr);
console.log("向数组下标为2插入 3个 元素 88,77,66: " + arr.splice(2, 0, 88, 77, 66));
console.log(arr);

// 从数组中删除元素
console.log("从数组 尾 删除一个元素: " + arr.pop());
console.log("从数组 头 删除一个元素: " + arr.shift());
console.log(arr);

console.log("删除数组 下标为4 的元素: " + arr.splice(4, 1));
console.log(arr);
console.log("删除数组中从 下标为2开始后3个 元素: " + arr.splice(2, 3));
console.log(arr);

// 清空数组的几种方式
arr.length = 0;
arr = []; // 只是更改数组的引用地址, 原数组内存将等待垃圾回收
arr.splice(0, arr.length);

arr = [1, 2, 3, 7, 5, 9, 1];
// 获取数组中某个元素
console.log("获取数组 下标为0 的元素: " + arr[0]);
console.log("获取数组 下标为100 的元素: " + arr[100]); // 当该元素没有的时候, 会返回一个undefined类型

// 修改数组中的某个元素
arr[0] = 111;
// arr[-10] = 0; // 我们可以给一个非法下标赋值, 且仍然不会报错, 但我们并不推荐这么做, 这样会把一个数组类型转换为复杂类型
arr[10] = 999; // 当给一个超出当前数组长度的下标元素赋值时, 该数组将会自动补齐数组长度, 且中间未赋值的数组元素为undefined类型
console.log(arr, arr[9], arr.length);

// 查找
console.log("从数组中查找 [元素为3] 的下标: ", arr.indexOf(3));
console.log("从数组中查找 [元素为53] 的下标: ", arr.indexOf(53)); // 当找不到时会返回-1

console.log("数组中是否包含 值为3 的元素", arr.includes(3));

// 查找数组中(从左往右)第一个大于5的数组元素
let item: number = arr.find((val: number, index: number, array: Array<number>) => {
    return val > 5;
});
console.log("数组中(从左往右)第一个大于5的数组元素", item);
// 查找数组中(从左往右)第一个大于5的数组元素的下标
let index: number = arr.findIndex((val: number, index: number, array: Array<number>) => {
    return val > 5;
});
console.log("数组中(从左往右)第一个大于5的数组元素的下标", index);

// 排序
console.log("升序排列后的数组: ", arr.sort(function (a, b) { return a - b }));
console.log("降序排列后的数组: ", arr.sort(function (a, b) { return b - a }));

arr = [1, 99, 88, 111, 200, 4, 5, 7, 222, 555, 30];
console.log("按字符排序后的数组: ", arr.sort());
console.log("反转数组顺序", arr.reverse());

// 遍历数组
for (let i: number = 0; i < arr.length; i++) {
    console.log(arr[i]);
}

for (let i in arr) {
    console.log(`下标${i}, 元素${arr[i]}`);
}

for (let val of arr) {
    console.log(`元素${val}`);
}

// 迭代数组
arr.forEach((val: number, index: number, arr: Array<number>) => {
    console.log(`元素:${val}, 下标${index}, 数组本身${arr}`);
});

let delArr: Array<number> = [1, 3, 5, 7, 12, 14, 19, 20, 25, 30];
// 删除数组中符合条件的某几个元素
let delTempArr: Array<number> = [];
for (let i: number = 0; i < delArr.length; i++) {
    if (delArr[i] % 2 == 0) {
        delTempArr.push(delArr[i]);
    }
}
delArr = delTempArr;
console.log(delArr);

// 获取一个新数组, 该数组中的元素值原数组值的10倍
let arr1: Array<number> = arr.map((val: number, index: number, array: Array<number>) => {
    return val * 10;
});
console.log(arr1);

// 获取一个新数组, 该数组中的元素为原数组中所有大于10的值
let arr2: Array<number> = arr.filter((val: number, index: number, array: Array<number>) => {
    return val > 10;
});
console.log(arr2);

console.log(arr);
let sum = arr.reduce((total, val: number, index: number, array: Array<number>) => {
    console.log(index, val, total);
    return total + val;
});
console.log("总和", sum);

// 判断该数组是否 每个元素都大于0, 该方法一旦判定为假将不会再迭代
let isHas: boolean = arr.every((val: number, index: number, array: Array<number>) => {
    return val > 0;
});
console.log("判断该数组是否 每个元素都大于0", isHas);
// 判断该数组是否 每个元素都小于200
let isHas0: boolean = arr.every((val: number, index: number, array: Array<number>) => {
    console.log(val < 200);
    return val < 200;
});
console.log("判断该数组是否 每个元素都小于200", isHas0);

// 判断该数组中是否 存在大于100的元素, 该方法一旦判定为真将不会再迭代
let isHas1: boolean = arr.some((val: number, index: number, array: Array<number>) => {
    console.log(val > 100);
    return val > 100;
});
console.log("该数组中是否存在大于100的元素", isHas1);

// 连接两个数组到一个新数组
let arr0 = [-1, 2, 4, 6, 7, 999];
let arrConcat: Array<number> = arr.concat(arr0);
console.log("arr和arr0连接后的新数组", arrConcat);

// 将数组元素填充为某个值, 可用作初始化数组
let arr3: Array<number> = new Array<number>(10);
arr3.fill(1);
console.log("数组每个元素初始化为1", arr3);

// 将数组元素转换为字符串
let arr4: Array<string> = ["hello", ",", "my", " ", "word!"];
console.log("将数组元素转换为字符串", arr4.join(""), "用逗号连接成字符串", arr4.join(","));
// 将字符串转换为数组
let str: string = "你好,我的朋友!";
console.log(str.split(""));

// 获取子数组
let arr6 = ["111", "333", "str", "is", "fff", "sos"];
console.log("下标从0到3的子数组", arr6.slice(0, 4));