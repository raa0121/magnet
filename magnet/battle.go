package magnet

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Battle struct {
	tick int
	backgroundX float64
}

var (
	player *Player
	jumpTick int
	slideTick int
	stage int
	score int = 100
)

func (s *Battle) Update(g *Game)  {
	s.tick++
	s.backgroundX -= 4
	if g.backgroundX == -float64(ScreenWidth) {
		g.backgroundX = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) && !player.isJump && !player.isSlide {
		jumpTick = s.tick
		player.isJump = true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyX) && !player.isSlide && !player.isJump {
		player.frame0.Y = 60
		player.frameSize.Y =  playerFrameHeight - 60
		player.leftUp.Y = playerFootY - playerFrameHeight + 60
		player.isSlide = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyX) {
		player.frame0.Y = 0
		player.frameSize.Y = playerFrameHeight
		player.leftUp.Y = playerFootY - player.frameSize.Y
		player.isSlide = false
	}
	if player.isJump {
		x := float64(s.tick - jumpTick) / 60.0
		if x < 0.8 {
			player.leftUp.Y -= 6
		}
		if x >= 1.7 {
			player.leftUp.Y  += 4
			if player.leftUp.Y >  playerFootY - player.frameSize.Y {
				player.isJump = false
				player.leftUp.Y = playerFootY - player.frameSize.Y
			}
		}
	}
	m := maps.Maps[stage]
	for i, o := range m.Objects {
		maps.Maps[stage].Objects[i].positionX = o.X + float64(ScreenWidth / 2 - (s.tick * 4))
		if isCollision(
			Point{o.collisionLeftUp.X + o.positionX, playerFootY - o.Y + o.collisionLeftUp.Y},
			Point{o.collisionRightDown.X + o.positionX, playerFootY - o.Y + o.collisionRightDown.Y},
		) {
			m.Objects[i].isHit = true
			fmt.Printf("tick:%d Object[%d] is Hit\n", s.tick, i)
		}
	}
	for i, o := range m.Objects {
		if !o.isHit && player.leftUp.X > o.positionX {
			fmt.Printf("score: %d\n", score)
			score += 100
			m.Objects[i].isHit = true
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
		objOption.GeoM.Translate(o.positionX, playerFootY - o.Y)
		collisionOption := &ebiten.DrawImageOptions{}
		collisionOption.GeoM.Translate(o.collisionLeftUp.X + o.positionX, playerFootY - o.Y + o.collisionLeftUp.Y)
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
	playerOption.GeoM.Translate(player.leftUp.X, player.leftUp.Y)

	i := (s.tick / 5) % playerFrameNum
	sx, sy := i * playerFrameWidth, int(player.frame0.Y)
	enemySx, enemySy := i * playerFrameWidth, playerFrame0Y

	enemyOption := &ebiten.DrawImageOptions{}
	enemyOption.GeoM.Translate(0, playerFootY - playerFrameHeight)
	screen.DrawImage(
		EnemyRunImage.SubImage(
			image.Rect(enemySx, enemySy, enemySx + playerFrameWidth, enemySy + playerFrameHeight),
		).(*ebiten.Image),
		enemyOption,
	)

	if player.isJump {
		screen.DrawImage(
			PlayerRunImage.SubImage(
				image.Rect(4 * playerFrameWidth, 0, 5 * playerFrameWidth, playerFrameHeight),
			).(*ebiten.Image),
			playerOption,
		)
	} else if player.isSlide {
		screen.DrawImage(
			PlayerSlideImage.SubImage(
				image.Rect(sx, sy, sx + int(player.frameSize.X), sy + int(player.frameSize.Y)),
			).(*ebiten.Image),
			playerOption,
		)
	} else {
		screen.DrawImage(
			PlayerRunImage.SubImage(
				image.Rect(sx, sy, sx + int(player.frameSize.X), sy + int(player.frameSize.Y)),
			).(*ebiten.Image),
			playerOption,
		)
	}
	screen.DrawImage(
		CollisionImage.SubImage(
			image.Rect(0, 0, int(player.frameSize.X), int(player.frameSize.Y)),
		).(*ebiten.Image),
		playerOption,
	)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("tick:%d\nFPS: %.2f\nPlayeFrameSizeY: %.2f\nPlayerLeftUp: %.2f", s.tick, ebiten.CurrentFPS(), player.frameSize.Y, player.leftUp.Y))

	displayRectangle := text.BoundString(Font, fmt.Sprintf("Score: %d", score))
	text.Draw(screen, fmt.Sprintf("Score: %d", score), Font, ScreenWidth - displayRectangle.Dx() - 10, displayRectangle.Dy() + 10, color.White)
}
