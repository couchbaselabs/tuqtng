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
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/tuqtng/parser/goyacc"
)

var debugLevel = flag.Int("debug", 0, "yacc debug level")

func main() {

	flag.Parse()

	clog.EnableKey("PARSER")
	clog.EnableKey("SCANNER")

	parser := goyacc.NewN1qlParserWithDebug(*debugLevel)

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		if line == "" {
			continue
		}

		ast, err := parser.Parse(line)
		if err != nil {
			clog.Error(err)
		} else {
			fmt.Printf("%v\n", ast)
		}
	}

}
