# k8s ingress

类似service, 但是service是四层转发, ingress是七层转发

## 场景

- 多个Service共享同一个IP地址和端口。使用Ingress可以将流量路由到不同的Service中，而无需为每个Service分配独立的IP地址和端口。
- 多个域名共享同一个IP地址和端口。使用Ingress可以根据主机名将流量路由到不同的Service中，而无需为每个域名分配独立的IP地址和端口。
- 路径基础的流量路由。使用Ingress可以根据HTTP请求的路径将流量路由到不同的Service中。例如，您可以将所有以/api开头的请求路由到一个Service中，将所有以/web开头的请求路由到另一个Service中。
- SSL终止。使用Ingress可以将SSL终止放在Ingress控制器中，而不是在每个Service中。这样可以简化SSL证书的管理，并提高应用程序的安全性。

## yaml

```sh
export out="--dry-run=client -o yaml"
kubectl create ing ngx-ing --rule="ngx.test/=ngx-svc:80" --class=ngx-ink $out
```

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: ngx-ing
  
spec:
  ingressClassName: ngx-ink # 命令行参数
  
  rules: # 规则
  - host: ngx.test
    http:
      paths:
      - path: /
        pathType: Exact
        backend:
          service:
            name: ngx-svc
            port:
              number: 80
```