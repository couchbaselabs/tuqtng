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
)

type From struct {
	Pool       string
	Bucket     string
	Projection Expression
	As         string
	Over       *From
}

func (this *From) GetAliases() []string {
	// important, keep the order correct
	// top-down, first alias is bucket
	rv := []string{this.As}
	if this.Over != nil {
		otherAliases := this.Over.GetAliases()
		for _, alias := range otherAliases {
			rv = append(rv, alias)
		}
	}
	return rv
}

func (this *From) GenerateAlias() {
	// if there was no alias
	// try to generate one
	if this.As == "" {
		switch proj := this.Projection.(type) {
		case nil:
			// empty projection
			// bucket name is the alias
			// FROM bucket
			// becomes FROM bucket AS bucket
			this.As = this.Bucket
		case *Property:
			// in this case the property path is the alias
			// FROM bucket.prop
			// becomes FROM bucket.prop AS prop
			this.As = proj.Path
		case *DotMemberOperator:
			// in this case the right-most property path is the alias
			// FROM bucket.propa.propb
			// becomes FROM bucket.propa.propb AS propb
			this.As = proj.Right.Path
		case *BracketMemberOperator:
		case *BracketSliceMemberOperator:
			// we decided to NOT assign an alias for these type at this time
		default:
			panic(fmt.Sprintf("unexpected type %T in FROM", proj))
		}
		// in all other cases there is no alias generated
	}
	// if there is an over, reccurse this call
	if this.Over != nil {
		this.Over.GenerateAlias()
	}
}

// FROM is a generic structure capturing both top-level FROMs and OVER constructs
// in top level FROMs the start of the Projection path is the bucket name
// this should be set correct and also removed from the Projection
// for example:
// FROM person.friends OVER contacts AS contact
// the top-level FROM should have bucket "person" and projection "friends"
// the second level FROM has no bucket and projection "contacts"
func (this *From) ConvertToBucketFrom() {
	if this.Bucket != "" {
		// already converted
		return
	}
	// walk the Projection Expression
	// there are only 4 valid types in this limited case
	// dot memeber, bracket memeber, and property
	// for dot/slice + bracket member we keep walking to the left
	// until we eventually get the base property
	var prevprev Expression
	var prev Expression
	var curr Expression = this.Projection
	found := false
	for !found {
		switch curri := curr.(type) {
		case *DotMemberOperator:
			prevprev = prev
			prev = curri
			curr = curri.Left
		case *BracketMemberOperator:
			prevprev = prev
			prev = curri
			curr = curri.Left
		case *BracketSliceMemberOperator:
			prevprev = prev
			prev = curri
			curr = curri.Left
		case *Property:
			this.Bucket = curri.Path
			switch prev := prev.(type) {
			case nil:
				// if there was no previous node
				// then it was a single property
				// this became the bucket, set to nil
				this.Projection = nil
			case *BracketMemberOperator:
				// find the RHS we need push up
				rhs := prev.Right
				switch prevprev := prevprev.(type) {
				case nil:
					// if there was no previous previous node
					// then this RHS is now the projection
					this.Projection = NewBracketMemberOperator(NewProperty(curri.Path), rhs)
				case *BracketMemberOperator:
					// replace the LHS of the previous previous
					prevprev.Left = NewBracketMemberOperator(NewProperty(curri.Path), rhs)
				case *DotMemberOperator:
					// replace the LHS of the previous previous
					prevprev.Left = NewBracketMemberOperator(NewProperty(curri.Path), rhs)
				}
			case *DotMemberOperator:
				// find the RHS we need push up
				rhs := prev.Right
				switch prevprev := prevprev.(type) {
				case nil:
					// if there was no previous previous node
					// then this RHS is now the projection
					this.Projection = rhs
				case *BracketMemberOperator:
					// replace the LHS of the previous previous
					prevprev.Left = rhs
				case *DotMemberOperator:
					// replace the LHS of the previous previous
					prevprev.Left = rhs
				case *BracketSliceMemberOperator:
					// replace the LHS of the previous previous
					prevprev.Left = rhs
				}

			}

			found = true
		}
	}
}
