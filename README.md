# ArgoCD studies

Configure CI/CD pipeline + k8s (kind) + ArgoCD + GitOps (in same repo OR in separated repo)

```bash
# Create k8s cluster
kind create cluster
kubectl cluster-info --context kind-kind

# Install ArgoCD inside k8s cluster +
# Expose ArgoCD UI to port 3009 +
# Get default UI login credentials
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
kubectl port-forward svc/argocd-server -n argocd 3009:443 --address='0.0.0.0'
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo

# Access https://localhost:3009 (with HTTPS, oh yeah! xD)
Credentials:
  - Username: admin
  - Password: Value from k8s secret argocd-initial-admin-secret

# Create an app inside ArgoCD UI +
# Configure sync pasting the content of this file in "Edit as YAML" option
# Choose one of the options (GitOps using same repo or separated repo)
argocd_app_with_gitops_in_the_same_repo.yaml
or
argocd_app_with_gitops_in_separated_repo.yaml

# Delete the k8s cluster
kubectl delete all --all
kubectl delete "$(kubectl api-resources --namespaced=true --verbs=delete -o name | tr "\n" "," | sed -e 's/,$//')" --all
kubectl delete all --all --all-namespaces
kind delete cluster
```
