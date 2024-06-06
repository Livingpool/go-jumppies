package game

import "github.com/hajimehoshi/ebiten/v2"

type Environment struct {
	position Vector
	sprite   *ebiten.Image
	size     string
	theme    string
}

const (
	upperSpawnBound = screenHeight / 2
)

func NewEnvironment() *Environment {
	return &Environment{}
}

func (e *Environment) Draw(screen *ebiten.Image) {
	bounds := e.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)

}
