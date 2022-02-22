
make docker_build

docker-compose stop -t 1

docker-compose rm -f
docker volume rm device-color-changer_consul-config device-color-changer_consul-data \
  device-color-changer_consul-scripts device-color-changer_db-data\
  device-color-changer_log-data

docker-compose up --force-recreate



