![](docs/tuq.png)

![](docs/tng.png)

## Running

### Users

    $ go get github.com/couchbaselabs/tuqtng
    $ tuqtng -couchbase http://localhost:8091

### Developers

    $ mkdir -p $GOPATH/src/github.com/couchbaselabs
    $ cd $GOPATH/src/github.com/couchbaselabs
    $ git clone git@github.com:couchbaselabs/tuqtng.git
    $ cd tuqtng
    $ go get -d -v ./...
    $ ./build.sh
    $ ./cbq-engine -couchbase http://localhost:8091

To run against the test json files as your "site":

    $ ./cbq-engine -couchbase dir:./test -pool=json

Then in your cbq:

    $ ./cbq/cbq
    cbq> select * from orders;

To run unit tests:

    $ go test ./...

## Querying via interactive command-line tool

    $ cbq
    cbq> SELECT * FROM beer-sample LIMIT 5

The --engine parameter let's you specify the cbq-engine server:

    $ cbq --engine="http://localhost:8093/"

## Querying via HTTP

### HTTP Get

    curl 'http://localhost:8093/query?q=URL_ENCODED_QUERY_STRING'

### HTTP Post

    curl -HContent-Type:text/plain -XPOST http://localhost:8093/query -d 'QUERY_STRING'
