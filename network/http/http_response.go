//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/couchbaselabs/tuqtng/misc"
	"github.com/couchbaselabs/tuqtng/query"
)

type HttpResponse struct {
	query    *HttpQuery
	w        http.ResponseWriter
	results  chan interface{}
	warnings []query.Error
	info     []query.Error
	err      query.Error
	count    int
}

func (this *HttpResponse) SendError(err query.Error) {
	switch err.Level() {
	case query.EXCEPTION:
		this.err = err
	case query.WARNING:
		this.warnings = append(this.warnings, err)
	case query.INFO:
		this.info = append(this.info, err)
	}

	if err.IsFatal() {
		close(this.results)
	}
}

func (this *HttpResponse) SendResult(val interface{}) {
	this.results <- val
}

func (this *HttpResponse) NoMoreResults() {
	close(this.results)
}

func (this *HttpResponse) Process() error {

	_, err := this.openResponse()
	if err != nil {
		return err
	}

	_, err = this.ProcessResults()
	if err != nil {
		return err
	}

	if this.err == nil {
		_, err = this.ProcessWarningsAndInfo()
		if err != nil {
			return err
		}
	} else {
		// an error occured
		_, err = this.ProcessError()
		if err != nil {
			return err
		}
	}

	_, err = this.closeResponse()
	if err != nil {
		return err
	}

	return nil
}

func (this *HttpResponse) ProcessResults() (int, error) {
	for val := range this.results {
		if this.count == 0 {
			_, err := this.openArray("resultset")
			if err != nil {
				return 0, err
			}
		} else {
			_, err := this.continueResponse()
			if err != nil {
				return 0, err
			}
		}
		valSanitized := misc.SanitizeUnrepresentableJSON(val)
		_, err := this.printObj(valSanitized)
		if err != nil {
			return 0, err
		}
		this.count++

		// flush response so client receives rows as soon as possible
		if f, ok := this.w.(http.Flusher); ok {
			f.Flush()
		}
	}

	// close resultset

	if this.count == 0 && this.err == nil {
		_, err := this.openArray("resultset")
		if err != nil {
			return 0, err
		}
		_, err = this.closeArray()
		if err != nil {
			return 0, err
		}
	} else if this.count != 0 {
		_, err := this.continueResponseLast()
		if err != nil {
			return 0, err
		}
		_, err = this.closeArray()
		if err != nil {
			return 0, err
		}
	}

	return 0, nil
}

func (this *HttpResponse) ProcessWarningsAndInfo() (int, error) {
	_, err := this.continueResponse()
	if err != nil {
		return 0, err
	}

	if len(this.warnings) > 0 {
		_, err := this.printArrayOfErrors("warnings", this.warnings)
		if err != nil {
			return 0, err
		}
	}

	// forcibly generate some infos
	this.SendError(query.NewTotalRowsInfo(this.count))
	elapsed_duration := this.query.Duration()
	this.SendError(query.NewTotalElapsedTimeInfo(elapsed_duration.String()))

	if len(this.info) > 0 {
		_, err := this.printArrayOfErrors("info", this.info)
		if err != nil {
			return 0, err
		}
	}

	return 0, nil
}

func (this *HttpResponse) ProcessError() (int, error) {
	if this.count != 0 {
		_, err := this.continueResponse()
		if err != nil {
			return 0, err
		}
	}
	return this.printError(this.err)
}

func (this *HttpResponse) openResponse() (int, error) {
	return fmt.Fprint(this.w, "{\n")
}

func (this *HttpResponse) closeResponse() (int, error) {
	return fmt.Fprint(this.w, "\n}\n")
}

func (this *HttpResponse) continueResponse() (int, error) {
	return fmt.Fprint(this.w, ",\n")
}

func (this *HttpResponse) continueResponseLast() (int, error) {
	return fmt.Fprint(this.w, "\n")
}

func (this *HttpResponse) openArray(name string) (int, error) {
	return fmt.Fprint(this.w, "    \"", name, "\": [\n")
}

func (this *HttpResponse) openKey(name string, indent string) (int, error) {
	return fmt.Fprint(this.w, indent, "\"", name, "\":\n")
}

func (this *HttpResponse) closeArray() (int, error) {
	return fmt.Fprint(this.w, "    ]")
}

func (this *HttpResponse) printObj(obj interface{}) (int, error) {
	objBytes, err := json.MarshalIndent(obj, "        ", "    ")
	if err != nil {
		return 0, err
	}
	return fmt.Fprint(this.w, "        ", string(objBytes))
}

func (this *HttpResponse) printError(err query.Error) (int, error) {
	this.openKey("error", "    ")
	return this.printObj(err)
}

func (this *HttpResponse) printArrayOfErrors(name string, arr []query.Error) (int, error) {
	this.openArray(name)
	for i, item := range arr {
		if i != 0 {
			this.continueResponse()
		}
		this.printObj(item)
	}
	this.continueResponseLast()
	this.closeArray()
	return 0, nil
}
