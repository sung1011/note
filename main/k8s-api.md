# k8s API

## subcommands

### create

Create a resource from a file or from stdin

```sh
kubectl create -f <k8s-yaml-file>
```

### apply

Apply a configuration to a resource by file name or stdin
更新资源, 如果不存在则创建, 也可以更新部分资源配置

```sh
kubectl apply -f <k8s yaml file>
kubectl apply -f deployment.yaml --patch '{"spec": {"replicas": 3}}' # 更新副本数
kubectl apply -f service.yaml --patch '{"spec": {"ports": [{"name": "http", "port": 8080, "targetPort": 80}]}}' # 更新服务的端口
```

### get

Display one or many resources

```sh
kubectl get <any>
kubectl get all
kubectl get pod -w # 实时获取pods状态
```

### logs

Print the logs for a container in a pod

```sh
kubectl logs <pod name>
kubectl logs <pod name> -c <container name>
```

### label

Update the labels on a resource

```sh
kubectl get po -l app=ui,rel=beta get
kubectl label node <node name> gpu=true set
```

### annotate

Update the annotations on a resource

```sh
kubectl annotate pod <pod name> foo=bar
```

### delete

Delete resources by file names, stdin, resources and names, or by resources and label selector

```sh
kubectl delete po <pod name>
kubectl delete po -l <label key=val>
kubectl delete po --all
kubectl delete all --all
```

### describe

Show details of a specific resource or group of resources

```sh
kubectl describe pod <pod name> # 显示特定Pod的详细信息
kubectl describe service <service name> # 显示特定服务的详细信息
kubectl describe deployment <deployment name> # 显示特定部署的详细信息
```

### exec

Execute a command in a container

```sh
kubectl exec <pod name> -- curl -s http://10.111.249.153
kubectl exec -it ngx-pod -- sh
```

### debug

Create debugging sessions for troubleshooting workloads and nodes

```sh
kubectl get po <pod name> -o yaml 获取pod yaml
kubectl describe <any>
kubectl edit <any>
kubectl port-forward <pod name> <local port>:<pod port> 本地端口映射pod端口
```

### explain

Get documentation for a resource

```sh
kubectl explain pod
kubectl explain pod.metadata
kubectl explain pod.spec
```

## action

### replication-controller

```sh
kubectl scale rc <rc name> --replicas=10
kubectl delete rc <rc name> # 删除rc并且删除pod
kubectl delete rc <rc name> --cascade=false # 删除rc但不删除pod
```

### namespace

```sh
kubectl create ns <namespace name>
kubectl get po -n <namespace name>
kubectl delete ns <namespace name>
```

## syntax

```sh
kubectl [command] [TYPE] [NAME] [flags]
```

## Operations

