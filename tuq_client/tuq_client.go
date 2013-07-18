//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	cbgb "github.com/couchbaselabs/cbgb"

	tnetwork "github.com/couchbaselabs/tuqtng/network"
	thttp "github.com/couchbaselabs/tuqtng/network/http"
	tserver "github.com/couchbaselabs/tuqtng/server"
)

var tiServer = flag.String("tuqtng", "http://localhost:8093/", "URL to tuqtng")

func main() {
	flag.Usage = usage
	flag.Parse()

	tiServer := *tiServer
	if tiServer == "play" {
		cbgb.Main("embedded cbgb")
		time.Sleep(2 * time.Second) // cbgb needs some time to start.
		go tuqtng_play(":8093", "http://localhost:8091")
		time.Sleep(1 * time.Second) // tuqtng needs some time to start.
		tiServer = "http://localhost:8093/"
	}

	HandleInteractiveMode(tiServer)
}

func usage() {
	fmt.Fprintf(os.Stderr, "%s <flags>\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\nflags:\n")
	fmt.Fprintf(os.Stderr, "  --tuqtng=URL\n")
	fmt.Fprintf(os.Stderr, "    example: --tuqtng=http://localhost:8093/\n")
	fmt.Fprintf(os.Stderr, "    if the URL is \"play\", then tuq_client runs a local\n")
	fmt.Fprintf(os.Stderr, "    embedded server for developer playing convenience\n")
}

func tuqtng_play(addr, couchbase string) {
	// create a QueryChannel
	queryChannel := make(tnetwork.QueryChannel)

	// create one or more network endpoints
	httpEndpoint := thttp.NewHttpEndpoint(addr)
	httpEndpoint.SendQueriesTo(queryChannel)

	err := tserver.Server(couchbase, "default", queryChannel)
	if err != nil {
		log.Fatalf("Unable to run server, err: %v", err)
	}
}

func execute_internal(tiServer, line string, w io.Writer) error {

	url := tiServer + "query"
	resp, err := http.Post(url, "text/plain", strings.NewReader(line))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)

	return nil
}
