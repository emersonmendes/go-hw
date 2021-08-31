# Project for study purpose only (Go + Postgres + Docker + Kubernetes)

## Build

``` sh
docker-compose -f docker-compose-build.yml up
```


## Run

``` sh
docker-compose up
```

### About Pull secret
    
    https://docs.microsoft.com/en-us/azure/container-registry/container-registry-auth-kubernetes

### About ingress basic nginx

    https://docs.microsoft.com/pt-br/azure/aks/ingress-basic


``` sh

# Create a namespace for your ingress resources
kubectl create namespace ingress-basic

# Add the ingress-nginx repository
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx

# Use Helm to deploy an NGINX ingress controller
helm install nginx-ingress ingress-nginx/ingress-nginx \
    --namespace ingress-basic \
    --set controller.replicaCount=2 \
    --set controller.nodeSelector."beta\.kubernetes\.io/os"=linux \
    --set defaultBackend.nodeSelector."beta\.kubernetes\.io/os"=linux \
    --set controller.admissionWebhooks.patch.nodeSelector."beta\.kubernetes\.io/os"=linux

# show
kubectl --namespace ingress-basic get services -o wide -w nginx-ingress-ingress-nginx-controller

# apply
kubectl apply -f aks-helloworld-one.yaml --namespace ingress-basic
kubectl apply -f aks-helloworld-two.yaml --namespace ingress-basic
kubectl apply -f go-hw.yaml --namespace ingress-basic
kubectl apply -f ingress.yaml --namespace ingress-basic

```