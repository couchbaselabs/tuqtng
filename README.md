![](docs/tuq.png)

![](docs/tng.png)

## Running

### Users

    $ go get github.com/couchbaselabs/tuqtng
    $ tuqtnq -couchbase http://localhost:8091

### Developers

    $ mkdir -p $GOPATH/src/github.com/couchbaselabs
    $ cd $GOPATH/src/github.com/couchbaselabs
    $ git clone git@github.com:couchbaselabs/tuqtng.git
    $ cd tuqtng
    $ go get code.google.com/p/go.exp/locale/collate
    $ go get github.com/gorilla/mux
    $ go build
    $ ./tuqtng -couchbase http://localhost:8091

## Querying via interactive command-line tool

    $ tuq_client
    tuq> SELECT * FROM beer-sample LIMIT 5

The tuq_client takes an optional parameter pointing to the tuqtng server:

* --tuqtng="http://localhost:8093/"

## Querying via HTTP

### HTTP Get

    curl 'http://localhost:8093/query?q=URL_ENCODED_QUERY_STRING'

### HTTP Post

    curl -HContent-Type:text/plain -XPOST http://localhost:8093/query -d 'QUERY_STRING'
