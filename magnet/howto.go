package magnet

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	resources "github.com/raa0121/magnet/magnet/internal"

	"image"
	"log"
)

type HowTo struct {
}

var (
	howToBackgroundImg *ebiten.Image
)

func init() {
	howToBackgroundInit()
}

func howToBackgroundInit() {// {{{
	b, err := resources.Images.Open("howto.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(b)
	if err != nil {
		log.Fatal(err)
	}
	howToBackgroundImg = ebiten.NewImageFromImage(img)
}// }}}

func (s *HowTo) Update(m *Game)  {
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

func (s *HowTo) Draw(screen *ebiten.Image)  {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(howToBackgroundImg,op)
}
