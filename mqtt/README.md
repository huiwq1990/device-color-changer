
cd ~/github/edgex-compose/

cd compose-builder
make gen ds-mqtt mqtt-broker no-secty ui

ls | grep 'docker-compose.yml'



docker compose -f docker-compose.yml -f docker-compose.override.yml up -d





docker run --rm --name=mqtt-scripts \
    -v /Users/wangqinghui9/github/device-color-changer/mqtt/mqtt-scripts:/scripts  --network host \
    dersimn/mqtt-scripts --dir /scripts




curl --request DELETE http://localhost:59881/api/v2/device/name/my-custom-device





## 创建设备

curl -X 'POST' 'http://localhost:59881/api/v2/device' -d '[{"apiVersion":"v2","device":{"name":"test-device-01","adminState":"UNLOCKED","operatingState":"UP","labels":["test"],"serviceName":"device-mqtt","profileName":"my-custom-device-profile","protocols":{"mqtt":{"CommandTopic":"command/test-device-01"}}}}]'


docker run --rm --name=test-device-01 \
    -v /Users/wangqinghui9/github/device-color-changer/mqtt/test-device-01:/scripts  --network host \
    dersimn/mqtt-scripts --dir /scripts



curl -X GET localhost:59880/api/v2/event/device/name/test-device-01

