# k8s deployment

部署app到可用的node上, 并确保pod的副本数符合配置, 任意扩容缩容

## yaml

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
spec:
  replicas: 3 # 副本数; 可通过scale命令快捷修改, 立刻生效; kubectl scale deploy myapp --replicas=5
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

## usage

```sh
kubectl create deploy myapp --image=redis:5.0.4 --replicas=3 --dry-run=client -o yaml > deploy.yaml
kubectl apply -f deploy.yaml

kubectl get pod -l app=myapp
```

> `deployment`只能控制自己创建的pod, 不能控制手动创建的pod