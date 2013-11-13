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
    $ go build
    $ ./tuqtng -couchbase http://localhost:8091

To run against the test json files as your "site":

    $ ./tuqtng -couchbase dir:./test -pool=json

Then in your tuq_client:

    $ ./tuq_client/tuq_client
    tuq> select * from orders;

To run unit tests:

    $ go test ./...

## Querying via interactive command-line tool

    $ tuq_client
    tuq> SELECT * FROM beer-sample LIMIT 5

The --tuqtng parameter let's you specify the tuqtng server:

    $ tuq_client --tuqtng="http://localhost:8093/"

## Querying via HTTP

### HTTP Get

    curl 'http://localhost:8093/query?q=URL_ENCODED_QUERY_STRING'

### HTTP Post

    curl -HContent-Type:text/plain -XPOST http://localhost:8093/query -d 'QUERY_STRING'
