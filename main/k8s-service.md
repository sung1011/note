# k8s service

负载均衡, 服务发现, 服务暴露, 四层转发

## 场景

pod会被deployment或daemonSet控制器管理, 但是pod的ip是动态变化的, 无法直接访问, 需要通过service来访问

## 架构

```js
                apiserver
                    |
--------------------------------------------------------
| `Node`            |                                
|                   kube-proxy
|                   |
|                   clusterIP (iptables)                  
|                 | | |
--------------------------------------------------------
                 |  |  |
--------------------------------------------------------
| `Pod` | `Pod` | `Pod` | `Pod` | `Pod` | `Pod` | `Pod` |
--------------------------------------------------------
```

## yaml

```sh
export out="--dry-run=client -o yaml"
kubectl expose deploy ngx-dep --port=80 --target-port=80 $out

kubectl describe svc ngx-svc
kubectl get pod -o wide
```

```yaml
apiVersion: v1
kind: Service
metadata:
  name: ngx-svc
  namespace: my-namespace # 指定命名空间
  
spec:
  selector:
    app: ngx-dep
    
  ports:
  - port: 80 # 端口
    targetPort: 80 # 容器端口
    protocol: TCP
```

> 先创建deployment, 再创建service, service会自动关联deployment (pod, daemonSet同理)

> kubectl export 而非 create, 可能表达的意思是: 通过service暴露的是已经存在的pod, 而非创建新的pod

## 域名 (DNS)

- 域名规则
  - Pod的DNS名称: `<pod-ip-address>.<pod-name>.<namespace>.pod.cluster.local`
  - Service的DNS名称: `<service-name>.<namespace>.svc.cluster.local`
  - Service的端口DNS名称: `<port-name>.<service-name>.<namespace>.svc.cluster.local`

## namespace 名字空间

```sh
kubectl create namespace my-namespace
kubectl get ns
```

> 默认namespace=default

> 一般是在yaml文件中指定namespace, 也可以在命令行中指定

## 对外暴露服务

### ClusterIp

默认模式, 只能在集群内部访问

### NodePort

在每个node生成固定的IP和端口, 任意请求其中一个, kube-proxy会路由到真正的node上, 再由 ClusterIP 转发到pod

> 缺点: 端口数有限(30000~32767); 暴露的端口是静态的, 不能动态分配; 转发浪费流量

### LoadBalancer

使用云提供商的负载局衡器, 可以向外部暴露服务.外部的负载均衡器可以路由到 NodePort 服务和 ClusterIP 服务.

### ExternalName

通过返回 CNAME 和它的值, 可以将服务映射到 externalName 字段的内容(例如,  foo.bar.example.com). 没有任何类型代理被创建, 这只有 Kubernetes 1.7 或更高版本的 kube-dns 才支持.