| Operation      | Syntax                                                                                                                                    | Description                                                                                            |
| -------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------ |
| annotate       | kubectl annotate (-f FILENAME \                                                                                                           | TYPE NAME \                                                                                            | TYPE/NAME) KEY_1=VAL_1 ... KEY_N=VAL_N [--overwrite] [--all] [--resource-version=version] [flags]  | Add or update the annotations of one or more resources.                                                        |
| api-versions   | kubectl api-versions [flags]                                                                                                              | List the API versions that are available.                                                              |
| apply          | kubectl apply -f FILENAME [flags]                                                                                                         | Apply a configuration change to a resource from a file or stdin.                                       |
| attach         | kubectl attach POD -c CONTAINER [-i] [-t] [flags]                                                                                         | Attach to a running container either to view the output stream or interact with the container (stdin). |
| autoscale      | kubectl autoscale (-f FILENAME \                                                                                                          | TYPE NAME \                                                                                            | TYPE/NAME) [--min=MINPODS] --max=MAXPODS [--cpu-percent=CPU] [flags]                               | Automatically scale the set of pods that are managed by a replication controller.                              |
| cluster-info   | kubectl cluster-info [flags]                                                                                                              | Display endpoint information about the master and services in the cluster.                             |
| config         | kubectl config SUBCOMMAND [flags]                                                                                                         | Modifies kubeconfig files. See the individual subcommands for details.                                 |
| create         | kubectl create -f FILENAME [flags]                                                                                                        | Create one or more resources from a file or stdin.                                                     |
| delete         | kubectl delete (-f FILENAME \                                                                                                             | TYPE [NAME \                                                                                           | /NAME \                                                                                            | -l label \                                                                                                     | --all]) [flags]                                                               | Delete resources either from a file, stdin, or specifying label selectors, names, resource selectors, or resources. |
| describe       | kubectl describe (-f FILENAME \                                                                                                           | TYPE [NAME_PREFIX \                                                                                    | /NAME \                                                                                            | -l label]) [flags]                                                                                             | Display the detailed state of one or more resources.                          |
| diff           | kubectl diff -f FILENAME [flags]                                                                                                          | Diff file or stdin against live configuration (BETA)                                                   |
| edit           | kubectl edit (-f FILENAME \                                                                                                               | TYPE NAME \                                                                                            | TYPE/NAME) [flags]                                                                                 | Edit and update the definition of one or more resources on the server by using the default editor.             |
| exec           | kubectl exec POD [-c CONTAINER] [-i] [-t] [flags] [-- COMMAND [args...]]                                                                  | Execute a command against a container in a pod.                                                        |
| explain        | kubectl explain [--recursive=false] [flags]                                                                                               | Get documentation of various resources. For instance pods, nodes, services, etc.                       |
| expose         | kubectl expose (-f FILENAME \                                                                                                             | TYPE NAME \                                                                                            | TYPE/NAME) [--port=port] [--protocol=TCP\                                                          | UDP] [--target-port=number-or-name] [--name=name] [--external-ip=external-ip-of-service] [--type=type] [flags] | Expose a replication controller, service, or pod as a new Kubernetes service. |
| get            | kubectl get (-f FILENAME \                                                                                                                | TYPE [NAME \                                                                                           | /NAME \                                                                                            | -l label]) [--watch] [--sort-by=FIELD] [[-o \                                                                  | --output]=OUTPUT_FORMAT] [flags]                                              | List one or more resources.                                                                                         |
| label          | kubectl label (-f FILENAME \                                                                                                              | TYPE NAME \                                                                                            | TYPE/NAME) KEY_1=VAL_1 ... KEY_N=VAL_N [--overwrite] [--all] [--resource-version=version] [flags]  | Add or update the labels of one or more resources.                                                             |
| logs           | kubectl logs POD [-c CONTAINER] [--follow] [flags]                                                                                        | Print the logs for a container in a pod.                                                               |
| patch          | kubectl patch (-f FILENAME \                                                                                                              | TYPE NAME \                                                                                            | TYPE/NAME) --patch PATCH [flags]                                                                   | Update one or more fields of a resource by using the strategic merge patch process.                            |
| port-forward   | kubectl port-forward POD [LOCAL_PORT:]REMOTE_PORT [...[LOCAL_PORT_N:]REMOTE_PORT_N] [flags]                                               | Forward one or more local ports to a pod.                                                              |
| proxy          | kubectl proxy [--port=PORT] [--www=static-dir] [--www-prefix=prefix] [--api-prefix=prefix] [flags]                                        | Run a proxy to the Kubernetes API server.                                                              |
| replace        | kubectl replace -f FILENAME                                                                                                               | Replace a resource from a file or stdin.                                                               |
| rolling-update | kubectl rolling-update OLD_CONTROLLER_NAME ([NEW_CONTROLLER_NAME] --image=NEW_CONTAINER_IMAGE \                                           | -f NEW_CONTROLLER_SPEC) [flags]                                                                        | Perform a rolling update by gradually replacing the specified replication controller and its pods. |
| run            | kubectl run NAME --image=image [--env="key=value"] [--port=port] [--replicas=replicas] [--dry-run=bool] [--overrides=inline-json] [flags] | Run a specified image on the cluster.                                                                  |
| scale          | kubectl scale (-f FILENAME \                                                                                                              | TYPE NAME \                                                                                            | TYPE/NAME) --replicas=COUNT [--resource-version=version] [--current-replicas=count] [flags]        | Update the size of the specified replication controller.                                                       |
| stop           | kubectl stop                                                                                                                              | Deprecated: Instead, see kubectl delete.                                                               |
| version        | kubectl version [--client] [flags]                                                                                                        | Display the Kubernetes version running on the client and server.                                       |                                                                                                    |

