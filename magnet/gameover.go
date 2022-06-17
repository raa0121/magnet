package magnet

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameOver struct {
}

func (s *GameOver) Update(g *Game)  {
	if ebiten.IsKeyPressed(ebiten.KeyZ) {
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.SceneType.Type = SceneTitle
	}
}

func (s *GameOver) Draw(screen *ebiten.Image)  {
}
