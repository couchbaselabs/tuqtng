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
	"log"
	"sort"

	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/query"
)

type Order struct {
	Source         Operator
	OrderBy        []*ast.SortExpression
	itemChannel    query.ItemChannel
	supportChannel PipelineSupportChannel
	buffer         query.ItemCollection
}

func NewOrder(orderBy []*ast.SortExpression) *Order {
	return &Order{
		OrderBy:        orderBy,
		itemChannel:    make(query.ItemChannel),
		supportChannel: make(PipelineSupportChannel),
		buffer:         make(query.ItemCollection, 0),
	}
}

func (this *Order) SetSource(source Operator) {
	this.Source = source
}

func (this *Order) GetChannels() (query.ItemChannel, PipelineSupportChannel) {
	return this.itemChannel, this.supportChannel
}

func (this *Order) Run() {
	defer close(this.itemChannel)
	defer close(this.supportChannel)

	go this.Source.Run()

	var item query.Item
	var obj interface{}
	sourceItemChannel, supportChannel := this.Source.GetChannels()
	ok := true
	for ok {
		select {
		case item, ok = <-sourceItemChannel:
			if ok {
				this.processItem(item)
			}
		case obj, ok = <-supportChannel:
			if ok {
				switch obj := obj.(type) {
				case query.Error:
					this.supportChannel <- obj
					if obj.IsFatal() {
						return
					}
				default:
					this.supportChannel <- obj
				}
			}
		}
	}

	// sort
	sort.Sort(this)

	// write the output
	for _, item := range this.buffer {
		this.itemChannel <- item
	}
}

func (this *Order) processItem(item query.Item) {
	this.buffer = append(this.buffer, item)
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
			case *query.Undefined:
			default:
				log.Printf("Error evaluating expression: %v", lerr)
				return false
			}
		}
		rightVal, rerr := oe.Expr.Evaluate(right)
		if rerr != nil {
			switch rerr := rerr.(type) {
			case *query.Undefined:
			default:
				log.Printf("Error evaluating expression: %v", rerr)
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
			// both not missing, compare values
			result := ast.CollateJSON(leftVal, rightVal)
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
