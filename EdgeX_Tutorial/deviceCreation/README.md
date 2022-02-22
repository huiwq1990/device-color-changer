
docker build -t hub.jdcloud.com/tpaas/device-create:v0.0.0 .


docker run --net=host -it hub.jdcloud.com/tpaas/device-create:v0.0.0  /bin/sh


python ./createRESTDevice.py -ip 127.0.0.1 -devip edgex-color-server


pip freeze requirements.txt

cat requirements.txt





pip install requests
pip install requests_toolbelt



curl -X POST -d '{"name":"camera control","protocol":"HTTP","address":"localhost","port":49977,"path":"/api/v1/callback","publisher":"none","user":"none","password":"none","topic":"none"}' localhost:48081/api/v1/addressable
curl -X POST -d '{"name":"camera1 address","protocol":"HTTP","address":"localhost","port":49999,"path":"/camera1","publisher":"none","user":"none","password":"none","topic":"none"}' localhost:48081/api/v1/addressable


curl  localhost:48081/api/v1/addressable

curl -X DELETE localhost:48081/api/v1/addressable/e0f4f643-7f3e-4b5e-8902-3b326c702b8f