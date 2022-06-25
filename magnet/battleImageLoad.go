package magnet

import (
	"image/png"
	"io/fs"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	resources "github.com/raa0121/magnet/magnet/internal"
	"golang.org/x/image/font"
)

var (
	backgroundImage *ebiten.Image
	playerWaitImage *ebiten.Image
	playerRunImage *ebiten.Image
	playerSlideImage *ebiten.Image
	enemyWaitImage *ebiten.Image
	enemyRunImage *ebiten.Image
	objct1Image *ebiten.Image
	objct2Image *ebiten.Image
	objct3Image *ebiten.Image
	objct4Image *ebiten.Image
	collisionImage *ebiten.Image
	Font font.Face
)

func init() {
	backgroundImageInit()
	playerRunImageInit()
	playerSlideImageInit()
	enemyRunImageInit()
	object1ImageInit()
	object2ImageInit()
	object3ImageInit()
	object4ImageInit()
	collisionImageInit()
	fontInit()
}

func backgroundImageInit() {
	b, err := resources.Images.Open("bg.png")
	if err != nil {
		log.Fatal(err)
	}
	p, err := png.Decode(b)
	if err != nil {
		log.Fatal(err)
	}
	backgroundImage = ebiten.NewImageFromImage(p)
}

func playerWaitImageInit() {
	b, err := resources.Images.Open("player_wait.png")
	if err != nil {
		log.Fatal(err)
	}
	p, err := png.Decode(b)
	if err != nil {
		log.Fatal(err)
	}
	playerWaitImage = ebiten.NewImageFromImage(p)
}

func playerRunImageInit() {
	b, err := resources.Images.Open("player_run.png")
	if err != nil {
		log.Fatal(err)
	}
	p, err := png.Decode(b)
	if err != nil {
		log.Fatal(err)
	}
	playerRunImage = ebiten.NewImageFromImage(p)
}

func playerSlideImageInit() {
	b, err := resources.Images.Open("player_slide.png")
	if err != nil {
		log.Fatal(err)
	}
	p, err := png.Decode(b)
	if err != nil {
		log.Fatal(err)
	}
	playerSlideImage = ebiten.NewImageFromImage(p)
}

func enemyWaitImageInit() {
	b, err := resources.Images.Open("enemy_wait.png")
	if err != nil {
		log.Fatal(err)
	}
	p, err := png.Decode(b)
	if err != nil {
		log.Fatal(err)
	}
	enemyWaitImage = ebiten.NewImageFromImage(p)
}

func enemyRunImageInit() {
	b, err := resources.Images.Open("enemy_run.png")
	if err != nil {
		log.Fatal(err)
	}
	p, err := png.Decode(b)
	if err != nil {
		log.Fatal(err)
	}
	enemyRunImage = ebiten.NewImageFromImage(p)
}

func object1ImageInit() {
	b, err := resources.Images.Open("object_usb.png")
	if err != nil {
		log.Fatal(err)
	}
	p, err := png.Decode(b)
	if err != nil {
		log.Fatal(err)
	}
	objct1Image = ebiten.NewImageFromImage(p)
}

func object2ImageInit() {
	b, err := resources.Images.Open("object_lan.png")
	if err != nil {
		log.Fatal(err)
	}
	p, err := png.Decode(b)
	if err != nil {
		log.Fatal(err)
	}
	objct2Image = ebiten.NewImageFromImage(p)
}

func object3ImageInit() {
	b, err := resources.Images.Open("object_lan2.png")
	if err != nil {
		log.Fatal(err)
	}
	p, err := png.Decode(b)
	if err != nil {
		log.Fatal(err)
	}
	objct3Image = ebiten.NewImageFromImage(p)
}

func object4ImageInit() {
	b, err := resources.Images.Open("object_display.png")
	if err != nil {
		log.Fatal(err)
	}
	p, err := png.Decode(b)
	if err != nil {
		log.Fatal(err)
	}
	objct4Image = ebiten.NewImageFromImage(p)
}

func collisionImageInit() {
	b, err := resources.Images.Open("obj.png")
	if err != nil {
		log.Fatal(err)
	}
	p, err := png.Decode(b)
	if err != nil {
		log.Fatal(err)
	}
	collisionImage = ebiten.NewImageFromImage(p)
}

func fontInit() {
	b, err := fs.ReadFile(resources.Fonts, "mplus-1p-regular.ttf")
	if err != nil {
		log.Fatal(err)
	}
	tt, err := truetype.Parse(b)
	if err != nil {
		log.Fatal(err)
	}
	Font = truetype.NewFace(tt, &truetype.Options{
		Size: 24,
		DPI: 72,
		Hinting: font.HintingFull,
	})
}
