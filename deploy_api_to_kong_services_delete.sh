#!/bin/bash

eval $(minikube docker-env)

KONG_ADMIN="$(minikube service kong-admin --url)"
echo $KONG_ADMIN

echo "Removing previous deployment "
echo "---------------------------- "
for d in */ ; do
  if [[ $d = *"service"* ]]; then
    echo "---------------------------- "
    echo "Processing: $d"
    echo "---------------------------- "
    if [[ ${d%-service/} != *"public"* && ${d%-service/} != *"subscrp"*  ]]; then
        curl -i -X DELETE --url $KONG_ADMIN/services/$d
    fi
  fi
done

