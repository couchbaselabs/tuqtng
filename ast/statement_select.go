//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package ast

import ()

type SelectStatement struct {
	Distinct                  bool                 `json:"distinct"`
	Select                    ResultExpressionList `json:"select"`
	From                      *From                `json:"from"`
	Where                     Expression           `json:"where"`
	GroupBy                   ExpressionList       `json:"group_by"`
	Having                    Expression           `json:"having"`
	OrderBy                   SortExpressionList   `json:"orderby"`
	Limit                     int                  `json:"limit"`
	Offset                    int                  `json:"offset"`
	ExplainOnly               bool                 `json:"explain"`
	explicitProjectionAliases []string
	aggregateReferences       ExpressionList
}

func NewSelectStatement() *SelectStatement {
	return &SelectStatement{
		Limit: -1,
	}
}

func (this *SelectStatement) GetFromAliases() []string {
	if this.From != nil {
		return this.From.GetAliases()
	}
	return []string{}
}

func (this *SelectStatement) GetFrom() *From {
	return this.From
}

func (this *SelectStatement) GetWhere() Expression {
	return this.Where
}

func (this *SelectStatement) GetGroupBy() ExpressionList {
	return this.GroupBy
}

func (this *SelectStatement) GetHaving() Expression {
	return this.Having
}

func (this *SelectStatement) GetOrderBy() SortExpressionList {
	return this.OrderBy
}

func (this *SelectStatement) GetLimit() int {
	return this.Limit
}

func (this *SelectStatement) GetOffset() int {
	return this.Offset
}

func (this *SelectStatement) IsDistinct() bool {
	return this.Distinct
}

func (this *SelectStatement) GetResultExpressionList() ResultExpressionList {
	return this.Select
}

func (this *SelectStatement) SetExplainOnly(explainOnly bool) {
	this.ExplainOnly = explainOnly
}

func (this *SelectStatement) IsExplainOnly() bool {
	return this.ExplainOnly
}

func (this *SelectStatement) GetExplicitProjectionAliases() []string {
	return this.explicitProjectionAliases
}

func (this *SelectStatement) GetAggregateReferences() ExpressionList {
	return this.aggregateReferences
}

func (this *SelectStatement) IsAggregate() bool {
	if this.GroupBy != nil {
		return true
	}

	// now check the projection for any aggregate functions
	return this.Select.ContainsAggregateFunctionCall()
}

