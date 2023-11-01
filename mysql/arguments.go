package mysql

import (
	"fmt"
	"io"
	"sync"

	"github.com/dwethmar/gosqle/expressions"
)

var (
	_ expressions.Expression = &Argument{}
)

// Argument is a single argument for a mysql statement.
type Argument struct {
	V interface{}
}

func (s *Argument) Value() interface{} { return s.V }

func (s *Argument) Write(sw io.StringWriter) error {
	if _, err := sw.WriteString("?"); err != nil {
		return fmt.Errorf("could not write to io.StringWriter")
	}

	return nil
}

func NewArgument(value interface{}) *Argument {
	return &Argument{
		V: value,
	}
}

// Arguments is a list of arguments.
type Arguments struct {
	mutex  sync.Mutex
	Values []interface{}
}

func NewArguments() *Arguments {
	return &Arguments{
		Values: []interface{}{},
	}
}

func (a *Arguments) NewArgument(value interface{}) *Argument {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	a.Values = append(a.Values, value)

	return NewArgument(value)
}
