package expressions

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/util"
)

var (
	_ Expression = (*Column)(nil)
)

// Column is a column in a table
type Column struct {
	From string // table name
	Name string // column name
}

// Write writes the column to the given writer
func (s Column) WriteTo(sw io.StringWriter) error {
	if s.Name == "" {
		return fmt.Errorf("column name is empty")
	}

	if s.From != "" {
		if err := util.WriteStrings(sw, s.From, "."); err != nil {
			return fmt.Errorf("could not write from: %v", err)
		}
	}

	if _, err := sw.WriteString(s.Name); err != nil {
		return fmt.Errorf("could not write name: %v", err)
	}

	return nil
}
