package credits

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/programatta/mummygo/states"
	"github.com/programatta/mummygo/utils"
)

//Credits contiene la funcionalidad del estado que muestra los creditos.
type Credits struct {
	nextStateID string
	uicredits   *UICredits
}

//NewCredits es un contructor
func NewCredits(fontsloader *utils.FontsLoader) states.IState {
	c := &Credits{}
	c.nextStateID = "credits"

	c.uicredits = NewUICredits(fontsloader)
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
	screen.Fill(color.NRGBA{0xCE, 0x9C, 0x72, 0xff})
	c.uicredits.Draw(screen)
}

//NextState ...
func (c *Credits) NextState() string {
	return c.nextStateID
}

//End ...
func (c *Credits) End() {

}
