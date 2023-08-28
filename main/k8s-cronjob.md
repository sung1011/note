# k8s cronjob

定期执行一次就退出的程序

## yaml

```sh
kubectl create cronjob echo-job --image=busybox --dry-run=client -o yaml > echo-job.yaml
kubectl apply -f echo-job.yaml
```

```yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: echo-cj
spec:
  schedule: '*/1 * * * *'
  jobTemplate:
    spec:
      template:
        spec:
          restartPolicy: OnFailure
          containers:
          - image: busybox
            name: echo-cj
            imagePullPolicy: IfNotPresent
            command: ["/bin/echo"]
            args: ["hello", "world"]
```