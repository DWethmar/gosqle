package statement

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/clauses"
)

var (
	_ Statement = (*Select)(nil)
)

// deleteClausesOrder is the order in which clauses are written.
var deleteClausesOrder = []clauses.ClauseType{
	clauses.WhereType,
	clauses.ReturningType, // only for postgres
}

// Delete writes a SQL From string to the given string writer.
type Delete struct {
	ClauseWriter
	table string
}

func WriteDelete(sw io.StringWriter, table string) error {
	if _, err := sw.WriteString("DELETE FROM "); err != nil {
		return fmt.Errorf("failed to write DELETE FROM: %v", err)
	}

	if _, err := sw.WriteString(table); err != nil {
		return fmt.Errorf("failed to write table: %v", err)
	}

	return nil
}

// ToSQL returns the query as a string and it's arguments and an error if any.
func (d *Delete) WriteTo(sw io.StringWriter) error {
	if err := WriteDelete(sw, d.table); err != nil {
		return err
	}

	if len(d.clauses) > 0 {
		if _, err := sw.WriteString(" "); err != nil {
			return fmt.Errorf("failed to write space: %v", err)
		}

		if err := d.ClauseWriter.WriteTo(sw); err != nil {
			return err
		}
	}

	return nil
}

// NewDelete creates a new Delete statement.
func NewDelete(
	table string,
) *Delete {
	return &Delete{
		ClauseWriter: ClauseWriter{
			clauses:         map[clauses.ClauseType]clauses.Clause{},
			order:           deleteClausesOrder,
			ClauseSeparator: SpaceSeparator,
		},
		table: table,
	}
}
