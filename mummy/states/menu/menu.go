package menu

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/programatta/mummygo/states"
	"github.com/programatta/mummygo/utils"
)

//Menu contiene la funcionalidad para gestionra el menu principal del juego
type Menu struct {
	nextStateID string
	uimenu      *UIMenu
	alfa        float64
	state       tstate
}

//NewMenu es un contructor
func NewMenu(fontsloader *utils.FontsLoader) states.IState {
	m := &Menu{}
	//	m.nextStateID = "menu"

	//Creamos el UI del juego (TODO: colocar iconos)
	m.uimenu = NewUIMenu(fontsloader)
	m.alfa = 0
	m.state = enter
	return m
}

/*===========================================================================*/
/*                               Interface IState                            */
/*===========================================================================*/

//Init ...
func (m *Menu) Init() {
	// m.nextStateID = "menu"
	m.alfa = 0
	m.state = enter
}

//ProcessEvents procesa los eventos del juego.
func (m *Menu) ProcessEvents() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		m.nextStateID = "gameplay"
		m.state = exit
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		m.nextStateID = "credits"
		m.state = exit
	}

}

//Update actualiza la lógica de los creditos.
func (m *Menu) Update(dt float64) {
	if m.state == enter {
		m.alfa += dt
		if m.alfa > 1 {
			m.alfa = 1
			m.state = normal
		}
	} else if m.state == exit {
		m.alfa -= dt
		if m.alfa < 0 {
			m.alfa = 0
			m.state = end
		}
	}
}

//Draw draws the game.
func (m *Menu) Draw(screen *ebiten.Image) {
	halfa := 0xff * m.alfa

	screen.Fill(color.NRGBA{0xCE, 0x9C, 0x72, uint8(halfa)})
	m.uimenu.Draw(screen)
}

//NextState ...
func (m *Menu) NextState() string {
	if m.state == end {
		return m.nextStateID
	}

	return "menu"
}

//End ...
func (m *Menu) End() {

}

type tstate int

const (
	enter  tstate = tstate(0)
	normal tstate = tstate(1)
	exit   tstate = tstate(2)
	end    tstate = tstate(3)
)
