package object

import (
	"fmt"
)

// ObjectType is the type for object
type ObjectType string

// Constants for object type
const (
	INTEGEROBJ = "INTEGER"
	BOOLEANOBJ = "BOOLEAN"
	NULLOBJ    = "NULL"
)

// Object provides the object functions
type Object interface {
	Type() ObjectType
	Inspect() string
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
