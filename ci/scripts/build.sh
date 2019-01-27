#!/bin/sh
set -e -u -x
# Install git for go get

echo ">> Install git"
apk add --no-cache git

# set up directory stuff for golang
echo ">> Setup Directories"
mkdir -p /go/src/github.com/shreddedbacon/
ln -s $PWD/uaa-webui-release /go/src/github.com/shreddedbacon/uaa-webui
ls -alh /go/src/github.com/shreddedbacon
cd  /go/src/github.com/shreddedbacon/uaa-webui
echo ">> Get"
go get -v .
cd -
echo ">> Build"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o built-release/uaa-webui github.com/shreddedbacon/uaa-webui

echo ">> Create artifact"
VERSION=$(cat ${VERSION_FROM})
cd built-release
tar czf uaa-webui-linux-$VERSION.tar.gz uaa-webui
