package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-asteroids/assets"
)

type Player struct {
	sprite *ebiten.Image
}

func newPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite

	player := &Player{
		sprite: sprite,
	}
	return player
}

func (player *Player) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	screen.DrawImage(player.sprite, options)
}

func (player *Player) Update() {
	
}
