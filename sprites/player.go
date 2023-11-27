package sprites

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tylerdimon/somethingfishy/utils"
	"math"
)

type Player struct {
	position utils.Vector
	sprite   *ebiten.Image
	rotation float64
}

func NewPlayer(screenWidth, screenHight float64) *Player {
	sprite := PlayerSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos := utils.Vector{
		X: screenWidth/2 - halfW,
		Y: screenHight/2 - halfH,
	}

	return &Player{
		position: pos,
		sprite:   sprite,
	}
}

func (p *Player) Update() error {
	rotateSpeed := math.Pi / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		p.rotation -= rotateSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		p.rotation += rotateSpeed
	}

	moveSpeed := 5.0

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.position.Y += moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.position.Y -= moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.position.X -= moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.position.X += moveSpeed
	}

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, op)
}
