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
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"

	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/network"
	"github.com/couchbaselabs/tuqtng/network/http"
	"github.com/couchbaselabs/tuqtng/server"
)

var VERSION = "0.0.0" // Build-time overriddable.

var addr = flag.String("addr", ":8093", "HTTP listen address")
var couchbaseSite = flag.String("couchbase", "", "Couchbase Cluster Address (http://...) or dir:PATH")
var defaultPoolName = flag.String("pool", "default", "Default Pool")
var logKeys = flag.String("log", "", "Log keywords, comma separated")
var devMode = flag.Bool("dev", false, "Developer Mode")
var profileMode = flag.Bool("profile", false, "Profile Mode")

func main() {
	flag.Parse()

	clog.ParseLogFlag(*logKeys)

	if *devMode {
		ast.EnableDeveloperFunctions()
		clog.Log("Developer Mode Enabled")
	}

	go dumpOnSignal(syscall.SIGUSR2)

	// create a QueryChannel
	queryChannel := make(network.QueryChannel)

	// create one or more network endpoints
	httpEndpoint := http.NewHttpEndpoint(*addr, *profileMode)
	httpEndpoint.SendQueriesTo(queryChannel)

	err := server.Server(VERSION, *couchbaseSite, *defaultPoolName, queryChannel)
	if err != nil {
		clog.Fatalf("Unable to run server, err: %v", err)
	}
}

func dumpOnSignal(signals ...os.Signal) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, signals...)
	for _ = range c {
		pprof.Lookup("goroutine").WriteTo(os.Stderr, 1)
	}
}
