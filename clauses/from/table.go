package from

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/expressions"
)

var (
	_ expressions.Expression = (*Table)(nil)
)

// Table represents a table in a query.
type Table string

// Write writes a table to the given Writer.
func (t Table) WriteTo(sw io.StringWriter) error {
	if t == "" {
		return fmt.Errorf("name is empty")
	}

	if _, err := sw.WriteString(string(t)); err != nil {
		return fmt.Errorf("error writing table: %w", err)
	}

	return nil
}
