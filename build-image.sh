#!/bin/sh
# build image
set -e

suffix="$1"
suffix=${suffix:=v1}

export GO111MODULE=on

cd cmd/manager
go build -o project-operator || exit 1
cd ../..

image="project-operator:$suffix"
echo -e "building image: $image\n"
tag="harbor.haodai.net/ops/$image"
docker build -t $tag .
docker push $tag
