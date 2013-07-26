//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package goyacc

import (
	"fmt"
	"strings"
	"sync"

	"github.com/couchbaselabs/tuqtng/ast"
)

var DebugTokens = false
var DebugGrammar = false

var parsingStack *Stack
var parsingStatement ast.Statement
var crashHard = false

type UnqlParser struct {
	mutex sync.Mutex
}

func NewUnqlParser() *UnqlParser {
	return &UnqlParser{}
}

func (u *UnqlParser) Parse(input string) (returnStatement ast.Statement, err error) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	parsingStack = new(Stack)
	parsingStatement = ast.NewSelectStatement()

	defer func() {
		r := recover()
		if r != nil && r == "syntax error" {
			// if we're panicing over a syntax error, chill
			err = fmt.Errorf("Parse Error - %v", r)
		} else if r != nil {
			// otherise continue to panic
			if crashHard {
				panic(r)
			} else {
				err = fmt.Errorf("Other Error - %v", r)
			}
		}
	}()

	yyParse(NewLexer(strings.NewReader(input)))
	returnStatement = parsingStatement
	return
}
