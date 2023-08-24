package util

import (
	"fmt"
	"io"
)

// WriteStrings writes strings to the given string writer.
func WriteStrings(sw io.StringWriter, str ...string) error {
	for _, st := range str {
		if _, err := sw.WriteString(st); err != nil {
			return fmt.Errorf("failed to write string to io.StringWriter: %w", err)
		}
	}

	return nil
}
