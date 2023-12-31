package gosqle

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/clauses/where"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/statement"
)

// Delete is a wrapper for a delete query statement.
type Delete struct {
	statement.Statement
}

// Where adds a where clause to the query.
func (d *Delete) Where(logic ...logic.Logic) *Delete {
	if len(logic) == 0 {
		return d
	}

	return d.SetClause(where.New(logic))
}

// SetClause sets the clause for the query.
func (d *Delete) SetClause(c clauses.Clause) *Delete {
	d.Statement.SetClause(c)
	return d
}

// WriteTo writes the delete statement to the given string writer.
func (d *Delete) Write(sw io.StringWriter) error {
	if err := d.Statement.Write(sw); err != nil {
		return fmt.Errorf("failed to write delete statement: %v", err)
	}

	// Add a semicolon to the end of the query.
	if _, err := sw.WriteString(";"); err != nil {
		return fmt.Errorf("failed to write semicolon: %v", err)
	}

	return nil
}

// NewDelete creates a new delete query.
func NewDelete(table string) *Delete {
	return &Delete{
		Statement: statement.NewDelete(table),
	}
}
