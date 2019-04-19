<?php
class A
{
    public static function foo()
    {
        static::who();
    }

    public static function who()
    {
        echo __CLASS__."\n";
    }
}

class B extends A
{
    public static function test()
    {
        A::foo();//A
        parent::foo();//C
        self::foo();//C
    }

    public static function who()
    {
        echo __CLASS__."\n";
    }
}
class C extends B
{
    public static function who()
    {
        echo __CLASS__."\n";
    }
}

C::test();
