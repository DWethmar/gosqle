package gosqle

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/clauses/join"
	"github.com/dwethmar/gosqle/clauses/orderby"
	"github.com/dwethmar/gosqle/clauses/set"
	"github.com/dwethmar/gosqle/clauses/where"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/statement"
)

// Update is a wrapper for a update query statement.
type Update struct {
	statement.Statement
}

// From adds a from clause to the select statement.
func (u *Update) Join(j ...join.Options) *Update {
	if len(j) == 0 {
		return u
	}

	return u.SetClause(join.New(j))
}

// Set adds a set clause to the select statement.
func (u *Update) Set(changes ...set.Change) *Update {
	return u.SetClause(set.New(changes))
}

// Where adds a where clause to the select statement.
func (u *Update) Where(conditions ...logic.Logic) *Update {
	if len(conditions) == 0 {
		return u
	}

	return u.SetClause(where.New(conditions))
}

// OrderBy adds a order by clause to the select statement.
func (u *Update) OrderBy(sorting ...orderby.Sort) *Update {
	if len(sorting) == 0 {
		return u
	}

	return u.SetClause(orderby.New(sorting))
}

// SetClause sets the clause for the select statement.
func (u *Update) SetClause(c clauses.Clause) *Update {
	u.Statement.SetClause(c)
	return u
}

// WriteTo writes the select statement to the given writer.
func (u *Update) Write(sw io.StringWriter) error {
	if err := u.Statement.Write(sw); err != nil {
		return fmt.Errorf("failed to write insert statement: %v", err)
	}

	// Add a semicolon to the end of the query.
	if _, err := sw.WriteString(";"); err != nil {
		return fmt.Errorf("failed to write semicolon: %v", err)
	}

	return nil
}

// NewUpdate creates a new update query.
func NewUpdate(table string) *Update {
	return &Update{
		Statement: statement.NewUpdate(table),
	}
}
