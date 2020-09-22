package object

import (
	"bytes"
	"fmt"
	"monkey/ast"
	"strings"
)

// ObjectType is the type for object
type ObjectType string

// Constants for object type
const (
	INTEGEROBJ     = "INTEGER"
	BOOLEANOBJ     = "BOOLEAN"
	NULLOBJ        = "NULL"
	RETURNVALUEOBJ = "RETURN_VALUE"
	ERROROBJ       = "ERROR"
	FUNCTIONOBJ    = "FUNCTION"
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

// NewEnvironment creates a new environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// Environment keep tracks of all names and maps to object
type Environment struct {
	store map[string]Object
}

// Get returns the object associated with the name
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

// Set sets env value
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
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
