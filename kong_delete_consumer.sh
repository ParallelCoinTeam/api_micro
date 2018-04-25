#!/bin/bash

eval $(minikube docker-env)

KONG_ADMIN="$(minikube service kong-admin --url)"
PUBLIC_SRVC="$(minikube service public-srvc --url)"
ROLES_SRVC="$(minikube service roles-srvc --url)"
USERS_SRVC="$(minikube service roles-srvc --url)"
echo $KONG_ADMIN
echo $PUBLIC_SRVC
echo $ROLES_SRVC
echo $USERS_SRVC

LENGTH="$(curl -X GET --url $KONG_ADMIN/consumers | jq ' .["data"] | length') "
echo "Length:"$LENGTH

ROUTES="$(curl -X GET --url $KONG_ADMIN/consumers )"
echo $ROUTES
for (( c=0; c<=$LENGTH-1; c++ ))
do  
   echo "--------------------"
   echo "Deleting route $c "
   echo "--------------------"
   ROUTE_ID="$(jq -n "$ROUTES" | jq -c '.["data"]['$c'].id' | awk -F '"' '{print $2}' | tr -dc '[[:print:]]' )"
   echo $ROUTE_ID
   
   URL=$KONG_ADMIN/consumers/$ROUTE_ID 
   echo $URL
   curl -i -X DELETE --url $URL
done

