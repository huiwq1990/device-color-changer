# Color Changer
Test app for REST commands. 
Web interface shows a 300x300px colored square on port 5000. 
REST API accepts changes to the color. 
Web interface uses JS and AJAX to auto refresh, reflecting any new updates pushed via the REST API. 

Example: 
Access web page over port 5000
Push REST call: 
Method: PUT
URI: <ip>:5000/api/v1/device/id/changeColor
BODY: 
{
    "color": "orange"
}



docker build -t hub.jdcloud.com/tpaas/docker-server-color-changer:v0.0.0-a .

docker push hub.jdcloud.com/tpaas/docker-server-color-changer:v0.0.0-a

docker run -it -v ${PWD}:/usr/src/app/ hub.jdcloud.com/tpaas/docker-server-color-changer:v0.0.0-a /bin/bash

curl  http://127.0.0.1:5000/api/v1/device/xx/changeColor

curl -X PUT -H "Content-Type: application/json" -d '{"xx":"xx"}' http://127.0.0.1:5000/api/v1/device/xx/changeColor




114.67.122.239:6443