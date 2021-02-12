package enemies

import "github.com/hajimehoshi/ebiten"

//IEnemy interfaz de los enemigos.
type IEnemy interface {
	Update(dt float64)
	Draw(screen *ebiten.Image)
	Position() (float64, float64)
}
