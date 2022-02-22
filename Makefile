.PHONY: build test clean docker

GO = CGO_ENABLED=0 GO111MODULE=on go

MICROSERVICES=cmd/device-color
# REPO ?= hub.jdcloud.com
REGISTRY ?= docker.io
REPO ?= huiwq1990

.PHONY: $(MICROSERVICES)

DOCKERS=docker_build_device_vehicle docker_build_device_raspirest docker_build_raspberry_mocker docker_build_raspberry_server
.PHONY: $(DOCKERS)

VERSION=v$(shell cat ./VERSION 2>/dev/null || echo 0.0.0-a)
GIT_SHA=$(shell git rev-parse HEAD)
GOFLAGS=-ldflags "-X github.com/edgexfoundry/device-color-changer.Version=$(VERSION)"

build: $(MICROSERVICES)
	$(GO) build ./...

cmd/device-color:
	$(GO) build $(GOFLAGS) -o $@ ./cmd

test:
	$(GO) test ./... -coverprofile=coverage.out
	$(GO) vet ./...
	gofmt -l .
	[ "`gofmt -l .`" = "" ]
	./bin/test-attribution-txt.sh
	./bin/test-go-mod-tidy.sh

clean:
	rm -f $(MICROSERVICES)

docker_build: docker_build_device_vehicle docker_build_device_raspberrypi docker_build_raspberrypi_mocker docker_build_raspberrypi_server

# docker_build_device_color_changer:
# 	docker build \
# 		--label "git_sha=$(GIT_SHA)" \
# 		-t ${REGISTRY}/${REPO}/device-color-changer:$(VERSION) \
# 		.
#
# docker_build_device_color_changer_python:
# 	cd deviceColorChangerPython && \
# 	docker build -t ${REGISTRY}/${REPO}/raspberry-server:$(VERSION) .

docker_build_device_vehicle:
	cd device-vehicle && \
	make docker_build

docker_build_device_raspberrypi:
	cd device-raspberrypi && \
	make docker_build

docker_build_raspberrypi_mocker:
	cd raspberrypi-mocker && \
	make docker_build

docker_build_raspberrypi_server:
	cd raspberrypi-server && \
	make docker_build

docker_push: docker_push_device_vehicle docker_push_device_raspberrypi docker_push_raspberrypi_mocker docker_push_raspberrypi_server

docker_push_device_vehicle:
	cd device-vehicle && \
	make docker_push

docker_push_device_raspberrypi:
	cd device-raspberrypi && \
	make docker_push

docker_push_raspberrypi_mocker:
	cd raspberrypi-mocker && \
	make docker_push

docker_push_raspberrypi_server:
	cd raspberrypi-server && \
	make docker_push


docker_compose_up:
	cd edgex-compose && \
	sh restart.sh