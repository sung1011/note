# php-basic
  
## 浮点数  

- 浮点数运算 `bcadd, bcdiv, bcmod, bcmul, bcpow, bcsub, ...`  

```php
$value = "2.1";
if (is_numeric($value) && intval($value) == $value) {
    echo 'int' . PHP_EOL;
    $value += 0;
} elseif (is_numeric($value) && strpos($value, '.') !== false) {
    echo 'float' . PHP_EOL;
    $value = bcadd($value, 0, 2);
}
var_dump($value);
```
