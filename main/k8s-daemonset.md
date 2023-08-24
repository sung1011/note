# k8s daemonSet

部署到可用的node上, 并确保每个node有且只有一个pod运行

## 场景

- `部署监控代理` 如: prometheus 必须每个节点都有一个 Pod 用来监控节点的状态，实时上报信息。
- `部署日志收集代理` 如: fluentd 必须在每个节点上运行一个 Pod，才能够搜集容器运行时产生的日志数据。
- `部署网络代理` 如: kube-proxy 必须每个节点都运行一个 Pod，否则节点就无法加入 Kubernetes 网络
- `安全应用` 每个节点都运行一个 Pod，用来监控节点的安全状态，如：检查恶意程序 异常登录 漏洞扫描等。

## yaml

```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: myapp
spec:
  selector:
    matchLabels:
      app: myapp # 通过label选择器, 选择属于myapp的pod
  template:
    metadata:
      labels:
        app: myapp
    spec:
      containers:
      - name: myapp
        image: redis:5.0.4
        resources:
          limits:
            memory: "128Mi"
            cpu: "500m"
        ports:
        - containerPort: 6379
```

> 非常类似于`deployment`, 只是没有replicas属性, (replicas变成了1)

## usage

```sh
kubectl apply -f ds.yaml
```

> `DaemonSet`只能控制自己创建的pod, 不能控制手动创建的pod

> 需要更精细化控制, 可以用到 `污点taint & 容忍toleration & 静态pod`

## 精细化控制pod

### 污点taint

拒绝pod调度到指定节点上运行

```sh
# 污点taint
kubectl taint nodes node1 key=value:NoSchedule # 给node1打上污点
kubectl describe node master # 查看master节点的污点 `node-role.kubernetes.io/master:NoSchedule`即拒绝pod调度到本node上运行, 因此master节点上不会运行其他pod
kubectl taint node master node-role.kubernetes.io/master:NoSchedule- # 删除master的污点, 允许pod调度到master节点上运行; 缺点是可能导致过多pod灌入master
kubectl taint node master node-role.kubernetes.io/master:NoSchedule # 给master打上污点, 拒绝pod调度到master节点上运行
```

> 类似label, 但是label是pod主动选择, 而污点是node主动选择

### 容忍toleration

允许pod调度到指定污点节点上运行

```yaml
# 容忍toleration (在daemonSet 的 yaml中设置)
tolerations:
- key: node-role.kubernetes.io/master
  effect: NoSchedule
  operator: Exists
```

> 类似label, 但是label是pod主动选择, 而容忍度是node主动选择

### 静态pod

- 静态pod是由kubelet管理的pod, 不是由apiserver管理的pod
- 静态pod的yaml默认存放在node的`etc/kubernetes/manifests`
- k8s的4个核心组件 `kube-apiserver` `kube-controller-manager` `kube-scheduler` `kubelet` 都是静态pod; 这也是它们能够先于k8s集群启动的原因


