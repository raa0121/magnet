package magnet

import (
	"fmt"
	"image"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	resources "github.com/raa0121/magnet/magnet/internal"
)

type Battle struct {
	frame int
	backgroundX float64
}

const (
	playerFrame0X = 0
	playerFrame0Y = 0
	playerFrameWidth = 256
	playerFrameHeight = 256
	playerFrameNum = 8
)

var (
	BackgroundImage *ebiten.Image
	PlayerWaitImage *ebiten.Image
	PlayerRunImage *ebiten.Image
	ObjImage *ebiten.Image
)

var (
	playerLeftUp = Point{
		(ScreenWidth - playerFrameWidth) / 2,
		(720 - playerFrameHeight),
	}
	playerSize = Point{playerFrameWidth, playerFrameHeight}
	obj1LeftUp = Point{900, 700}
	obj1Size = Point{300, 300}
)

func init() {
	backgroundImageInit()
	playerRunImageInit()
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

func (s *Battle) Update(m *Game)  {
	s.frame++
	s.backgroundX -= 4
	if m.backgroundX == -float64(ScreenWidth) {
		m.backgroundX = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		m.SceneType.Type = SceneBattle
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		m.SceneType.Type = SceneBattle
	}
	if 0 < len(inpututil.JustPressedTouchIDs()) {
		m.SceneType.Type = SceneBattle
	}
}

func (s *Battle) Draw(screen *ebiten.Image)  {
	bgFirstOption := &ebiten.DrawImageOptions{}
	bgSecondOption := &ebiten.DrawImageOptions{}
	bgFirstOption.GeoM.Translate(s.backgroundX, 0)
	bgSecondOption.GeoM.Translate(s.backgroundX + float64(ScreenWidth), 0)
	screen.DrawImage(BackgroundImage, bgFirstOption)
	screen.DrawImage(BackgroundImage, bgSecondOption)

	playerOption := &ebiten.DrawImageOptions{}
	playerOption.GeoM.Translate(playerLeftUp.X, playerLeftUp.Y)

	i := (s.frame / 5) % playerFrameNum
	sx, sy := i*playerFrameWidth, playerFrame0Y
	screen.DrawImage(
		PlayerRunImage.SubImage(
			image.Rect(sx, sy, sx+playerFrameWidth, sy+playerFrameHeight),
		).(*ebiten.Image),
		playerOption,
	)

	objOption := &ebiten.DrawImageOptions{}
	objOption.GeoM.Translate(obj1LeftUp.X, obj1LeftUp.Y)
	screen.DrawImage(ObjImage, objOption)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %f, isCollision: %+v", ebiten.CurrentFPS(), isCollision(obj1LeftUp, obj1Size)))
}
