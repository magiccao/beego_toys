#!/bin/bash

CURDIR=$(pwd)   
export GOPATH=$GOPATH:${CURDIR}

cd src && go build main.go
mv main ../
