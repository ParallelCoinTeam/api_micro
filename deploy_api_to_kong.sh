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
    curl -i -X DELETE --url $KONG_ADMIN/apis/${d%-service/}-api
  fi
done


echo "Deployment "
echo "---------------------------- "
for d in */ ; do
  if [[ $d = *"service"* ]]; then
    echo "---------------------------- "
    echo "Processing: $d"
    echo "---------------------------- "
    TEMP_URL="$(minikube service ${d%/} --url)"
    echo $TEMP_URL
    curl -i -X POST --url $KONG_ADMIN/apis/ --data 'name='${d%-service/}'-api' --data 'hosts='${d%-service/} --data 'upstream_url='$TEMP_URL'/v1/roles'
  
    if [[ ${d%-service/} != *"public"* ]]; then
       curl -i -X POST --url $KONG_ADMIN/apis/${d%-service/}-api/plugins --data "name=key-auth"
    fi
  fi
done
