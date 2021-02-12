package menu

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/programatta/mummygo/utils"
	"golang.org/x/image/font"
)

//UIMenu ...
type UIMenu struct {
	arcadeFontTitle font.Face
	arcadeFontDesc  font.Face
}

//NewUIMenu es el constructor.
func NewUIMenu(fontloader *utils.FontsLoader) *UIMenu {
	ui := UIMenu{}
	ui.arcadeFontTitle = fontloader.GetFont("BarcadeBrawl.ttf", 72, 40)
	ui.arcadeFontDesc = fontloader.GetFont("BarcadeBrawl.ttf", 72, 20)
	return &ui
}

//Draw ...
func (ui *UIMenu) Draw(screen *ebiten.Image) {
	screenWidth, _ := screen.Size()

	//Titulo.
	fztitle := 40
	uititle := fmt.Sprintf("Mummy GO")

	xt := (screenWidth - len(uititle)*fztitle) / 2
	yt := 4 * fztitle
	text.Draw(screen, uititle, ui.arcadeFontTitle, xt, yt, color.Black)

	//Opciones.
	fzdescription := 20
	descriptions := []string{"Press '1' key to play", "Press '2' key to see credits", "Press 'ESC' key to exit"}

	for i, desc := range descriptions {
		ll := len(desc) * fzdescription
		x := (screenWidth - ll) / 2
		y := (15 + 2*i) * fzdescription
		text.Draw(screen, desc, ui.arcadeFontDesc, x, y, color.Black)
	}
}
