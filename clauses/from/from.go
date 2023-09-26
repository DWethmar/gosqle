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

// From represents a FROM.
type From struct {
	Expr expressions.Expression
	As   string
}

// Write writes a SQL Write string to the given string writer.
//
// Examples:
//
//	FROM table AS alias
//	FROM table
//	FROM (SELECT * FROM table) AS alias
func Write(sw io.StringWriter, from From) error {
	if err := util.WriteStrings(sw, "FROM "); err != nil {
		return fmt.Errorf("from: %v", err)
	}

	if err := from.Expr.WriteTo(sw); err != nil {
		return fmt.Errorf("from: %v", err)
	}

	if from.As != "" {
		if err := util.WriteStrings(sw, " AS ", from.As); err != nil {
			return fmt.Errorf("from: %v", err)
		}
	}

	return nil
}

// Clause represents a FROM clause.
type Clause struct {
	from From
}

func (c *Clause) Type() clauses.ClauseType         { return clauses.FromType }
func (c *Clause) WriteTo(sw io.StringWriter) error { return Write(sw, c.from) }

// New creates a new from clause
func New(from From) *Clause {
	return &Clause{
		from: from,
	}
}
