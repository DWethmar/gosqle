package util

import (
	"fmt"
	"io"
)

// WriteStrings writes strings to the given string writer.
func WriteStrings(sw io.StringWriter, strs ...string) error {
	for _, str := range strs {
		if _, err := sw.WriteString(str); err != nil {
			return fmt.Errorf("failed to write string to io.StringWriter: %w", err)
		}
	}

	return nil
}
