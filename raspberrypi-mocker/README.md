## 说明

作为树莓派的模拟服务端，承载edgex对树莓派的服务。

## 调试命令

```shell
docker build -t hub.jdcloud.com/tpaas/docker-server-color-changer:v0.0.0-a .
docker push hub.jdcloud.com/tpaas/docker-server-color-changer:v0.0.0-a

docker run --network=host -d -p 5000:5000 hub.jdcloud.com/tpaas/device-color-changer-python:v0.0.0-a


```

```shell

docker run -it -v ${PWD}:/usr/src/app/ hub.jdcloud.com/tpaas/device-color-changer-python:v0.0.0-a /bin/bash

curl  http://127.0.0.1:5000/api/v1/device/xx/changeColor

curl -X PUT -H "Content-Type: application/json" -d '{"color":"black"}' http://192.168.68.101:5000/api/v1/device/xx/changeColor
curl -X PUT -H "Content-Type: application/json" -d '{"color":"red"}' http://192.168.68.101:5000/api/v1/device/xx/changeColor

curl -X PUT -H "Content-Type: application/json" -d '{"color":"xx"}' http://192.168.68.101:5000/api/v1/device/xx/changeColor

```

```shell


```
