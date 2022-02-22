.PHONY: build test clean docker

GO = CGO_ENABLED=0 GO111MODULE=on go

MICROSERVICES=cmd/device-color
REPO ?= hub.jdcloud.com

.PHONY: $(MICROSERVICES)

DOCKERS=docker_device_virtual_go
.PHONY: $(DOCKERS)

VERSION=v$(shell cat ./VERSION 2>/dev/null || echo 0.0.0)
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

docker_device_color_changer:
	docker build \
		--label "git_sha=$(GIT_SHA)" \
		-t ${REPO}/tpaas/docker-device-color-changer:$(VERSION) \
		.
# 		-t ${REPO}/tpaas/docker-device-color-changer:$(GIT_SHA) \

# 	docker push ${REPO}/tpaas/docker-device-color-changer:$(VERSION)
