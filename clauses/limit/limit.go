package limit

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
)

var (
	_ clauses.Clause = &Clause{}
)

// Write writes a SQL limit string to the given string writer.
//
// Example:
//
//	Mysql:
//		LIMIT <expression>
//	Postgres:
//		LIMIT <expression>
func Write(sw io.StringWriter, s expressions.Expression) error {
	if _, err := sw.WriteString("LIMIT "); err != nil {
		return fmt.Errorf("could not write limit clause")
	}

	if err := s.WriteTo(sw); err != nil {
		return fmt.Errorf("could not write limit expression")
	}

	return nil
}

// Clause represents a LIMIT clause.
type Clause struct {
	expressions.Expression
}

func (l *Clause) Type() clauses.ClauseType         { return clauses.LimitType }
func (l *Clause) WriteTo(sw io.StringWriter) error { return Write(sw, l.Expression) }

func New(limit expressions.Expression) *Clause {
	return &Clause{
		Expression: limit,
	}
}
