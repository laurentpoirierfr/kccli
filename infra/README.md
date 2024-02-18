# Keycloak

## keycloack 

```shell
minikube addons list
minikube addons enable ingress
minikube addons enable ingress-dns

export NAMESPACE=security
kubectl create ns $NAMESPACE
kubectl apply -k ./keycloak -n $NAMESPACE
```

## /etc/resolv.conf

Add in **/etc/resolv.conf** minikube ip

```shell
nameserver <minikube ip>
search local
timeout 5s
```