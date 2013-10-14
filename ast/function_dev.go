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

	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/dparval"
)

// these functions should NOT be registered by default
// they should only be registered in developer mode

func EnableDeveloperFunctions() {
	SystemFunctionRegistry["ENABLE_LOG"] = NewFunctionCallEnableLog
	SystemFunctionRegistry["DISABLE_LOG"] = NewFunctionCallDisableLog
	SystemFunctionRegistry["ERROR"] = NewFunctionCallError
	SystemFunctionRegistry["PANIC"] = NewFunctionCallPanic
}

func DisableDeveloperFunctions() {
	delete(SystemFunctionRegistry, "ENABLE_LOG")
	delete(SystemFunctionRegistry, "DISABLE_LOG")
	delete(SystemFunctionRegistry, "ERROR")
	delete(SystemFunctionRegistry, "PANIC")
}

type FunctionCallEnableLog struct {
	FunctionCall
}

func NewFunctionCallEnableLog(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallEnableLog{
		FunctionCall{
			Type:     "function",
			Name:     "ENABLE_LOG",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallEnableLog) Copy() Expression {
	return &FunctionCallEnableLog{
		FunctionCall{
			Type:     "function",
			Name:     "ENABLE_LOG",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallEnableLog) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := this.Operands[0].Expr.Evaluate(item)

	// the spec defines this functin to ONLY operate on numeric values
	// all other types result in NULL
	if err != nil {
		return nil, err
	}

	if av.Type() == dparval.STRING {
		avalue := av.Value()
		switch avalue := avalue.(type) {
		case string:
			clog.EnableKey(avalue)
		}
	}

	return dparval.NewValue(true), nil
}

func (this *FunctionCallEnableLog) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallDisableLog struct {
	FunctionCall
}

func NewFunctionCallDisableLog(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallDisableLog{
		FunctionCall{
			Type:     "function",
			Name:     "DISABLE_LOG",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallDisableLog) Copy() Expression {
	return &FunctionCallDisableLog{
		FunctionCall{
			Type:     "function",
			Name:     "DISABLE_LOG",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallDisableLog) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := this.Operands[0].Expr.Evaluate(item)

	// the spec defines this functin to ONLY operate on numeric values
	// all other types result in NULL
	if err != nil {
		return nil, err
	}

	if av.Type() == dparval.STRING {
		avalue := av.Value()
		switch avalue := avalue.(type) {
		case string:
			clog.DisableKey(avalue)
		}
	}

	return dparval.NewValue(true), nil
}

func (this *FunctionCallDisableLog) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallError struct {
	FunctionCall
}

func NewFunctionCallError(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallError{
		FunctionCall{
			Type:     "function",
			Name:     "ERROR",
			Operands: operands,
			minArgs:  -1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallError) Copy() Expression {
	return &FunctionCallError{
		FunctionCall{
			Type:     "function",
			Name:     "ERROR",
			Operands: this.Operands.Copy(),
			minArgs:  -1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallError) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// if there was an argument, see if it evaluates to true
	if len(this.Operands) == 1 {

		// first evaluate the argument
		av, err := this.Operands[0].Expr.Evaluate(item)

		// the spec defines this functin to ONLY operate on numeric values
		// all other types result in NULL
		if err != nil {
			return nil, err
		}

		boolVal := ValueInBooleanContext(av.Value())
		if boolVal == true {
			return nil, fmt.Errorf("ERROR() called, argument evaluated true")
		}

		return dparval.NewValue(boolVal), nil
	}
	// otherwise just error
	return nil, fmt.Errorf("ERROR() called")
}

func (this *FunctionCallError) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallPanic struct {
	FunctionCall
}

func NewFunctionCallPanic(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallPanic{
		FunctionCall{
			Type:     "function",
			Name:     "PANIC",
			Operands: operands,
			minArgs:  -1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallPanic) Copy() Expression {
	return &FunctionCallPanic{
		FunctionCall{
			Type:     "function",
			Name:     "PANIC",
			Operands: this.Operands.Copy(),
			minArgs:  -1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallPanic) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// if there was an argument, see if it evaluates to true
	if len(this.Operands) == 1 {

		// first evaluate the argument
		av, err := this.Operands[0].Expr.Evaluate(item)

		// the spec defines this functin to ONLY operate on numeric values
		// all other types result in NULL
		if err != nil {
			return nil, err
		}

		boolVal := ValueInBooleanContext(av.Value())
		if boolVal == true {
			panic("PANIC() called, argument evaluated true")
		}

		return dparval.NewValue(boolVal), nil
	}
	// otherwise just panic
	panic("PANIC() called")
	return dparval.NewValue(true), nil
}

func (this *FunctionCallPanic) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
