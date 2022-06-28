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
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("magnet")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
