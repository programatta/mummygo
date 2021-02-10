package gameplay

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

//UIGame ...
type UIGame struct {
	arcadeFont font.Face
	lives      int
	potions    int
	level      int
}

//NewUIGame is a constructor
func NewUIGame() *UIGame {
	ui := &UIGame{}

	f, e := os.Open("assets/fonts/ka1.ttf")
	if e != nil {
		log.Fatal(e)
	}
	reader := bufio.NewReader(f)

	fontData := make([]byte, 0)
	var data []byte
	data = make([]byte, 10240)
	for {
		n, err := reader.Read(data)
		if err != nil {
			log.Fatal(err)
		}
		fontData = append(fontData, data...)
		if n < 10240 {
			break
		}
	}
	ttfont, err := truetype.Parse(fontData)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	ui.arcadeFont = truetype.NewFace(ttfont, &truetype.Options{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

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

//Draw ...
func (ui *UIGame) Draw(screen *ebiten.Image) {
	uistring := fmt.Sprintf("Lives:%02d  Potions:%02d  Level:%02d", ui.lives, ui.potions, ui.level)

	screenWidth, screenHeight := screen.Size()
	text.Draw(screen, uistring, ui.arcadeFont, (screenWidth - len(uistring)*int(fontSize)), screenHeight-int(fontSize/2), color.Black)
}

const fontSize float64 = 18
