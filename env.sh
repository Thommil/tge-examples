#!/bin/sh

if [ "$1" == "dev" ] ; then
    export GOPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
    export GO111MODULE=off
else
    export GOPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/go"
    export GO111MODULE=on
fi

export PATH=$PATH:$GOPATH/bin