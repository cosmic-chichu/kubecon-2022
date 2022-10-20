#!/bin/bash

set -e

## build go-numa-effects image
printf "build go-numa-effects image\n"
docker build -t docker.io/cosmicchichu/go-numa-effects -f ./go-numa-effects/Dockerfile ./go-numa-effects
printf "\n\n\n"

## push go-numa-effects image
printf "push go-numa-effects image\n"
docker push docker.io/cosmicchichu/go-numa-effects:latest
printf "\n\n\n"

## build go-numa-stream image
printf "build go-numa-stream image\n"
docker build -t docker.io/cosmicchichu/go-numa-stream -f ./go-numa-stream/Dockerfile ./go-numa-stream
printf "\n\n\n"

## push go-numa-stream image
printf "push go-numa-stream image\n"
docker push docker.io/cosmicchichu/go-numa-stream:latest
printf "\n\n\n"
