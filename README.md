# kind-argocd-playground

```
make launch-k8s
make deploy-argocd
kubectl port-forward service/argocd-server -n argocd 8080:443 &
kubectl get secret -n argocd argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
argocd login localhost:8080 --grpc-web --insecure --username admin --password $(kubectl get secret -n argocd argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d)
```
Browser: https://localhost:8080 -> Sync apps in this order
```
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
```
