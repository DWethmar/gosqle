package postgres

import (
	"fmt"
	"io"
	"sync"

	"github.com/dwethmar/gosqle/expressions"
)

var (
	_ expressions.Expression = &Argument{}
)

// Argument is a value that provides a query with a parameter.
type Argument struct {
	args  *Arguments // A reference to the Arguments collection that created this argument.
	Index int
	Value interface{}
}

// WriteTo writes the argument to the StringWriter and records the write order.
func (a *Argument) Write(sw io.StringWriter) error {
	// if part of a set of arguments, write the index and record the write order.
	if a.Index == 0 && a.args != nil {
		a.args.mut.Lock()
		defer a.args.mut.Unlock()

		a.args.writeOrder = append(a.args.writeOrder, a)
		a.Index = len(a.args.writeOrder)
	}

	if _, err := sw.WriteString(fmt.Sprintf("$%d", a.Index)); err != nil {
		return fmt.Errorf("could not write to io.StringWriter: %w", err)
	}

	return nil
}

// Arguments is a list of arguments. It keeps track of the order in which Write was called.
type Arguments struct {
	mut        sync.Mutex
	writeOrder []*Argument // A slice to keep track of the write order.
}

func NewArguments() *Arguments {
	return &Arguments{
		writeOrder: []*Argument{},
	}
}

// Create adds a new argument to the collection and returns it.
func (a *Arguments) Create(value interface{}) *Argument {
	a.mut.Lock()
	defer a.mut.Unlock()
	arg := &Argument{
		args:  a,
		Value: value,
	}
	return arg
}

// Values returns the values in the order that Write on the arguments was called.
func (a *Arguments) Values() []interface{} {
	a.mut.Lock()
	defer a.mut.Unlock()
	orderedValues := make([]interface{}, len(a.writeOrder))
	for i, arg := range a.writeOrder {
		orderedValues[i] = arg.Value
	}
	return orderedValues
}

// NewArgument creates a new argument.
func NewArgument(index int, value interface{}) *Argument {
	return &Argument{
		Index: index,
		Value: value,
	}
}
