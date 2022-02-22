.PHONY: build test clean docker

VERSION=v$(shell cat ./VERSION 2>/dev/null || echo 0.0.0)
GIT_SHA=v$(shell git rev-parse HEAD || echo -notset)

REGISTRY ?= docker.io
REPO ?= huiwq1990

DOCKERS=docker_raspberry_server
.PHONY: $(DOCKERS)

clean:
	rm -f $(MICROSERVICES)

docker_build: $(DOCKERS)

docker_raspberry_server:
	docker build \
		--label "git_sha=$(GIT_SHA)" \
		-t ${REGISTRY}/${REPO}/raspberrypi-server:$(GIT_SHA) \
		-t ${REGISTRY}/${REPO}/raspberrypi-server:$(VERSION)-dev \
		.

docker_push:
	docker push ${REGISTRY}/${REPO}/raspberrypi-server:$(GIT_SHA)
	docker push ${REGISTRY}/${REPO}/raspberrypi-server:$(VERSION)-dev