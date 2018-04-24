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

LENGTH="$(curl -X GET --url $KONG_ADMIN/routes | jq ' .["data"] | length') "
echo "Length:"$LENGTH

ROUTES="$(curl -X GET --url $KONG_ADMIN/routes )"
for (( c=0; c<=$LENGTH-1; c++ ))
do  
   echo "--------------------"
   echo "Deleting route $c "
   echo "--------------------"
   ROUTE_ID="$(jq -n "$ROUTES" | jq -c '.[]['$c'].id' | awk -F '"' '{print $2}' | tr -dc '[[:print:]]' )"
   URL=$KONG_ADMIN/routes/$ROUTE_ID 
   echo $URL
   curl -i -X DELETE --url $URL
done

deploySecure () {
    echo "-----------------------------"
    echo Deploying $1
    echo
    echo "-----------------------------Deleting Service-----------------------------"
    curl -i -X DELETE --url $KONG_ADMIN/services/$1
    echo
    echo "-----------------------------Creating new Service-----------------------------"
    curl -i -X POST --url $KONG_ADMIN/services/ --data 'name='$1 --data 'url='$3$2
    echo
    echo "-----------------------------Creating new Route-----------------------------"
    curl -i -X POST --url $KONG_ADMIN/services/$1/routes --data 'paths[]='$2 --data 'methods[]='$4
    echo
    echo "-----------------------------Applying key-auth to service----------------------"
    curl -i -X POST --url $KONG_ADMIN/services/$1/plugins --data "name=key-auth"
    echo
}
deployPublic () {
    echo "-----------------------------"
    echo Deploying $1
    echo "-----------------------------"
    echo
    echo "-----------------------------Deleting Service-----------------------------"
    curl -i -X DELETE --url $KONG_ADMIN/services/$1
    echo
    echo "-----------------------------Creating new Service-----------------------------"
    curl -i -X POST --url $KONG_ADMIN/services/ --data 'name='$1 --data 'url='$3$2
    echo
    echo "-----------------------------Creating new Route-----------------------------"
    curl -i -X POST --url $KONG_ADMIN/services/$1/routes --data 'paths[]='$2 --data 'methods[]='$4
}

deployPublic "public-register" "/v1/register" $PUBLIC_SRVC "POST"
deployPublic "public-authenticate" "/v1/authenticate" $PUBLIC_SRVC "POST"
$ROLE
deploySecure "role-create" "/v1/roles" $ROLES_SRVC "POST"
deploySecure "role-get-all" "/v1/roles" $ROLES_SRVC "GET"
deploySecure "role-get" "/v1/roles/role_id" $ROLES_SRVC "GET"
deploySecure "role-patch" "/v1/roles/role_id" $ROLES_SRVC "PATCH"
deploySecure "role-delete" "/v1/roles/role_id" $ROLES_SRVC "DELETE"
#USER
deploySecure "user-get-all" "/v1/users" $USERS_SRVC "GET"
deploySecure "user-get" "/v1/users/user_id" $USERS_SRVC "GET"
deploySecure "user-patch" "/v1/users/user_id" $USERS_SRVC "PATCH"
deploySecure "user-delete" "/v1/users/user_id" $USERS_SRVC "DELETE"

