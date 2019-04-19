<?php
// 意图：将对象组合成树形结构以表示"部分-整体"的层次结构。组合模式使得用户对单个对象和组合对象的使用具有一致性。

// 主要解决：它在我们树型结构的问题中，模糊了简单元素和复杂元素的概念，客户程序可以向处理简单元素一样来处理复杂元素，从而使得客户程序与复杂元素的内部结构解耦。

// 何时使用： 1、您想表示对象的部分-整体层次结构（树形结构）。 2、您希望用户忽略组合对象与单个对象的不同，用户将统一地使用组合结构中的所有对象。

// 如何解决：树枝和叶子实现统一接口，树枝内部组合该接口。

// 关键代码：树枝内部组合该接口，并且含有内部属性 List，里面放 Component。

// 应用实例： 1、算术表达式包括操作数、操作符和另一个操作数，其中，另一个操作符也可以是操作数、操作符和另一个操作数。 2、在 JAVA AWT 和 SWING 中，对于 Button 和 Checkbox 是树叶，Container 是树枝。

// 优点： 1、高层模块调用简单。 2、节点自由增加。

// 缺点：在使用组合模式时，其叶子和树枝的声明都是实现类，而不是接口，违反了依赖倒置原则。

// 使用场景：部分、整体场景，如树形菜单，文件、文件夹的管理。

// 注意事项：定义时为具体类。
interface Component
{
    public function getComposite(); //返回自己的实例
    public function operation();
}
 
class Composite implements Component
{ // 树枝组件角色
    private $_composites;
    public function __construct()
    {
        $this->_composites = array();
    }
    public function getComposite()
    {
        return $this;
    }
    public function operation()
    {
        foreach ($this->_composites as $composite) {
            $composite->operation();
        }
    }
 
    public function add(Component $component)
    {  //聚集管理方法 添加一个子对象
        $this->_composites[] = $component;
    }
 
    public function remove(Component $component)
    { // 聚集管理方法 删除一个子对象
        foreach ($this->_composites as $key => $row) {
            if ($component == $row) {
                unset($this->_composites[$key]);
                return true;
            }
        }
        return false;
    }

    public function getChild()
    { // 聚集管理方法 返回所有的子对象
        return $this->_composites;
    }
}
 
class Leaf implements Component
{
    private $_name;
    public function __construct($name)
    {
        $this->_name = $name;
    }
    public function operation()
    {
    }
    public function getComposite()
    {
        return null;
    }
}
 
// client
$leaf1 = new Leaf('first');
$leaf2 = new Leaf('second');

$composite = new Composite();
$composite->add($leaf1);
$composite->add($leaf2);
$composite->operation();

$composite->remove($leaf2);
$composite->operation();



//透明式合成模式
interface Component
{ // 抽象组件角色
    public function getComposite(); // 返回自己的实例
    public function operation(); // 示例方法
    public function add(Component $component); // 聚集管理方法,添加一个子对象
    public function remove(Component $component); // 聚集管理方法 删除一个子对象
    public function getChild(); // 聚集管理方法 返回所有的子对象
}
 
class Composite implements Component
{ // 树枝组件角色
    private $_composites;
    public function __construct()
    {
        $this->_composites = array();
    }
    public function getComposite()
    {
        return $this;
    }
    public function operation()
    { // 示例方法，调用各个子对象的operation方法
        foreach ($this->_composites as $composite) {
            $composite->operation();
        }
    }
    public function add(Component $component)
    { // 聚集管理方法 添加一个子对象
        $this->_composites[] = $component;
    }
    public function remove(Component $component)
    { // 聚集管理方法 删除一个子对象
        foreach ($this->_composites as $key => $row) {
            if ($component == $row) {
                unset($this->_composites[$key]);
                return true;
            }
        }
        return false;
    }
    public function getChild()
    { // 聚集管理方法 返回所有的子对象
        return $this->_composites;
    }
}
 
class Leaf implements Component
{
    private $_name;
    public function __construct($name)
    {
        $this->_name = $name;
    }
    public function operation()
    {
        echo $this->_name."<br>";
    }
    public function getComposite()
    {
        return null;
    }
    public function add(Component $component)
    {
        return false;
    }
    public function remove(Component $component)
    {
        return false;
    }
    public function getChild()
    {
        return null;
    }
}
 
// client
$leaf1 = new Leaf('first');
$leaf2 = new Leaf('second');

$composite = new Composite();
$composite->add($leaf1);
$composite->add($leaf2);
$composite->operation();

$composite->remove($leaf2);
$composite->operation();
