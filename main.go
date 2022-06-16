package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/raa0121/magnet/magnet"
)

func main() {
	game, err := magnet.NewGame()
	if err != nil {
		panic(err)
	}
	fullScreenWidth, fullScreenHeight := ebiten.ScreenSizeInFullscreen()
	if fullScreenWidth <= magnet.ScreenWidth || fullScreenHeight <= magnet.ScreenHeight {
		ebiten.SetWindowSize(magnet.ScreenWidth / 2, magnet.ScreenHeight / 2)
	} else {
		ebiten.SetWindowSize(magnet.ScreenWidth, magnet.ScreenHeight)
	}
	ebiten.SetWindowTitle("magnet")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
