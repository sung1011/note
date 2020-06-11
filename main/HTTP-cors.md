# 跨域 cors

## 同源策略 SOP same origin policy

协议 + 域名 + 端口 相同即同源

> 限制行为: cookie, LocalStorage 和 IndexDB; DOM, JS; AJAX;

## 解决方案

1. 通过jsonp跨域
2. document.domain + iframe跨域
3. location.hash + iframe
4. window.name + iframe跨域
5. postMessage跨域
6. 跨域资源共享（CORS）

    Access-Control-Allow-Origin: { 允许跨域的host，若无端口结尾不加/ }
    Access-Control-Allow-Credentials: true // 允许前端带认证cookie：启用此项后，上面的域名不能为'*'
    Access-Control-Allow-Headers: Content-Type,X-Requested-With // 提示OPTIONS检测时，后端需要设置的常用自定义头

7. nginx代理跨域
8. nodejs中间件代理跨域
9. WebSocket协议跨域
