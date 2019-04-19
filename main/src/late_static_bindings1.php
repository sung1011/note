<?php
class A
{
    public static function who()
    {
        echo __CLASS__;
    }
    public static function test()
    {
        static::who(); // 后期静态绑定从这里开始
    }
}

class B extends A
{
    public static function who()
    {
        echo __CLASS__;
    }
}

B::test();//B
