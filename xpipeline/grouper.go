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
	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/query"
)

type Grouper struct {
	Base       *BaseOperator
	GroupBy    ast.ExpressionList
	groups     map[string]*dparval.Value
	Aggregates ast.ExpressionList
	groupAll   bool
}

func NewGrouper(groupBy ast.ExpressionList, aggs ast.ExpressionList) *Grouper {
	rv := &Grouper{
		Base:       NewBaseOperator(),
		GroupBy:    groupBy,
		groups:     make(map[string]*dparval.Value),
		Aggregates: aggs,
	}
	if len(groupBy) == 0 {
		// group all behavior
		rv.GroupBy = ast.ExpressionList{ast.NewLiteralBool(true)}
		rv.groupAll = true
	}
	return rv
}

func (this *Grouper) SetSource(source Operator) {
	this.Base.SetSource(source)
}

func (this *Grouper) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.Base.GetChannels()
}

func (this *Grouper) Run() {
	this.Base.RunOperator(this)
}

func (this *Grouper) processItem(item *dparval.Value) bool {
	groupkey := dparval.NewValue(make([]interface{}, len(this.GroupBy)))
	for i, groupElement := range this.GroupBy {
		groupkeyval, err := groupElement.Evaluate(item)
		if err == nil {
			groupkey.SetIndex(i, groupkeyval)
		} else {
			switch err := err.(type) {
			case *dparval.Undefined:
				// FIXME better way?
				groupkey.SetIndex(i, "__tuqtng__MISSING__")
			default:
				return this.Base.SendError(query.NewError(err, "error evaluating group by"))
			}
		}
	}
	// FIXME slow, but lets me use map to match same groups
	groupkeybytes := groupkey.Bytes()
	groupkeystring := string(groupkeybytes)

	group, ok := this.groups[groupkeystring]
	if !ok {
		// new group
		this.groups[groupkeystring] = item
		group = item
		this.setGroupDefaults(group)
	}
	this.updateGroup(group, item)
	return true
}

func (this *Grouper) setGroupDefaults(group *dparval.Value) {
	// here we should ask each aggregate function to initialize default values for the group
	for _, agg := range this.Aggregates {
		switch agg := agg.(type) {
		case ast.AggregateFunctionCallExpression:
			err := agg.DefaultAggregate(group)
			if err != nil {
				this.Base.SendError(query.NewError(err, "error setting group defaults"))
			}
		}
	}
}

func (this *Grouper) updateGroup(group *dparval.Value, val *dparval.Value) {
	// here we should update aggregate functions with this new item
	for _, agg := range this.Aggregates {
		switch agg := agg.(type) {
		case ast.AggregateFunctionCallExpression:
			err := agg.UpdateAggregate(group, val)
			if err != nil {
				this.Base.SendError(query.NewError(err, "error updating group info"))
			}
		}
	}
}

func (this *Grouper) afterItems() {
	for _, group := range this.groups {
		this.Base.SendItem(group)
	}

	if this.groupAll && len(this.groups) == 0 {
		// need to report correctly on the empty group
		group := dparval.NewValue(map[string]interface{}{})
		this.setGroupDefaults(group)
		this.Base.SendItem(group)
	}
}
