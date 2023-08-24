package statement

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
)

const (
	SpaceSeparator = " "
	CommaSeparator = ", "
)

type Statement interface {
	expressions.Expression
	SetClause(clauses.Clause)
}

// ClauseWriter represents a ClauseWriter.
type ClauseWriter struct {
	order           []clauses.ClauseType
	clauses         map[clauses.ClauseType]clauses.Clause
	ClauseSeparator string
}

// SetClause implements Statement SetClause.
func (s *ClauseWriter) SetClause(c clauses.Clause) {
	if c == nil {
		delete(s.clauses, c.Type())
		return
	}

	s.clauses[c.Type()] = c
}

// Write writes clauses to the string writer.
func (s *ClauseWriter) WriteTo(sw io.StringWriter) error {
	clauses := []clauses.Clause{}
	for _, t := range s.order {
		if c, ok := s.clauses[t]; ok {
			clauses = append(clauses, c)
		}
	}

	for i, clause := range clauses {
		// Write clause to the string builder.
		if err := clause.WriteTo(sw); err != nil {
			return fmt.Errorf("failed to write clause: %s", err)
		}

		// Add a separator after the clause if not last clause.
		if i < len(clauses)-1 {
			if _, err := sw.WriteString(s.ClauseSeparator); err != nil {
				return fmt.Errorf("failed to write separator: %s", err)
			}
		}
	}

	return nil
}
