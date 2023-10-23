package set

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

// Change represents a change to a column.
type Change struct {
	// Col is the column name.
	Col  string
	Expr expressions.Expression
}

// Write writes a SQL string to the given string writer.
func (c *Change) Write(sw io.StringWriter) error {
	if err := util.WriteStrings(sw, c.Col, " = "); err != nil {
		return fmt.Errorf("failed to write column=%s: %w", c.Col, err)
	}

	if err := c.Expr.Write(sw); err != nil {
		return fmt.Errorf("failed to write expression")
	}

	return nil
}

// Set writes a SQL string to the given string writer.
//
// Example:
//
//	Mysql:
//		SET field1 = ?, field2 = ?, field3 = ?
//	Postgres:
//		SET field1 = $1, field2 = $2, field3 = $3
func Write(sw io.StringWriter, changes []Change) error {
	if len(changes) == 0 {
		return fmt.Errorf("no columns given")
	}

	if _, err := sw.WriteString(`SET `); err != nil {
		return fmt.Errorf("failed to write string: %w", err)
	}

	for i, change := range changes {
		if err := change.Write(sw); err != nil {
			return fmt.Errorf("failed to write change: %w", err)
		}

		if i < len(changes)-1 {
			if err := util.WriteStrings(sw, ", "); err != nil {
				return fmt.Errorf("failed to write separator: %w", err)
			}
		}
	}

	return nil
}

// Clause writes a SQL string to the given string writer.
type Clause struct {
	changes []Change
}

func (s *Clause) Type() clauses.ClauseType       { return clauses.SetType }
func (s *Clause) Write(sw io.StringWriter) error { return Write(sw, s.changes) }

func New(changes []Change) *Clause {
	return &Clause{
		changes: changes,
	}
}
