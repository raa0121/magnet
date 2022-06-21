package magnet

import (
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	resources "github.com/raa0121/magnet/magnet/internal"
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
	CollisionImage *ebiten.Image
)



func init() {
	backgroundImageInit()
	playerRunImageInit()
	playerSlideImageInit()
	enemyRunImageInit()
	object1ImageInit()
	object2ImageInit()
	collisionImageInit()
}

func backgroundImageInit() {
	b, err := resources.Embed.Open("bg.png")
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
	b, err := resources.Embed.Open("player_wait.png")
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
	b, err := resources.Embed.Open("player_run.png")
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
	b, err := resources.Embed.Open("player_slide.png")
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
	b, err := resources.Embed.Open("enemy_wait.png")
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
	b, err := resources.Embed.Open("enemy_run.png")
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
	b, err := resources.Embed.Open("object_usb.png")
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
	b, err := resources.Embed.Open("object_lan.png")
	if err != nil {
		log.Fatal(err)
	}
	p, err := png.Decode(b)
	if err != nil {
		log.Fatal(err)
	}
	Objct2Image = ebiten.NewImageFromImage(p)
}

func collisionImageInit() {
	b, err := resources.Embed.Open("obj.png")
	if err != nil {
		log.Fatal(err)
	}
	p, err := png.Decode(b)
	if err != nil {
		log.Fatal(err)
	}
	CollisionImage = ebiten.NewImageFromImage(p)
}
