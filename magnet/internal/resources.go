package internal

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed resources
var rawEmbed embed.FS

var (
	Embed fs.FS
)

func init() {
	var err error
	Embed, err = fs.Sub(rawEmbed, "resources")
	if err != nil {
		log.Fatal(err)
	}
}
