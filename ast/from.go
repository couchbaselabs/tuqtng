//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package ast

type From struct {
	Bucket     string
	Projection Expression
	As         string
}

// FROM is a generic structure capturing both top-level FROMs and OVER constructs
// in top level FROMs the start of the Projection path is the bucket name
// this should be set correct and also removed from the Projection
// for example:
// FROM person.friends OVER contacts AS contact
// the top-level FROM should have bucket "person" and projection "friends"
// the second level FROM has no bucket and projection "contacts"
func (this *From) ConvertToBucketFrom() {
	// walk the Projection Expression
	// there are only 3 valid types in this limited cae
	// dot memeber, bracket memeber, and property
	// for dot/bracket member we keep walking to the left
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
		case *Property:
			this.Bucket = curri.Path
			switch prev := prev.(type) {
			case nil:
				// if there was no previous node
				// then it was a single property
				// this became the bucket
				// so now set projection to nil
				this.Projection = nil
			case *BracketMemberOperator:
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
				}
			}
			found = true
		}
	}
}
