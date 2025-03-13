package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"go-asteroids/assets"
	"math"
)

const (
	rotationPerSecond = math.Pi
	maxAcceleration   = 8.0
	ScreenWidth       = 1280
	ScreenHeight      = 720
)

var curAcceleration float64

type Player struct {
	game           *Game
	sprite         *ebiten.Image
	rotation       float64
	position       Vector
	playerVelocity float64
}

func newPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite

	// Center player on screen
	bounds := sprite.Bounds()
	halfWidth := float64(bounds.Dx()) / 2
	halfHeight := float64(bounds.Dy()) / 2

	position := Vector{
		X: ScreenWidth/2 - halfWidth,
		Y: ScreenHeight/2 - halfHeight,
	}

	player := &Player{
		sprite:   sprite,
		game:     game,
		position: position,
	}

	return player
}

func (player *Player) Draw(screen *ebiten.Image) {
	bounds := player.sprite.Bounds()
	halfWidth := float64(bounds.Dx()) / 2
	halfHeight := float64(bounds.Dy()) / 2

	options := &ebiten.DrawImageOptions{}

	options.GeoM.Translate(-halfWidth, -halfHeight)
	options.GeoM.Rotate(player.rotation)
	options.GeoM.Translate(+halfWidth, +halfHeight)

	options.GeoM.Translate(player.position.X, player.position.Y)

	screen.DrawImage(player.sprite, options)
}

func (player *Player) Update() {
	speed := rotationPerSecond / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		fmt.Println("A pressed")
		player.rotation -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		fmt.Println("D pressed")
		player.rotation += speed
	}

	player.Accelerate()
}

func (player *Player) Accelerate() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		player.keepOnScreen()

		if curAcceleration < maxAcceleration {
			curAcceleration = player.playerVelocity + 4
		}

		if curAcceleration >= 8 {
			curAcceleration = 8
		}
		player.playerVelocity = curAcceleration

		// Move in the direction we are pointing
		directionX := math.Sin(player.rotation) * curAcceleration
		directionY := math.Cos(player.rotation) * -curAcceleration

		// Move the player on the screen
		player.position.X += directionX
		player.position.Y += directionY
	}
}

func (player *Player) keepOnScreen() {
	if player.position.X >= float64(ScreenWidth) {
		player.position.X = 0
	}

	if player.position.X < 0 {
		player.position.X = ScreenWidth
	}

	if player.position.Y >= float64(ScreenHeight) {
		player.position.Y = 0
	}

	if player.position.Y < 0 {
		player.position.Y = ScreenHeight
	}

}
