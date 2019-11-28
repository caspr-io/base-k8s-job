FROM alpine:3.10

RUN apk add --no-cache curl git ca-certificates

RUN addgroup -S caspr && \
  adduser -S caspr -G caspr --home /caspr

WORKDIR /caspr
USER caspr
