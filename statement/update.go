package statement

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/util"
)

var (
	_ Statement = (*Select)(nil)
)

// SelectClausesOrder is the order in which the select clauses are written.
var updateClausesOrder = []clauses.ClauseType{
	clauses.SetType,
	clauses.JoinType,
	clauses.WhereType,
	clauses.OrderByType,
	clauses.LimitType,
	clauses.ReturningType, // only for postgres
}

// WriteUpdate writes a SQL update query to the given string writer.
//
// Example:
//
//	UPDATE table
func WriteUpdate(sw io.StringWriter, table string) error {
	if err := util.WriteStrings(sw, "UPDATE ", table); err != nil {
		return fmt.Errorf("failed to write UPDATE: %v", err)
	}

	return nil
}

// Update writes an update statement to the given string writer.
type Update struct {
	ClauseWriter
	table string
}

// ToSQL returns the query as a string and it's arguments and an error if any.
func (u *Update) Write(sw io.StringWriter) error {
	if err := WriteUpdate(sw, u.table); err != nil {
		return err
	}

	if len(u.clauses) > 0 {
		if _, err := sw.WriteString(" "); err != nil {
			return fmt.Errorf("failed to write space: %v", err)
		}

		if err := u.ClauseWriter.Write(sw); err != nil {
			return err
		}
	}

	return nil
}

// NewUpdateClauseWriter creates a new ClauseWriter for update statements.
func NewUpdateClauseWriter() ClauseWriter {
	return ClauseWriter{
		clauses:         map[clauses.ClauseType]clauses.Clause{},
		order:           updateClausesOrder,
		ClauseSeparator: SpaceSeparator,
	}
}

// NewUpdate creates a new update statement.
func NewUpdate(
	table string,
) *Update {
	return &Update{
		ClauseWriter: NewUpdateClauseWriter(),
		table:        table,
	}
}
