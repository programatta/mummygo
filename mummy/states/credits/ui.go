package credits

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/programatta/mummygo/utils"
	"golang.org/x/image/font"
)

//UICredits ...
type UICredits struct {
	arcadeFontTitle font.Face
	arcadeFontDesc  font.Face
	descPosY        float64
	imgScroll       *ebiten.Image
	imgHeader       *ebiten.Image
	imgFooter       *ebiten.Image
	waitTime        float64
}

//NewUICredits es el constructor.
func NewUICredits(fontloader *utils.FontsLoader) *UICredits {
	ui := UICredits{}
	ui.arcadeFontTitle = fontloader.GetFont("BarcadeBrawl.ttf", 72, 30)
	ui.arcadeFontDesc = fontloader.GetFont("BarcadeBrawl.ttf", 72, 16)
	ui.descPosY = 80 //tamaño de la cabecera.
	ui.waitTime = 5  //segundos.
	return &ui
}

//Update ...
func (ui *UICredits) Update(dt float64) {
	_, screenHeight := ebiten.WindowSize()
	if ui.waitTime > 0 {
		ui.waitTime -= dt
	} else {
		ui.descPosY -= dt * 30
		if math.Abs(ui.descPosY) >= float64(maxScrollView-80) {
			ui.descPosY = float64(screenHeight) - 54
		}
	}
}

//Draw ...
func (ui *UICredits) Draw(screen *ebiten.Image) {
	//Dibujamos primero el scroll
	ui.drawScroll(screen)

	//Dubujamos las partes estaticas.
	ui.drawHeader(screen)
	ui.drawFooter(screen)
}

//Reset ...
func (ui *UICredits) Reset() {
	ui.descPosY = 80 //tamaño de la cabecera.
	ui.waitTime = 5  //segundos
}

/*===========================================================================*/
/*                               Private Section                             */
/*===========================================================================*/
func (ui *UICredits) drawHeader(screen *ebiten.Image) {
	screenWidth, _ := screen.Size()

	//Titulo.
	fztitle := 30
	uititle := fmt.Sprintf("-Credits-")

	xt := (screenWidth - len(uititle)*fztitle) / 2 // + 20
	yt := 2 * fztitle

	//Creamos la imagen de la cabecera.
	op := &ebiten.DrawImageOptions{}
	sx, _ := ebiten.WindowSize()
	if ui.imgHeader == nil {
		ui.imgHeader, _ = ebiten.NewImage(sx, yt+20, ebiten.FilterDefault)
		ui.imgHeader.Fill(color.NRGBA{0xCE, 0x9C, 0x72, 0xff})
	}
	op.GeoM.Translate(0, 0)

	text.Draw(ui.imgHeader, uititle, ui.arcadeFontTitle, xt, yt, color.White)
	screen.DrawImage(ui.imgHeader, op)
}

func (ui *UICredits) drawFooter(screen *ebiten.Image) {
	screenWidth, _ := screen.Size()
	fzdescription := 16

	//.:Volver al menu  principal.
	backDesc := "Press 'ESC' key to back main menu"
	x := (screenWidth - len(backDesc)*fzdescription) / 2
	y := fzdescription * 2

	op := &ebiten.DrawImageOptions{}
	sx, sy := ebiten.WindowSize()
	if ui.imgFooter == nil {
		ui.imgFooter, _ = ebiten.NewImage(sx, y+12, ebiten.FilterDefault)
		ui.imgFooter.Fill(color.NRGBA{0xCE, 0x9C, 0x72, 0xff})
	}
	op.GeoM.Translate(0, float64(sy)-44)

	text.Draw(ui.imgFooter, backDesc, ui.arcadeFontDesc, x, y, color.Black)
	screen.DrawImage(ui.imgFooter, op)
}

func (ui *UICredits) drawScroll(screen *ebiten.Image) {
	screenWidth, _ := screen.Size()
	op := &ebiten.DrawImageOptions{}
	sx, _ := ebiten.WindowSize()
	if ui.imgScroll == nil {
		ui.imgScroll, _ = ebiten.NewImage(sx, maxScrollView, ebiten.FilterDefault)
		ui.imgScroll.Fill(color.NRGBA{0xff, 0xff, 0xff, 0x00}) //Imagen con alfa = 0 (transparente)
	}
	op.GeoM.Translate(0, ui.descPosY)

	//Descriptions.
	//.:Codigo, sonido y fx.
	fzdescription := 16
	descriptions := []string{
		"Code: Programatta",
		"*",
		"Sound/FX: freesound.org",
		"gabisaraceni",
		"nebyoolae",
		"Mativve",
		"Haramir",
		"EminYILDIRIM",
		"MATRIXXX_",
		"jalastram",
		"Mrthenoronha",
		"MakoFox",
		"bongmoth",
		"Jamius",
		"spookymodem",
		"LittleRobotSoundFactory",
		"Michel88",
		"Antimsounds",
		"d761747",
		"*",
		"Music: freesound.org",
		"OllieOllie",
		"Jedo",
		"LittleRobotSoundFactory",
		"*",
		"Graphics: - ",
	}

	for i, desc := range descriptions {
		ll := len(desc) * fzdescription

		x := (screenWidth - ll) / 2
		y := (2 + 2*i) * fzdescription
		text.Draw(ui.imgScroll, desc, ui.arcadeFontDesc, x, y, color.Black)
	}
	screen.DrawImage(ui.imgScroll, op)
}

const maxScrollView int = 908
