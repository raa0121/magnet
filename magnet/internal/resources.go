package internal

import (
	"embed"
	"io/fs"
	"log"
	"path"
)

//go:embed resources
var rawEmbed embed.FS

var (
	Embed fs.FS
	Images fs.FS
	Fonts fs.FS
)

func init() {
	var err error
	Embed, err = fs.Sub(rawEmbed, "resources")
	if err != nil {
		log.Fatal(err)
	}
	Images, err = fs.Sub(rawEmbed, path.Join("resources", "images"))
	if err != nil {
		log.Fatal(err)
	}
	Fonts, err = fs.Sub(rawEmbed, path.Join("resources", "fonts"))
	if err != nil {
		log.Fatal(err)
	}
}
