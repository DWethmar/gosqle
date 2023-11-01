package alias

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/util"
)

var _ expressions.Expression = &Alias{}

// Alias represents an alias.
// example:
//
//	expressions AS a
type Alias struct {
	Expr expressions.Expression
	As   string // optional
}

func (a *Alias) SetAs(as string) *Alias {
	a.As = as
	return a
}

// Write implements expressions.Expression.
func (a *Alias) Write(sw io.StringWriter) error {
	if err := a.Expr.Write(sw); err != nil {
		return fmt.Errorf("error writing alias expression: %w", err)
	}

	if a.As == "" {
		return nil
	}

	if err := util.WriteStrings(sw, " AS ", a.As); err != nil {
		return fmt.Errorf("error writing alias: %w", err)
	}

	return nil
}

func New(expr expressions.Expression) *Alias {
	return &Alias{
		Expr: expr,
	}
}
