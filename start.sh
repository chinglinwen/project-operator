#!/bin/sh
cd cmd/manager
echo "compiling..."
go build -o project-operator || exit 1
cd ../..
pkill project-operator

WATCH_NAMESPACE="" cmd/manager/project-operator