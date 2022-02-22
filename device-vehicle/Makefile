.PHONY: build test clean docker

GO=CGO_ENABLED=1 GO111MODULE=on go

# Don't need CGO_ENABLED=1 on Windows w/o ZMQ.
# If it is enabled something is invoking gcc and causing errors
ifeq ($(OS),Windows_NT)
  GO=CGO_ENABLED=0 GO111MODULE=on go
endif

MICROSERVICES=cmd/device-virtual

.PHONY: $(MICROSERVICES)

DOCKERS=docker_device_vehicle
.PHONY: $(DOCKERS)

VERSION=v$(shell cat ./VERSION 2>/dev/null || echo 0.0.0)
GIT_SHA=v$(shell git rev-parse HEAD || echo -notset)
GOFLAGS=-ldflags "-X github.com/huiwq1990/device-vehicle.Version=$(VERSION)"

REGISTRY ?= docker.io
REPO ?= huiwq1990

tidy:
	go mod tidy

build: $(MICROSERVICES)

cmd/device-virtual:
	$(GO) build $(GOFLAGS) -o $@ ./cmd

test:
	$(GO) test ./... -coverprofile=coverage.out
	$(GO) vet ./...
	gofmt -l $$(find . -type f -name '*.go'| grep -v "/vendor/")
	[ "`gofmt -l $$(find . -type f -name '*.go'| grep -v "/vendor/")`" = "" ]
	./bin/test-attribution-txt.sh

clean:
	rm -f $(MICROSERVICES)

docker_build: $(DOCKERS)

docker_device_vehicle:
	docker build \
		--label "git_sha=$(GIT_SHA)" \
		-t ${REGISTRY}/${REPO}/device-vehicle:$(GIT_SHA) \
		-t ${REGISTRY}/${REPO}/device-vehicle:$(VERSION)-dev \
		.

docker_push: docker_push_device_vehicle

docker_push_device_vehicle:
	docker push ${REGISTRY}/${REPO}/device-vehicle:$(GIT_SHA)
	docker push ${REGISTRY}/${REPO}/device-vehicle:$(VERSION)-dev

vendor:
	CGO_ENABLED=0 GO111MODULE=on go mod vendor
