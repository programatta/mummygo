package credits

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/programatta/mummygo/utils"
	"golang.org/x/image/font"
)

//UICredits ...
type UICredits struct {
	arcadeFontTitle font.Face
	arcadeFontDesc  font.Face
}

//NewUICredits es el constructor.
func NewUICredits(fontloader *utils.FontsLoader) *UICredits {
	ui := UICredits{}
	ui.arcadeFontTitle = fontloader.GetFont("BarcadeBrawl.ttf", 72, 30)
	ui.arcadeFontDesc = fontloader.GetFont("BarcadeBrawl.ttf", 72, 16)
	return &ui
}

//Draw ...
func (ui *UICredits) Draw(screen *ebiten.Image) {
	screenWidth, screenHeight := screen.Size()

	//Titulo.
	fztitle := 30
	uititle := fmt.Sprintf("-Credits-")

	xt := (screenWidth - len(uititle)*fztitle) / 2 // + 20
	yt := 2 * fztitle
	text.Draw(screen, uititle, ui.arcadeFontTitle, xt, yt, color.White)

	//Descriptions.
	//.:Codigo, sonido y fx.
	fzdescription := 16
	descriptions := []string{"Code: Programatta", "Sound/FX: - ", "Music: - ", "Graphics: - "}
	//offsets := []int{24 * 2, 24 * 2, 24 * 2, 24 * 2}
	for i, desc := range descriptions {
		ll := len(desc) * fzdescription

		x := (screenWidth - ll) / 2 //+ offsets[i]
		y := (11 + 2*i) * fzdescription
		text.Draw(screen, desc, ui.arcadeFontDesc, x, y, color.Black)
	}

	//.:Volver al menu  principal.
	backDesc := "Press 'ESC' key to back main menu"
	x := (screenWidth - len(backDesc)*fzdescription) / 2 //+ 16*5
	y := screenHeight - fzdescription*2
	text.Draw(screen, backDesc, ui.arcadeFontDesc, x, y, color.Black)
}
