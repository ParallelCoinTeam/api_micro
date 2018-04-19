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


#curl -i -X DELETE --url $KONG_ADMIN/services/public-register 
#curl -i -X POST --url $KONG_ADMIN/services/ --data 'name=public-register' --data 'url='$PUBLIC_SRVC'/v1/register'
#curl -i -X POST --url $KONG_ADMIN/services/public-register/routes --data 'hosts[]=register'
#
#curl -i -X DELETE --url $KONG_ADMIN/services/public-authenticate 
#curl -i -X POST --url $KONG_ADMIN/services/ --data 'name=public-authenticate' --data 'url='$PUBLIC_SRVC'/v1/authenticate'
#curl -i -X POST --url $KONG_ADMIN/services/public-authenticate/routes --data 'hosts[]=authenticate'


deploySecure () {
    echo "-----------------------------"
    echo Deploying $1
    echo "-----------------------------"
    curl -i -X DELETE --url $KONG_ADMIN/services/$1
    curl -i -X POST --url $KONG_ADMIN/services/ --data 'name='$1 --data 'url='$3$2
    curl -i -X POST --url $KONG_ADMIN/services/$1/routes --data 'hosts[]='$1
    curl -i -X POST --url $KONG_ADMIN/services/$1/plugins --data "name=key-auth"
}
deployPublic () {
    echo "-----------------------------"
    echo Deploying $1
    echo "-----------------------------"
    curl -i -X DELETE --url $KONG_ADMIN/services/$1
    curl -i -X POST --url $KONG_ADMIN/services/ --data 'name='$1 --data 'url='$3$2
    curl -i -X POST --url $KONG_ADMIN/services/$1/routes --data 'hosts[]='$1
}

deployPublic "public-register" "/v1/register" $PUBLIC_SRVC
deployPublic "public-authenticate" "/v1/authenticate" $PUBLIC_SRVC
deploySecure "role-create" "/v1/roles" $ROLES_SRVC

#echo "Removing previous deployment "
#echo "---------------------------- "
#for d in */ ; do
#  if [[ $d = *"service"* ]]; then
#    echo "---------------------------- "
#    echo "Processing: $d"
#    echo "---------------------------- "
#    if [[ ${d%-service/} != *"public"* && ${d%-service/} != *"subscrp"*  ]]; then
#        curl -i -X DELETE --url $KONG_ADMIN/services/$d
#    fi
#  fi
#done
#
#
#echo "Deployment "
#echo "---------------------------- "
#for d in */ ; do
#  if [[ $d = *"service"* ]]; then
#    echo "---------------------------- "
#    echo "Processing: $d"
#    echo "---------------------------- "
#    TEMP_URL="$(minikube service ${d%/} --url)"
#    echo $TEMP_URL
#
#    if [[ ${d%-service/} != *"public"* && ${d%-service/} != *"subscrp"*  ]]; then
#       curl -i -X POST --url $KONG_ADMIN/services/ --data 'name='${d%/}  --data 'url='$TEMP_URL'/v1/'${d%-service/}
#       curl -i -X POST --url $KONG_ADMIN/services/${d%/}/routes --data "hosts[]="${d%-service/}
#       curl -i -X POST --url $KONG_ADMIN/services/${d%/}/plugins --data "name=key-auth"
#    fi
#  fi
####done
