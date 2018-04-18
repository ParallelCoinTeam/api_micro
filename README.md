# api_micro
A distributed API service example



$ grep -rl 'Role' . | xargs sed -i 's/Role/User/g'
$ grep -rl 'role' . | xargs sed -i 's/role/user/g'
$ grep -ir -o -w 'Role' .  | wc -w


docker rmi -f $(docker images | grep "^<none>" | awk "{print $3}" )


kubectl get jobs --all-namespaces | sed '1d' | awk '{ print $2, "--namespace", $1 }' | while read line; do kubectl delete jobs $line; done


awk '{print $0","}' temp1 > temp2

cut -c1-2 temp > temp1

#
dep init && dep ensure
