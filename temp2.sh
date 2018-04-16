#!/bin/bash

eval $(minikube docker-env)

KONG_ADMIN="$(minikube service kong-admin --url)"
echo $KONG_ADMIN

ROUTE_ID="$(curl -X GET --url $KONG_ADMIN/routes | jq -c '.[][0].service.id' | awk -F '"' '{print $2}' )"

echo $ROUTE_ID
