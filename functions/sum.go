package functions

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/util"
)

var _ expressions.Expression = &Sum{}

// Sum is a sum aggregate function.
// example:
//
//	SUM(column)
type Sum struct {
	Col *expressions.Column
}

func (s *Sum) Write(sw io.StringWriter) error {
	if err := util.WriteStrings(sw, "SUM("); err != nil {
		return fmt.Errorf("error writing SUM function: %v", err)
	}

	if err := s.Col.Write(sw); err != nil {
		return fmt.Errorf("error writing SUM expression: %v", err)
	}

	if err := util.WriteStrings(sw, ")"); err != nil {
		return fmt.Errorf("error writing SUM closing parenthesis: %v", err)
	}

	return nil
}

// NewSum returns a new Sum aggregate functions expression.
func NewSum(column *expressions.Column) *Sum {
	return &Sum{
		Col: column,
	}
}
