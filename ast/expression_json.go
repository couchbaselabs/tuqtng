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
	"encoding/json"
	"fmt"
)

func UnmarshalExpression(bytes []byte) (Expression, error) {
	var temp struct {
		Type string `json:"type"`
	}

	err := json.Unmarshal(bytes, &temp)
	if err != nil {
		return nil, err
	}

	switch temp.Type {
	case "literal_number":
		var numb LiteralNumber
		err := json.Unmarshal(bytes, &numb)
		if err != nil {
			return nil, err
		}
		return &numb, nil
	case "property":
		var prop Property
		err := json.Unmarshal(bytes, &prop)
		if err != nil {
			return nil, err
		}
		return &prop, nil
	case "dot_member":
		var dm DotMemberOperator
		err := json.Unmarshal(bytes, &dm)
		if err != nil {
			return nil, err
		}
		return &dm, nil
	case "bracket_member":
		var bm BracketMemberOperator
		err := json.Unmarshal(bytes, &bm)
		if err != nil {
			return nil, err
		}
		return &bm, nil
	case "bracket_slice_member":
		var bsm BracketSliceMemberOperator
		err := json.Unmarshal(bytes, &bsm)
		if err != nil {
			return nil, err
		}
		return &bsm, nil
	default:
		panic(fmt.Sprintf("Unable to Unmarshall this type of Expression: %v", temp.Type))
	}
}
