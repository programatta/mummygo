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
	alfa        float64
	state       tstate
}

//NewCredits es un contructor
func NewCredits(fontsloader *utils.FontsLoader) states.IState {
	c := &Credits{}
	c.nextStateID = "credits"

	c.uicredits = NewUICredits(fontsloader)
	c.alfa = 0
	c.state = enter
	return c
}

/*===========================================================================*/
/*                               Interface IState                            */
/*===========================================================================*/

//Init ...
func (c *Credits) Init() {
	c.nextStateID = "credits"
	c.alfa = 0
	c.state = enter
	c.uicredits.Reset()
}

//ProcessEvents procesa los eventos del juego.
func (c *Credits) ProcessEvents() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		c.nextStateID = "menu"
		c.state = exit
	}
}

//Update actualiza la lógica de los creditos.
func (c *Credits) Update(dt float64) {

	if c.state == enter {
		c.alfa += dt
		if c.alfa > 1 {
			c.alfa = 1
			c.state = normal
		}
	} else if c.state == exit {
		c.alfa -= dt
		if c.alfa < 0 {
			c.alfa = 0
			c.state = end
		}
	}
	c.uicredits.Update(dt)
}

//Draw draws the game.
func (c *Credits) Draw(screen *ebiten.Image) {
	halfa := 0xff * c.alfa

	screen.Fill(color.NRGBA{0xCE, 0x9C, 0x72, uint8(halfa)})
	c.uicredits.Draw(screen, c.alfa)
}

//NextState ...
func (c *Credits) NextState() string {
	if c.state == end {
		return c.nextStateID
	}
	return "credits"
}

//End ...
func (c *Credits) End() {

}

type tstate int

const (
	enter  tstate = tstate(0)
	normal tstate = tstate(1)
	exit   tstate = tstate(2)
	end    tstate = tstate(3)
)
