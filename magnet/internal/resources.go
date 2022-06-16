package internal

import (
	"embed"
	"image/png"
	"io/fs"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed resources
var rawEmbed embed.FS

var (
	Embed fs.FS
	BackgroundImage *ebiten.Image
)

func init() {
	var err error
	Embed, err = fs.Sub(rawEmbed, "resources")
	if err != nil {
		log.Fatal(err)
	}
	backgroundImageInit()
}

func backgroundImageInit() {
	b, err := Embed.Open("bg.png")
	if err != nil {
		log.Fatal(err)
	}
	p, err := png.Decode(b)
	if err != nil {
		log.Fatal(err)
	}
	BackgroundImage = ebiten.NewImageFromImage(p)
}
