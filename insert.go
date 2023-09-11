package gosqle

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/clauses/values"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/statement"
)

// Insert is a wrapper for a insert query statement.
type Insert struct {
	statement.Statement
}

// Values adds values to the insert query.
func (i *Insert) Values(arguments ...expressions.Expression) *Insert {
	i.Statement.SetClause(values.New(arguments))
	return i
}

// SetClause sets the clause for the query.
func (i *Insert) SetClause(c clauses.Clause) *Insert {
	i.Statement.SetClause(c)
	return i
}

// Write writes the insert query to the given writer.
// It also adds a semicolon to the end of the query.
func (i *Insert) WriteTo(sw io.StringWriter) error {
	if err := i.Statement.WriteTo(sw); err != nil {
		return fmt.Errorf("failed to write insert statement: %v", err)
	}

	// Add a semicolon to the end of the query.
	if _, err := sw.WriteString(";"); err != nil {
		return fmt.Errorf("failed to write semicolon: %v", err)
	}

	return nil
}

// NewInsert creates a new insert query.
func NewInsert(
	into string,
	columns ...string,
) *Insert {
	return &Insert{
		Statement: statement.NewInsert(into, columns),
	}
}
