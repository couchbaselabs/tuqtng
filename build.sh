#!/bin/bash

if [ -z "$1" ]
    then
    echo "echo building cbq-engine"
    go build -o cbq-engine
elif [ $1 == "clean" ]
    then
    echo "cleaning cbq-engine"
    go clean
    rm -f cbq-engine
else
    echo "No comprehendo senor !"
fi

