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

// Argument is a value that provides a query with parameter.
type Argument struct {
	Index int // Index starts at 1
	Value interface{}
}

// WriteTo implements Expression.
func (s *Argument) WriteTo(sw io.StringWriter) error {
	if _, err := sw.WriteString(fmt.Sprintf("$%d", s.Index)); err != nil {
		return fmt.Errorf("could not write to io.StringWriter")
	}

	return nil
}

// Arguments is a list of arguments. It keeps track of the index of the arguments.
type Arguments struct {
	mut   sync.Mutex
	Index int
	Args  []interface{}
}

func NewArguments() *Arguments {
	return &Arguments{
		Index: 0,
		Args:  []interface{}{},
	}
}

func (a *Arguments) NewArgument(value interface{}) *Argument {
	a.mut.Lock()
	defer a.mut.Unlock()

	a.Index++
	a.Args = append(a.Args, value)

	return NewArgument(value, a.Index)
}

// NewArgument creates a new argument with a given index.
func NewArgument(value interface{}, index int) *Argument {
	return &Argument{
		Index: index,
		Value: value,
	}
}
