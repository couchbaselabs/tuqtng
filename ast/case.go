//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package ast

import (
	"fmt"
	"log"

	"github.com/couchbaselabs/tuqtng/query"
)

type WhenThen struct {
	When Expression `json:"when"`
	Then Expression `json:"then"`
}

type CaseOperator struct {
	Type      string      `json:"type"`
	WhenThens []*WhenThen `json:"whenthens"`
	Else      Expression  `json:"else"`
}

func NewCaseOperator() *CaseOperator {
	return &CaseOperator{
		Type: "case",
	}
}

func (this *CaseOperator) Evaluate(item query.Item) (query.Value, error) {
	// walk through the WhenThens in order
	for _, WhenThen := range this.WhenThens {
		// evalute the when
		whenVal, err := WhenThen.When.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *query.Undefined:
				// MISSING is not true, keep looking
				continue
			default:
				// any other error should be returned to caller
				return nil, err
			}
		}

		whenBoolVal := ValueInBooleanContext(whenVal)
		if whenBoolVal == true {
			// evalauate the then
			thenVal, err := WhenThen.Then.Evaluate(item)
			if err != nil {
				return nil, err
			}
			return thenVal, nil
		}
	}

	// if we got this point, none of the WhenThen pairs were satisfied
	// if there was an else, evaluate it and return
	if this.Else != nil {
		elseVal, err := this.Else.Evaluate(item)
		if err != nil {
			return nil, err
		}
		return elseVal, nil
	}
	// if there was no else, return null
	return nil, nil
}

func (this *CaseOperator) Validate() error {
	for _, WhenThen := range this.WhenThens {
		err := WhenThen.When.Validate()
		if err != nil {
			return err
		}
		err = WhenThen.Then.Validate()
		if err != nil {
			return err
		}
	}
	if this.Else != nil {
		err := this.Else.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *CaseOperator) VerifyFormalNotation(aliases []string, defaultAlias string) (Expression, error) {

	for _, WhenThen := range this.WhenThens {
		newwhen, err := WhenThen.When.VerifyFormalNotation(aliases, defaultAlias)
		if err != nil {
			return nil, err
		}
		if newwhen != nil {
			WhenThen.When = newwhen
		}
		newthen, err := WhenThen.Then.VerifyFormalNotation(aliases, defaultAlias)
		if err != nil {
			return nil, err
		}
		if newthen != nil {
			WhenThen.Then = newthen
			//this.WhenThens[i] = WhenThen
		} else {
			log.Printf("not update then")
		}
	}
	if this.Else != nil {
		newelse, err := this.Else.VerifyFormalNotation(aliases, defaultAlias)
		if err != nil {
			return nil, err
		}
		if newelse != nil {
			this.Else = newelse
		}
	}
	return nil, nil
}

func (this *CaseOperator) String() string {
	inside := ""
	for i, wt := range this.WhenThens {
		if i != 0 {
			inside = inside + " "
		}
		inside = inside + fmt.Sprintf("WHEN %v THEN %v", wt.When, wt.Then)
	}
	if this.Else != nil {
		inside = inside + fmt.Sprintf(" ELSE %v", this.Else)
	}
	return fmt.Sprintf("CASE %v END", inside)
}
