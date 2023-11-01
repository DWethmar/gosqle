package where

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/logic"
)

var (
	_ clauses.Clause = &Where{}
)

// WriteWhere writes a where clause to the given string writer and returns the args.
func WriteWhere(sw io.StringWriter, conditions []logic.Logic) error {
	if len(conditions) == 0 {
		return fmt.Errorf("conditions is empty")
	}

	if _, err := sw.WriteString("WHERE "); err != nil {
		return fmt.Errorf("failed to write WHERE: %w", err)
	}

	if err := logic.Where(sw, conditions); err != nil {
		return fmt.Errorf("failed to write WHERE logic: %w", err)
	}

	return nil
}

// UpdateWriter writes a SQL Update string.
type Where struct {
	conditions []logic.Logic
}

func (w *Where) Type() clauses.ClauseType       { return clauses.WhereType }
func (w *Where) Write(sw io.StringWriter) error { return WriteWhere(sw, w.conditions) }

// New
func New(conditions []logic.Logic) *Where {
	return &Where{
		conditions: conditions,
	}
}
