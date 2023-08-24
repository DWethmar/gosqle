package having

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/predicates"
)

// WriteHaving writes the HAVING clause to the given writer.
func WriteHaving(sw io.StringWriter, preds []predicates.Predicate) error {
	if _, err := sw.WriteString("HAVING "); err != nil {
		return fmt.Errorf("error writing HAVING: %v", err)
	}

	if err := predicates.WriteAll(sw, preds); err != nil {
		return fmt.Errorf("error writing HAVING predicates: %v", err)
	}

	return nil
}

var _ clauses.Clause = &Clause{}

type Clause struct {
	Predicates []predicates.Predicate
}

func (*Clause) Type() clauses.ClauseType           { return clauses.HavingType }
func (c *Clause) WriteTo(sw io.StringWriter) error { return WriteHaving(sw, c.Predicates) }

func New(predicates []predicates.Predicate) *Clause {
	return &Clause{
		Predicates: predicates,
	}
}
