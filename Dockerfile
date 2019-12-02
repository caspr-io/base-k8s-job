FROM alpine:3.10

RUN apk add --no-cache curl git ca-certificates jq

# Install yq (jq for Yaml)
ARG YQ_VERSION=2.4.1
RUN curl -LO https://github.com/mikefarah/yq/releases/download/${YQ_VERSION}/yq_linux_amd64 && \
  mv yq_linux_amd64 /usr/bin/yq && \
  chmod +x /usr/bin/yq

RUN addgroup -S caspr && \
  adduser -S caspr -G caspr --home /caspr

WORKDIR /caspr
USER caspr
