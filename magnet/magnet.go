package magnet

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth = 1920
	ScreenHeight = 1080
)

type Game struct{
	SceneType
	Title
	HowTo
	Battle
	GameOver
}

type Point struct {
	X, Y float64
}

func (g *Game) Update() error {
	switch g.SceneType.Type {
	case SceneTitle:
		g.Title.Update(g)
	case SceneHowTo:
		g.HowTo.Update(g)
	case SceneBattle:
		g.Battle.Update(g)
	case SceneGameOver:
		g.GameOver.Update(g)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.SceneType.Type {
	case SceneTitle:
		g.Title.Draw(screen)
	case SceneHowTo:
		g.HowTo.Draw(screen)
	case SceneBattle:
		g.Battle.Draw(screen)
	case SceneGameOver:
		g.GameOver.Draw(screen)
	}
}

func (g *Game) Layout(width, height int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) init() {
	g.SceneType = SceneType{
		Type: SceneTitle,
	}
	g.Title = Title{}
	g.HowTo = HowTo{}
	g.Battle = Battle{}
	g.GameOver = GameOver{}
}

func NewGame() (*Game, error) {
	g := &Game{}
	g.init()
	return g, nil
}

func isCollision(objLeftUp, objSize Point) bool {
	if math.Abs((player.collisionLeftUp.X + player.collisionRightDown.X / 2) - (objLeftUp.X + objSize.X / 2)) < player.collisionRightDown.X / 2 + objSize.X / 2 &&
		math.Abs((player.collisionLeftUp.Y + player.collisionRightDown.Y / 2) - (objLeftUp.Y + objSize.Y / 2)) < player.collisionRightDown.Y / 2 + objSize.Y / 2 {
			return true
		}
	return false
}
