# k8s service

## 创建service

    expose
    yaml
    服务暴露多端口
    命名端口

## 服务发现

    Env
    DNS

- K8s集群会内置一个dns服务器, service创建成功后, 会在dns服务器里导入一些记录, 想要访问某个服务, 通过dns服务器解析出对应的ip和port, 从而实现服务访问

## endpoint

    endpoint是k8s集群中的一个资源对象, 存储在etcd中, 用来记录一个service对应的所有pod的访问地址.  

- service配置selector: endpoint controller会自动创建对应的endpoint对象.

- service不配置selector: endpoint controller不会生成endpoint对象. 可手动创建endpoint(name必须与svc name相同)

      endpoint controller: 是k8s集群控制器的其中一个组件, 其功能如下:

- 负责生成和维护所有endpoint对象的控制器
- 负责监听service和对应pod的变化
- 监听到service被删除, 则删除和该service同名的endpoint对象
- 监听到新的service被创建, 则根据新建service信息获取相关pod列表, 然后创建对应endpoint对象
- 监听到service被更新, 则根据更新后的service信息获取相关pod列表, 然后更新对应endpoint对象
- 监听到pod事件, 则更新对应的service的endpoint对象, 将podIp记录到endpoint中

## 暴露服务 - ClusterIp

    默认模式, 只能在集群内部访问

## 暴露服务 - NodePort

    通过每个 Node 上的 IP 和静态端口(NodePort)暴露服务.NodePort 服务会路由到 ClusterIP 服务, 这个 ClusterIP 服务会自动创建.通过请求 <  NodeIP >:< NodePort >, 可以从集群的外部访问一个 NodePort 服务.

## 暴露服务 - LoadBalancer

    使用云提供商的负载局衡器, 可以向外部暴露服务.外部的负载均衡器可以路由到 NodePort 服务和 ClusterIP 服务.

## 暴露服务 - ExternalName

    通过返回 CNAME 和它的值, 可以将服务映射到 externalName 字段的内容(例如,  foo.bar.example.com). 没有任何类型代理被创建, 这只有 Kubernetes 1.7 或更高版本的 kube-dns 才支持.

## Ingress路由到服务

## 就绪探针 ReadinessProbe

## headless发现独立pod

## 排除服务故障
