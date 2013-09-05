#!/bin/sh -e

project=github.com/couchbaselabs/tuqtng
top=`go list -f '{{.Dir}}' $project`
version=`git describe`

cd $top

DIST=$top/dist

testpkg() {
    go test $project/...
    go vet $project/...
}

mkversion() {
    echo "{\"version\": \"$version\"}" > $DIST/version.json
}

build() {
    pkg=$project
    goflags="-v -ldflags '-X main.VERSION $version'"

    eval env GOARCH=386   GOOS=linux CGO_ENABLED=0 go build $goflags -o $DIST/tuqtng.lin32 $pkg &
    eval env GOARCH=arm   GOOS=linux CGO_ENABLED=0 go build $goflags -o $DIST/tuqtng.arm $pkg &
    eval env GOARCH=arm   GOARM=5 GOOS=linux CGO_ENABLED=0 go build $goflags -o $DIST/tuqtng.arm5 $pkg &
    eval env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build $goflags -o $DIST/tuqtng.lin64 $pkg &
    eval env GOARCH=amd64 GOOS=freebsd CGO_ENABLED=0 go build $goflags -o $DIST/tuqtng.fbsd $pkg &&
    eval env GOARCH=386   GOOS=windows go build $goflags -o $DIST/tuqtng.win32.exe $pkg &
    eval env GOARCH=amd64 GOOS=windows go build $goflags -o $DIST/tuqtng.win64.exe $pkg &
    eval env GOARCH=amd64 GOOS=darwin go build $goflags -o $DIST/tuqtng.mac $pkg &

    wait
}

buildclient() {
    pkg=$project/tuq_client
    goflags="-v -ldflags '-X main.VERSION $version'"

    eval env GOARCH=386   GOOS=linux CGO_ENABLED=0 go build $goflags -o $DIST/tuq_client.lin32 $pkg &
    eval env GOARCH=arm   GOOS=linux CGO_ENABLED=0 go build $goflags -o $DIST/tuq_client.arm $pkg &
    eval env GOARCH=arm   GOARM=5 GOOS=linux CGO_ENABLED=0 go build $goflags -o $DIST/tuq_client.arm5 $pkg &
    eval env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build $goflags -o $DIST/tuq_client.lin64 $pkg &
    eval env GOARCH=amd64 GOOS=freebsd CGO_ENABLED=0 go build $goflags -o $DIST/tuq_client.fbsd $pkg &&
    eval env GOARCH=386   GOOS=windows go build $goflags -o $DIST/tuq_client.win32.exe $pkg &
    eval env GOARCH=amd64 GOOS=windows go build $goflags -o $DIST/tuq_client.win64.exe $pkg &
    eval env GOARCH=amd64 GOOS=darwin go build $goflags -o $DIST/tuq_client.mac $pkg &

    wait
}

buildtutorial() {
    pkg=$project/tutorial
    goflags="-v -ldflags '-X main.VERSION $version'"

    eval env GOARCH=386   GOOS=linux CGO_ENABLED=0 go build $goflags -o $DIST/tuq_tutorial.lin32 $pkg &
    eval env GOARCH=arm   GOOS=linux CGO_ENABLED=0 go build $goflags -o $DIST/tuq_tutorial.arm $pkg &
    eval env GOARCH=arm   GOARM=5 GOOS=linux CGO_ENABLED=0 go build $goflags -o $DIST/tuq_tutorial.arm5 $pkg &
    eval env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build $goflags -o $DIST/tuq_tutorial.lin64 $pkg &
    eval env GOARCH=amd64 GOOS=freebsd CGO_ENABLED=0 go build $goflags -o $DIST/tuq_tutorial.fbsd $pkg &
    eval env GOARCH=386   GOOS=windows go build $goflags -o $DIST/tuq_tutorial.win32.exe $pkg &
    eval env GOARCH=amd64 GOOS=windows go build $goflags -o $DIST/tuq_tutorial.win64.exe $pkg &
    eval env GOARCH=amd64 GOOS=darwin go build $goflags -o $DIST/tuq_tutorial.mac $pkg &

    wait
}

