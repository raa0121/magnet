package magnet

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	resources "github.com/raa0121/magnet/magnet/internal"
)

const (
	ScreenWidth = 1920
	ScreenHeight = 1080
)

type Magnet struct{
	frame float64
	backgroundX float64
}

func (m *Magnet) Update() error {
	m.frame += 1
	m.backgroundX -= 4
	if m.backgroundX == -float64(ScreenWidth) {
		m.backgroundX = 0
	}
	return nil
}

func (m *Magnet) Draw(screen *ebiten.Image) {
	firstOption := &ebiten.DrawImageOptions{}
	secondOption := &ebiten.DrawImageOptions{}
	firstOption.GeoM.Translate(m.backgroundX, 0)
	secondOption.GeoM.Translate(m.backgroundX + float64(ScreenWidth), 0)
	screen.DrawImage(resources.BackgroundImage, firstOption)
	screen.DrawImage(resources.BackgroundImage, secondOption)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %f", ebiten.CurrentFPS()))
}

func (m *Magnet) Layout(width, height int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() (*Magnet, error) {
	return &Magnet{}, nil
}
