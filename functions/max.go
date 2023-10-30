package functions

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/util"
)

var _ expressions.Expression = &Max{}

// Max is a max aggregate function.
// example:
//
//	MAX(column)
type Max struct {
	Col *expressions.Column
}

func (m *Max) Write(sw io.StringWriter) error {
	if err := util.WriteStrings(sw, "MAX("); err != nil {
		return fmt.Errorf("error writing MAX function: %v", err)
	}

	if err := m.Col.Write(sw); err != nil {
		return fmt.Errorf("error writing MAX expression: %v", err)
	}

	if err := util.WriteStrings(sw, ")"); err != nil {
		return fmt.Errorf("error writing MAX closing parenthesis: %v", err)
	}

	return nil
}

// NewMax returns a new Max aggregate functions expression.
func NewMax(column *expressions.Column) *Max {
	return &Max{
		Col: column,
	}
}
