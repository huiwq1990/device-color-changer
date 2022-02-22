

dockerComposeFile=docker-compose-arm.yml

docker-compose -f ${dockerComposeFile} down

# https://stackoverflow.com/questions/54807762/docker-compose-meaning-of-in-volume-definition
# docker-compose -f ${dockerComposeFile} stop -t 1

docker-compose -f ${dockerComposeFile} up
