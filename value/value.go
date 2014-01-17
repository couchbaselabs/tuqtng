//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package value

import (
	"fmt"
	"strconv"

	jsonpointer "github.com/dustin/go-jsonpointer"
	json "github.com/dustin/gojson"
)

// When you try to access a nested property or index that does not exist,
// the return value will be nil, and the return error will be Undefined.
type Undefined string

// Description of which property or index was undefined (if known).
func (this Undefined) Error() string {
	if string(this) != "" {
		return fmt.Sprintf("%s is not defined.", string(this))
	}
	return "Not defined."
}

var _UNDEFINED = Undefined("")

const _MARSHAL_ERROR = "Unexpected marshal error on valid data."

// A channel of Value objects
type ValueChannel chan Value

// A collection of Value objects
type ValueCollection []Value

// An interface for storing and manipulating a (possibly JSON) value.
type Value interface {
	Type() int
	Actual() interface{}
	Duplicate() Value
	Bytes() []byte
	Field(field string) (Value, error)
	SetField(field string, val interface{})
	Index(index int) (Value, error)
	SetIndex(index int, val interface{})
}

type AnnotatedValue interface {
	Value
	GetAttachment(key string) interface{}
	SetAttachment(key string, val interface{})
	RemoveAttachment(key string) interface{}
}

// Bring a data object into the Value type system
func NewValue(val interface{}) Value {
	switch val := val.(type) {
	case Value:
		return val
	case float64:
		return floatValue(val)
	case string:
		return stringValue(val)
	case bool:
		return boolValue(val)
	case nil:
		return nullValue{}
	case []interface{}:
		return &arrayValue{val}
	case map[string]interface{}:
		return objectValue(val)
	default:
		panic(fmt.Sprintf("Cannot create value for type %T", val))
	}
}

// Create a new Value from a slice of bytes. (this need not be valid JSON)
func NewValueFromBytes(bytes []byte) Value {
	var parsedType int
	err := json.Validate(bytes)

	if err == nil {
		parsedType = identifyType(bytes)

		switch parsedType {
		case NUMBER, STRING, BOOLEAN, NULL:
			var p interface{}
			err := json.Unmarshal(bytes, &p)
			if err != nil {
				panic("Unexpected parse error on valid JSON")
			}

			return NewValue(p)
		}
	}

	rv := parsedValue{
		raw: bytes,
	}

	if err != nil {
		rv.parsedType = NOT_JSON
	} else {
		rv.parsedType = parsedType
	}

	return &rv
}

// Create an AnnotatedValue to hold attachments
func NewAnnotatedValue(val interface{}) AnnotatedValue {
	switch val := val.(type) {
	case AnnotatedValue:
		return val
	case Value:
		av := annotatedValue{
			Value:    val,
			attacher: attacher{nil},
		}

		return &av
	case []byte:
		return NewAnnotatedValue(NewValueFromBytes(val))
	default:
		return NewAnnotatedValue(NewValue(val))
	}
}

// The data types supported by Value
const (
	NOT_JSON = iota
	NULL
	BOOLEAN
	NUMBER
	STRING
	ARRAY
	OBJECT
)

type floatValue float64

func (this floatValue) Type() int {
	return NUMBER
}

func (this floatValue) Actual() interface{} {
	return float64(this)
}

func (this floatValue) Duplicate() Value {
	return this
}

func (this floatValue) Bytes() []byte {
	bytes, err := json.Marshal(this.Actual())
	if err != nil {
		panic(_MARSHAL_ERROR)
	}
	return bytes
}

func (this floatValue) Field(field string) (Value, error) {
	return nil, Undefined(field)
}

func (this floatValue) SetField(field string, val interface{}) {
}

func (this floatValue) Index(index int) (Value, error) {
	return nil, _UNDEFINED
}

func (this floatValue) SetIndex(index int, val interface{}) {
}

type stringValue string

func (this stringValue) Type() int {
	return STRING
}

func (this stringValue) Actual() interface{} {
	return string(this)
}

func (this stringValue) Duplicate() Value {
	return this
}

func (this stringValue) Bytes() []byte {
	bytes, err := json.Marshal(this.Actual())
	if err != nil {
		panic(_MARSHAL_ERROR)
	}
	return bytes
}

