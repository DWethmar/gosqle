package from

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/util"
)

var (
	_ clauses.Clause = &Clause{}
)

// Write writes a SQL Write string to the given string writer.
//
// Examples:
//
//	FROM table AS alias
//	FROM table
//	FROM (SELECT * FROM table) AS alias
func Write(sw io.StringWriter, expr expressions.Expression) error {
	if err := util.WriteStrings(sw, "FROM "); err != nil {
		return fmt.Errorf("from: %v", err)
	}

	if err := expr.WriteTo(sw); err != nil {
		return fmt.Errorf("from: %v", err)
	}

	return nil
}

// Clause represents a FROM clause.
type Clause struct {
	expressions.Expression
}

func (c *Clause) Type() clauses.ClauseType         { return clauses.FromType }
func (c *Clause) WriteTo(sw io.StringWriter) error { return Write(sw, c.Expression) }

// New creates a new from clause
func New(expr expressions.Expression) *Clause {
	return &Clause{
		Expression: expr,
	}
}
