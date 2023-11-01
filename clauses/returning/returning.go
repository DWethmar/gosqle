package returning

import (
	"errors"
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/clauses"
)

var (
	_ clauses.Clause = &Clause{}
)

// Write writes a RETURNING clause to a strings.Builder.
//
// Example:
//
//	RETURNING field1, field2, field3
func Write(sw io.StringWriter, selectColumns []alias.Alias) error {
	if len(selectColumns) == 0 {
		return errors.New("no columns to return")
	}

	if _, err := sw.WriteString(`RETURNING `); err != nil {
		return fmt.Errorf("error on writing RETURNING: %w", err)
	}

	for i, column := range selectColumns {
		if err := column.Write(sw); err != nil {
			return fmt.Errorf("error on writing RETURNING column: %w", err)
		}

		if i != len(selectColumns)-1 {
			if _, err := sw.WriteString(", "); err != nil {
				return fmt.Errorf("error on writing RETURNING column separator: %w", err)
			}
		}
	}

	return nil
}

// Clause represents a RETURNING clause.
type Clause struct {
	selectColumns []alias.Alias
}

func (r *Clause) Type() clauses.ClauseType       { return clauses.ReturningType }
func (r *Clause) Write(sw io.StringWriter) error { return Write(sw, r.selectColumns) }

func New(selectColumns []alias.Alias) *Clause {
	return &Clause{
		selectColumns: selectColumns,
	}
}
