//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package xpipeline

import (
	"sort"

	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/misc"
)

type Order struct {
	Base            *BaseOperator
	OrderBy         []*ast.SortExpression
	buffer          dparval.ValueCollection
	explicitAliases []string
}

func NewOrder(orderBy []*ast.SortExpression, explicitAliases []string) *Order {
	return &Order{
		Base:            NewBaseOperator(),
		OrderBy:         orderBy,
		buffer:          make(dparval.ValueCollection, 0),
		explicitAliases: explicitAliases,
	}
}

func (this *Order) SetSource(source Operator) {
	this.Base.SetSource(source)
}

func (this *Order) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.Base.GetChannels()
}

func (this *Order) Run(stopChannel misc.StopChannel) {
	this.Base.RunOperator(this, stopChannel)
}

func (this *Order) processItem(item *dparval.Value) bool {
	if this.explicitAliases != nil {
		projection := item.GetAttachment("projection")
		projectionValue, ok := projection.(*dparval.Value)
		if ok {
			for _, explicitAlias := range this.explicitAliases {
				// put the explicit alias values from the projection into the item
				aliasedProjectedValue, err := projectionValue.Path(explicitAlias)
				if err == nil {
					item.SetPath(explicitAlias, aliasedProjectedValue)
				}
			}
		}
	}
	this.buffer = append(this.buffer, item)
	return true
}

func (this *Order) afterItems() {
	// sort
	sort.Sort(this)

	// write the output
	for _, item := range this.buffer {
		this.Base.SendItem(item)
	}
}

// sort.Interface interface

func (this *Order) Len() int      { return len(this.buffer) }
func (this *Order) Swap(i, j int) { this.buffer[i], this.buffer[j] = this.buffer[j], this.buffer[i] }
func (this *Order) Less(i, j int) bool {
	left := this.buffer[i]
	right := this.buffer[j]

	for _, oe := range this.OrderBy {
		leftVal, lerr := oe.Expr.Evaluate(left)
		if lerr != nil {
			switch lerr := lerr.(type) {
			case *dparval.Undefined:
			default:
				clog.Error(lerr)
				return false
			}
		}
		rightVal, rerr := oe.Expr.Evaluate(right)
		if rerr != nil {
			switch rerr := rerr.(type) {
			case *dparval.Undefined:
			default:
				clog.Error(rerr)
				return false
			}
		}

		// at this point, the only errors left should be MISSING/UNDEFINED
		if oe.Ascending && lerr != nil && rerr == nil {
			// ascending, left missing, right not, left is less
			return true
		} else if !oe.Ascending && rerr != nil && lerr == nil {
			// descending right missing, left not, left is more
			return true
		} else if !oe.Ascending && lerr != nil && rerr == nil {
			// descending, left missing, right not, left is less
			return false
		} else if oe.Ascending && rerr != nil && lerr == nil {
			//ascending, left not, left is more
			return false
		} else if lerr == nil && rerr == nil {
			lv := leftVal.Value()
			rv := rightVal.Value()

			// both not missing, compare values
			result := ast.CollateJSON(lv, rv)
			if result != 0 {
				if oe.Ascending && result < 0 {
					return true
				} else if !oe.Ascending && result > 0 {
					return true
				} else {
					return false
				}
			}
		}
		// at this level they are the same, keep going
	}

	// if we go to this point the order expressions could not differentiate between the elements
	return false
}
