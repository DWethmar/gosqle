package offset

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
)

var (
	_ clauses.Clause = &Clause{}
)

// Write writes a SQL offset string to the given string writer.
//
// Example:
//
//	Mysql:
//		OFFSET <expression>
//	Postgres:
//		OFFSET <expression>
func Write(sw io.StringWriter, s expressions.Expression) error {
	if _, err := sw.WriteString("OFFSET "); err != nil {
		return fmt.Errorf("could not write offset clause")
	}

	if err := s.Write(sw); err != nil {
		return fmt.Errorf("could not write offset expression")
	}

	return nil
}

// Clause represents a OFFSET clause.
type Clause struct {
	expressions.Expression
}

func (o *Clause) Type() clauses.ClauseType       { return clauses.OffsetType }
func (o *Clause) Write(sw io.StringWriter) error { return Write(sw, o.Expression) }

func New(offset expressions.Expression) *Clause {
	return &Clause{
		Expression: offset,
	}
}
