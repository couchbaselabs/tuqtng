#!/bin/sh

for file in test/json/*
    do
      if [ -d "$file" ];then
       bn=`basename "$file"`
       curl -s -XPOST "http://localhost:8091/_api/buckets" -d "name=$bn"
       loadfile -bucket $bn -v test/json/$bn/*
      fi
    done
