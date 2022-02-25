## 说明

作为边和设备的中介，
1）接收edgex的命令，传送给设备；
2）获取设备状态，传送给edgex；

## 调试命令

```shell
docker build -t hub.jdcloud.com/tpaas/docker-server-color-changer:v0.0.0-a .
docker push hub.jdcloud.com/tpaas/docker-server-color-changer:v0.0.0-a

docker run --network=host -d -p 5000:5000 hub.jdcloud.com/tpaas/device-color-changer-python:v0.0.0-a


```

```shell

docker run -it -v ${PWD}:/usr/src/app/ hub.jdcloud.com/tpaas/device-color-changer-python:v0.0.0-a /bin/bash

curl  http://127.0.0.1:5000/api/v1/device/xx/changeColor

curl -X PUT -H "Content-Type: application/json" -d '{"xx":"xx"}' http://127.0.0.1:5000/api/v1/device/xx/changeColor

```
