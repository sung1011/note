<?php
// 意图：提供一种方法顺序访问一个聚合对象中各个元素, 而又无须暴露该对象的内部表示。

// 主要解决：不同的方式来遍历整个整合对象。

// 何时使用：遍历一个聚合对象。

// 如何解决：把在元素之间游走的责任交给迭代器，而不是聚合对象。

// 关键代码：定义接口：hasNext, next。

// 应用实例：JAVA 中的 iterator。

// 优点： 1、它支持以不同的方式遍历一个聚合对象。 2、迭代器简化了聚合类。 3、在同一个聚合上可以有多个遍历。 4、在迭代器模式中，增加新的聚合类和迭代器类都很方便，无须修改原有代码。

// 缺点：由于迭代器模式将存储数据和遍历数据的职责分离，增加新的聚合类需要对应增加新的迭代器类，类的个数成对增加，这在一定程度上增加了系统的复杂性。

// 使用场景： 1、访问一个聚合对象的内容而无须暴露它的内部表示。 2、需要为聚合对象提供多种遍历方式。 3、为遍历不同的聚合结构提供一个统一的接口。

// 注意事项：迭代器模式就是分离了集合对象的遍历行为，抽象出一个迭代器类来负责，这样既可以做到不暴露集合的内部结构，又可让外部代码透明地访问集合内部的数据。
class sample implements Iterator
{
    private $_items ;
 
    public function __construct(&$data)
    {
        $this->_items = $data;
    }
    public function current()
    {
        return current($this->_items);
    }
 
    public function next()
    {
        next($this->_items);
    }
 
    public function key()
    {
        return key($this->_items);
    }
 
    public function rewind()
    {
        reset($this->_items);
    }
 
    public function valid()
    {
        return ($this->current() !== false);
    }
}
 
// client
$data = array(1, 2, 3, 4, 5);
$sa = new sample($data);
foreach ($sa as $key => $row) {
    echo $key, ' ', $row, '<br />';
}


//Yii FrameWork Demo
class CMapIterator implements Iterator
{
    /**
    * @var array the data to be iterated through
    */
    private $_d;
    /**
    * @var array list of keys in the map
    */
    private $_keys;
    /**
    * @var mixed current key
    */
    private $_key;
 
    /**
    * Constructor.
    * @param array the data to be iterated through
    */
    public function __construct(&$data)
    {
        $this->_d=&$data;
        $this->_keys=array_keys($data);
    }
 
    /**
    * Rewinds internal array pointer.
    * This method is required by the interface Iterator.
    */
    public function rewind()
    {
        $this->_key=reset($this->_keys);
    }
 
    /**
    * Returns the key of the current array element.
    * This method is required by the interface Iterator.
    * @return mixed the key of the current array element
    */
    public function key()
    {
        return $this->_key;
    }
 
    /**
    * Returns the current array element.
    * This method is required by the interface Iterator.
    * @return mixed the current array element
    */
    public function current()
    {
        return $this->_d[$this->_key];
    }
 
    /**
    * Moves the internal pointer to the next array element.
    * This method is required by the interface Iterator.
    */
    public function next()
    {
        $this->_key=next($this->_keys);
    }
 
    /**
    * Returns whether there is an element at current position.
    * This method is required by the interface Iterator.
    * @return boolean
    */
    public function valid()
    {
        return $this->_key!==false;
    }
}
 
$data = array('s1' => 11, 's2' => 22, 's3' => 33);
$it = new CMapIterator($data);
foreach ($it as $row) {
    echo $row, '<br />';
}
