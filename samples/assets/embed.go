package assets

import (
	_ "embed"
)

var (
	//go:embed sample_spritesheet.json
	Spritesheet_JSON []byte
	//go:embed sample_spritesheet.png
	Spritesheet_PNG []byte
)
