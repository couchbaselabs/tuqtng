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
	"fmt"
	"os"
	"os/signal"
	"os/user"
	"syscall"

	"github.com/couchbaselabs/clog"
	"github.com/sbinet/liner"
)

func HandleInteractiveMode(tiServer string) {

	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Unable to determine current user, history file disabled: %v", err)
	}

	var liner = liner.NewLiner()
	defer liner.Close()

	LoadHistory(liner, currentUser)

	go signalCatcher(liner)

	for {
		line, err := liner.Prompt("tuq> ")
		if err != nil {
			break
		}

		if line == "" {
			continue
		}

		UpdateHistory(liner, currentUser, line)
		err = execute_internal(tiServer, line, os.Stdout)
		if err != nil {
			clog.Error(err)
		}
	}

}

/**
 *  Attempt to clean up after ctrl-C otherwise
 *  terminal is left in bad shape
 */
func signalCatcher(liner *liner.State) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT)
	<-ch
	liner.Close()
	os.Exit(0)
}