func (this stringValue) Field(field string) (Value, error) {
	return nil, Undefined(field)
}

func (this stringValue) SetField(field string, val interface{}) {
}

func (this stringValue) Index(index int) (Value, error) {
	return nil, _UNDEFINED
}

func (this stringValue) SetIndex(index int, val interface{}) {
}

type boolValue bool

func (this boolValue) Type() int {
	return BOOLEAN
}

func (this boolValue) Actual() interface{} {
	return bool(this)
}

func (this boolValue) Duplicate() Value {
	return this
}

func (this boolValue) Bytes() []byte {
	bytes, err := json.Marshal(this.Actual())
	if err != nil {
		panic(_MARSHAL_ERROR)
	}
	return bytes
}

func (this boolValue) Field(field string) (Value, error) {
	return nil, Undefined(field)
}

func (this boolValue) SetField(field string, val interface{}) {
}

func (this boolValue) Index(index int) (Value, error) {
	return nil, _UNDEFINED
}

func (this boolValue) SetIndex(index int, val interface{}) {
}

type nullValue struct {
}

func (this nullValue) Type() int {
	return NULL
}

func (this nullValue) Actual() interface{} {
	return nil
}

func (this nullValue) Duplicate() Value {
	return this
}

func (this nullValue) Bytes() []byte {
	bytes, err := json.Marshal(this.Actual())
	if err != nil {
		panic(_MARSHAL_ERROR)
	}
	return bytes
}

func (this nullValue) Field(field string) (Value, error) {
	return nil, Undefined(field)
}

func (this nullValue) SetField(field string, val interface{}) {
}

func (this nullValue) Index(index int) (Value, error) {
	return nil, _UNDEFINED
}

func (this nullValue) SetIndex(index int, val interface{}) {
}

type arrayValue struct {
	actual []interface{}
}

func (this *arrayValue) Type() int {
	return ARRAY
}

func (this *arrayValue) Actual() interface{} {
	return this.actual
}

func (this *arrayValue) Duplicate() Value {
	return &arrayValue{duplicateSlice(this.actual)}
}

func (this *arrayValue) Bytes() []byte {
	bytes, err := json.Marshal(this.Actual())
	if err != nil {
		panic(_MARSHAL_ERROR)
	}
	return bytes
}

func (this *arrayValue) Field(field string) (Value, error) {
	return nil, Undefined(field)
}

func (this *arrayValue) SetField(field string, val interface{}) {
}

func (this *arrayValue) Index(index int) (Value, error) {
	if index >= 0 && index < len(this.actual) {
		return NewValue(this.actual[index]), nil
	}

	// consistent with parsedValue
	return nil, _UNDEFINED
}

func (this *arrayValue) SetIndex(index int, val interface{}) {
	if index >= 0 && index < len(this.actual) {
		this.actual[index] = val
	} else if index >= 0 {
		av := make([]interface{}, index+1)
		copy(av, this.actual)
		av[index] = val
		this.actual = av
	}
}

type objectValue map[string]interface{}

func (this objectValue) Type() int {
	return OBJECT
}

func (this objectValue) Actual() interface{} {
	return (map[string]interface{})(this)
}

func (this objectValue) Duplicate() Value {
	return objectValue(duplicateMap(this))
}

func (this objectValue) Bytes() []byte {
	bytes, err := json.Marshal(this.Actual())
	if err != nil {
		panic(_MARSHAL_ERROR)
	}
	return bytes
}

func (this objectValue) Field(field string) (Value, error) {
	result, ok := this[field]
	if ok {
		return NewValue(result), nil
	}

	// consistent with parsedValue
	return nil, _UNDEFINED
}

func (this objectValue) SetField(field string, val interface{}) {
	this[field] = val
}

func (this objectValue) Index(index int) (Value, error) {
	return nil, _UNDEFINED
}

func (this objectValue) SetIndex(index int, val interface{}) {
}

// A structure for storing and manipulating a (possibly JSON) value.
type parsedValue struct {
	raw        []byte
	parsedType int
	parsed     Value
}

func (this *parsedValue) Type() int {
	return this.parsedType
}

func (this *parsedValue) Actual() interface{} {
	if this.parsedType == NOT_JSON {
		return nil
	}

	return this.parse().Actual()
}

