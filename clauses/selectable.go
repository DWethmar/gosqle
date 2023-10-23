package clauses

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/util"
)

// Selectable represents a something that can be selected in a select statement or returning, etc.
type Selectable struct {
	Expr expressions.Expression
	As   string
}

// Write writes a SQL select column to the given string writer.
func (s Selectable) Write(sw io.StringWriter) error {
	if err := s.Expr.Write(sw); err != nil {
		return fmt.Errorf("failed to write column: %v", err)
	}

	if s.As != "" {
		if err := util.WriteStrings(sw, " AS ", s.As); err != nil {
			return fmt.Errorf("failed to write AS: %v", err)
		}
	}

	return nil
}
