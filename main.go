package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct {
	player *Player
}

func (game *Game) Update() error {
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.player.Draw(screen)
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (ScreenWidth, ScreenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := &Game{}
	game.player = newPlayer(game)

	err := ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}
