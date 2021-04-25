#!/bin/bash

DOCKER_IMAGE=manigandanjeff/hfctl:latest
SERVER_BIN=hfctl
GO_OS=linux

# make build-server
# make GOOS=$GO_OS deps test build-server
make GOOS=$GO_OS deps build-server

echo "==> building docker image $DOCKER_IMAGE ..."
docker build -t $DOCKER_IMAGE .

echo "==> pushing docker image $DOCKER_IMAGE to docker hub ..."
docker push $DOCKER_IMAGE

echo "==> romoving $SERVER_BIN ..."
rm $SERVER_BIN
