#!/bin/sh

if [ "$1" == "dev" ] ; then
    export GOPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
    export GO111MODULE=off
    go get -u github.com/thommil/tge@develop
else
    export GOPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/go"
    export GO111MODULE=on
    go get -u github.com/thommil/tge@master
fi

export PATH=$PATH:$GOPATH/bin