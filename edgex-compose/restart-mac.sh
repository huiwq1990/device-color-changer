
# https://stackoverflow.com/questions/54807762/docker-compose-meaning-of-in-volume-definition
docker compose -f docker-compose-no-secty-with-app-sample.yml stop -t 1


docker compose -f docker-compose-no-secty-with-app-sample.yml rm -f

docker rm -v $(docker ps -aq -f status=exited)

prefix=edgex-compose
docker volume rm ${prefix}_consul-config ${prefix}_consul-data \
 ${prefix}_db-data\

#docker compose -f docker-compose-no-secty-with-app-sample.yml up --force-recreate
docker compose -f 2.2.0.yaml up --force-recreate



