package expressions

import (
	"fmt"
	"io"
)

// Expression is a generic SQL expression.
type Expression interface {
	WriteTo(io.StringWriter) error
}

// List is a list of expressions separated by a comma.
type List []Expression

func (e List) WriteTo(writer io.StringWriter) error {
	for i, argument := range e {
		if err := argument.WriteTo(writer); err != nil {
			return fmt.Errorf("could not write expression at index %d: %w", i, err)
		}

		if i < len(e)-1 {
			if _, err := writer.WriteString(", "); err != nil {
				return fmt.Errorf("error writing comma: %w", err)
			}
		}
	}

	return nil
}
