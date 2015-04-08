#!/usr/bin/env bash

if [ ! -f make ]; then
    echo 'make must be run within its container folder' 1>&2
    exit 1
fi

if [ "$1" == "1.1.2" ]; then
    OLDGOROOT="$GOROOT"
    GOROOT="/usr/local/go1.1.2"
    echo $GOROOT
fi

if [ "$1" == "1.2.1" ]; then
    OLDGOROOT="$GOROOT"
    GOROOT="/usr/local/go1.2.1"
    echo $GOROOT
fi

if [ "$1" == "1.3" ]; then
    OLDGOROOT="$GOROOT"
    GOROOT="/usr/local/go1.3"
    echo $GOROOT
fi

if [ "$1" == "1.3rc2" ]; then
    OLDGOROOT="$GOROOT"
    GOROOT="/usr/local/go1.3rc2"
    echo $GOROOT
fi

if [ "$1" == "1.4" ]; then
    OLDGOROOT="$GOROOT"
    GOROOT="/usr/local/go1.4"
    echo $GOROOT
fi

CURDIR=`pwd`
export GOPATH="$CURDIR"
$GOROOT/bin/gofmt -w src


$GOROOT/bin/go install agent

echo 'finished'
