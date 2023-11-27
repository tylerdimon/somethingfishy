package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tylerdimon/somethingfishy/sprites"
	"github.com/tylerdimon/somethingfishy/utils"
	"time"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Game struct {
	PlayerPosition utils.Vector
	AttackTimer    *utils.Timer
	Player         *sprites.Player
	OppSpawnTimer  *utils.Timer
	Opps           []*sprites.Opp
}

func NewGame() *Game {
	return &Game{
		PlayerPosition: utils.Vector{X: 100, Y: 100},
		AttackTimer:    utils.NewTimer(5 * time.Second),
		OppSpawnTimer:  utils.NewTimer(5 * time.Second),
		Player:         sprites.NewPlayer(ScreenWidth, ScreenHeight),
	}
}

func (g *Game) Update() error {
	err := g.Player.Update()
	if err != nil {
		panic(err)
	}

	g.OppSpawnTimer.Update()
	if g.OppSpawnTimer.IsReady() {
		g.OppSpawnTimer.Reset()

		m := sprites.NewOpp(ScreenWidth, ScreenHeight)
		g.Opps = append(g.Opps, m)
	}

	for _, o := range g.Opps {
		o.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.Draw(screen)

	for _, o := range g.Opps {
		o.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
