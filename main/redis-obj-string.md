# redis 字符串对象

## encoding

int, embstr, raw

实例

```c
// int
redisObject {
    type: REDIS_STRING,
    encoding: REDIS_ENCODING_INT,
    ...
    ptr:  12345
}
//embstr 略
// raw
redisObject {
    type: REDIS_STRING,
    encoding: REDIS_ENCODING_RAW,
    ...
    ptr: &sdshdr{
        len: 5,
        malloc: 5,
        flag: 1,
        buf: [h,e,l,l,o,\0],
    }
}
```

## encoding转换

```js
set foo 123 // int
append foo over // raw (123over)

set foo abc // int
append foo d// raw (abcd)

set foo 123 // int
incr foo // int (124)
```

- embstr是只读的（无修改的方法）当对**embstr进行append**时，会先变为raw再修改。
- 当对**int进行append**时，也会变为raw再修改。
- 当对**int进行incr**等操作，是对int的修改，不会转化类型。

## 实现

| cmd         | int           | embstr               | raw                  |
| ----------- | ------------- | -------------------- | -------------------- |
| set         | int           | embstr               | raw                  |
| get         | copy ->string |
| append      | ->raw         | ->raw                | sdscatlen()          |
| incrbyfloat | ->long double | ->long double or err | ->long double or err |
| incrby      | +             | err                  | err                  |
| decrby      | -             | err                  | err                  |
| strlen      | copy ->string | sdslen()             | sdslen()             |
