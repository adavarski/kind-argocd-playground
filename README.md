## KinD: Argo CD, Grafana, Prometheus, Loki, Tempo, Phlare and VictoriaMetrics.

```
make launch-k8s
make deploy-argocd
kubectl port-forward service/argocd-server -n argocd 8080:443 &
kubectl get secret -n argocd argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
argocd login localhost:8080 --grpc-web --insecure --username admin --password $(kubectl get secret -n argocd argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d)

$ kubectl get secrets -n prometheus prometheus-grafana -o jsonpath="{.data.admin-password}" | base64 -d
$ kubectl port-forward svc/prometheus-grafana 3000:80 -n prometheus

```
Browser: https://localhost:8080 -> Sync apps in this order via argocd UI
```
$ kustomize build ./manifests/applications/ | yq ea [.] -o json | jq -r '. | sort_by(.metadata.annotations."argocd.argoproj.io/sync-wave" // "0" | tonumber) | .[] | .metadata.name'
namespaces
cert-manager
loki
phlare
prometheus-adapter
tempo
victoriametrics
monitoring
prometheus
sandbox
```
Check
```
$ kubectl get po --all-namespaces|grep -v sand
NAMESPACE            NAME                                                         READY   STATUS             RESTARTS   AGE
argocd               argocd-application-controller-0                              1/1     Running            0          32m
argocd               argocd-applicationset-controller-6477f4dc9-d24w5             1/1     Running            0          32m
argocd               argocd-dex-server-587855cf49-zbzfm                           1/1     Running            0          32m
argocd               argocd-notifications-controller-5f88985887-4nqjr             1/1     Running            0          32m
argocd               argocd-redis-59687468f9-n9ndb                                1/1     Running            0          34m
argocd               argocd-repo-server-6594ddf4f4-2zh7r                          1/1     Running            0          32m
argocd               argocd-server-7f9cd56796-gwrgl                               1/1     Running            0          32m
cert-manager         cert-manager-55b858df44-w79c8                                1/1     Running            0          27m
cert-manager         cert-manager-cainjector-7f47598f9b-pjwf9                     1/1     Running            0          27m
cert-manager         cert-manager-webhook-7d694cd764-xnxfg                        1/1     Running            0          27m
kube-system          coredns-787d4945fb-jjk84                                     1/1     Running            0          41m
kube-system          coredns-787d4945fb-xf7t4                                     1/1     Running            0          41m
kube-system          etcd-hands-on-control-plane                                  1/1     Running            0          41m
kube-system          kindnet-5c57c                                                1/1     Running            0          41m
kube-system          kindnet-5fdl5                                                1/1     Running            0          40m
kube-system          kindnet-d6xpv                                                1/1     Running            0          40m
kube-system          kindnet-pw9sk                                                1/1     Running            0          40m
kube-system          kube-apiserver-hands-on-control-plane                        1/1     Running            0          41m
kube-system          kube-controller-manager-hands-on-control-plane               1/1     Running            0          41m
kube-system          kube-proxy-gx2kx                                             1/1     Running            0          40m
kube-system          kube-proxy-l8nhg                                             1/1     Running            0          40m
kube-system          kube-proxy-qrxkd                                             1/1     Running            0          41m
kube-system          kube-proxy-qstfh                                             1/1     Running            0          40m
kube-system          kube-scheduler-hands-on-control-plane                        1/1     Running            0          41m
local-path-storage   local-path-provisioner-c8855d4bb-wbk5l                       1/1     Running            0          41m
loki                 loki-0                                                       1/1     Running            0          26m
loki                 loki-promtail-cbm9g                                          1/1     Running            0          26m
loki                 loki-promtail-ddmr7                                          1/1     Running            0          26m
loki                 loki-promtail-mx7s7                                          1/1     Running            0          26m
loki                 loki-promtail-z249r                                          1/1     Running            0          26m
phlare               phlare-0                                                     1/1     Running            0          21m
prometheus-adapter   prometheus-adapter-7c6bbdd68b-5qzdt                          1/1     Running            0          14m
prometheus           alertmanager-prometheus-kube-prometheus-alertmanager-0       2/2     Running            0          9m43s
prometheus           prometheus-grafana-66f47cb6fc-hxjpp                          3/3     Running            0          10m
prometheus           prometheus-kube-prometheus-operator-68b694d86f-4gc9z         1/1     Running            0          10m
prometheus           prometheus-kube-state-metrics-cdf984bd9-zkzq5                1/1     Running            0          10m
prometheus           prometheus-prometheus-kube-prometheus-prometheus-0           2/2     Running            0          9m4s
prometheus           prometheus-prometheus-node-exporter-6tsjx                    1/1     Running            0          10m
prometheus           prometheus-prometheus-node-exporter-b6spq                    1/1     Running            0          10m
prometheus           prometheus-prometheus-node-exporter-hj9vq                    1/1     Running            0          10m
prometheus           prometheus-prometheus-node-exporter-w8cwh                    1/1     Running            0          10m
prometheus           promlens-69596fbb57-gdwhc                                    1/1     Running            0          7m20s
tempo                tempo-0                                                      1/1     Running            0          11m
victoriametrics      victoriametrics-victoria-metrics-operator-786cbbd895-vcgzn   1/1     Running            0          11m
victoriametrics      vmagent-vmagent-5d4bc68b54-bvms8                             2/2     Running            0          7m16s
victoriametrics      vmsingle-database-6d4bbfffc4-tnxwm                           1/1     Running            0          7m19s

$ kubectl get svc --all-namespaces|grep -v sand
NAMESPACE            NAME                                                 TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                                                                                                   AGE
argocd               argocd-application-controller-metrics                ClusterIP   10.96.213.107   <none>        8082/TCP                                                                                                  32m
argocd               argocd-applicationset-controller                     ClusterIP   10.96.47.210    <none>        7000/TCP                                                                                                  34m
argocd               argocd-dex-server                                    ClusterIP   10.96.223.85    <none>        5556/TCP,5557/TCP                                                                                         34m
argocd               argocd-redis                                         ClusterIP   10.96.75.148    <none>        6379/TCP                                                                                                  34m
argocd               argocd-repo-server                                   ClusterIP   10.96.119.233   <none>        8081/TCP                                                                                                  34m
argocd               argocd-repo-server-metrics                           ClusterIP   10.96.55.20     <none>        8084/TCP                                                                                                  32m
argocd               argocd-server                                        ClusterIP   10.96.151.31    <none>        80/TCP,443/TCP                                                                                            34m
argocd               argocd-server-metrics                                ClusterIP   10.96.221.171   <none>        8083/TCP                                                                                                  32m
cert-manager         cert-manager                                         ClusterIP   10.96.202.197   <none>        9402/TCP                                                                                                  28m
cert-manager         cert-manager-webhook                                 ClusterIP   10.96.60.32     <none>        443/TCP                                                                                                   28m
default              kubernetes                                           ClusterIP   10.96.0.1       <none>        443/TCP                                                                                                   41m
kube-system          kube-dns                                             ClusterIP   10.96.0.10      <none>        53/UDP,53/TCP,9153/TCP                                                                                    41m
kube-system          prometheus-kube-prometheus-coredns                   ClusterIP   None            <none>        9153/TCP                                                                                                  10m
kube-system          prometheus-kube-prometheus-kube-controller-manager   ClusterIP   None            <none>        10257/TCP                                                                                                 10m
kube-system          prometheus-kube-prometheus-kube-etcd                 ClusterIP   None            <none>        2381/TCP                                                                                                  10m
kube-system          prometheus-kube-prometheus-kube-proxy                ClusterIP   None            <none>        10249/TCP                                                                                                 10m
kube-system          prometheus-kube-prometheus-kube-scheduler            ClusterIP   None            <none>        10259/TCP                                                                                                 10m
kube-system          prometheus-kube-prometheus-kubelet                   ClusterIP   None            <none>        10250/TCP,10255/TCP,4194/TCP                                                                              9m52s
loki                 loki                                                 ClusterIP   10.96.199.38    <none>        3100/TCP                                                                                                  26m
loki                 loki-headless                                        ClusterIP   None            <none>        3100/TCP                                                                                                  26m
loki                 loki-memberlist                                      ClusterIP   None            <none>        7946/TCP                                                                                                  26m
phlare               phlare                                               ClusterIP   10.96.109.36    <none>        4100/TCP                                                                                                  21m
phlare               phlare-headless                                      ClusterIP   None            <none>        4100/TCP                                                                                                  21m
phlare               phlare-memberlist                                    ClusterIP   None            <none>        7946/TCP                                                                                                  21m
prometheus-adapter   prometheus-adapter                                   ClusterIP   10.96.198.161   <none>        443/TCP                                                                                                   14m
prometheus           alertmanager-operated                                ClusterIP   None            <none>        9093/TCP,9094/TCP,9094/UDP                                                                                9m52s
prometheus           prometheus-grafana                                   ClusterIP   10.96.153.250   <none>        80/TCP                                                                                                    10m
prometheus           prometheus-kube-prometheus-alertmanager              ClusterIP   10.96.242.125   <none>        9093/TCP,8080/TCP                                                                                         10m
prometheus           prometheus-kube-prometheus-operator                  ClusterIP   10.96.203.97    <none>        443/TCP                                                                                                   10m
prometheus           prometheus-kube-prometheus-prometheus                ClusterIP   10.96.106.134   <none>        9090/TCP,8080/TCP                                                                                         10m
prometheus           prometheus-kube-state-metrics                        ClusterIP   10.96.126.29    <none>        8080/TCP                                                                                                  10m
prometheus           prometheus-operated                                  ClusterIP   None            <none>        9090/TCP                                                                                                  9m13s
prometheus           prometheus-prometheus-node-exporter                  ClusterIP   10.96.136.159   <none>        9100/TCP                                                                                                  10m
tempo                tempo                                                ClusterIP   10.96.0.2       <none>        3100/TCP,6831/UDP,6832/UDP,14268/TCP,14250/TCP,9411/TCP,55680/TCP,55681/TCP,4317/TCP,4318/TCP,55678/TCP   12m
victoriametrics      victoriametrics-victoria-metrics-operator            ClusterIP   10.96.105.52    <none>        8080/TCP,443/TCP                                                                                          11m
victoriametrics      vmagent-vmagent                                      ClusterIP   10.96.205.86    <none>        8429/TCP                                                                                                  7m25s
victoriametrics      vmsingle-database                                    ClusterIP   10.96.163.98    <none>        8429/TCP                                                                                                  7m28s
davar@carbon:~/Documents/TRAININGS-Summer-2023/ARGO-MON/kind-argocd-playground$ 

```
<img src="pictures/ArgoCD-applications.png?raw=true" width="1000">

<img src="pictures/Grafana-DataSources.png?raw=true" width="1000">

<img src="pictures/Grafana-UI.png?raw=true" width="1000">
