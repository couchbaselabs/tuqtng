To run the tutorial...

(a) Install Go, http://golang.org/doc/install

Please be sure to export GOROOT as outlined before continuing.

(b) Run "go get github.com/couchbaselabs/tuqtng"

(c) Run the below commands:

    cd $GOROOT/src/github.com/couchbaselabs/tuqtng
    ./tuqtng -couchbase dir:test/ -pool json &
    cd tutorial; ./tutorial &

(d) Visit http://localhost:8000 with your browser.