builddistpackages() {

    mkdir -p $DIST/tutorial_tmp
    cd tutorial
    go build
    cd $top
    tutorial/tutorial -src tutorial/content/ -dst $DIST/tutorial_tmp/
    # FIXME see if we can eliminate this step
    cp $DIST/tutorial_tmp/tutorial/content/tutorial.html $DIST/tutorial_tmp/tutorial/content/index.html

    # mac build
    mkdir -p $DIST/stage
    cp $DIST/start_tutorial.sh $DIST/stage
    cp $DIST/tuqtng.mac $DIST/stage/tuqtng
    cp $DIST/tuq_client.mac $DIST/stage/tuq_client
    cp -r static/ $DIST/stage/static
    mkdir -p $DIST/stage/static/tutorial
    cp -r $DIST/tutorial_tmp/tutorial/content/ $DIST/stage/static/tutorial
    mkdir -p $DIST/stage/data/default/tutorial
    cp -r test/json/tutorial/ $DIST/stage/data/default/tutorial
    cd $DIST/stage
    zip $DIST/cb-query_preview_x86_64_mac.zip -r .
    cd $top
    rm -rf $DIST/stage

    #linux 32
    mkdir -p $DIST/stage
    cp $DIST/start_tutorial.sh $DIST/stage
    cp $DIST/tuqtng.lin32 $DIST/stage/tuqtng
    cp $DIST/tuq_client.lin32 $DIST/stage/tuq_client
    cp -r static/ $DIST/stage/static
    mkdir -p $DIST/stage/static/tutorial
    cp -r $DIST/tutorial_tmp/tutorial/content/ $DIST/stage/static/tutorial
    mkdir -p $DIST/stage/data/default/tutorial
    cp -r test/json/tutorial/ $DIST/stage/data/default/tutorial
    cd $DIST/stage
    tar zcvf $DIST/cb-query_preview_x86_linux.tar.gz .
    cd $top
    rm -rf $DIST/stage

    #linux 64
    mkdir -p $DIST/stage
    cp $DIST/start_tutorial.sh $DIST/stage
    cp $DIST/tuqtng.lin64 $DIST/stage/tuqtng
    cp $DIST/tuq_client.lin64 $DIST/stage/tuq_client
    cp -r static/ $DIST/stage/static
    mkdir -p $DIST/stage/static/tutorial
    cp -r $DIST/tutorial_tmp/tutorial/content/ $DIST/stage/static/tutorial
    mkdir -p $DIST/stage/data/default/tutorial
    cp -r test/json/tutorial/ $DIST/stage/data/default/tutorial
    cd $DIST/stage
    tar zcvf $DIST/cb-query_preview_x86_64_linux.tar.gz .
    cd $top
    rm -rf $DIST/stage

    # windows packages commented out as build seems to be failing

    #win 32
    # mkdir -p $DIST/stage
    # cp $DIST/tuqtng.win32.exe $DIST/stage/tuqtng.exe
    # cp $DIST/tuq_client.win32.exe $DIST/stage/tuq_client.exe
    # cp -r static/ $DIST/stage/static
    # mkdir -p $DIST/stage/static/tutorial
    # cp -r $DIST/tutorial_tmp/tutorial/content/ $DIST/stage/static/tutorial
    # mkdir -p $DIST/stage/data/default/tutorial
    # cp -r test/json/tutorial/ $DIST/stage/data/default/tutorial
    # cd $DIST/stage
    # zip $DIST/cb-query_preview_x86_win.zip -r .
    # cd $top
    # rm -rf $DIST/stage

    #win 64
    # mkdir -p $DIST/stage
    # cp $DIST/tuqtng.win64.exe $DIST/stage/tuqtng.exe
    # cp $DIST/tuq_client.win64 $DIST/stage/tuq_client.exe
    # cp -r static/ $DIST/stage/static
    # mkdir -p $DIST/stage/static/tutorial
    # cp -r $DIST/tutorial_tmp/tutorial/content/ $DIST/stage/static/tutorial
    # mkdir -p $DIST/stage/data/default/tutorial
    # cp -r test/json/tutorial/ $DIST/stage/data/default/tutorial
    # cd $DIST/stage
    # zip $DIST/cb-query_preview_x86_64_win.zip -r .
    # cd $top
    # rm -rf $DIST/stage

    rm -rf $DIST/tutorial_tmp
    
}

compress() {
    rm -f $DIST/tuqtng.*.gz $DIST/tuq_client.*.gz $DIST/tuq_tutorial.*.gz || true

    for i in lin32 arm arm5 lin64 fbsd win32 win64 mac
    do
        if [ -f $DIST/tuq_tutorial.$i ]; then
            tar zcf $DIST/tuq_tutorial.$i.tar.gz    \
              -C $top test/json/tutorial \
              -C $top/tutorial content \
              -C $DIST tuq_tutorial.$i
            rm -f $DIST/tuq_tutorial.$i
        fi &
    done

    for i in $DIST/tuqtng.* $DIST/tuq_client.*
    do
        gzip -9v $i &
    done

    wait
}

benchmark() {
    go test -test.bench . > $DIST/benchmark.txt
}

coverage() {
    for sub in ast misc plan test xpipeline
    do
        gocov test $project/$sub | gocov-html > $DIST/cov-$sub.html
    done
    cd $top/test
    gocov test -deps -exclude-goroot > $DIST/integ.json
    cat $DIST/integ.json | jq '{"Packages": [.Packages[] | if .Name > "github.com/couchbaselabs/tuqtng" and .Name < "github.com/couchbaselabs/tuqtnh" then . else empty end]}' > $DIST/integ2.json
    cat $DIST/integ2.json |gocov-html > $DIST/integ-cov.html
    cd $top
}

upload() {
    cbfsclient ${cbfsserver:-http://cbfs.hq.couchbase.com:8484/} upload \
        -ignore=$DIST/.cbfsclient.ignore -delete -v \
        $DIST/ tuqtng/

    cbfsclient ${cbfsserver:-http://cbfs.hq.couchbase.com:8484/} upload \
        -ignore=$DIST/.cbfsclient.ignore -delete -v \
        $DIST/redirect.html tuqtng
}

testpkg
mkversion
build
buildclient
buildtutorial
builddistpackages
compress
benchmark
coverage
upload
