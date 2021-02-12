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
}

//NewMenu es un contructor
func NewMenu(fontsloader *utils.FontsLoader) states.IState {
	m := &Menu{}
	m.nextStateID = "menu"

	//Creamos el UI del juego (TODO: colocar iconos)
	m.uimenu = NewUIMenu(fontsloader)
	return m
}

/*===========================================================================*/
/*                               Interface IState                            */
/*===========================================================================*/

//Init ...
func (m *Menu) Init() {
	m.nextStateID = "menu"
}

//ProcessEvents procesa los eventos del juego.
func (m *Menu) ProcessEvents() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		m.nextStateID = "gameplay"
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		m.nextStateID = "credits"
	}

}

//Update actualiza la lógica de los creditos.
func (m *Menu) Update(dt float64) {

}

//Draw draws the game.
func (m *Menu) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0xCE, 0x9C, 0x72, 0xff})
	m.uimenu.Draw(screen)
}

//NextState ...
func (m *Menu) NextState() string {
	return m.nextStateID
}
