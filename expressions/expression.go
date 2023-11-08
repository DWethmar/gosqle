package expressions

import (
	"fmt"
	"io"
)

// Expression is a generic SQL expression.
type Expression interface {
	Write(io.StringWriter) error
}

// List is a list of expressions separated by a comma.
type List []Expression

func (e List) Write(writer io.StringWriter) error {
	for i, argument := range e {
		if err := argument.Write(writer); err != nil {
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

var _ Expression = ExpressionFunc(nil)

type ExpressionFunc func(io.StringWriter) error

func (e ExpressionFunc) Write(writer io.StringWriter) error {
	return e(writer)
}

func WrapInParenthesis(sw io.StringWriter, e Expression) error {
	if _, err := sw.WriteString("("); err != nil {
		return fmt.Errorf("failed to write (")
	}

	if err := e.Write(sw); err != nil {
		return fmt.Errorf("failed to write expression: %w", err)
	}

	if _, err := sw.WriteString(")"); err != nil {
		return fmt.Errorf("failed to write )")
	}

	return nil
}

// Prepend writes a string and then an expression.
func Prepend(sw io.StringWriter, str string, e Expression) error {
	if _, err := sw.WriteString(str); err != nil {
		return fmt.Errorf("failed to prepend string: %w", err)
	}

	if err := e.Write(sw); err != nil {
		return fmt.Errorf("failed to write expression: %w", err)
	}

	return nil
}

// Append writes an expression and then a string.
func Append(sw io.StringWriter, e Expression, str string) error {
	if err := e.Write(sw); err != nil {
		return fmt.Errorf("failed to append string expression: %w", err)
	}

	if _, err := sw.WriteString(str); err != nil {
		return fmt.Errorf("failed to write %s", str)
	}

	return nil
}
