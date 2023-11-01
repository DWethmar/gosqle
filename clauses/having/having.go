package having

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/logic"
)

// WriteHaving writes the HAVING clause to the given writer.
func WriteHaving(sw io.StringWriter, conditions []logic.Logic) error {
	if sw == nil {
		return fmt.Errorf("string writer is nil")
	}

	if len(conditions) == 0 {
		return fmt.Errorf("conditions is empty")
	}

	if _, err := sw.WriteString("HAVING "); err != nil {
		return fmt.Errorf("error writing HAVING: %v", err)
	}

	if err := logic.Where(sw, conditions); err != nil {
		return fmt.Errorf("error writing HAVING logic: %v", err)
	}

	return nil
}

var _ clauses.Clause = &Clause{}

type Clause struct {
	conditions []logic.Logic
}

func (*Clause) Type() clauses.ClauseType         { return clauses.HavingType }
func (c *Clause) Write(sw io.StringWriter) error { return WriteHaving(sw, c.conditions) }

func New(conditions []logic.Logic) *Clause {
	return &Clause{
		conditions: conditions,
	}
}
