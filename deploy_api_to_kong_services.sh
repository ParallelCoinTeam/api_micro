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


echo "Deployment "
echo "---------------------------- "
for d in */ ; do
  if [[ $d = *"service"* ]]; then
    echo "---------------------------- "
    echo "Processing: $d"
    echo "---------------------------- "
    TEMP_URL="$(minikube service ${d%/} --url)"
    echo $TEMP_URL

    if [[ ${d%-service/} != *"public"* && ${d%-service/} != *"subscrp"*  ]]; then
       curl -i -X POST --url $KONG_ADMIN/services/ --data 'name='${d%/}  --data 'url='$TEMP_URL'/v1/'${d%-service/}
       curl -i -X POST --url $KONG_ADMIN/services/${d%/}/routes --data "hosts[]="${d%-service/}
       curl -i -X POST --url $KONG_ADMIN/services/${d%/}/plugins --data "name=key-auth"
    fi
  fi
done
