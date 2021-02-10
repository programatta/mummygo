package credits

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/programatta/mummygo/states"
)

//Credits contiene la funcionalidad del estado que muestra los creditos.
type Credits struct {
	nextStateID string
}

//NewCredits es un contructor
func NewCredits() states.IState {
	c := &Credits{}
	c.nextStateID = "credits"
	return c
}

/*===========================================================================*/
/*                               Interface IState                            */
/*===========================================================================*/

//Init ...
func (c *Credits) Init() {
	c.nextStateID = "credits"
}

//ProcessEvents procesa los eventos del juego.
func (c *Credits) ProcessEvents() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		c.nextStateID = "menu"
	}
}

//Update actualiza la lógica de los creditos.
func (c *Credits) Update(dt float64) {

}

//Draw draws the game.
func (c *Credits) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})
}

//NextState ...
func (c *Credits) NextState() string {
	return c.nextStateID
}
