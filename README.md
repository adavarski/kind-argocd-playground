# kind-argocd-playground

```
make launch-k8s
make deploy-argocd
kubectl port-forward service/argocd-server -n argocd 8080:443 &
kubectl get secret -n argocd argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
argocd login localhost:8080 --grpc-web --insecure --username admin --password $(kubectl get secret -n argocd argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d)
```
