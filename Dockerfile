FROM golang:1.13-alpine as module_base

WORKDIR /go/src/caspr-result

# Force the go compiler to use modules
ENV GO111MODULE=on

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

FROM module_base as builder

COPY . .

RUN go build -o caspr-result ./main.go

FROM alpine:3.10

RUN apk add --no-cache curl git ca-certificates jq

# Install yq (jq for Yaml)
ARG YQ_VERSION=2.4.1
RUN curl -LO https://github.com/mikefarah/yq/releases/download/${YQ_VERSION}/yq_linux_amd64 && \
  mv yq_linux_amd64 /usr/bin/yq && \
  chmod +x /usr/bin/yq


COPY --from=builder /go/src/caspr-result/caspr-result /usr/bin

RUN addgroup -S caspr && \
  adduser -S caspr -G caspr --home /caspr

WORKDIR /caspr
USER caspr

# Set environment variables for provisioning result replier
ENV NATS_URL=""
ENV STAN_CLUSTER_ID="streaming-stan"
ENV STAN_CLIENT_ID=""
ENV STAN_QUEUE_GROUP="provisioner"
