ROOTPROJECT ?= ../root
CI_REGISTRY = docker.io
CI_REGISTRY_IMAGE = caspr/base-provisioner
PROTOBUF_FILES=api/provisioning/provisioning_events.pb.go

include ${ROOTPROJECT}/include.mk

# Builds Docker image
.PHONY: build install
build: protobuf/generate docker/build
install: build docker/push

# Dummy targets for cluster/up and cluster/teardown
.PHONY: up down
up:
down:
