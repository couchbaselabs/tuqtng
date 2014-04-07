#!/bin/sh

echo
echo "In your web browser open http://localhost:8093/tutorial/"
echo

WD="`dirname $0`"
$WD/cbq-engine -couchbase dir:$WD/data -staticPath $WD/static
