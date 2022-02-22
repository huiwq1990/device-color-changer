

make docker_device_color_changer


kubectl create secret docker-registry tpaas-itg --docker-server=hub-pub.jdcloud.com --docker-username=vessel-product --docker-password=qlzYnokM86 --docker-email=wangqinghui9@jd.com -n default




./device-color --cp=consul://edgex-core-consul:8500 --registry --confdir=/res
