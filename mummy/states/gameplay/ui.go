package gameplay

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/programatta/mummygo/utils"
	"golang.org/x/image/font"
)

//UIGame ...
type UIGame struct {
	arcadeFont font.Face
	lives      int
	potions    int
	level      int
	score      int
}

//NewUIGame is a constructor
func NewUIGame(fontloader *utils.FontsLoader) *UIGame {
	ui := &UIGame{}

	ui.arcadeFont = fontloader.GetFont("BarcadeBrawl.ttf", 72, 12)
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

//Draw ...
func (ui *UIGame) Draw(screen *ebiten.Image) {
	uistring := fmt.Sprintf("Lives:%02d Potions:%d Level:%02d Score:%05d ", ui.lives, ui.potions, ui.level, ui.score)

	fontSize := 12
	screenWidth, screenHeight := screen.Size()
	x := (screenWidth - len(uistring)*fontSize) / 2
	y := screenHeight - fontSize
	text.Draw(screen, uistring, ui.arcadeFont, x, y, color.Black)
}
