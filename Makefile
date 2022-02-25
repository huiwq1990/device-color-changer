.PHONY: build test clean docker

GO = CGO_ENABLED=0 GO111MODULE=on go

MICROSERVICES=cmd/device-color
# REPO ?= hub.jdcloud.com
REGISTRY ?= docker.io
REPO ?= huiwq1990

.PHONY: $(MICROSERVICES)

DOCKERS=docker_device_virtual_go
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

docker: $(DOCKERS)

docker_build_device_color_changer:
	docker build \
		--label "git_sha=$(GIT_SHA)" \
		-t ${REGISTRY}/${REPO}/device-color-changer:$(VERSION) \
		.
# 		-t ${REPO}/tpaas/docker-device-color-changer:$(GIT_SHA) \

docker_push:
	docker push ${REGISTRY}/${REPO}/device-color-changer:$(VERSION)

docker_build_device_color_changer_python:
	cd deviceColorChangerPython && \
	docker build -t ${REGISTRY}/${REPO}/device-color-changer-python:$(VERSION) .




