package template

import "embed"

//go:embed base/* base/**/*
var Files embed.FS
