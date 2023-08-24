package util

import (
	"errors"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/mock"
)

func TestWriteStrings(t *testing.T) {
	t.Run("should write strings", func(t *testing.T) {
		sb := new(strings.Builder)

		if err := WriteStrings(sb, "hello", " ", "world"); err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if sb.String() != "hello world" {
			t.Errorf("expected 'hello world', got %s", sb.String())
		}
	})

	t.Run("should return error on io.StringWriter.WriteString error", func(t *testing.T) {
		sb := mock.StringWriterFn(func(_ string) (int, error) {
			return 0, errors.New("error")
		})

		err := WriteStrings(sb, "hello", " ", "world")
		if err == nil {
			t.Errorf("expected error, got nil")
		}

		if err.Error() != "failed to write string to io.StringWriter: error" {
			t.Errorf("expected 'failed to write string to io.StringWriter: : error', got %q", err.Error())
		}
	})
}
