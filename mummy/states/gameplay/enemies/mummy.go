package enemies

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/programatta/mummygo/utils"
)

//Mummy ...
type Mummy struct {
	spriteSheet   *utils.SpriteSheet
	posX, posY    float64
	animations    map[string]*utils.Animation
	currentDir    string
	currentDirPos int
	//demosteps     int
	dirs  []string
	state tmummyState
	sc    float64
	toY   float64
}

//NewMummy is a constructor.
func NewMummy(spriteSheet *utils.SpriteSheet, x, y int) *Mummy {
	mummy := &Mummy{}

	mummy.spriteSheet = spriteSheet
	mummy.posX = float64(x)
	mummy.posY = float64(y)

	//Prepare animations.
	mummy.animations = make(map[string]*utils.Animation)

	//Down.
	animaDown := utils.NewAnimation()
	animaDown.AddFrame(spriteSheet.GetFrameByName("mummy-0.png"))
	animaDown.AddFrame(spriteSheet.GetFrameByName("mummy-1.png"))
	animaDown.AddFrame(spriteSheet.GetFrameByName("mummy-2.png"))
	animaDown.AddFrame(spriteSheet.GetFrameByName("mummy-1.png"))
	mummy.animations["down"] = animaDown

	//Left.
	animaLeft := utils.NewAnimation()
	animaLeft.AddFrame(spriteSheet.GetFrameByName("mummy-3.png"))
	animaLeft.AddFrame(spriteSheet.GetFrameByName("mummy-4.png"))
	animaLeft.AddFrame(spriteSheet.GetFrameByName("mummy-5.png"))
	animaLeft.AddFrame(spriteSheet.GetFrameByName("mummy-4.png"))
	mummy.animations["left"] = animaLeft

	//Right.
	animaRight := utils.NewAnimation()
	animaRight.AddFrame(spriteSheet.GetFrameByName("mummy-6.png"))
	animaRight.AddFrame(spriteSheet.GetFrameByName("mummy-7.png"))
	animaRight.AddFrame(spriteSheet.GetFrameByName("mummy-8.png"))
	animaRight.AddFrame(spriteSheet.GetFrameByName("mummy-7.png"))
	mummy.animations["right"] = animaRight

	//Up.
	animaUp := utils.NewAnimation()
	animaUp.AddFrame(spriteSheet.GetFrameByName("mummy-9.png"))
	animaUp.AddFrame(spriteSheet.GetFrameByName("mummy-10.png"))
	animaUp.AddFrame(spriteSheet.GetFrameByName("mummy-11.png"))
	animaUp.AddFrame(spriteSheet.GetFrameByName("mummy-10.png"))
	mummy.animations["up"] = animaUp

	mummy.currentDir = "down"
	//mummy.demosteps = 0
	mummy.currentDirPos = 0
	mummy.dirs = []string{"down", "left", "right", "up"}

	mummy.sc = 0
	mummy.state = mummyShowing
	return mummy
}

//Update updates a mummy logic
func (m *Mummy) Update(dt float64) {
	// m.demosteps++
	// if m.demosteps%600 == 0 {
	// 	m.demosteps = 0
	// 	m.currentDirPos = (m.currentDirPos + 1) % 4
	// 	m.currentDir = m.dirs[m.currentDirPos]
	// }
	m.animations[m.currentDir].UpdateFrame()
	if m.state == mummyShowing {
		if m.sc < 1 {
			m.sc += dt
		} else {
			m.sc = 1.0
			m.state = mummyLeaving
			m.toY = m.posY + 32
		}
	} else if m.state == mummyLeaving {
		if m.posY < m.toY {
			m.posY = m.posY + 32*dt //0.016
		} else if m.posY >= m.toY {
			m.posY = m.toY
			m.toY = 0
			m.state = mummyLookingfor
		}
	}
}

//Draw renders a current frame.
func (m *Mummy) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(m.sc, m.sc)
	op.GeoM.Translate(m.posX, m.posY)
	op.GeoM.Translate(16.0-(32*m.sc)*0.5, 16.0-(32*m.sc)*0.5)

	texture := m.spriteSheet.GetTexture()

	frameData := m.animations[m.currentDir].GetFrame()
	screen.DrawImage(texture.SubImage(image.Rect(frameData.X, frameData.Y, frameData.X+frameData.W, frameData.Y+frameData.H)).(*ebiten.Image), op)
}

//Position devuelve las coordenadas X e Y de la momia.
func (m *Mummy) Position() (float64, float64) {
	return m.posX, m.posY
}

type tmummyState int

const (
	mummyShowing    tmummyState = tmummyState(0)
	mummyLeaving    tmummyState = tmummyState(1)
	mummyLookingfor tmummyState = tmummyState(2)
)
