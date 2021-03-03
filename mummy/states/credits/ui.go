package credits

import (
	"fmt"
	"image"
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
	scrollNormal    bool
	scrollViewHight int
}

//NewUICredits es el constructor.
func NewUICredits(fontloader *utils.FontsLoader) *UICredits {
	ui := UICredits{}
	ui.arcadeFontTitle = fontloader.GetFont("BarcadeBrawl.ttf", 72, 30)
	ui.arcadeFontDesc = fontloader.GetFont("BarcadeBrawl.ttf", 72, 16)
	ui.descPosY = 0
	ui.waitTime = maxWaitTimeInSec
	ui.scrollNormal = true
	return &ui
}

//Update ...
func (ui *UICredits) Update(dt float64) {
	if ui.waitTime > 0 {
		ui.waitTime -= dt
	} else {
		ui.descPosY += dt * 30
		if ui.scrollNormal {
			if math.Abs(ui.descPosY) >= float64(maxScrollSize) {
				ui.descPosY = 0
				ui.scrollNormal = false
			}
		} else {
			if float64(ui.scrollViewHight)-ui.descPosY <= 0 {
				ui.descPosY = 0
				ui.scrollNormal = true
			}
		}
	}
}

//Draw ...
func (ui *UICredits) Draw(screen *ebiten.Image, alfa float64) {
	ui.drawHeader(screen, alfa)
	ui.drawFooter(screen, alfa)
	ui.drawScroll(screen, alfa)
}

//Reset ...
func (ui *UICredits) Reset() {
	ui.descPosY = 0
	ui.waitTime = maxWaitTimeInSec
	ui.scrollNormal = true
}

/*===========================================================================*/
/*                               Private Section                             */
/*===========================================================================*/
func (ui *UICredits) drawHeader(screen *ebiten.Image, alfa float64) {
	halfa := 0xff * alfa
	screenWidth, _ := screen.Size()

	//Titulo.
	fztitle := 30
	uititle := fmt.Sprintf("-Credits-")

	xt := (screenWidth - len(uititle)*fztitle) / 2
	yt := 2 * fztitle

	//Creamos la imagen de la cabecera.
	op := &ebiten.DrawImageOptions{}
	sx, _ := ebiten.WindowSize()
	if ui.imgHeader == nil {
		ui.imgHeader, _ = ebiten.NewImage(sx, yt+20, ebiten.FilterDefault)
	}
	ui.imgHeader.Fill(color.NRGBA{0xCE, 0x9C, 0x72, uint8(halfa)})
	op.GeoM.Translate(0, 0)

	text.Draw(ui.imgHeader, uititle, ui.arcadeFontTitle, xt, yt, color.NRGBA{0xff, 0xff, 0xff, uint8(halfa)})
	screen.DrawImage(ui.imgHeader, op)
}

func (ui *UICredits) drawFooter(screen *ebiten.Image, alfa float64) {
	halfa := 0xff * alfa
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
	}
	ui.imgFooter.Fill(color.NRGBA{0xCE, 0x9C, 0x72, uint8(halfa)})
	op.GeoM.Translate(0, float64(sy)-44)

	text.Draw(ui.imgFooter, backDesc, ui.arcadeFontDesc, x, y, color.NRGBA{0x00, 0x00, 0x00, uint8(halfa)})
	screen.DrawImage(ui.imgFooter, op)
}

func (ui *UICredits) drawScroll(screen *ebiten.Image, alfa float64) {
	halfa := 0xff * alfa

	screenWidth, _ := screen.Size()
	op := &ebiten.DrawImageOptions{}
	sx, sy := ebiten.WindowSize()
	if ui.imgScroll == nil {
		ui.imgScroll, _ = ebiten.NewImage(sx, maxScrollSize, ebiten.FilterDefault)
	}
	ui.imgScroll.Fill(color.NRGBA{0xCE, 0x9C, 0x72, uint8(halfa)})

	_, headerHeight := ui.imgHeader.Size()
	_, footerHeight := ui.imgFooter.Size()

	// posición debajo de la cabecera.
	op.GeoM.Translate(0, float64(headerHeight))

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
		text.Draw(ui.imgScroll, desc, ui.arcadeFontDesc, x, y, color.NRGBA{0x00, 0x00, 0x00, uint8(halfa)})
	}

	//Ventana del scroll
	ui.scrollViewHight = sy - headerHeight - footerHeight
	if ui.scrollNormal {
		screen.DrawImage(ui.imgScroll.SubImage(image.Rect(0, 0+int(ui.descPosY), sx, ui.scrollViewHight+int(ui.descPosY))).(*ebiten.Image), op)

		if ui.scrollViewHight+int(ui.descPosY) > maxScrollSize {
			//completamos con una imagen temporal el hueco dejado por la imagen
			//de scroll.
			p := ui.imgScroll.SubImage(image.Rect(0, 0+int(ui.descPosY), sx, ui.scrollViewHight+int(ui.descPosY))).Bounds().Size()

			op2 := &ebiten.DrawImageOptions{}
			offset := ui.scrollViewHight - p.Y
			imgtmp, _ := ebiten.NewImage(p.X, offset, ebiten.FilterDefault)
			imgtmp.Fill(color.NRGBA{0xCE, 0x9C, 0x72, uint8(halfa)})
			op2.GeoM.Translate(0, float64(p.Y)+float64(headerHeight))
			screen.DrawImage(imgtmp, op2)
		}
	} else {
		//completamos con una imagen temporal ya que hemos bajado la imagen de
		//scroll a la parte alta del footer.
		op2 := &ebiten.DrawImageOptions{}
		sx2, _ := ebiten.WindowSize()
		offset := ui.scrollViewHight - int(ui.descPosY)
		imgtmp, _ := ebiten.NewImage(sx2, offset, ebiten.FilterDefault)
		imgtmp.Fill(color.NRGBA{0xCE, 0x9C, 0x72, uint8(halfa)})
		op2.GeoM.Translate(0, float64(headerHeight))
		screen.DrawImage(imgtmp, op2)

		//movemos la imagen de scroll a la parte alta del footer para ir su-
		//biendo y mostrando texto.
		op.GeoM.Translate(0, float64(offset))
		screen.DrawImage(ui.imgScroll.SubImage(image.Rect(0, 0, sx, int(ui.descPosY))).(*ebiten.Image), op)
	}

}

const (
	maxScrollSize    int     = 908 //908 sin texto debajo de graphics.
	maxWaitTimeInSec float64 = 3
)
