package magnet

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	score int
	enemyX float64
	isStart bool
)

func (s *Battle) Update(g *Game)  {
	s.tick++
	if !countDownPlayer.IsPlaying() && !isStart {
		countDownPlayer.SetVolume(0.8)
		countDownPlayer.Rewind()
		countDownPlayer.Play()
		isStart = true
	}
	if s.tick / 60 > 3 {
		if !bgmPlayer.IsPlaying() {
			bgmPlayer.SetVolume(0.1)
			bgmPlayer.Rewind()
			bgmPlayer.Play()
		}
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
			if !slidePlayer.IsPlaying() {
				slidePlayer.SetVolume(0.8)
				slidePlayer.Rewind()
				slidePlayer.Play()
			}
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
				if !jumpPlayer.IsPlaying() {
					jumpPlayer.SetVolume(0.8)
					jumpPlayer.Rewind()
					jumpPlayer.Play()
				}
			}
			if x >= 1.7 {
				player.leftUp.Y  += 4.5
				if player.leftUp.Y >=  playerFootY - player.frameSize.Y {
					player.isJump = false
					player.leftUp.Y = playerFootY - player.frameSize.Y
				}
				if !failPlayer.IsPlaying() && player.isJump {
					failPlayer.SetVolume(0.8)
					failPlayer.Rewind()
					failPlayer.Play()
				}
			}
		}
		m := maps.Maps[stage]
		for i, o := range m.Objects {
			maps.Maps[stage].Objects[i].positionX = o.X + float64(ScreenWidth / 2 - (s.tick * 4))
			if isCollision(
				Point{o.collisionLeftUp.X + maps.Maps[stage].Objects[i].positionX, playerFootY - o.Y + o.collisionLeftUp.Y},
				o.collisionRightDown,
			) && !player.objectHit[stage][i] {
				player.objectHit[stage][i] = true
				enemyX += 90
				if !damagePlayer.IsPlaying() {
					damagePlayer.SetVolume(0.8)
					damagePlayer.Rewind()
					damagePlayer.Play()
				}
			}
		}
		for i, o := range m.Objects {
			if !player.objectHit[stage][i] && !o.isScored && player.leftUp.X > o.positionX + o.collisionRightDown.X {
				score += 100
				maps.Maps[stage].Objects[i].isScored = true
			}
		}
	}
}

func (s *Battle) Draw(screen *ebiten.Image)  {
	bgFirstOption := &ebiten.DrawImageOptions{}
	bgSecondOption := &ebiten.DrawImageOptions{}
	bgFirstOption.GeoM.Translate(s.backgroundX, 0)
	bgSecondOption.GeoM.Translate(s.backgroundX + float64(ScreenWidth), 0)
	screen.DrawImage(backgroundImage, bgFirstOption)
	screen.DrawImage(backgroundImage, bgSecondOption)

	m := maps.Maps[stage]
	for _, o := range m.Objects {
		objOption := &ebiten.DrawImageOptions{}
		objOption.GeoM.Translate(o.positionX, playerFootY - o.Y)
		collisionOption := &ebiten.DrawImageOptions{}
		collisionOption.GeoM.Translate(o.collisionLeftUp.X + o.positionX, playerFootY - o.Y + o.collisionLeftUp.Y)
		switch (o.ObjectType) {
		case 1:
			screen.DrawImage(objct1Image, objOption)
		case 2:
			screen.DrawImage(objct2Image, objOption)
		case 3:
			screen.DrawImage(objct3Image, objOption)
		case 4:
			screen.DrawImage(objct4Image, objOption)
		}
		//screen.DrawImage(
		//	collisionImage.SubImage(
		//		image.Rect(0, 0, int(o.collisionRightDown.X), int(o.collisionRightDown.Y)),
		//	).(*ebiten.Image),
		//	collisionOption,
		//)
	}
	playerOption := &ebiten.DrawImageOptions{}
	playerOption.GeoM.Translate(player.leftUp.X, player.leftUp.Y)

	i := (s.tick / 5) % playerFrameNum
	sx, sy := i * playerFrameWidth, int(player.frame0.Y)
	enemySx, enemySy := i * playerFrameWidth, playerFrame0Y

	enemyOption := &ebiten.DrawImageOptions{}
	enemyOption.GeoM.Translate(enemyX, playerFootY - playerFrameHeight)
	screen.DrawImage(
		enemyRunImage.SubImage(
			image.Rect(enemySx, enemySy, enemySx + playerFrameWidth, enemySy + playerFrameHeight),
		).(*ebiten.Image),
		enemyOption,
	)

	if player.isJump {
		screen.DrawImage(
			playerRunImage.SubImage(
				image.Rect(4 * playerFrameWidth, 0, 5 * playerFrameWidth, playerFrameHeight),
			).(*ebiten.Image),
			playerOption,
		)
	} else if player.isSlide {
		screen.DrawImage(
			playerSlideImage.SubImage(
				image.Rect(sx, sy, sx + int(player.frameSize.X), sy + int(player.frameSize.Y)),
			).(*ebiten.Image),
			playerOption,
		)
	} else {
		screen.DrawImage(
			playerRunImage.SubImage(
				image.Rect(sx, sy, sx + int(player.frameSize.X), sy + int(player.frameSize.Y)),
			).(*ebiten.Image),
			playerOption,
		)
	}

	// playerCollisionOption := &ebiten.DrawImageOptions{}
	// playerCollisionOption.GeoM.Translate(player.leftUp.X, player.leftUp.Y)
	// playerCollisionOption.GeoM.Translate(player.collisionLeftUp.X, player.collisionLeftUp.Y)
	// screen.DrawImage(
	// 	collisionImage.SubImage(
	// 		image.Rect(0, 0, int(player.collisionRightDown.X - player.collisionLeftUp.X), int(player.collisionRightDown.Y)),
	// 	).(*ebiten.Image),
	// 	playerCollisionOption,
	// )

	countDownOption := &ebiten.DrawImageOptions{}
	if s.tick / 60 < 1 {
		screen.DrawImage(countDown3Image, countDownOption)
	} else if s.tick / 60 < 2 {
		screen.DrawImage(countDown2Image, countDownOption)
	} else if s.tick / 60 < 3 {
		screen.DrawImage(countDown1Image, countDownOption)
	}

	//ebitenutil.DebugPrint(screen, fmt.Sprintf("tick:%d\nFPS: %.2f\nPlayeFrameSizeY: %.2f\nPlayerLeftUp: %.2f", s.tick, ebiten.CurrentFPS(), player.frameSize.Y, player.leftUp.Y))

	displayRectangle := text.BoundString(Font, fmt.Sprintf("Score: %d", score))
	text.Draw(screen, fmt.Sprintf("Score: %d", score), Font, ScreenWidth - displayRectangle.Dx() - 10, displayRectangle.Dy() + 10, color.White)
}
