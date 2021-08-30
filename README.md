# Project for study purpose only (Go + Postgres + Docker)

## Build

``` sh
docker-compose -f docker-compose-build.yml up
```


## Run

``` sh
docker-compose up
```

# Azure Kubernets

``` sh

# Creating a Resource Group
az group create --location brazilsouth --name gohw_resource_group

# Create an AKS cluster with one worker node
az aks create --resource-group gohw_resource_group --name gohw-aks-cluster --node-count 1 --kubernetes-version 1.21.2 --generate-ssh-keys

# Show the existing AKS cluster JSON output
az aks show --name gohw-aks-cluster --resource-group gohw_resource_group

# Run the command below to copy the kubeconfig to a local computer
az aks get-credentials --name gohw-aks-cluster --resource-group gohw_resource_group

# view the nodes
kubectl get nodes

# Deploy example nginx
kubectl create -f https://raw.githubusercontent.com/kubernetes/website/master/content/en/examples/controllers/nginx-deployment.yaml

# view deployments
kubectl get deployments -A

# remove deploy
kubectl delete deploy deploymentname -n namespacename

# build image
docker build -t emersonRegistry.azurecr.io/go-hw:v1 .

# push docker image
az acr login --name emersonRegistry

docker login emersonRegistry.azurecr.io 

docker push emersonRegistry.azurecr.io/go-hw:v1

# deploy to kubernetes
kubectl create -f go-hw.yaml

# deploy load balancer
kubectl create namespace ingress-basic
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm install nginx-ingress ingress-nginx/ingress-nginx \
    --namespace ingress-basic \
    --set controller.replicaCount=2 \
    --set controller.nodeSelector."beta\.kubernetes\.io/os"=linux \
    --set defaultBackend.nodeSelector."beta\.kubernetes\.io/os"=linux \
    --set controller.admissionWebhooks.patch.nodeSelector."beta\.kubernetes\.io/os"=linux
kubectl --namespace ingress-basic get services -o wide -w nginx-ingress-ingress-nginx-controller

kubectl apply -f go-hw.yaml --namespace ingress-basic
kubectl apply -f go-hw-ingress.yaml --namespace ingress-basic

```