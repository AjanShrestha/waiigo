package object

// NewEnclosedEnvironment returns env that encapsulates the inner env
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// NewEnvironment creates a new environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// Environment keep tracks of all names and maps to object
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get returns the object associated with the name
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set sets env value
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
