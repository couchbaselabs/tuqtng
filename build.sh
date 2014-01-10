#!/bin/bash

if [ -z "$1" ]
    then
    echo "building cbq-engine and cbq/cbq"
    go build -o cbq-engine
    cd cbq; go build; cd ..
elif [ $1 == "debug" ]
    then
    echo "building cbq-engine and cbq/cbq with debug flags"
    go build -gcflags "-N -l" -o cbq-engine
    cd cbq; go build -gcflags "-N -l"; cd ..
elif [ $1 == "release" ]
    then
    echo "building cbq-engine and cbq/cbq without debug symbols"
    go build -ldflags "-s" -o cbq-engine
    cd cbq; go build -ldflags "-s"; cd ..
elif [ $1 == "clean" ]
    then
    echo "cleaning.. "
    go clean
    rm -f cbq-engine
    cd cbq; rm -f cbq; cd ..
elif [ $1 == "fmt" ]
    then
    go fmt ./...
elif [ $1 == "tags" ]
then
    find . -name "*.go" > cscope.files
    find . -name "*.go" | xargs gotags > tags 
else
    echo "No comprehendo senor !"
fi

