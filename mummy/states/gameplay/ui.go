package gameplay

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/programatta/mummygo/utils"
	"golang.org/x/image/font"
)

//UIGame ...
type UIGame struct {
	arcadeFont  font.Face
	spriteSheet *utils.SpriteSheet
	lives       int
	potions     int
	level       int
	score       int
	haskey      bool
	haspapyre   bool
	percent     float64
	warnlevel   int
}

//NewUIGame is a constructor
func NewUIGame(fontloader *utils.FontsLoader, spriteSheet *utils.SpriteSheet) *UIGame {
	ui := &UIGame{}

	ui.arcadeFont = fontloader.GetFont("BarcadeBrawl.ttf", 72, 12)
	ui.spriteSheet = spriteSheet
	return ui
}

//SetLives ...
func (ui *UIGame) SetLives(lives int) {
	ui.lives = lives
}

//SetPotions ...
func (ui *UIGame) SetPotions(potions int) {
	ui.potions = potions
}

//SetLevel ...
func (ui *UIGame) SetLevel(level int) {
	ui.level = level
}

//SetScore ...
func (ui *UIGame) SetScore(score int) {
	ui.score = score
}

//SetPapyre ...
func (ui *UIGame) SetPapyre(hasPapyre bool) {
	ui.haspapyre = hasPapyre
}

//SetKey ...
func (ui *UIGame) SetKey(hasKey bool) {
	ui.haskey = hasKey
}

//SetPercentOxigen ...
func (ui *UIGame) SetPercentOxigen(percent float64, warnlevel int) {
	ui.percent = percent
	ui.warnlevel = warnlevel
}

//Draw ...
func (ui *UIGame) Draw(screen *ebiten.Image) {
	ui.drawPlayerInfo(screen)
	ui.drawPotionsInfo(screen)
	ui.drawCollectablesInfo(screen)
	ui.drawLevelScoreInfo(screen)

	ui.drawOxigenInfo(screen)
}

/*===========================================================================*/
/*                               Private Section                             */
/*===========================================================================*/

func (ui *UIGame) drawPlayerInfo(screen *ebiten.Image) {
	_, screenHeight := screen.Size()
	fontSize := 12

	framaData := ui.spriteSheet.GetFrameByName("player-1.png")
	texture := ui.spriteSheet.GetTexture()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(45, float64(screenHeight-framaData.H))
	screen.DrawImage(texture.SubImage(
		image.Rect(framaData.X, framaData.Y, framaData.X+framaData.W, framaData.Y+framaData.H)).(*ebiten.Image),
		op)

	uistring := fmt.Sprintf("x%02d", ui.lives)
	y := screenHeight - fontSize
	x := 45 + framaData.W
	text.Draw(screen, uistring, ui.arcadeFont, x, y, color.Black)
}

func (ui *UIGame) drawPotionsInfo(screen *ebiten.Image) {
	_, screenHeight := screen.Size()
	fontSize := 12

	framaData := ui.spriteSheet.GetFrameByName("object-0.png")
	texture := ui.spriteSheet.GetTexture()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(125, float64(screenHeight-framaData.H))
	screen.DrawImage(texture.SubImage(
		image.Rect(framaData.X, framaData.Y, framaData.X+framaData.W, framaData.Y+framaData.H)).(*ebiten.Image),
		op)

	uistring := fmt.Sprintf("x%02d", ui.potions)
	y := screenHeight - fontSize
	x := 125 + framaData.W
	text.Draw(screen, uistring, ui.arcadeFont, x, y, color.Black)
}

func (ui *UIGame) drawCollectablesInfo(screen *ebiten.Image) {
	_, screenHeight := screen.Size()

	keyFramaData := ui.spriteSheet.GetFrameByName("object-3.png")
	papyreFramaData := ui.spriteSheet.GetFrameByName("object-14.png")
	texture := ui.spriteSheet.GetTexture()

	//Llave.
	keyalpha := 0.45
	if ui.haskey {
		keyalpha = 1.0
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(215, float64(screenHeight-keyFramaData.H))
	op.ColorM.Scale(1.0, 1.0, 1.0, keyalpha)
	screen.DrawImage(texture.SubImage(
		image.Rect(keyFramaData.X, keyFramaData.Y, keyFramaData.X+keyFramaData.W, keyFramaData.Y+keyFramaData.H)).(*ebiten.Image),
		op)

	//Papiro.
	papyrealpha := 0.45
	if ui.haspapyre {
		papyrealpha = 1.0
	}
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(255, float64(screenHeight-papyreFramaData.H))
	op2.ColorM.Scale(1.0, 1.0, 1.0, papyrealpha)
	screen.DrawImage(texture.SubImage(
		image.Rect(papyreFramaData.X, papyreFramaData.Y, papyreFramaData.X+papyreFramaData.W, papyreFramaData.Y+papyreFramaData.H)).(*ebiten.Image),
		op2)
}

func (ui *UIGame) drawLevelScoreInfo(screen *ebiten.Image) {
	uistring := fmt.Sprintf("Level:%02d Score:%05d ", ui.level, ui.score)

	fontSize := 12
	_, screenHeight := screen.Size()
	x := 304
	y := screenHeight - fontSize

	//Texto.
	text.Draw(screen, uistring, ui.arcadeFont, x, y, color.Black)
}

func (ui *UIGame) drawOxigenInfo(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	sx, sy := ebiten.WindowSize()

	oxigenBarSize := sx - 90

	if ui.percent > 0 {
		oxigenBarSize = int(float64(oxigenBarSize) - (float64(oxigenBarSize)*ui.percent)/100)
	}

	oximg, _ := ebiten.NewImage(oxigenBarSize, 16, ebiten.FilterDefault)

	if ui.warnlevel == 0 {
		oximg.Fill(color.RGBA{0x5D, 0xCB, 0xEC, 0xff})
	} else if ui.warnlevel == 1 {
		oximg.Fill(color.RGBA{0xFF, 0xF6, 0x00, 0xff})
	} else if ui.warnlevel == 2 {
		oximg.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	}

	//op.GeoM.Scale(float64(sx), float64(sy))
	op.GeoM.Translate(45, float64(sy)-56)

	op.ColorM.Scale(1.0, 1.0, 1.0, 1.0)
	screen.DrawImage(oximg, op)
}
