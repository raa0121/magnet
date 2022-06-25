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
	BackgroundImage *ebiten.Image
	PlayerWaitImage *ebiten.Image
	PlayerRunImage *ebiten.Image
	PlayerSlideImage *ebiten.Image
	EnemyWaitImage *ebiten.Image
	EnemyRunImage *ebiten.Image
	Objct1Image *ebiten.Image
	Objct2Image *ebiten.Image
	Objct3Image *ebiten.Image
	Objct4Image *ebiten.Image
	CollisionImage *ebiten.Image
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
	BackgroundImage = ebiten.NewImageFromImage(p)
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
	PlayerWaitImage = ebiten.NewImageFromImage(p)
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
	PlayerRunImage = ebiten.NewImageFromImage(p)
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
	PlayerSlideImage = ebiten.NewImageFromImage(p)
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
	EnemyWaitImage = ebiten.NewImageFromImage(p)
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
	EnemyRunImage = ebiten.NewImageFromImage(p)
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
	Objct1Image = ebiten.NewImageFromImage(p)
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
	Objct2Image = ebiten.NewImageFromImage(p)
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
	Objct3Image = ebiten.NewImageFromImage(p)
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
	Objct4Image = ebiten.NewImageFromImage(p)
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
	CollisionImage = ebiten.NewImageFromImage(p)
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
