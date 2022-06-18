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
	EnemyWaitImage *ebiten.Image
	EnemyRunImage *ebiten.Image
	ObjImage *ebiten.Image
)



func init() {
	backgroundImageInit()
	playerRunImageInit()
	enemyRunImageInit()
	objImageInit()
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

func objImageInit() {
	b, err := resources.Embed.Open("obj.png")
	if err != nil {
		log.Fatal(err)
	}
	p, err := png.Decode(b)
	if err != nil {
		log.Fatal(err)
	}
	ObjImage = ebiten.NewImageFromImage(p)
}
