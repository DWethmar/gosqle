package values

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
)

var (
	_ clauses.Clause = &Clause{}
)

// DefaultValue is a default value clause.
//
// Only supported in PostgreSQL and MySQL.
type DefaultValue struct{}

// validateValuesInput validates the input for the Values clause.
// otherwise the linter will complain about code complexity.
func validateValuesInput(values []expressions.Expression) error {
	if len(values) == 0 {
		return fmt.Errorf("no values given")
	}

	return nil
}

// Write writes a values clause to the given io.Writer.
//
// Example:
//
//	VALUES ($1, $2, $3)
//	VALUES (DEFAULT, $1, $2)
func Write(sw io.StringWriter, exprs []expressions.Expression) error {
	if err := validateValuesInput(exprs); err != nil {
		return err
	}

	if _, err := sw.WriteString(`VALUES (`); err != nil {
		return fmt.Errorf("failed to write string: %w", err)
	}

	for i, e := range exprs {
		if err := e.Write(sw); err != nil {
			return fmt.Errorf("failed to write expression %d: %w", i, err)
		}

		if i < len(exprs)-1 {
			if _, err := sw.WriteString(", "); err != nil {
				return fmt.Errorf("failed to write comma separator: %w", err)
			}
		}
	}

	if _, err := sw.WriteString(`)`); err != nil {
		return fmt.Errorf("failed to write closing parenthesis: %w", err)
	}

	return nil
}

// Clause represents a values clause.
type Clause struct {
	expressions []expressions.Expression
}

// Type implements Clause.
func (*Clause) Type() clauses.ClauseType         { return clauses.ValuesType }
func (v *Clause) Write(sw io.StringWriter) error { return Write(sw, v.expressions) }

// New creates a new Values clause.
func New(expressions []expressions.Expression) *Clause {
	return &Clause{
		expressions: expressions,
	}
}
