package from

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/util"
)

var (
	_ expressions.Expression = (*Table)(nil)
)

// Table represents a table in a query.
type Table struct {
	// Name of the table.
	Name string
	// As of the table: "table AS alias".
	As string
}

// Write writes a table to the given Writer.
func (t Table) WriteTo(sw io.StringWriter) error {
	if t.Name == "" {
		return fmt.Errorf("name is empty")
	}

	if t.As != "" {
		if err := util.WriteStrings(sw, t.Name, " AS ", t.As); err != nil {
			return fmt.Errorf("error writing table: %w", err)
		}
	} else {
		if _, err := sw.WriteString(t.Name); err != nil {
			return fmt.Errorf("error writing table: %w", err)
		}
	}

	return nil
}
