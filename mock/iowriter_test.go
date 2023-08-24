package mock

import "testing"

func TestStringWriterFn_WriteString(t *testing.T) {
	t.Run("should write string", func(t *testing.T) {
		fn := StringWriterFn(func(s string) (n int, err error) {
			return len(s), nil
		})

		n, err := fn.WriteString("test")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if n != 4 {
			t.Errorf("expected to write 4 bytes, got %d", n)
		}
	})
}
