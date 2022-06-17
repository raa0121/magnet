package magnet

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	resources "github.com/raa0121/magnet/magnet/internal"
)

const (
	ScreenWidth = 1920
	ScreenHeight = 1080
	playerFrame0X = 0
	playerFrame0Y = 0
	playerFrameWidth = 512
	playerFrameHeight = 512
	playerFrameNum = 8
)

type Magnet struct{
	frame int
	backgroundX float64
}

func (m *Magnet) Update() error {
	m.frame++
	m.backgroundX -= 4
	if m.backgroundX == -float64(ScreenWidth) {
		m.backgroundX = 0
	}
	return nil
}

func (m *Magnet) Draw(screen *ebiten.Image) {
	bgFirstOption := &ebiten.DrawImageOptions{}
	bgSecondOption := &ebiten.DrawImageOptions{}
	bgFirstOption.GeoM.Translate(m.backgroundX, 0)
	bgSecondOption.GeoM.Translate(m.backgroundX + float64(ScreenWidth), 0)
	screen.DrawImage(resources.BackgroundImage, bgFirstOption)
	screen.DrawImage(resources.BackgroundImage, bgSecondOption)

	playerOption := &ebiten.DrawImageOptions{}
	playerOption.GeoM.Translate(-float64(playerFrameWidth)/2, -float64(playerFrameHeight)/2)
	playerOption.GeoM.Translate(ScreenWidth/2, ScreenHeight/2)

	i := (m.frame / 5) % playerFrameNum
	sx, sy := i*playerFrameWidth, playerFrame0Y
	screen.DrawImage(
		resources.PlayerImage.SubImage(
			image.Rect(sx, sy, sx+playerFrameWidth, sy+playerFrameHeight),
		).(*ebiten.Image),
		playerOption,
	)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %f", ebiten.CurrentFPS()))
}

func (m *Magnet) Layout(width, height int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() (*Magnet, error) {
	return &Magnet{}, nil
}
