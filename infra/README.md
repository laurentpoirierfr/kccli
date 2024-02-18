# Infrastructure

## /etc/resolv.conf

Add in **/etc/resolv.conf** minikube ip

```shell
nameserver <minikube ip>
search local
timeout 5s
```

## Cert Manager

```
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.14.2/cert-manager.yaml
```

## Keycloack 

```shell
minikube addons list
minikube addons enable ingress
minikube addons enable ingress-dns

export NAMESPACE=security
kubectl create ns $NAMESPACE
kubectl apply -k ./keycloak -n $NAMESPACE
```


```shell
open https://keycloak.k8s.local
```