func (this *SelectStatement) VerifySemantics() error {

	var err error
	// get the list of explicit projection aliases, and check it for duplicates
	this.explicitProjectionAliases, err = this.Select.CheckForDuplicateAliases()
	if err != nil {
		return err
	}

	// now apply default naming function
	err = this.Select.AssignDefaultNames(this.explicitProjectionAliases)
	if err != nil {
		return err
	}

	// fix up all the froms (FIXME need to refactor this into FROM/OVER clenaup)
	if this.From != nil {
		this.From.ConvertToBucketFrom()
		this.From.GenerateAlias()
	}

	// verify formal notations
	err = this.verifyFormalNotation(this.explicitProjectionAliases)
	if err != nil {
		return err
	}

	// validate the expressions
	err = this.validate()
	if err != nil {
		return err
	}

	// if this is an aggregate query we need to perform some additional checks
	if this.IsAggregate() {
		// separate two cases, with group by and without
		if this.GroupBy == nil {
			// this means an aggregate function was used, but there was no group by
			// in this case, all projection expressions MUST be aggregate function calls
			err := this.Select.VerifyAllAggregateFunctionsOrInThisList(ExpressionList{})
			if err != nil {
				return err
			}
			if this.Having != nil {
				fdc := NewExpressionFunctionalDependencyChecker(ExpressionList{})
				_, err := this.Having.Accept(fdc)
				if err != nil {
					return err
				}
			}
			if this.OrderBy != nil {
				err := this.OrderBy.VerifyAllAggregateFunctionsOrInThisList(ExpressionList{})
				if err != nil {
					return err
				}
			}
			// set group by to an empty expresison list
			// the group will use this to behave correctly
			this.GroupBy = ExpressionList{}
		} else {
			// ensure that all projection expressions are either aggregate functipn calls
			// or in the group by expression list
			err := this.Select.VerifyAllAggregateFunctionsOrInThisList(this.GroupBy)
			if err != nil {
				return err
			}
			if this.Having != nil {
				fdc := NewExpressionFunctionalDependencyChecker(this.GroupBy)
				_, err := this.Having.Accept(fdc)
				if err != nil {
					return err
				}
			}
			if this.OrderBy != nil {
				err := this.OrderBy.VerifyAllAggregateFunctionsOrInThisList(this.GroupBy)
				if err != nil {
					return err
				}
			}
		}

		// gather the aggregate references and store these
		this.aggregateReferences = this.findAggregateFunctionReferences()
	}

	// if you combine DISTINCT with ORDER BY we have to do an additional validation
	if this.Distinct && this.OrderBy != nil {
		// every Order By expression MUST be equivalent to one in the projection
		err := this.OrderBy.VerifyAllEquivalentToThisList(this.Select.ExpressionList())
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *SelectStatement) validate() error {

	// validate the projection
	err := this.Select.Validate()
	if err != nil {
		return err
	}

	// validate the where
	if this.Where != nil {
		whereValidator := NewExpressionValidatorNoAggregates()
		this.Where, err = this.Where.Accept(whereValidator)
		if err != nil {
			return err
		}
	}

	// validate the order by
	err = this.OrderBy.Validate()
	if err != nil {
		return err
	}

	// validate the group by
	if this.GroupBy != nil {
		err = this.GroupBy.Validate()
		if err != nil {
			return err
		}
	}

	// validate the having
	if this.Having != nil {
		havingValidator := NewExpressionValidator()
		this.Having, err = this.Having.Accept(havingValidator)
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *SelectStatement) verifyFormalNotation(explicitProjectionAliases []string) error {

	// gather the list of aliases
	aliases := this.GetFromAliases()
	defaultAlias := ""
	if len(aliases) == 1 {
		defaultAlias = aliases[0]
	}

	formalNotation := NewExpressionFormalNotationConverter(explicitProjectionAliases, aliases, defaultAlias)

	// verify the projection (references to projection aliases not allowed)
	err := this.Select.VerifyFormalNotation(explicitProjectionAliases, aliases, defaultAlias)
	if err != nil {
		return err
	}

	// verify the where (references to projection aliases not allowed)
	if this.Where != nil {
		this.Where, err = this.Where.Accept(formalNotation)
		if err != nil {
			return err
		}
	}

	// verify the order by(references to projection aliases ARE allowed)
	// since order by CAN reference the explicit aliases, these must be added to the list
	// passed into this phase
	orderByAliases := make([]string, 0)
	for _, alias := range aliases {
		orderByAliases = append(orderByAliases, alias)
	}
	for _, alias := range explicitProjectionAliases {
		orderByAliases = append(orderByAliases, alias)
	}
	err = this.OrderBy.VerifyFormalNotation([]string{}, orderByAliases, defaultAlias)
	if err != nil {
		return err
	}

	// verify the group by (references to projection aliases not allowed)
	if this.GroupBy != nil {
		err = this.GroupBy.VerifyFormalNotation(explicitProjectionAliases, aliases, defaultAlias)
		if err != nil {
			return err
		}
	}

	// verify the having (references to projection aliases not allowed)
	if this.Having != nil {
		this.Having, err = this.Having.Accept(formalNotation)
		if err != nil {
			return err
		}
	}

	return nil
}

// SELECT, HAVING, and ORDER BY may all reference aggregate functions
// we need to find these so that we can produce the appropriate stats
// during the grouping phase
func (this *SelectStatement) findAggregateFunctionReferences() ExpressionList {
	ar := this.Select.findAggregateFunctionReferences()

	if this.Having != nil {
		af := NewExpressionAggregateFinder()
		this.Having.Accept(af)
		havingRefs := af.GetAggregates()
		if len(havingRefs) > 0 {
			ar = append(ar, havingRefs...)
		}
	}

	if this.OrderBy != nil {
		orderRefs := this.OrderBy.findAggregateFunctionReferences()
		if len(orderRefs) > 0 {
			ar = append(ar, orderRefs...)
		}
	}

	//finally we need to remove duplicates from this list
	dar := make(map[string]Expression)
	for _, agg := range ar {
		dar[agg.String()] = agg
	}

	rv := make(ExpressionList, 0)
	for _, agg := range dar {
		rv = append(rv, agg)
	}

	return rv
}

func (this *SelectStatement) Simplify() error {
	err := this.Select.Simplify()
	if err != nil {
		return err
	}

	if this.Where != nil {
		es := NewExpressionSimplifier()
		this.Where, err = this.Where.Accept(es)
		if err != nil {
			return err
		}
	}

	// validate the order by
	err = this.OrderBy.Simplify()
	if err != nil {
		return err
	}

	// validate the group by
	if this.GroupBy != nil {
		err = this.GroupBy.Simplify()
		if err != nil {
			return err
		}
	}

	// validate the having
	if this.Having != nil {
		es := NewExpressionSimplifier()
		this.Having, err = this.Having.Accept(es)
		if err != nil {
			return err
		}
	}

	return nil
}
