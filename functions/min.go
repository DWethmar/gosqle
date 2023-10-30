package functions

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/util"
)

var _ expressions.Expression = &Min{}

// Min is a min aggregate function.
// example:
//
//	MIN(column)
type Min struct {
	Col *expressions.Column
}

func (m *Min) Write(sw io.StringWriter) error {
	if err := util.WriteStrings(sw, "MIN("); err != nil {
		return fmt.Errorf("error writing MIN function: %v", err)
	}

	if err := m.Col.Write(sw); err != nil {
		return fmt.Errorf("error writing MIN expression: %v", err)
	}

	if err := util.WriteStrings(sw, ")"); err != nil {
		return fmt.Errorf("error writing MIN closing parenthesis: %v", err)
	}

	return nil
}

// NewMin returns a new Min aggregate functions expression.
func NewMin(column *expressions.Column) *Min {
	return &Min{
		Col: column,
	}
}
