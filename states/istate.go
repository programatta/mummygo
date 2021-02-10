package states

import "github.com/hajimehoshi/ebiten"

//IState ...
type IState interface {
	Init()
	ProcessEvents()
	Update(dt float64)
	Draw(screen *ebiten.Image)
	NextState() string
}
