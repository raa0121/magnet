package magnet

import "github.com/hajimehoshi/ebiten/v2"

const (
	SceneTitle = iota
	SceneHowTo
	SceneBattle
	SceneGameOver
)

type SceneType struct {
	Type int
}
type Scene interface {
	Update(g *Game) error
	Draw(screen *ebiten.Image)
}

