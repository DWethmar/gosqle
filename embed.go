package gosqle

import (
	"embed"
)

//go:embed examples/*.go
var GoExampleFiles embed.FS

//go:embed README.tmpl.md
var ReadMeTemplate []byte
