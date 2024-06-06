package game

import "github.com/hajimehoshi/ebiten/v2"

const (
	screenHeight = 800
	screenWidth  = 600
)

type Game struct {
	environments []*Environment
}

func NewGame() *Game {
	g := &Game{}

	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, e := range g.environments {
		e.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
