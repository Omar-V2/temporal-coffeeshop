package migration

import "embed"

//go:embed *.sql
var Customer embed.FS
