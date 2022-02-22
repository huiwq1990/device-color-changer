

dockerComposeFile=docker-compose-arm.yml
# https://stackoverflow.com/questions/54807762/docker-compose-meaning-of-in-volume-definition
docker-compose -f ${dockerComposeFile} stop -t 1


docker-compose -f ${dockerComposeFile} rm -f

docker rm -v $(docker ps -aq -f status=exited)

prefix=edgex-compose
docker volume rm ${prefix}_consul-config ${prefix}_consul-data \
 ${prefix}_db-data\

docker-compose -f ${dockerComposeFile} up --force-recreate



