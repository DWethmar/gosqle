package logic

import (
	"errors"
	"fmt"
	"io"
)

var ErrWriterNil = errors.New("writer is nil")

// Where writes a slice of logic to the given string writer and skips the first logic operator.
// Example:
//
//	x == 1 AND y == 2
func Where(sw io.StringWriter, logic []Logic) error {
	if sw == nil {
		return ErrWriterNil
	}

	for i, l := range logic {
		if i > 0 {
			if err := l.Write(sw); err != nil {
				return err
			}
		} else { // write first logic without logic operator
			if err := l.Condition.Write(sw); err != nil {
				return fmt.Errorf("failed to write logic condition: %w", err)
			}
		}

		if i < len(logic)-1 {
			if _, err := sw.WriteString(" "); err != nil {
				return fmt.Errorf("failed to write space: %w", err)
			}
		}
	}

	return nil
}
