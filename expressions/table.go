package expressions

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/util"
)

var (
	_ Expression = (*Table)(nil)
)

// Table represents a table in a query.
type Table struct {
	// Name of the table.
	Name string
	// Alias of the table.
	Alias string
}

func (t Table) SetParamOffset(int)  {}
func (t Table) Args() []interface{} { return nil }

// Write writes a table to the given Writer.
func (t Table) WriteTo(sw io.StringWriter) error {
	if t.Name == "" {
		return fmt.Errorf("name is empty")
	}

	if t.Alias != "" {
		if err := util.WriteStrings(sw, t.Name, " ", t.Alias); err != nil {
			return fmt.Errorf("error writing table: %w", err)
		}
	} else {
		if _, err := sw.WriteString(t.Name); err != nil {
			return fmt.Errorf("error writing table: %w", err)
		}
	}

	return nil
}
