# php 异常错误相关

    异常处理通常是防止未知错误产生所采取的处理措施.

## 组成

    检测(try)  
    抛出(throw)  
    捕获(catch)  

## 错误处理函数

    设置用户自定义的错误处理函数
    `set_error_handler`

### 作用范围

    可捕捉: E_NOTICE 、E_USER_ERROR、E_USER_WARNING、E_USER_NOTICE
    不可捕捉: E_ERROR, E_PARSE, E_CORE_ERROR, E_CORE_WARNING, E_COMPILE_ERROR and E_COMPILE_WARNING.

## 异常处理函数

    设置默认的异常处理程序, 用于没有用 try/catch 块来捕获的异常. 在 exception_handler 调用后异常会中止.
    `set_exception_handler`

## 终止处理函数

    注册一个会在php中止时执行的函数
    `register_shutdown_function`
