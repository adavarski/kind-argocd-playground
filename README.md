## KinD: ArgoCD, Grafana, Prometheus, Loki, Tempo, Phlare and VictoriaMetrics.

- `Loki` is a horizontally-scalable, highly-available, multi-tenant log aggregation system inspired by Prometheus. It is designed to be very cost effective and easy to operate, as it does not index the contents of the logs, but rather a set of labels for each log stream.

- `Grafana Tempo` is an open source, easy-to-use, and high-scale distributed tracing backend. Tempo is cost-efficient, requiring only object storage to operate, and is deeply integrated with Grafana, Prometheus, and Loki. Tempo can ingest common open source tracing protocols, including Jaeger, Zipkin, and OpenTelemetry.

- `Grafana Phlare` lets you aggregate continuous profiling data with high availability, multi-tenancy, and durable storage. This helps you get a better understanding of resource usage in your applications down to the line number.

- `VictoriaMetrics` is a fast, cost-effective and scalable monitoring solution and time series database.

- `Prometheus`

- `Grafana` 

## Requirements

- Linux OS
- [Docker](https://docs.docker.com/)
- [KinD](https://kind.sigs.k8s.io/)
- [kubectl](https://kubernetes.io/docs/reference/kubectl/)
- [helm](https://helm.sh/docs/intro/install/)
- [yq](https://github.com/mikefarah/yq)
- [argocd CLI](https://argo-cd.readthedocs.io/en/stable/cli_installation/)


### Usage:
```
make launch-k8s
make deploy-argocd

kubectl port-forward service/argocd-server -n argocd 8080:443 &
kubectl get secret -n argocd argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
argocd login localhost:8080 --grpc-web --insecure --username admin --password $(kubectl get secret -n argocd argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d)

kubectl get secrets -n prometheus prometheus-grafana -o jsonpath="{.data.admin-password}" | base64 -d
kubectl port-forward svc/prometheus-grafana 3000:80 -n prometheus

```
### Browser (ArgoCD) : https://localhost:8080 -> Sync apps in below order, based on Argo sync-wave annotation via ArgoCD UI! (Todo: fix `make sync-applications`)
```
### Sort applications by Argo sync-wave annotation 
$ kustomize build ./manifests/applications/ | yq ea [.] -o json | jq -r '. | sort_by(.metadata.annotations."argocd.argoproj.io/sync-wave" // "0" | tonumber) | .[] | .metadata.name'
namespaces
cert-manager
loki
phlare
prometheus-adapter
tempo
victoriametrics
prometheus
monitoring
sandbox

###Using argocd CLI example
kustomize build ./manifests/applications/ | yq ea [.] -o json | jq -r '. | sort_by(.metadata.annotations."argocd.argoproj.io/sync-wave" // "0" | tonumber) | .[] | .metadata.name' > apps-sync.sort
for app in `cat apps-sync.sort`; do argocd app sync $app --retry-limit 3 --timeout 300; done

```
### Check apps
```
$ kubectl get po --all-namespaces
NAMESPACE            NAME                                                         READY   STATUS    RESTARTS        AGE
argocd               argocd-application-controller-0                              1/1     Running   0               27m
argocd               argocd-applicationset-controller-6477f4dc9-7lvpz             1/1     Running   0               27m
argocd               argocd-dex-server-587855cf49-ndqpl                           1/1     Running   0               27m
argocd               argocd-notifications-controller-5f88985887-zpq6w             1/1     Running   0               27m
argocd               argocd-redis-59687468f9-6qpg4                                1/1     Running   0               30m
argocd               argocd-repo-server-6594ddf4f4-tws8x                          1/1     Running   0               27m
argocd               argocd-server-7f9cd56796-m68cc                               1/1     Running   0               27m
cert-manager         cert-manager-55b858df44-4xzvw                                1/1     Running   0               22m
cert-manager         cert-manager-cainjector-7f47598f9b-g24wc                     1/1     Running   0               22m
cert-manager         cert-manager-webhook-7d694cd764-hmh45                        1/1     Running   0               22m
kube-system          coredns-787d4945fb-d74hw                                     1/1     Running   0               34m
kube-system          coredns-787d4945fb-jsb68                                     1/1     Running   0               34m
kube-system          etcd-hands-on-control-plane                                  1/1     Running   0               34m
kube-system          kindnet-8rz6f                                                1/1     Running   0               34m
kube-system          kindnet-9kd9z                                                1/1     Running   0               33m
kube-system          kindnet-jxfgm                                                1/1     Running   0               33m
kube-system          kindnet-pw2dw                                                1/1     Running   0               33m
kube-system          kube-apiserver-hands-on-control-plane                        1/1     Running   0               34m
kube-system          kube-controller-manager-hands-on-control-plane               1/1     Running   0               34m
kube-system          kube-proxy-5k7zn                                             1/1     Running   0               33m
kube-system          kube-proxy-dcfkp                                             1/1     Running   0               33m
kube-system          kube-proxy-g7r9z                                             1/1     Running   0               33m
kube-system          kube-proxy-tlpmq                                             1/1     Running   0               34m
kube-system          kube-scheduler-hands-on-control-plane                        1/1     Running   0               34m
local-path-storage   local-path-provisioner-c8855d4bb-6c6cw                       1/1     Running   0               34m
loki                 loki-0                                                       1/1     Running   0               20m
loki                 loki-promtail-dplmz                                          1/1     Running   0               20m
loki                 loki-promtail-kw98s                                          1/1     Running   0               20m
loki                 loki-promtail-rsn7m                                          1/1     Running   0               20m
loki                 loki-promtail-x4tsf                                          1/1     Running   0               20m
phlare               phlare-0                                                     1/1     Running   0               18m
prometheus-adapter   prometheus-adapter-7c6bbdd68b-ljldh                          1/1     Running   0               17m
prometheus           alertmanager-prometheus-kube-prometheus-alertmanager-0       2/2     Running   0               13m
prometheus           prometheus-grafana-66f47cb6fc-t2sq4                          3/3     Running   0               15m
prometheus           prometheus-kube-prometheus-operator-68b694d86f-pztcs         1/1     Running   0               15m
prometheus           prometheus-kube-state-metrics-cdf984bd9-j2lpv                1/1     Running   0               15m
prometheus           prometheus-prometheus-kube-prometheus-prometheus-0           2/2     Running   0               10m
prometheus           prometheus-prometheus-node-exporter-5bmzh                    1/1     Running   0               15m
prometheus           prometheus-prometheus-node-exporter-db664                    1/1     Running   0               15m
prometheus           prometheus-prometheus-node-exporter-kvncl                    1/1     Running   0               15m
prometheus           prometheus-prometheus-node-exporter-xlsdf                    1/1     Running   0               15m
prometheus           promlens-69596fbb57-n8b8b                                    1/1     Running   0               11m
sandbox              dummy-metrics-d994b565d-hzg2h                                1/1     Running   0               11s
sandbox              request                                                      1/1     Running   0               8m26s
sandbox              todo-795757947b-lvmnf                                        1/1     Running   0               8m25s
tempo                tempo-0                                                      1/1     Running   0               16m
victoriametrics      victoriametrics-victoria-metrics-operator-786cbbd895-wzqmf   1/1     Running   0               16m
victoriametrics      vmagent-vmagent-5d4bc68b54-nrlxn                             2/2     Running   0               11m
victoriametrics      vmsingle-database-6d4bbfffc4-72kcm                           1/1     Running   0               11m


$ kubectl get svc --all-namespaces
NAMESPACE            NAME                                                 TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                                                                                                   AGE
argocd               argocd-application-controller-metrics                ClusterIP   10.96.243.43    <none>        8082/TCP                                                                                                  28m
argocd               argocd-applicationset-controller                     ClusterIP   10.96.200.187   <none>        7000/TCP                                                                                                  30m
argocd               argocd-dex-server                                    ClusterIP   10.96.171.200   <none>        5556/TCP,5557/TCP                                                                                         30m
argocd               argocd-redis                                         ClusterIP   10.96.176.242   <none>        6379/TCP                                                                                                  30m
argocd               argocd-repo-server                                   ClusterIP   10.96.182.245   <none>        8081/TCP                                                                                                  30m
argocd               argocd-repo-server-metrics                           ClusterIP   10.96.177.251   <none>        8084/TCP                                                                                                  28m
argocd               argocd-server                                        ClusterIP   10.96.235.0     <none>        80/TCP,443/TCP                                                                                            30m
argocd               argocd-server-metrics                                ClusterIP   10.96.137.59    <none>        8083/TCP                                                                                                  28m
cert-manager         cert-manager                                         ClusterIP   10.96.90.32     <none>        9402/TCP                                                                                                  22m
cert-manager         cert-manager-webhook                                 ClusterIP   10.96.168.207   <none>        443/TCP                                                                                                   22m
default              kubernetes                                           ClusterIP   10.96.0.1       <none>        443/TCP                                                                                                   34m
kube-system          kube-dns                                             ClusterIP   10.96.0.10      <none>        53/UDP,53/TCP,9153/TCP                                                                                    34m
kube-system          prometheus-kube-prometheus-coredns                   ClusterIP   None            <none>        9153/TCP                                                                                                  15m
kube-system          prometheus-kube-prometheus-kube-controller-manager   ClusterIP   None            <none>        10257/TCP                                                                                                 15m
kube-system          prometheus-kube-prometheus-kube-etcd                 ClusterIP   None            <none>        2381/TCP                                                                                                  15m
kube-system          prometheus-kube-prometheus-kube-proxy                ClusterIP   None            <none>        10249/TCP                                                                                                 15m
kube-system          prometheus-kube-prometheus-kube-scheduler            ClusterIP   None            <none>        10259/TCP                                                                                                 15m
kube-system          prometheus-kube-prometheus-kubelet                   ClusterIP   None            <none>        10250/TCP,10255/TCP,4194/TCP                                                                              14m
loki                 loki                                                 ClusterIP   10.96.97.176    <none>        3100/TCP                                                                                                  21m
loki                 loki-headless                                        ClusterIP   None            <none>        3100/TCP                                                                                                  21m
loki                 loki-memberlist                                      ClusterIP   None            <none>        7946/TCP                                                                                                  21m
phlare               phlare                                               ClusterIP   10.96.217.164   <none>        4100/TCP                                                                                                  18m
phlare               phlare-headless                                      ClusterIP   None            <none>        4100/TCP                                                                                                  18m
phlare               phlare-memberlist                                    ClusterIP   None            <none>        7946/TCP                                                                                                  18m
prometheus-adapter   prometheus-adapter                                   ClusterIP   10.96.200.242   <none>        443/TCP                                                                                                   17m
prometheus           alertmanager-operated                                ClusterIP   None            <none>        9093/TCP,9094/TCP,9094/UDP                                                                                14m
prometheus           prometheus-grafana                                   ClusterIP   10.96.4.160     <none>        80/TCP                                                                                                    15m
prometheus           prometheus-kube-prometheus-alertmanager              ClusterIP   10.96.133.168   <none>        9093/TCP,8080/TCP                                                                                         15m
prometheus           prometheus-kube-prometheus-operator                  ClusterIP   10.96.3.78      <none>        443/TCP                                                                                                   15m
prometheus           prometheus-kube-prometheus-prometheus                ClusterIP   10.96.205.35    <none>        9090/TCP,8080/TCP                                                                                         15m
prometheus           prometheus-kube-state-metrics                        ClusterIP   10.96.73.200    <none>        8080/TCP                                                                                                  15m
prometheus           prometheus-operated                                  ClusterIP   None            <none>        9090/TCP                                                                                                  10m
prometheus           prometheus-prometheus-node-exporter                  ClusterIP   10.96.64.146    <none>        9100/TCP                                                                                                  15m
sandbox              todo                                                 ClusterIP   10.96.202.85    <none>        80/TCP                                                                                                    8m43s
tempo                tempo                                                ClusterIP   10.96.0.2       <none>        3100/TCP,6831/UDP,6832/UDP,14268/TCP,14250/TCP,9411/TCP,55680/TCP,55681/TCP,4317/TCP,4318/TCP,55678/TCP   17m
victoriametrics      victoriametrics-victoria-metrics-operator            ClusterIP   10.96.141.235   <none>        8080/TCP,443/TCP                                                                                          16m
victoriametrics      vmagent-vmagent                                      ClusterIP   10.96.157.233   <none>        8429/TCP                                                                                                  12m
victoriametrics      vmsingle-database                                    ClusterIP   10.96.178.211   <none>        8429/TCP                                                                                                  12m


```
### Screenshots:

### Browser (ArgoCD) : https://localhost:8080

<img src="pictures/ArgoCD-applications.png?raw=true" width="1000">

### Browser (Grafana): http://localhost:3000 

<img src="pictures/Grafana-DataSources.png?raw=true" width="1000">

<img src="pictures/Grafana-UI.png?raw=true" width="1000">

<img src="pictures/Grafana-UI-ArgoCD-performance.png?raw=true" width="1000">


### Clean environment
```
make shutdown-k8s
```

Credits: https://github.com/zoetrope/k8s-hands-on
