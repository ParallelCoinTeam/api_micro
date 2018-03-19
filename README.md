# api_micro
A distributed API service example



$ grep -rl 'Role' . | xargs sed -i 's/Role/User/g'
$ grep -rl 'role' . | xargs sed -i 's/role/user/g'
$ grep -ir -o -w 'Role' .  | wc -w


docker rmi -f $(docker images | grep "^<none>" | awk "{print $3}" )
