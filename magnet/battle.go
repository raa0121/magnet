package magnet

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	playerLeftUp = Point{
		(ScreenWidth - playerFrameWidth) / 2,
		(720 - playerFrameHeight),
	}
	playerSize = Point{playerFrameWidth, playerFrameHeight}
	obj1LeftUp = Point{900, 700}
	obj1Size = Point{300, 300}
)

func (s *Battle) Update(m *Game)  {
	s.frame++
	s.backgroundX -= 4
	if m.backgroundX == -float64(ScreenWidth) {
		m.backgroundX = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		m.SceneType.Type = SceneGameOver
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		m.SceneType.Type = SceneGameOver
	}
	if 0 < len(inpututil.JustPressedTouchIDs()) {
		m.SceneType.Type = SceneGameOver
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

	for _, m := range maps.Maps {
		for _, o := range m.Objects {
			objOption := &ebiten.DrawImageOptions{}
			objOption.GeoM.Translate(o.X, -o.Y)
			objOption.GeoM.Translate(float64(ScreenWidth / 2 - (s.frame * 4)), 720)
			switch (o.ObjectType) {
			case 1:
				screen.DrawImage(ObjImage, objOption)
			}
		}
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %f", ebiten.CurrentFPS()))
}
