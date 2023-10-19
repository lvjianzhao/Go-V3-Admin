package resources

import "embed"

//go:embed dist/index.html
var Html []byte

//go:embed dist
var Static embed.FS
