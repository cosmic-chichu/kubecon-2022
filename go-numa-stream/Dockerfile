FROM golang:1.21.3 as build

WORKDIR /go/src/go-numa-stream

COPY go.mod go.mod
COPY go.sum go.sum

# Set CGO_ENABLED to enable cgo
RUN go env -w CGO_ENABLED=1

RUN go mod download
COPY ./ /go/src/go-numa-stream

RUN go build -o go-numa-stream stream.go

COPY entrypoint.sh entrypoint.sh
RUN chmod +x entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]

#FROM debian:buster
#
#USER root
#
#RUN apt-get update \
#  && apt-get install -y --no-install-recommends jq \
#  && rm -rf /var/lib/apt/lists/*
#
#COPY --from=build /go/src/go-numa-stream/go-numa-stream .
#COPY entrypoint.sh entrypoint.sh
#RUN chmod +x entrypoint.sh
#
#ENTRYPOINT ["/entrypoint.sh"]
