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
	"log"

	"github.com/couchbaselabs/tuqtng/network"
	"github.com/couchbaselabs/tuqtng/network/http"
	"github.com/couchbaselabs/tuqtng/qpipeline/static"
)

var addr = flag.String("addr", ":8093", "HTTP listen address")

func main() {

	// create a StaticQueryPipeline we use to process queries
	queryPipeline := static.NewStaticPipeline()

	// create a QueryChannel
	queryChannel := make(network.QueryChannel)

	// create one or more network endpoints
	httpEndpoint := http.NewHttpEndpoint(*addr)
	httpEndpoint.SendQueriesTo(queryChannel)

	log.Printf("tuqtng started...")

	// dispatch each query that comes in
	for query := range queryChannel {
		log.Printf("got a query %v", query)
		queryPipeline.DispatchQuery(query)
	}
}
