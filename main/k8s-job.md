# k8s job

执行一次就退出的程序

## yaml

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: echo-job
spec:
  template:
    spec:
      restartPolicy: OnFailure
      containers:
      - image: busybox
        name: echo-job
        imagePullPolicy: IfNotPresent
        command: ["/bin/echo"]
        args: ["hello", "world"]
```

## usage

```sh
# job
kubectl create job echo-job --image=busybox --dry-run=client -o yaml > echo-job.yaml
kubectl apply -f echo-job.yaml
```