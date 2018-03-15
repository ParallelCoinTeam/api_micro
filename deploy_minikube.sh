#!/bin/bash

eval $(minikube docker-env)

AWS_ECR="755455355830.dkr.ecr.us-east-2.amazonaws.com"

echo "Removing previous deployment "
echo "---------------------------- "
for d in */ ; do
  if [[ $d = *"service"* ]]; then
    echo "Processing: $d"
    kubectl delete deployment ${d%-service/}-deployment 
    kubectl delete service ${d%-service/}-service 
  fi
done
kubectl delete service database-service
kubectl delete secret db-secret

echo "Deploying "
echo "--------- "
for d in */ ; do
  if [[ $d = *"service"* ]]; then
    echo "Processing: $d"
    cd $d
    pwd 
    make build
    make reverse
    make AWS_ECR="$AWS_ECR" REPOSITORY_NAME="${d%/}" docker 
    kubectl create -f ${d%-service/}.yaml
    cd ..
  fi
done

kubectl create -f db.yaml
kubectl create -f secret.yaml
