#!/bin/bash
rm -rf build/
if [ ! -d build ]
then
    mkdir build
fi

CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o build/mydocker *.go
# 开启CGO后不支持交叉编译
# CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o build/mydocker-arm *.go