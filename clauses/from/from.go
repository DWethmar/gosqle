package from

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/clauses"
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
func Write(sw io.StringWriter, from alias.Alias) error {
	if err := util.WriteStrings(sw, "FROM "); err != nil {
		return fmt.Errorf("from: %v", err)
	}

	if err := from.Write(sw); err != nil {
		return fmt.Errorf("from: %v", err)
	}

	return nil
}

// Clause represents a FROM clause.
type Clause struct {
	from alias.Alias
}

func (c *Clause) Type() clauses.ClauseType       { return clauses.FromType }
func (c *Clause) Write(sw io.StringWriter) error { return Write(sw, c.from) }

// New creates a new from clause
func New(from alias.Alias) *Clause {
	return &Clause{
		from: from,
	}
}
