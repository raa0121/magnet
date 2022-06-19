package magnet

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	playerFrame0X = 0
	playerFrame0Y = 0
	playerFrameWidth = 256
	playerFrameHeight = 256
	playerFrameNum = 8
)

type Battle struct {
	tick int
	backgroundX float64
}

var (
	playerY = 0.0
	playerLeftUp = Point{
		(ScreenWidth - playerFrameWidth) / 2,
		(720 - playerFrameHeight - playerY),
	}
	playerSize = Point{playerFrameWidth, playerFrameHeight}
	jumpTick int
	isJump bool
	stage int
	score int
)

func (s *Battle) Update(g *Game)  {
	s.tick++
	s.backgroundX -= 4
	if g.backgroundX == -float64(ScreenWidth) {
		g.backgroundX = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) && !isJump {
		jumpTick = s.tick
		isJump = true
	}
	if isJump {
		x := float64(s.tick - jumpTick) / 60.0
		if x < 0.8 {
			playerY += 5
		}
		if x >= 1.2 {
			playerY -= 4
			if playerY < 0 {
				isJump = false
				playerY = 0
			}
		}
	}
	m := maps.Maps[stage]
	for i, o := range m.Objects {
		maps.Maps[stage].Objects[i].positionX = o.X + float64(ScreenWidth / 2 - (s.tick * 4))
		if isCollision(
			Point{o.collisionLeftUp.X + o.positionX, 720 - o.Y + o.collisionLeftUp.Y},
			Point{o.collisionRightDown.X + o.positionX, 720 - o.Y + o.collisionRightDown.Y},
		) {
			m.Objects[i].isHit = true
		}
	}
	for _, o := range m.Objects {
		if !o.isHit && playerLeftUp.X > o.positionX {
			fmt.Printf("score: %d\n", score)
			score += 100
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

	m := maps.Maps[stage]
	for _, o := range m.Objects {
		objOption := &ebiten.DrawImageOptions{}
		objOption.GeoM.Translate(o.positionX, 720 - o.Y)
		collisionOption := &ebiten.DrawImageOptions{}
		collisionOption.GeoM.Translate(o.collisionLeftUp.X + o.positionX, 720 - o.Y + o.collisionLeftUp.Y)
		switch (o.ObjectType) {
		case 1:
			screen.DrawImage(Objct1Image, objOption)
		case 2:
			screen.DrawImage(Objct2Image, objOption)
		}
		screen.DrawImage(
			CollisionImage.SubImage(
				image.Rect(0, 0, int(o.collisionRightDown.X), int(o.collisionRightDown.Y)),
			).(*ebiten.Image),
			collisionOption,
		)
	}
	playerOption := &ebiten.DrawImageOptions{}
	playerOption.GeoM.Translate(playerLeftUp.X, playerLeftUp.Y - playerY)

	i := (s.tick / 5) % playerFrameNum
	sx, sy := i * playerFrameWidth, playerFrame0Y

	enemyOption := &ebiten.DrawImageOptions{}
	enemyOption.GeoM.Translate(0, playerLeftUp.Y)
	screen.DrawImage(
		EnemyRunImage.SubImage(
			image.Rect(sx, sy, sx + playerFrameWidth, sy + playerFrameHeight),
		).(*ebiten.Image),
		enemyOption,
	)

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
	screen.DrawImage(CollisionImage, playerOption)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.2f\nPlayeY: %.2f", ebiten.CurrentFPS(), playerY))
}
