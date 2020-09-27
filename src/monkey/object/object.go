package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"monkey/ast"
	"strings"
)

// BuiltinFunction is the type defintion of callable Go Func
type BuiltinFunction func(args ...Object) Object

// ObjectType is the type for object
type ObjectType string

// Constants for object type
const (
	NULLOBJ  = "NULL"
	ERROROBJ = "ERROR"

	INTEGEROBJ = "INTEGER"
	BOOLEANOBJ = "BOOLEAN"
	STRINGOBJ  = "STRING"

	RETURNVALUEOBJ = "RETURN_VALUE"

	FUNCTIONOBJ = "FUNCTION"
	BUILTINOBJ  = "BUILTIN"

	ARRAYOBJ = "ARRAY"
	HASHOBJ  = "HASH"
)

// Object provides the object functions
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Hashable interface for hash object implementation
type Hashable interface {
	HashKey() HashKey
}

// HashKey defines the hash structure
type HashKey struct {
	Type  ObjectType
	Value uint64
}

// HashKey returns the boolean HashKey
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

// HashKey returns the integer HashKey
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// HashKey returns the string HashKey
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// Integer is for int data type
type Integer struct {
	Value int64
}

// Inspect provides the integer value repr
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Type returns the object type
func (i *Integer) Type() ObjectType {
	return INTEGEROBJ
}

// Boolean is for bool object type
type Boolean struct {
	Value bool
}

// Type returns the boolean type
func (b *Boolean) Type() ObjectType {
	return BOOLEANOBJ
}

// Inspect returns the boolean value repr
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// Null is for nil type
type Null struct{}

// Type returns the null type
func (n *Null) Type() ObjectType {
	return NULLOBJ
}

// Inspect returns the null value repr
func (n *Null) Inspect() string {
	return "null"
}

// ReturnValue for return
type ReturnValue struct {
	Value Object
}

// Type returns Return Type
func (rv *ReturnValue) Type() ObjectType {
	return RETURNVALUEOBJ
}

// Inspect returns the value of return
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// Error represents error
type Error struct {
	Message string
}

// Type returns error type
func (e *Error) Type() ObjectType {
	return ERROROBJ
}

// Inspect returns error message
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

// Function type for func
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type returns the func type
func (f *Function) Type() ObjectType {
	return FUNCTIONOBJ
}

// Inspect returns the func repr
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

// String object
type String struct {
	Value string
}

// Type returns string object type
func (s *String) Type() ObjectType {
	return STRINGOBJ
}

// Inspect returns the string value
func (s *String) Inspect() string {
	return s.Value
}

// Builtin is the wrapper for Builtin Functions
type Builtin struct {
	Fn BuiltinFunction
}

// Type returns the Builtin Type
func (b *Builtin) Type() ObjectType {
	return BUILTINOBJ
}

// Inspect returns the Builtin repr
func (b *Builtin) Inspect() string {
	return "builtin function"
}

// Array is the wrapper for array
type Array struct {
	Elements []Object
}

// Type returns the array type
func (ao *Array) Type() ObjectType {
	return ARRAYOBJ
}

// Inspect returns the array repr
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// HashPair defines hash pair structure
type HashPair struct {
	Key   Object
	Value Object
}

// Hash defines the wrapper for hash
type Hash struct {
	Pairs map[HashKey]HashPair
}

// Type for hash object
func (h *Hash) Type() ObjectType {
	return HASHOBJ
}

// Inspect for hash repr
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
