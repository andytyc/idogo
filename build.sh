#!/bin/bash

if [ "$#" != 1 ];then
    echo "执行错误, 必须有1个参数:\n  编译平台{amd,arm,darwin,windows}"
    exit 0
fi

if [ "$1" != "amd" -a "$1" != "arm" -a "$1" != "darwin" -a "$1" != "windows" ]
then
    echo "编译平台不合法"
    exit 0
fi

if [ "$1" = "amd" ]
then
    export GOOS=linux
    export GOARCH=amd64
elif [ "$1" = "arm" ]
then
    export GOOS=linux
    export GOARCH=arm64
elif [ "$1" = "darwin" ]
then
    export GOOS=darwin
    export GOARCH=amd64
elif [ "$1" = "windows" ]
then
    export CGO_ENABLED=0
    export GOOS=windows
    export GOARCH=amd64
fi

echo ">>>> build do $1 {$GOOS, $GOARCH}"

git checkout .

go mod tidy

function build_do() {
    echo "**** build_do start {$1} ****"
    if [ "$1" = "windows" ]; then
        go build -o idogo-$1 -ldflags "-s -w"
        md5sum idogo-$1.exe
    else
        go build -o idogo-$1 -ldflags "-s -w"
        md5sum idogo-$1
    fi
    echo "**** build_do end {$1} ****"
}

build_do $1

echo ">>>> build done $1 {$GOOS, $GOARCH}"
