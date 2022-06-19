package magnet

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Battle struct {
	tick int
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
	playerY = 0.0
	playerLeftUp = Point{
		(ScreenWidth - playerFrameWidth) / 2,
		(720 - playerFrameHeight),
	}
	playerSize = Point{playerFrameWidth, playerFrameHeight}
	obj1LeftUp = Point{900, 700}
	obj1Size = Point{300, 300}
	jumpTick int
	isJump bool
)

func (s *Battle) Update(m *Game)  {
	s.tick++
	s.backgroundX -= 4
	if m.backgroundX == -float64(ScreenWidth) {
		m.backgroundX = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		jumpTick = s.tick
		isJump = true
	}
	if isJump {
		x := float64(s.tick - jumpTick) / 60.0
		if x < 1 {
			playerY += 3
		}
		if x >= 1 {
			playerY -= 3
			if playerY < 0 {
				isJump = false
				playerY = 0
			}
		}
	}
}

func (s *Battle) Draw(screen *ebiten.Image)  {
	bgFirstOption := &ebiten.DrawImageOptions{}
	bgSecondOption := &ebiten.DrawImageOptions{}
	bgFirstOption.GeoM.Translate(s.backgroundX, 0)
	bgSecondOption.GeoM.Translate(s.backgroundX + float64(ScreenWidth), 0)
	screen.DrawImage(BackgroundImage, bgFirstOption)
	screen.DrawImage(BackgroundImage, bgSecondOption)

	for _, m := range maps.Maps {
		for _, o := range m.Objects {
			objOption := &ebiten.DrawImageOptions{}
			objOption.GeoM.Translate(o.X, -o.Y)
			objOption.GeoM.Translate(float64(ScreenWidth / 2 - (s.tick * 4)), 720)
			switch (o.ObjectType) {
			case 1:
				screen.DrawImage(Objct1Image, objOption)
			case 2:
				screen.DrawImage(Objct2Image, objOption)
			}
		}
	}
	playerOption := &ebiten.DrawImageOptions{}
	playerOption.GeoM.Translate(playerLeftUp.X, playerLeftUp.Y - playerY)

	i := (s.tick / 5) % playerFrameNum
	sx, sy := i*playerFrameWidth, playerFrame0Y
	if isJump {
		screen.DrawImage(
			PlayerRunImage.SubImage(
				image.Rect(4 * playerFrameWidth, 0, 5 * playerFrameWidth, playerFrameHeight),
			).(*ebiten.Image),
			playerOption,
		)
	} else {
		screen.DrawImage(
			PlayerRunImage.SubImage(
				image.Rect(sx, sy, sx + playerFrameWidth, sy + playerFrameHeight),
			).(*ebiten.Image),
			playerOption,
		)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.2f\nPlayeX: %.2f", ebiten.CurrentFPS(), playerY))
}
