# Project for study purpose only (Go + Postgres + Docker + Azure + Kubernetes)

## Build

``` sh
docker-compose -f docker-compose-build.yml up
```


## Run

``` sh
docker-compose up
```

### About Pull secret ( In caso of Image pull back off error  )
    
    https://docs.microsoft.com/en-us/azure/container-registry/container-registry-auth-kubernetes

### About ingress basic nginx

    https://docs.microsoft.com/pt-br/azure/aks/ingress-basic


### Remove resources

``` sh

az acr delete -n MyRegistry

az aks delete --name MyManagedCluster --resource-group MyResourceGroup

az group delete --name xxxxx

```


### Create resources Azure Kubernetes

    Create a resource group

``` sh
# Create Azure Container Registry
az acr create --resource-group EmersonResourceGroup \
  --name EmersonContainerRegistry --sku Basic

# Create AKS cluster 
az aks create \
    --resource-group EmersonResourceGroup \
    --name EmersonAKSCluster \
    --node-count 2 \
    --generate-ssh-keys \
    --attach-acr EmersonContainerRegistry

# ACR login
az acr login --name EmersonContainerRegistry --expose-token

# Docker login
docker login emersoncontainerregistry.azurecr.io -u 00000000-0000-0000-0000-000000000000

# Push Docker image to ACR

docker build -t emersoncontainerregistry.azurecr.io/go-hw:v1 .

docker push emersoncontainerregistry.azurecr.io/go-hw:v1

# Connect to AKS if not already
az aks get-credentials --name EmersonAKSCluster --resource-group EmersonResourceGroup

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

kubectl apply -f go-hw.yaml --namespace ingress-basic # don't forget to change image name/version in .yaml

kubectl apply -f aks-helloworld-one.yaml --namespace ingress-basic
kubectl apply -f aks-helloworld-two.yaml --namespace ingress-basic

kubectl apply -f ingress.yaml --namespace ingress-basic


```

