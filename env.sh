#!/bin/sh

LOCAL_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

if [ ! -d "$LOCAL_PATH/tge-builder" ];then
    git clone https://github.com/thommil/tge-builder
fi

if [ "$GO111MODULE" == "on" ] ; then
    export GOPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/go"
    echo "Retrieving dependencies ..."
    go get -u github.com/thommil/tge@develop
else
    export GOPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
    go get -u github.com/thommil/tge
    cd src/github.com/thommil/tge && git checkout develop && cd -
fi