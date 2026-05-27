// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package schema

import (
	"reflect"
)

type Converter func(string) reflect.Value

var (
	invalidValue = reflect.Value{}
	boolType     = reflect.Bool
	float32Type  = reflect.Float32
	float64Type  = reflect.Float64
	intType      = reflect.Int
	int8Type     = reflect.Int8
	int16Type    = reflect.Int16
	int32Type    = reflect.Int32
	int64Type    = reflect.Int64
	stringType   = reflect.String
	uintType     = reflect.Uint
	uint8Type    = reflect.Uint8
	uint16Type   = reflect.Uint16
	uint32Type   = reflect.Uint32
	uint64Type   = reflect.Uint64
)

// Default converters for basic types.
var builtinConverters = map[reflect.Kind]Converter{
	boolType:    convertBool,
	float32Type: convertFloat32,
	float64Type: convertFloat64,
	intType:     convertInt,
	int8Type:    convertInt8,
	int16Type:   convertInt16,
	int32Type:   convertInt32,
	int64Type:   convertInt64,
	stringType:  convertString,
	uintType:    convertUint,
	uint8Type:   convertUint8,
	uint16Type:  convertUint16,
	uint32Type:  convertUint32,
	uint64Type:  convertUint64,
}

func convertBool(value string) reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

func convertFloat32(value string) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func convertFloat64(value string) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func convertInt(value string) reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

func convertInt8(value string) reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

func convertInt16(value string) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func convertInt32(value string) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func convertInt64(value string) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func convertString(value string) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func convertUint(value string) reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

func convertUint8(value string) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func convertUint16(value string) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func convertUint32(value string) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func convertUint64(value string) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func convertPointer(k reflect.Kind, value string) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}
