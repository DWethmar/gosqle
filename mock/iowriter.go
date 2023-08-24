package mock

import "io"

// type check
var _ io.StringWriter = StringWriterFn(nil)

type StringWriterFn func(s string) (n int, err error)

func (fn StringWriterFn) WriteString(s string) (n int, err error) {
	return fn(s)
}
