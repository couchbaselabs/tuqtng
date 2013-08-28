//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package simple

import (
	"strings"

	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/plan"
)

// determines if the given expression is sargable with respect to a given index expression
// for example
// is abv > 7 sargable with respect to an index on abv? yes
type ExpressionSargable struct {
	indexExpression    ast.Expression
	equivalenceChecker *ast.ExpressionEquivalenceChecker
	constantChecker    *ast.ExpressionFunctionalDependencyChecker
	sargable           bool
	scanRanges         plan.ScanRanges
}

func NewExpressionSargable(indexExpression ast.Expression) *ExpressionSargable {
	return &ExpressionSargable{
		indexExpression:    indexExpression,
		equivalenceChecker: ast.NewExpressionEquivalenceChecker(ast.ExpressionList{indexExpression}),
		constantChecker:    ast.NewExpressionFunctionalDependencyCheckerFull(ast.ExpressionList{}),
		sargable:           false,
		scanRanges:         make(plan.ScanRanges, 0),
	}
}

func (this *ExpressionSargable) IsSargable() bool {
	return this.sargable
}

func (this *ExpressionSargable) ScanRanges() plan.ScanRanges {
	return this.scanRanges
}

func (this *ExpressionSargable) Visit(e ast.Expression) (ast.Expression, error) {
	switch e := e.(type) {
	case *ast.GreaterThanOperator:
		if this.matchesIndex(e.Left) && this.isConstant(e.Right) {
			val, err := this.evaluateConstant(e.Right)
			if err != nil {
				return nil, err
			}
			this.sargable = true
			newScanRange := &plan.ScanRange{Low: catalog.LookupValue{val}, Inclusion: catalog.High}
			this.scanRanges = append(this.scanRanges, newScanRange)
		} else if this.matchesIndex(e.Right) && this.isConstant(e.Left) {
			val, err := this.evaluateConstant(e.Left)
			if err != nil {
				return nil, err
			}
			this.sargable = true
			newScanRange := &plan.ScanRange{High: catalog.LookupValue{val}, Inclusion: catalog.Low}
			this.scanRanges = append(this.scanRanges, newScanRange)
		}
	case *ast.GreaterThanOrEqualOperator:
		if this.matchesIndex(e.Left) && this.isConstant(e.Right) {
			val, err := this.evaluateConstant(e.Right)
			if err != nil {
				return nil, err
			}
			this.sargable = true
			newScanRange := &plan.ScanRange{Low: catalog.LookupValue{val}, Inclusion: catalog.Both}
			this.scanRanges = append(this.scanRanges, newScanRange)
		} else if this.matchesIndex(e.Right) && this.isConstant(e.Left) {
			val, err := this.evaluateConstant(e.Left)
			if err != nil {
				return nil, err
			}
			this.sargable = true
			newScanRange := &plan.ScanRange{High: catalog.LookupValue{val}, Inclusion: catalog.Both}
			this.scanRanges = append(this.scanRanges, newScanRange)
		}
	case *ast.LessThanOperator:
		if this.matchesIndex(e.Left) && this.isConstant(e.Right) {
			val, err := this.evaluateConstant(e.Right)
			if err != nil {
				return nil, err
			}
			this.sargable = true
			newScanRange := &plan.ScanRange{High: catalog.LookupValue{val}, Inclusion: catalog.Low}
			this.scanRanges = append(this.scanRanges, newScanRange)
		} else if this.matchesIndex(e.Right) && this.isConstant(e.Left) {
			val, err := this.evaluateConstant(e.Left)
			if err != nil {
				return nil, err
			}
			this.sargable = true
			newScanRange := &plan.ScanRange{Low: catalog.LookupValue{val}, Inclusion: catalog.High}
			this.scanRanges = append(this.scanRanges, newScanRange)
		}
	case *ast.LessThanOrEqualOperator:
		if this.matchesIndex(e.Left) && this.isConstant(e.Right) {
			val, err := this.evaluateConstant(e.Right)
			if err != nil {
				return nil, err
			}
			this.sargable = true
			newScanRange := &plan.ScanRange{High: catalog.LookupValue{val}, Inclusion: catalog.Both}
			this.scanRanges = append(this.scanRanges, newScanRange)
		} else if this.matchesIndex(e.Right) && this.isConstant(e.Left) {
			val, err := this.evaluateConstant(e.Left)
			if err != nil {
				return nil, err
			}
			this.sargable = true
			newScanRange := &plan.ScanRange{Low: catalog.LookupValue{val}, Inclusion: catalog.Both}
			this.scanRanges = append(this.scanRanges, newScanRange)
		}
	case *ast.EqualToOperator:
		if this.matchesIndex(e.Left) && this.isConstant(e.Right) {
			val, err := this.evaluateConstant(e.Right)
			if err != nil {
				return nil, err
			}
			this.sargable = true
			newScanRange := &plan.ScanRange{Low: catalog.LookupValue{val}, High: catalog.LookupValue{val}, Inclusion: catalog.Both}
			this.scanRanges = append(this.scanRanges, newScanRange)
		} else if this.matchesIndex(e.Right) && this.isConstant(e.Left) {
			val, err := this.evaluateConstant(e.Left)
			if err != nil {
				return nil, err
			}
			this.sargable = true
			newScanRange := &plan.ScanRange{Low: catalog.LookupValue{val}, High: catalog.LookupValue{val}, Inclusion: catalog.Both}
			this.scanRanges = append(this.scanRanges, newScanRange)
		}
	case *ast.NotEqualToOperator:
		if this.matchesIndex(e.Left) && this.isConstant(e.Right) {
			val, err := this.evaluateConstant(e.Right)
			if err != nil {
				return nil, err
			}
			this.sargable = true
			newScanRange := &plan.ScanRange{High: catalog.LookupValue{val}, Inclusion: catalog.Low}
			this.scanRanges = append(this.scanRanges, newScanRange)
			newScanRange = &plan.ScanRange{Low: catalog.LookupValue{val}, Inclusion: catalog.High}
			this.scanRanges = append(this.scanRanges, newScanRange)
		} else if this.matchesIndex(e.Right) && this.isConstant(e.Left) {
			val, err := this.evaluateConstant(e.Left)
			if err != nil {
				return nil, err
			}
			this.sargable = true
			newScanRange := &plan.ScanRange{High: catalog.LookupValue{val}, Inclusion: catalog.Low}
			this.scanRanges = append(this.scanRanges, newScanRange)
			newScanRange = &plan.ScanRange{Low: catalog.LookupValue{val}, Inclusion: catalog.High}
			this.scanRanges = append(this.scanRanges, newScanRange)
		}

	case *ast.LikeOperator:
		if this.matchesIndex(e.Left) && this.isConstant(e.Right) {
			val, err := this.evaluateConstant(e.Right)
			if err != nil {
				return nil, err
			}
			pattern, ok := val.Value().(string)
			if ok {
				startKey, endKey := this.computePatternRange(pattern)
				this.sargable = true
				newScanRange := &plan.ScanRange{
					Low:       catalog.LookupValue{dparval.NewValue(startKey)},
					High:      catalog.LookupValue{dparval.NewValue(endKey)},
					Inclusion: catalog.Low}
				this.scanRanges = append(this.scanRanges, newScanRange)
			}
		}

	case *ast.NotLikeOperator:
		if this.matchesIndex(e.Left) && this.isConstant(e.Right) {
			val, err := this.evaluateConstant(e.Right)
			if err != nil {
				return nil, err
			}
			pattern, ok := val.Value().(string)
			if ok {
				startKey, endKey := this.computePatternRange(pattern)
				this.sargable = true
				newScanRange := &plan.ScanRange{
					High:      catalog.LookupValue{dparval.NewValue(startKey)},
					Inclusion: catalog.Low}
				this.scanRanges = append(this.scanRanges, newScanRange)
				newScanRange = &plan.ScanRange{
					Low:       catalog.LookupValue{dparval.NewValue(endKey)},
					Inclusion: catalog.High}
				this.scanRanges = append(this.scanRanges, newScanRange)
			}
		}
	}

	return e, nil
}

func (this *ExpressionSargable) matchesIndex(e ast.Expression) bool {
	_, err := e.Accept(this.equivalenceChecker)
	if err == nil {
		return true
	}
	return false
}

func (this *ExpressionSargable) isConstant(e ast.Expression) bool {
	_, err := e.Accept(this.constantChecker)
	if err == nil {
		return true
	}
	return false
}

func (this *ExpressionSargable) evaluateConstant(e ast.Expression) (*dparval.Value, error) {
	return e.Evaluate(dparval.NewValue(map[string]interface{}{}))
}

func (this *ExpressionSargable) computePatternRange(pattern string) (string, string) {
	patternParts := strings.SplitN(pattern, "%", 2)
	patternParts = strings.SplitN(patternParts[0], "_", 2)
	patternPrefix := patternParts[0]
	patternPrefixEndBytes := []byte(patternPrefix)
	patternPrefixEndBytes[len(patternPrefixEndBytes)-1] = patternPrefixEndBytes[len(patternPrefixEndBytes)-1] + 1
	patternPrefixEnd := string(patternPrefixEndBytes)
	return patternPrefix, patternPrefixEnd
}
