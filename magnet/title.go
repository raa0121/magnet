package magnet

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	resources "github.com/raa0121/magnet/magnet/internal"

	"image"
	"log"
)

type Title struct {
}

var (
	titleBackgroundImg *ebiten.Image
)

func init() {
	titleBackgroundInit()
}

func titleBackgroundInit() {
	b, err := resources.Images.Open("title.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(b)
	if err != nil {
		log.Fatal(err)
	}
	titleBackgroundImg = ebiten.NewImageFromImage(img)
}

func (s *Title) Update(m *Game)  {
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		m.SceneType.Type = SceneHowTo
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		m.SceneType.Type = SceneHowTo
	}
	if 0 < len(inpututil.JustPressedTouchIDs()) {
		m.SceneType.Type = SceneHowTo
	}
}

func (s *Title) Draw(screen *ebiten.Image)  {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(titleBackgroundImg, op)
}

