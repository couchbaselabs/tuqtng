#!/bin/bash

if [ -z "$1" ]
    then
    echo "building cbq-engine"
    go build -o cbq-engine
elif [ $1 == "debug" ]
    then
    echo "building cbq-engine with debug flags"
    go build -gcflags "-N -l" -o cbq-engine
elif [ $1 == "release" ]
    then
    echo "building cbq-engine without debug symbols"
    go build -ldflags "-s" -o cbq-engine
elif [ $1 == "clean" ]
    then
    echo "cleaning.. "
    go clean
    rm -f cbq-engine
elif [ $1 == "fmt" ]
    then
    go fmt ./...
elif [ $1 == "tags" ]
then
    find . -name "*.go" > cscope.files
    find . -name "*.go" | xargs gotags > ctags 
else
    echo "No comprehendo senor !"
fi

