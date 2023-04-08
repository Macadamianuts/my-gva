package internal

import _ "embed"

//go:embed gen.go.tpl
var Gen []byte

//go:embed  main.go.tpl
var Main []byte
