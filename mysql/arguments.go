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

type Argument struct {
	V interface{}
}

func (s *Argument) Value() interface{} { return s.V }

func (s *Argument) WriteTo(sw io.StringWriter) error {
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
	mutex sync.Mutex
	Args  []interface{}
}

func NewArguments() *Arguments {
	return &Arguments{
		Args: []interface{}{},
	}
}

func (a *Arguments) NewArgument(value interface{}) *Argument {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	a.Args = append(a.Args, value)

	return NewArgument(value)
}