func (this *parsedValue) Duplicate() Value {
	rv := parsedValue{
		raw:        this.raw,
		parsedType: this.parsedType,
	}

	if this.parsed != nil {
		rv.parsed = this.parsed.Duplicate()
	}

	return &rv
}

func (this *parsedValue) Bytes() []byte {
	switch this.parsedType {
	case ARRAY, OBJECT:
		return this.parse().Bytes()
	default:
		return this.raw
	}
}

func (this *parsedValue) Field(field string) (Value, error) {
	if this.parsed != nil {
		return this.parsed.Field(field)
	}

	if this.parsedType != OBJECT {
		return nil, Undefined(field)
	}

	if this.raw != nil {
		res, err := jsonpointer.Find(this.raw, "/"+field)
		if err != nil {
			return nil, err
		}
		if res != nil {
			return NewValueFromBytes(res), nil
		}
	}

	return nil, Undefined(field)
}

func (this *parsedValue) SetField(field string, val interface{}) {
	if this.parsedType != OBJECT {
		return
	}

	this.parse().SetField(field, val)
}

func (this *parsedValue) Index(index int) (Value, error) {
	if this.parsed != nil {
		return this.parsed.Index(index)
	}

	if this.parsedType != ARRAY {
		return nil, _UNDEFINED
	}

	if this.raw != nil {
		res, err := jsonpointer.Find(this.raw, "/"+strconv.Itoa(index))
		if err != nil {
			return nil, err
		}
		if res != nil {
			return NewValueFromBytes(res), nil
		}
	}

	return nil, _UNDEFINED
}

func (this *parsedValue) SetIndex(index int, val interface{}) {
	if this.parsedType != ARRAY {
		return
	}

	this.parse().SetIndex(index, val)
}

func (this *parsedValue) parse() Value {
	if this.parsed == nil {
		if this.parsedType == NOT_JSON {
			return nil
		}

		var p interface{}
		err := json.Unmarshal(this.raw, &p)
		if err != nil {
			panic("Unexpected parse error on valid JSON")
		}
		this.parsed = NewValue(p)
	}

	return this.parsed
}

func duplicateSlice(source []interface{}) []interface{} {
	if source == nil {
		return nil
	}

	result := make([]interface{}, len(source))
	for i, v := range source {
		result[i] = NewValue(v).Duplicate()
	}

	return result
}

func duplicateMap(source map[string]interface{}) map[string]interface{} {
	if source == nil {
		return nil
	}

	result := make(map[string]interface{}, len(source))
	for k, v := range source {
		result[k] = NewValue(v).Duplicate()
	}

	return result
}

func copyMap(source map[string]interface{}) map[string]interface{} {
	if source == nil {
		return nil
	}

	result := make(map[string]interface{}, len(source))
	for k, v := range source {
		result[k] = v
	}

	return result
}

func identifyType(bytes []byte) int {
	for _, b := range bytes {
		switch b {
		case '{':
			return OBJECT
		case '[':
			return ARRAY
		case '"':
			return STRING
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return NUMBER
		case 't', 'f':
			return BOOLEAN
		case 'n':
			return NULL
		}
	}
	panic("Unable to identify type of valid JSON")
	return -1
}

type annotatedValue struct {
	Value
	attacher
}

func (this *annotatedValue) Duplicate() Value {
	av := annotatedValue{
		Value:    this.Value.Duplicate(),
		attacher: attacher{copyMap(this.attacher.attachments)},
	}

	return &av
}

type attacher struct {
	attachments map[string]interface{}
}

// Return the object attached to this Value with this key.
// If no object is attached with this key, nil is returned.
func (this *attacher) GetAttachment(key string) interface{} {
	if this.attachments != nil {
		return this.attachments[key]
	}
	return nil
}

// Attach an arbitrary object to this Value with the specified key.
// Any existing value attached with this same key will be overwritten.
func (this *attacher) SetAttachment(key string, val interface{}) {
	if this.attachments == nil {
		this.attachments = make(map[string]interface{})
	}
	this.attachments[key] = val
}

// Remove an object attached to this Value with this key.
// If there had been an object attached to this Value with this key it is returned, otherwise nil.
func (this *attacher) RemoveAttachment(key string) interface{} {
	var rv interface{}
	if this.attachments != nil {
		rv = this.attachments[key]
		delete(this.attachments, key)
	}
	return rv
}
