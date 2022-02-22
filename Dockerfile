# Copyright (c) 2020 IOTech Ltd
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

ARG BASE=golang:1.15-alpine
FROM ${BASE} AS builder

ARG ALPINE_PKG_BASE="build-base git openssh-client"
ARG ALPINE_PKG_EXTRA=""

LABEL license='SPDX-License-Identifier: Apache-2.0' \
  copyright='Copyright (c) 2019-2020: IOTech'

# Replicate the APK repository override.
# If it is no longer necessary to avoid the CDN mirros we should consider dropping this as it is brittle.
RUN sed -e 's/dl-cdn[.]alpinelinux.org/nl.alpinelinux.org/g' -i~ /etc/apk/repositories
# Install our build time packages.
RUN apk add --update --no-cache ${ALPINE_PKG_BASE} ${ALPINE_PKG_EXTRA}

# set the working directory
WORKDIR $GOPATH/src/github.com/edgexfoundry/device-color-changer

COPY . .

# To run tests in the build container:
#   docker build --build-arg 'MAKE=build test' .
# This is handy of you do your Docker business on a Mac
#ARG MAKE='make build'
#RUN $MAKE
RUN CGO_ENABLED=0 GO111MODULE=on go build -ldflags "-X github.com/edgexfoundry/device-color-changer.Version=" -o cmd/device-color ./cmd

#FROM alpine
FROM centos:7
ENV APP_PORT=5000
EXPOSE $APP_PORT

COPY --from=builder /go/src/github.com/edgexfoundry/device-color-changer/cmd /
COPY --from=builder /go/src/github.com/edgexfoundry/device-color-changer/Attribution.txt /
COPY --from=builder /go/src/github.com/edgexfoundry/device-color-changer/LICENSE /

ENTRYPOINT ["/device-color"]
CMD ["--cp=consul://edgex-core-consul:8500", "--registry", "--confdir=/res"]
