<?php

# 多重锁示例
if(!lock1()){
    throw new Exception("");
}
if(!lock2()){
    unlock1();
    throw new Exception("");
}
try {
    # do something
} catch (Exception $ex) {
    throw new Exception("");
} finally {
    unlock1();
    unlock2();
}