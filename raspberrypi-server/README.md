## 说明

封装对树莓派硬件的操作服务，主要原因是golang操作gpio不太方便。

## 树莓派编译

```shell
scp -r ./raspberrypi-server/ ubuntu@192.168.68.101:/home/ubuntu
```

## 调试命令

```shell

make docker_build
make docker_push

docker run --privileged=true --network=host -p 9991:9991 -d huiwq1990/raspberrypi-server:v0.0.0-a 
# https://blog.csdn.net/u013042928/article/details/86484777
docker run --privileged=true --network=host -p 9991:9991 huiwq1990/raspberrypi-server:v0.0.0-dev

```

```shell

docker run -it -v ${PWD}:/usr/src/app/ hub.jdcloud.com/tpaas/device-color-changer-python:v0.0.0-a /bin/bash

curl  http://127.0.0.1:9991/api/v1/device/xx/changeColor

curl -X PUT -H "Content-Type: application/json" -d '{"color":"black"}' http://192.168.68.101:9991/api/v1/device/xx/changeColor
curl -X PUT -H "Content-Type: application/json" -d '{"color":"red"}' http://192.168.68.101:9991/api/v1/device/xx/changeColor

curl -X PUT -H "Content-Type: application/json" -d '{"color":"xx"}' http://192.168.68.101:5000/api/v1/device/xx/changeColor

```
