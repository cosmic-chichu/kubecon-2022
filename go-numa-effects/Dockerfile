FROM golang:1.17.1-buster as build

WORKDIR /go/src/go-numa-effects

COPY go.mod go.mod
COPY go.sum go.sum

RUN mkdir output

RUN go mod download
COPY ./ /go/src/go-numa-effects

RUN go build -o go-numa-effects main.go

FROM debian:buster

USER root

RUN apt-get update \
  && apt-get install -y --no-install-recommends jq \
  && rm -rf /var/lib/apt/lists/*

COPY --from=build /go/src/go-numa-effects/go-numa-effects .
COPY entrypoint.sh entrypoint.sh
RUN chmod +x entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
