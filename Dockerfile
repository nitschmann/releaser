FROM golang:1.18-alpine AS base

RUN apk add --upgrade bash git g++ gcc make

RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.45.2

FROM base AS builder

WORKDIR /src

CMD ["bash"]
