package where

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/predicates"
)

var (
	_ clauses.Clause = &Where{}
)

// WriteWhere writes a where clause to the given string writer and returns the args.
func WriteWhere(
	sw io.StringWriter,
	p []predicates.Predicate,
) error {
	if len(p) == 0 {
		return fmt.Errorf("no conditions given")
	}

	if _, err := sw.WriteString(`WHERE `); err != nil {
		return fmt.Errorf("failed to write WHERE: %w", err)
	}

	if err := predicates.WriteAll(sw, p); err != nil {
		return fmt.Errorf("error writing predicates: %v", err)
	}

	return nil
}

// UpdateWriter writes a SQL Update string.
type Where struct {
	predicates []predicates.Predicate
}

func (w *Where) Type() clauses.ClauseType         { return clauses.WhereType }
func (w *Where) WriteTo(sw io.StringWriter) error { return WriteWhere(sw, w.predicates) }

// New
func New(predicates []predicates.Predicate) *Where {
	return &Where{
		predicates: predicates,
	}
}