## Resource types

| Resource Name                   | Short Names | API Group                    | Namespaced | Resource Kind                  |
| ------------------------------- | ----------- | ---------------------------- | ---------- | ------------------------------ |
| componentstatuses               | cs          |                              | false      | ComponentStatus                |
| configmaps                      | cm          |                              | true       | ConfigMap                      |
| endpoints                       | ep          |                              | true       | Endpoints                      |
| limitranges                     | limits      |                              | true       | LimitRange                     |
| namespaces                      | ns          |                              | false      | Namespace                      |
| nodes                           | no          |                              | false      | Node                           |
| persistentvolumeclaims          | pvc         |                              | true       | PersistentVolumeClaim          |
| persistentvolumes               | pv          |                              | false      | PersistentVolume               |
| pods                            | po          |                              | true       | Pod                            |
| podtemplates                    |             |                              | true       | PodTemplate                    |
| replicationcontrollers          | rc          |                              | true       | ReplicationController          |
| resourcequotas                  | quota       |                              | true       | ResourceQuota                  |
| secrets                         |             |                              | true       | Secret                         |
| serviceaccounts                 | sa          |                              | true       | ServiceAccount                 |
| services                        | svc         |                              | true       | Service                        |
| mutatingwebhookconfigurations   |             | admissionregistration.k8s.io | false      | MutatingWebhookConfiguration   |
| validatingwebhookconfigurations |             | admissionregistration.k8s.io | false      | ValidatingWebhookConfiguration |
| customresourcedefinitions       | crd, crds   | apiextensions.k8s.io         | false      | CustomResourceDefinition       |
| apiservices                     |             | apiregistration.k8s.io       | false      | APIService                     |
| controllerrevisions             |             | apps                         | true       | ControllerRevision             |
| daemonsets                      | ds          | apps                         | true       | DaemonSet                      |
| deployments                     | deploy      | apps                         | true       | Deployment                     |
| replicasets                     | rs          | apps                         | true       | ReplicaSet                     |
| statefulsets                    | sts         | apps                         | true       | StatefulSet                    |
| tokenreviews                    |             | authentication.k8s.io        | false      | TokenReview                    |
| localsubjectaccessreviews       |             | authorization.k8s.io         | true       | LocalSubjectAccessReview       |
| selfsubjectaccessreviews        |             | authorization.k8s.io         | false      | SelfSubjectAccessReview        |
| selfsubjectrulesreviews         |             | authorization.k8s.io         | false      | SelfSubjectRulesReview         |
| subjectaccessreviews            |             | authorization.k8s.io         | false      | SubjectAccessReview            |
| horizontalpodautoscalers        | hpa         | autoscaling                  | true       | HorizontalPodAutoscaler        |
| cronjobs                        | cj          | batch                        | true       | CronJob                        |
| jobs                            |             | batch                        | true       | Job                            |
| certificatesigningrequests      | csr         | certificates.k8s.io          | false      | CertificateSigningRequest      |
| leases                          |             | coordination.k8s.io          | true       | Lease                          |
| events                          | ev          | events.k8s.io                | true       | Event                          |
| ingresses                       | ing         | extensions                   | true       | Ingress                        |
| networkpolicies                 | netpol      | networking.k8s.io            | true       | NetworkPolicy                  |
| poddisruptionbudgets            | pdb         | policy                       | true       | PodDisruptionBudget            |
| podsecuritypolicies             | psp         | policy                       | false      | PodSecurityPolicy              |
| clusterrolebindings             |             | rbac.authorization.k8s.io    | false      | ClusterRoleBinding             |
| clusterroles                    |             | rbac.authorization.k8s.io    | false      | ClusterRole                    |
| rolebindings                    |             | rbac.authorization.k8s.io    | true       | RoleBinding                    |
| roles                           |             | rbac.authorization.k8s.io    | true       | Role                           |
| priorityclasses                 | pc          | scheduling.k8s.io            | false      | PriorityClass                  |
| storageclasses                  | sc          | storage.k8s.io               | false      | StorageClass                   |
| volumeattachments               |             | storage.k8s.io               | false      | VolumeAttachment               |
