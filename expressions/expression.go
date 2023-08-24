package expressions

import "io"

// Expression is a generic SQL expression.
type Expression interface {
	WriteTo(io.StringWriter) error
}
