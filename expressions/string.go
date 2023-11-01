package expressions

import (
	"fmt"
	"io"
)

type String string

func (e String) Write(writer io.StringWriter) error {
	if _, err := writer.WriteString(string(e)); err != nil {
		return fmt.Errorf("error writing string: %w", err)
	}

	return nil
}
