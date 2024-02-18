# kccli


## ingress-dns

```shell
minikube addons enable ingress
minikube addons enable ingress-dns
```

Add in **/etc/resolv.conf** minikube ip

```shell
nameserver <minikube ip>
search local
timeout 5s
```

## References 

* https://github.com/spf13/cobra-cli/blob/main/README.md


