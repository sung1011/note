# k8s config

## ConfigMap

明文配置, 可任务查看修改. 如: 服务端口, 运行参数, 文件路径

### yaml

```sh
# create
kubectl create cm info --from-literal=k=v --dry-run=client -o yaml > info.yaml
kubectl apply -f info.yaml
# get
kubectl get cm
kubectl get cm info -o yaml
kubectl describe cm info
# update
kubectl apply -f info.yaml # 修改后重新apply
# delete
kubectl delete cm info
```

```yaml
# configMap
apiVersion: v1
kind: ConfigMap
metadata:
  name: info
data:
  count: '10'
  debug: 'on'
  path: '/etc/systemd'
  greeting: |
    say hello to kubernetes.

# pod
apiVersion: v1
kind: Pod
metadata:
  name: test1
spec:
    containers:
        - name: test1
        image: busybox
        command: ["/bin/sh"]
        args: ["-c", "echo $(count) $(debug) $(path) $(greeting)"]
        envFrom: # 从configMap中获取全部配置
            - configMapRef: # secret为secretKeyRef
                name: info
        - env # 从不同的configMap中获取配置
            - name: count
              valueFrom:
                configMapKeyRef:
                  name: info
                  key: count
            - name: debug
              valueFrom:
                configMapKeyRef:
                  name: info
                  key: debug
            - name: path
              valueFrom:
                configMapKeyRef:
                  name: info
                  key: path
            - name: greeting
              valueFrom:
                configMapKeyRef:
                  name: info
                  key: greeting
```

> pod需要重启, 才能获取到configMap的新配置

### e.g

- 用configMap配置redis <https://kubernetes.io/zh-cn/docs/tutorials/configuration/configure-redis-using-configmap/>

## Secret

密文配置, 涉及敏感信息, 如: 密码, 秘钥, 证书

### yaml

```sh
# create
kubectl create secret generic user --from-literal=name=root --dry-run=client -o yaml > test2.yaml
kubectl apply -f test2.yaml
```

```yaml
# secret
apiVersion: v1
kind: Secret
metadata:
  name: user
data:
  name: cm9vdA==  # root
  pwd: MTIzNDU2   # 123456
  db: bXlzcWw=    # mysql

# pod 同configMap, 只是敏感配置将`configMapKeyRef`换成`secretKeyRef`
apiVersion: v1
kind: Pod
metadata:
  name: env-pod
spec:
  containers:
  - env:
      - name: COUNT
        valueFrom:
          configMapKeyRef:
            name: info
            key: count
      - name: GREETING
        valueFrom:
          configMapKeyRef:
            name: info
            key: greeting
      - name: USERNAME
        valueFrom:
          secretKeyRef: # secretKeyRef
            name: user
            key: name
      - name: PASSWORD
        valueFrom:
          secretKeyRef:
            name: user
            key: pwd
    image: busybox
    name: busy
    imagePullPolicy: IfNotPresent
    command: ["/bin/sleep", "300"]
```

> secret类型的data值必须为base64后的数据 `echo -n 'root' | base64`

### 预定义类型

| 内置类型                            | 用法                                   |
| ----------------------------------- | -------------------------------------- |
| Opaque                              | 户定义的任意数据                       |
| kubernetes.io/service-account-token | 服务账号令牌                           |
| kubernetes.io/dockercfg             | ~/.dockercfg 文件的序列化形式          |
| kubernetes.io/dockerconfigjson      | ~/.docker/config.json 文件的序列化形式 |
| kubernetes.io/basic-auth            | 于基本身份认证的凭据                   |
| kubernetes.io/ssh-auth              | 于 SSH 身份认证的凭据                  |
| kubernetes.io/tls                   | 于 TLS 客户端或者服务器端的数据        |
| bootstrap.kubernetes.io/token       | 启动引导令牌数据                       |

## Volume方式使用ConfigMap和Secret

为pod挂在多个Volume, 里面存放供应用使用的配置文件

### yaml

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: vol-pod
spec:
  volumes: # pod级别, 体现了pod共享存储
  - name: cm-vol
    configMap:
      name: info
  - name: sec-vol
    secret:
      secretName: user
  containers:
  - volumeMounts:
    - mountPath: /tmp/cm-items
      name: cm-vol
    - mountPath: /tmp/sec-items
      name: sec-vol
    image: busybox
    name: busy
    imagePullPolicy: IfNotPresent
    command: ["/bin/sleep", "300"]
```