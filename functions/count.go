package functions

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/util"
)

var _ expressions.Expression = &Count{}

// Count is a count aggregate function.
// example:
//
//	COUNT(column)
type Count struct {
	Col *expressions.Column
}

// WriteTo implements Expression.
func (c *Count) Write(sw io.StringWriter) error {
	if err := util.WriteStrings(sw, "COUNT("); err != nil {
		return fmt.Errorf("error writing COUNT function: %v", err)
	}

	if err := c.Col.Write(sw); err != nil {
		return fmt.Errorf("error writing COUNT expression: %v", err)
	}

	if err := util.WriteStrings(sw, ")"); err != nil {
		return fmt.Errorf("error writing COUNT closing parenthesis: %v", err)
	}

	return nil
}

// NewCount returns a new Count aggregate functions expression.
func NewCount(column *expressions.Column) *Count {
	return &Count{
		Col: column,
	}
}
