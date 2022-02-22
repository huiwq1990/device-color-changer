


curl -X PUT http://127.0.0.1:5000/api/v1/device/abc/changeColor \
-H 'Content-Type: application/json' \
-d '{"color":"red"}'



docker build -t hub.jdcloud.com/tpaas/docker-server-color-changer:v0.0.0 .






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