package root

import "embed"

//go:embed web
var EmbedWeb embed.FS

//go:embed domains
var EmbedDomains embed.FS
