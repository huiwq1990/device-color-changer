ARG BASE=golang:1.17-alpine3.15
FROM ${BASE} AS builder

ARG ALPINE_PKG_BASE="make git openssh-client gcc libc-dev zeromq-dev libsodium-dev"
ARG ALPINE_PKG_EXTRA=""

# set the working directory
WORKDIR /device-vehicle

# Replicate the APK repository override.
# If it is no longer necessary to avoid the CDN mirros we should consider dropping this as it is brittle.
RUN sed -e 's/dl-cdn[.]alpinelinux.org/nl.alpinelinux.org/g' -i~ /etc/apk/repositories
# Install our build time packages.
RUN apk add --update --no-cache ${ALPINE_PKG_BASE} ${ALPINE_PKG_EXTRA}

COPY go.mod vendor* ./
RUN [ ! -d "vendor" ] && go mod download all || echo "skipping..."

COPY . .
# To run tests in the build container:
#   docker build --build-arg 'MAKE=build test' .
# This is handy of you do your Docker business on a Mac
ARG MAKE='make build'
RUN $MAKE

FROM alpine:3.14
RUN sed -e 's/dl-cdn[.]alpinelinux.org/nl.alpinelinux.org/g' -i~ /etc/apk/repositories
RUN apk add --update --no-cache zeromq dumb-init
#FROM nicolaka/netshoot

WORKDIR /
COPY --from=builder /device-vehicle/cmd /
COPY --from=builder /device-vehicle/driver /

EXPOSE 59966

ENTRYPOINT ["/device-vehicle"]
CMD ["--cp=consul.http://edgex-core-consul:8500", "--registry", "--confdir=/res"]
