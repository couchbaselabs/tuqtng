//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package server

import (
	"fmt"
	"strings"

	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/catalog/couchbase"
	"github.com/couchbaselabs/tuqtng/catalog/file"
	"github.com/couchbaselabs/tuqtng/catalog/mock"
	"github.com/couchbaselabs/tuqtng/network"
	"github.com/couchbaselabs/tuqtng/qpipeline/static"

	"github.com/couchbaselabs/clog"
)

func Site(s string) (catalog.Site, error) {
	if strings.HasPrefix(s, ".") || strings.HasPrefix(s, "/") {
		return file.NewSite(s)
	}
	if strings.HasPrefix(s, "dir:") {
		return file.NewSite(s[4:])
	}
	if strings.HasPrefix(s, "mock:") {
		return mock.NewSite(s)
	}
	return couchbase.NewSite(s)
}

func Server(version, couchbaseSite, poolName string,
	queryChannel network.QueryChannel) error {
	site, err := Site(couchbaseSite)
	if err != nil {
		return fmt.Errorf("Unable to access site %s, err: %v", couchbaseSite, err)
	}

	var pool catalog.Pool
	if site != nil {
		pool, err = site.Pool(poolName)
		if err != nil {
			return fmt.Errorf("Unable to access pool %v in the site: %v", poolName, err)
		}
	}

	// create a StaticQueryPipeline we use to process queries
	queryPipeline := static.NewStaticPipeline(pool)

	clog.Log("tuqtng started...")
	clog.Log("version: %s", version)
	clog.Log("site: %s", couchbaseSite)

	// dispatch each query that comes in
	for query := range queryChannel {
		go queryPipeline.DispatchQuery(query)
	}

	return nil
}
