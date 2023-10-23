package statement

import (
	"errors"
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/util"
)

var (
	_ Statement = (*Select)(nil)
)

// SelectClausesOrder is the order in which the select clauses are written.
var insertClausesOrder = []clauses.ClauseType{
	clauses.ValuesType,
	clauses.ReturningType, // postgres only
}

// WriteInsert writes a SQL WriteInsert statement to the given string writer.
//
// Example:
//
//	INSERT INTO users (username, age)
func WriteInsert(sw io.StringWriter, into string, columns []string) error {
	if err := util.WriteStrings(sw, "INSERT INTO ", into); err != nil {
		return fmt.Errorf("failed to write table: %v", err)
	}

	if len(columns) > 0 {
		if _, err := sw.WriteString(" ("); err != nil {
			return fmt.Errorf("failed to write opening parenthesis: %v", err)
		}

		for i, e := range columns {
			if _, err := sw.WriteString(e); err != nil {
				return fmt.Errorf("failed to write column: %v", err)
			}

			if i < len(columns)-1 {
				if _, err := sw.WriteString(CommaSeparator); err != nil {
					return fmt.Errorf("failed to write comma separator: %v", err)
				}
			}
		}

		if _, err := sw.WriteString(")"); err != nil {
			return fmt.Errorf("failed to write closing parenthesis: %v", err)
		}
	}

	return nil
}

// SelectWriter writes a SQL From string to the given string writer.
type Insert struct {
	ClauseWriter
	table   string
	columns []string
}

// ToSQL returns the query as a string and it's arguments and an error if any.
func (i *Insert) Write(sw io.StringWriter) error {
	if i.table == "" {
		return errors.New("no table specified")
	}

	if err := WriteInsert(sw, i.table, i.columns); err != nil {
		return err
	}

	if len(i.clauses) > 0 {
		if _, err := sw.WriteString(" "); err != nil {
			return fmt.Errorf("failed to write space: %v", err)
		}

		if err := i.ClauseWriter.Write(sw); err != nil {
			return err
		}
	}

	return nil
}

// NewSelectClause creates a new SelectClause.
func NewInsert(table string, columns []string) *Insert {
	return &Insert{
		ClauseWriter: ClauseWriter{
			clauses:         map[clauses.ClauseType]clauses.Clause{},
			order:           insertClausesOrder,
			ClauseSeparator: SpaceSeparator,
		},
		table:   table,
		columns: columns,
	}
}
