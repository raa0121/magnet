package magnet

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type GameOver struct {
}

func (s *GameOver) Update(g *Game)  {
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		score = 0
		enemyX = 0
		tick = 0
		backgroundX = 0
		isStart = false
		playerInit()
		mapInit()
		g.SceneType.Type = SceneTitle
	}
}

func (s *GameOver) Draw(screen *ebiten.Image)  {
	displayRectangle := text.BoundString(Font, fmt.Sprintf("Your Score: %d\nPress Z.", score))
	text.Draw(screen, fmt.Sprintf("Your Score: %d\nPress Z.", score), Font, ScreenWidth / 2 - displayRectangle.Dx() / 2, ScreenHeight / 2 - displayRectangle.Dy() / 2, color.White)
}
