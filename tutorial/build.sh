#!/bin/bash

rm -f tutorial.zip
(cd .. && go build .)
(cd ../tuq_client && go build .)
go build .
7z a -tzip tutorial.zip ../tuqtng ../tuq_client/tuq_client ../test tutorial content 1>/dev/null
