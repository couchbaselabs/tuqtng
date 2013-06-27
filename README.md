![](docs/tuq.png)

![](docs/tng.png)

## Running

### Users

    go get github.com/couchbaselabs/tuqtng
    ./tuqtnq

### Developers

    mkdir -p $GOPATH/src/github.com/couchbaselabs
    git clone git@github.com:couchbaselabs/tuqtng.git
    cd tuqtng
    go build
    ./tuqtng

## Querying

### HTTP Get

    curl 'http://localhost:8093/query?q=URL_ENCODED_QUERY_STRING'

### HTTP Post

    curl -HContent-Type:text/plain -XPOST http://localhost:8093/query -d 'QUERY_STRING'
