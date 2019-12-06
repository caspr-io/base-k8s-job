ROOTPROJECT ?= ../root
CI_REGISTRY = docker.io
CI_REGISTRY_IMAGE = caspr/base-provisioner
include ${ROOTPROJECT}/include.mk

# Builds Docker image
.PHONY: build install
build: docker/build
install: build docker/push

# Dummy targets for cluster/up and cluster/teardown
.PHONY: up down
up:
down:
