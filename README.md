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
kubectl get deployments

```