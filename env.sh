#!/bin/sh
    
export GOPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export PATH=$PATH:$GOPATH/bin
export GO111MODULE=off

go get -u github.com/thommil/tge