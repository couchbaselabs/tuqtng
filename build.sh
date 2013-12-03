#!/bin/bash

if [ -z "$1" ]
    then
    echo "echo building cbq_engine"
    go build -o cbq_engine
elif [ $1 == "clean" ]
    then
    echo "cleaning cbq_engine"
    go clean
    rm -f cbq_engine
else
    echo "No comprehendo senor !"
fi

