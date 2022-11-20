package templates

import _ "embed"

//go:embed go.tpl
var Go string

//go:embed cs.tpl
var CS string
