package statement

import (
	"errors"
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/clauses"
)

var (
	_ Statement = (*Select)(nil)
)

// SelectClausesOrder is the order in which the select clauses are written.
var selectClausesOrder = []clauses.ClauseType{
	clauses.FromType,
	clauses.JoinType,
	clauses.WhereType,
	clauses.GroupByType,
	clauses.LimitType,
	clauses.OffsetType,
	clauses.HavingType,
	clauses.OrderByType,
	clauses.ReturningType, // only for postgres
}

// WriteSelect writes a SQL select query to the given string writer.
//
// Example:
//
//	SELECT field1, field2, field3
func WriteSelect(sw io.StringWriter, selectColumns []clauses.Selectable) error {
	if len(selectColumns) == 0 {
		return errors.New("no select columns specified")
	}

	if _, err := sw.WriteString("SELECT "); err != nil {
		return fmt.Errorf("failed to write SELECT: %v", err)
	}

	for i, e := range selectColumns {
		if err := e.Write(sw); err != nil {
			return fmt.Errorf("failed to write expression: %v", err)
		}

		if i != len(selectColumns)-1 {
			if _, err := sw.WriteString(CommaSeparator); err != nil {
				return fmt.Errorf("failed to write separator: %v", err)
			}
		}
	}

	return nil
}

// SelectWriter writes a SQL From string to the given string writer.
type Select struct {
	ClauseWriter
	selectColumns []clauses.Selectable
}

// ToSQL returns the query as a string and it's arguments and an error if any.
func (s *Select) Write(sw io.StringWriter) error {
	if err := WriteSelect(sw, s.selectColumns); err != nil {
		return err
	}

	if len(s.clauses) > 0 {
		if _, err := sw.WriteString(" "); err != nil {
			return fmt.Errorf("failed to write space: %v", err)
		}

		if err := s.ClauseWriter.Write(sw); err != nil {
			return err
		}
	}

	return nil
}

// NewSelectClauseWriter creates a new ClauseWriter for select queries.
func NewSelectClauseWriter() ClauseWriter {
	return ClauseWriter{
		clauses:         map[clauses.ClauseType]clauses.Clause{},
		order:           selectClausesOrder,
		ClauseSeparator: SpaceSeparator,
	}
}

// NewSelectClause creates a new SelectClause.
func NewSelect(selectColumns []clauses.Selectable) *Select {
	return &Select{
		ClauseWriter:  NewSelectClauseWriter(),
		selectColumns: selectColumns,
	}
}
