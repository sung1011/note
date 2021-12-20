<?php 
/**
 *简单工厂又叫静态工厂方法模式, 这样理解可以确定, 简单工厂模式是通过一个静态方法创建对象的.
 */
// 区分: 用来生产同一等级结构中的任意产品.(在遵循开闭原则(对拓展开放, 对修改关闭)的条件下,  不支持拓展增加产品)
// 特点
// 	1、它是一个具体的类, 而非接口或抽象类.有一个重要的create()方法, 利用if或者 switch创建产品并返回.
// 	2、create()方法通常是静态的, 所以也称之为静态工厂.
// 缺点
// 	1、扩展性差(我想增加一个产品, 除了新增一个产品类, 还需要修改工厂类方法)
// 	2、不同的产品需要不同额外参数的时候 不支持.


class BMWWheel
{
}
class MercedesWheel
{
}
class BYDWheel
{
}

class WheelFactoty // 限定了Wheel产品
{
    public static function createBMWWheel()
    {
        return new BMWWheel();
    }
    public static function createMercedesWheel()
    {
        return new MercedesWheel();
    }
    public static function createBYDWheel()
    {
        return new BYDWheel();
    }
}
 
WheelFactoty::createBMWWheel();
WheelFactoty::createMercedesWheel();
WheelFactoty::createBYDWheel();
