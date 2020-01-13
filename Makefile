ROOTPROJECT ?= ../root
CI_REGISTRY=docker.io
CI_REGISTRY_IMAGE=caspr/base-k8s-job
PROTOBUF_FILES=api/provisioning/provisioning_services.pb.go
PROJECT_PACKAGE=github.com/caspr-io/caspr
GENERATE_GROUPS=traefik:v1alpha1

include ${ROOTPROJECT}/include.mk

# Builds Docker image
.PHONY: build install
build: kube/generate protobuf/generate docker/build
install: build docker/push

# Dummy targets for cluster/up and cluster/teardown
.PHONY: up down
up:
down:
