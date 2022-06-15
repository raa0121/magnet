package magnet

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth = 1920
	ScreenHeight = 1080
)

type Magnet struct{}

func (m *Magnet) Update() error {
	return nil
}

func (m *Magnet) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %f", ebiten.CurrentFPS()))
}

func (m *Magnet) Layout(width, height int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() (*Magnet, error) {
	return &Magnet{}, nil
}
