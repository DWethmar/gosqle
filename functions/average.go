package functions

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/util"
)

var _ expressions.Expression = &Avg{}

// Avg is a avg aggregate function.
// example:
//
//	AVG(column)
type Avg struct {
	Col *expressions.Column
}

func (a *Avg) Write(sw io.StringWriter) error {
	if err := util.WriteStrings(sw, "AVG("); err != nil {
		return fmt.Errorf("error writing AVG function: %v", err)
	}

	if err := a.Col.Write(sw); err != nil {
		return fmt.Errorf("error writing AVG expression: %v", err)
	}

	if err := util.WriteStrings(sw, ")"); err != nil {
		return fmt.Errorf("error writing AVG closing parenthesis: %v", err)
	}

	return nil
}

// NewAvg returns a new Avg aggregate functions expression.
func NewAvg(column *expressions.Column) *Avg {
	return &Avg{
		Col: column,
	}
}
