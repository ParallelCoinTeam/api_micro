#!/bin/bash

eval $(minikube docker-env)

KONG_ADMIN="$(minikube service kong-admin --url)"
echo $KONG_ADMIN

LENGTH="$(curl -X GET --url $KONG_ADMIN/routes | jq ' .["data"] | length') "
echo "Lenght:"$LENGTH

for (( c=0; c<=$LENGTH -1; c++ ))
do  
   echo "----------------"
   echo "Welcome $c times"
   ROUTE_ID="$(curl -X GET --url $KONG_ADMIN/routes | jq -c '.[]['$c'].id' | awk -F '"' '{print $2}' )"
   URL=$KONG_ADMIN/routes/$ROUTE_ID 
   URL=$(tr -dc '[[:print:]]' <<< "$URL")
   echo $URL
   curl -i -X DELETE --url $URL
   echo "----------------"
done

