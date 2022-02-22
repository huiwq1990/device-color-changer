


curl -X PUT http://127.0.0.1:5000/api/v1/device/abc/changeColor \
-H 'Content-Type: application/json' \
-d '{"color":"red"}'


docker-compose rm -f
docker volume rm device-color-changer_consul-config device-color-changer_consul-data \
  device-color-changer_consul-scripts device-color-changer_db-data\
  device-color-changer_log-data

docker-compose up --force-recreate


https://stackoverflow.com/questions/54807762/docker-compose-meaning-of-in-volume-definition



#docker-compose pull
docker-compose up --build -d


# docker-compose stop -t 1



docker build -t hub.jdcloud.com/tpaas/docker-server-color-changer:v0.0.0 .

docker rm $( docker ps -a -q)

docker ip列表
docker ps -q | xargs -n 1 docker inspect --format '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} {{ .Name }}' | sed 's/ \// /'


curl -X POST http://127.0.0.1:48081/api/v1/device \
-H 'Content-Type: application/json' \
-d '
{
"name": "TestApp2",
"description": "Test application",
"adminState": "unlocked",
"operatingState": "enabled",
"protocols": {
"example": {
"host": "localhost",
"port": "0",
"unitID": "1"
}
},
"labels": [
"color",
"testapp"
],
"location": "tokyo",
"service": {
"name": "color-changer"
},
"profile": {
"name": "colorChanger"
}
}' 