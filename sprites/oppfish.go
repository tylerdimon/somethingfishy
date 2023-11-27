package sprites

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tylerdimon/somethingfishy/utils"
	"math"
	"math/rand"
)

type Opp struct {
	position      utils.Vector
	movement      utils.Vector
	sprite        *ebiten.Image
	rotationSpeed float64
	rotation      float64
}

func NewOpp(screenWidth, screenHight float64) *Opp {
	sprite := OppSprites[rand.Intn(len(OppSprites))]

	// Figure out the target position — the screen center, in this case
	target := utils.Vector{
		X: screenWidth / 2,
		Y: screenHight / 2,
	}

	// The distance from the center the meteor should spawn at — half the width
	r := screenWidth / 2.0

	// Pick a random angle — 2π is 360° — so this returns 0° to 360°
	angle := rand.Float64() * 2 * math.Pi

	// Figure out the spawn position by moving r pixels from the target at the chosen angle
	pos := utils.Vector{
		X: target.X + math.Cos(angle)*r,
		Y: target.Y + math.Sin(angle)*r,
	}

	// Randomized velocity
	velocity := 0.25 + rand.Float64()*1.5

	// Direction is the target minus the current position
	direction := utils.Vector{
		X: target.X - pos.X,
		Y: target.Y - pos.Y,
	}

	// Normalize the vector — get just the direction without the length
	normalizedDirection := direction.Normalize()

	// Multiply the direction by velocity
	movement := utils.Vector{
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}

	rotationSpeed := -0.02 + rand.Float64()*0.04

	return &Opp{
		position:      utils.Vector{},
		sprite:        sprite,
		movement:      movement,
		rotationSpeed: rotationSpeed,
	}
}

func (m *Opp) Update() {
	m.position.X += m.movement.X
	m.position.Y += m.movement.Y
	m.rotation += m.rotationSpeed
}

func (o *Opp) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(o.position.X, o.position.Y)
	screen.DrawImage(o.sprite, op)
}
