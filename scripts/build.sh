#!/bin/sh

apk update && apk add make git

if [ ! -e ./go.mod ]; then
    go mod init go-testing-practice
fi

make
