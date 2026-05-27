package schema

import (
	"reflect"
)

type encoderFunc func(reflect.Value) string

// Encoder encodes values from a struct into url.Values.
type Encoder struct {
	cache  *cache
	regenc map[reflect.Type]encoderFunc
}

// NewEncoder returns a new Encoder with defaults.
func NewEncoder() *Encoder { _ = "STUB: not implemented"; return nil }

// Encode encodes a struct into map[string][]string.
//
// Intended for use with url.Values.
func (e *Encoder) Encode(src interface{}, dst map[string][]string) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterEncoder registers a converter for encoding a custom type.
func (e *Encoder) RegisterEncoder(value interface{}, encoder func(reflect.Value) string) {
	_ = "STUB: not implemented"
	return
}

// SetAliasTag changes the tag used to locate custom field aliases.
// The default tag is "schema".
func (e *Encoder) SetAliasTag(tag string) {
	_ = "STUB: not implemented"

	// isValidStructPointer test if input value is a valid struct pointer.
	return
}

func isValidStructPointer(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

func isZero(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

// Compare other types directly:

func (e *Encoder) encode(v reflect.Value, dst map[string][]string) error {
	_ = "STUB: not implemented"
	return nil
}

// Encode struct pointer types if the field is a valid pointer and a struct.

// Encode non-slice types and custom implementations immediately.

// Encode a slice.

func (e *Encoder) hasCustomEncoder(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func typeEncoder(t reflect.Type, reg map[reflect.Type]encoderFunc) encoderFunc {
	_ = "STUB: not implemented"
	return *new(encoderFunc)
}

func encodeBool(v reflect.Value) string { _ = "STUB: not implemented"; return "" }

func encodeInt(v reflect.Value) string { _ = "STUB: not implemented"; return "" }

func encodeUint(v reflect.Value) string { _ = "STUB: not implemented"; return "" }

func encodeFloat(v reflect.Value, bits int) string { _ = "STUB: not implemented"; return "" }

func encodeFloat32(v reflect.Value) string { _ = "STUB: not implemented"; return "" }

func encodeFloat64(v reflect.Value) string { _ = "STUB: not implemented"; return "" }

func encodeString(v reflect.Value) string { _ = "STUB: not implemented"; return "" }
