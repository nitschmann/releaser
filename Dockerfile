FROM golang:1.25-alpine AS base

RUN apk add --no-cache --upgrade bash git g++ gcc make curl

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin v1.59.1

FROM base AS builder

WORKDIR /src

CMD ["bash"]
